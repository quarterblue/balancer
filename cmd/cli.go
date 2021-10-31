package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"strings"
)

// Configuration settings
type Settings struct {
	// Represents the port this node will listen on
	Port string

	// Boolean flag to indicate whether to create a new ring
	Ring bool

	// IP address of the node to contact to join the ring initially
	Join string

	// IP address of the current node
	Address string
}

type Request struct {
	Key   string
	Value string
}

// Returns the IP address of the current machine
func GetLocalAddress() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func call(address string, method string, request Request, response interface{}) error {
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Printf("Error connecting: %v", err)
		return err
	}
	defer client.Close()

	if err = client.Call(method, request, response); err != nil {
		log.Printf("Client call: %s, %v", method, err)
		return err
	}

	return nil
}

func Looper(s Settings) {

	if s.Address == "" {
		s.Address = GetLocalAddress()
	}

	fmt.Printf("Listening on: %s\n", s.Port)
	fmt.Printf("Ring Join: %t\n", s.Ring)
	fmt.Printf("Join address: %s\n", s.Join)
	fmt.Printf("Node address: %s\n", s.Address)

	// New RPC Server with the configured settings
	r := RPCServer{
		Settings: s,
	}

	// Listen for RPC requests in a separate go routine
	go r.init("127.0.0.1:3001")

	log.Printf("Interactive shell")
	log.Printf("Commands: ping, get, post")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		args := strings.SplitN(line, " ", 4)
		if len(args) > 1 {
			args[1] = strings.TrimSpace(args[1])
		}

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "ping":
			targetAddr := args[1]
			msg := args[2]

			response := ""
			request := &Request{
				Key:   msg,
				Value: "",
			}
			if err := call(targetAddr, "Store.Ping", *request, &response); err != nil {
				log.Fatalf("Calling Store.Ping: %v", err)
			}
			fmt.Println(response)
		case "get":
			targetAddr := args[1]
			key := args[2]

			response := ""
			request := &Request{
				Key:   key,
				Value: "",
			}
			if err := call(targetAddr, "Store.Get", *request, &response); err != nil {
				log.Fatalf("Calling Store.Get: %v", err)
			}
			fmt.Println(response)
		case "put":
			targetAddr := args[1]
			key := args[2]
			value := args[3]

			request := &Request{
				Key:   key,
				Value: value,
			}

			response := ""
			if err := call(targetAddr, "Store.Put", *request, &response); err != nil {
				log.Fatalf("Calling Store.Put: %v", err)
			}
			fmt.Println(response)
		default:
			fmt.Println("Invalid command.")
		}
	}
}
