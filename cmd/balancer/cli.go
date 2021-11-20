package balancer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type Ident [20]byte

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

type SRequest struct {
}

type SResponse struct {
	Predecessor   Entry
	Successor     Entry
	SuccessorList []Entry
}

type Response struct {
	Entry Entry
	Value string
	Ok    bool
}

func StabilizeLoop(done chan interface{}, chord *Node) {
Loop:
	for {
		select {
		case <-done:
			break Loop
		default:
			time.Sleep(1 * time.Second)
			chord.stabilize()
		}
	}
}

func Looper(s Settings) {

	if s.Address == "" {
		s.Address = GetLocalAddress()
	}

	fmt.Printf("Listening on: %s\n", s.Port)
	fmt.Printf("Ring Join: %t\n", s.Ring)
	fmt.Printf("Join address: %s\n", s.Join)
	fmt.Printf("Node address: %s\n", s.Address)

	addr := fmt.Sprint(s.Address, ":", s.Port)
	fmt.Println("Address full:")
	fmt.Println(addr)

	var successor *Entry

	// Create a new ring
	if s.Ring {
		successor = &Entry{
			IpAddr:     s.Address,
			Port:       s.Port,
			Identifier: hashString(AddrToIpPort(s.Address, s.Port)),
		}
	} else {
		// Join a ring specified
		joinAddrSplit := strings.SplitN(s.Join, ":", 2)

		successor = &Entry{
			IpAddr:     joinAddrSplit[0],
			Port:       joinAddrSplit[1],
			Identifier: hashString(s.Join),
		}
	}

	chord := NewChordNode(s.Address, s.Port, successor)

	// New RPC Server with the configured settings
	rpc := RPCServer{
		Settings: s,
	}

	kv := make(map[string]string)
	kvMutex := sync.RWMutex{}

	store := &Store{
		KeyValue: kv,
		mutex:    &kvMutex,
	}

	// Listen for RPC requests in a separate go routine
	go rpc.init("127.0.0.1:3001", store, chord)

	done := make(chan interface{})

	go StabilizeLoop(done, chord)

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

		command := args[0]

		switch command {
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
		case "dump":
			fmt.Println("Dumping Results:")
			store.mutex.RLock()
			for k, v := range store.KeyValue {
				fmt.Println(k, v)
			}
			store.mutex.RUnlock()

		case "join":
			targetAddr := strings.SplitN(args[1], ":", 2)
			fmt.Printf("Joining Address: %s\n", targetAddr)
			newEntry := NewEntry(targetAddr[0], targetAddr[1])
			chord.SuccessorList[1] = newEntry
			chord.SuccessorList = append(chord.SuccessorList, chord.SuccessorList[1])
			fmt.Printf("Successfully joined Address: %s\n", chord.SuccessorList[1].IpAddr)

		default:
			fmt.Println("Invalid command.")
		}
	}
}
