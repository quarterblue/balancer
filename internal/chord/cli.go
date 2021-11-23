package chord

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/quarterblue/balancer/proto"
)

// Configuration settings
type Settings struct {
	// IP address of the current node
	Address string

	// Represents the port this node will listen on
	Port string

	// IP address of the node to contact to join the ring initially
	Join string

	// Boolean flag to indicate whether to create a new ring
	Ring bool

	// Boolean flag to indicate whether to start interactive CLI
	Cli bool
}

func StabilizeLoop(done chan interface{}, chord *Chord) {
Loop:
	for {
		select {
		case <-done:
			break Loop
		default:
			time.Sleep(1 * time.Second)
			chord.Stabilize()
		}
	}
}

func InitializeChord(s Settings) (*Chord, *RPCServer, *Store) {
	if s.Address == "" {
		s.Address = GetLocalAddress()
	}

	fmt.Printf("Listening on: %s:%s\n", s.Address, s.Port)
	fmt.Printf("Ring Join: %t\n", s.Ring)
	fmt.Printf("Join address: %s\n", s.Join)

	var successor *Node
	var chord *Chord

	// Create a new ring
	if s.Ring {
		successor = &Node{
			IpAddr:     s.Address,
			Port:       s.Port,
			Identifier: hashString(AddrToIpPort(s.Address, s.Port)),
		}
		chord = NewChordNode(s.Address, s.Port, successor)
	} else {
		// Join a ring specified
		joinAddrSplit := strings.SplitN(s.Join, ":", 2)

		chord = NewChordNode(s.Address, s.Port, nil)

		successor = &Node{
			IpAddr:     joinAddrSplit[0],
			Port:       joinAddrSplit[1],
			Identifier: hashString(s.Join),
		}

		chord.Join(successor)
	}

	// New RPC Server with the configured settings
	rpc := &RPCServer{
		Settings: s,
	}

	kv := make(map[string]string)
	kvMutex := sync.RWMutex{}

	store := &Store{
		KeyValue: kv,
		mutex:    &kvMutex,
	}

	return chord, rpc, store
}

func ExecutionLoop(s Settings) {

	chord, rpc, store := InitializeChord(s)
	// Listen for RPC requests in a separate go routine
	go rpc.init("127.0.0.1:3001", store, chord)

	// done := make(chan interface{})

	// go StabilizeLoop(done, chord)

	log.Printf("Interactive shell")
	log.Printf("Commands: ping, get, post")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("[balancer] CMD-> ")
	for scanner.Scan() {
		fmt.Printf("[balancer] OUT-> ")
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

		var response *proto.KVResponse
		var err error

		switch command {
		case "ping":
			targetAddr := args[1]

			request := &proto.KVRequest{
				Key:   "",
				Value: "",
			}

			if response, err = gRpcCall(targetAddr, "ping", request); err != nil {
				log.Fatalf("Calling Store.Ping: %v", err)
			}

			log.Println(response)

		case "get":
			targetAddr := args[1]
			key := args[2]

			request := &proto.KVRequest{
				Key:   key,
				Value: "",
			}
			if response, err = gRpcCall(targetAddr, "get", request); err != nil {
				log.Fatalf("Calling Store.Get: %v", err)
			}

			log.Println(response)

		case "put":
			targetAddr := args[1]
			key := args[2]
			value := args[3]

			request := &proto.KVRequest{
				Key:   key,
				Value: value,
			}

			if response, err = gRpcCall(targetAddr, "put", request); err != nil {
				log.Fatalf("Calling Store.Put: %v", err)
			}

			log.Println(response)

		case "dump":
			fmt.Println("Dumping Results:")
			store.mutex.RLock()
			for k, v := range store.KeyValue {
				fmt.Println(k, v)
			}
			store.mutex.RUnlock()

		case "succ":
			chord.SuccMutex.RLock()
			fmt.Printf("Successor: %s\n", chord.successor().IpAddrString())
			chord.SuccMutex.RUnlock()

		case "pred":
			chord.PredMutex.Lock()
			if pred := chord.predecessor(); pred == nil {
				fmt.Printf("Predecessor: nil\n")
			} else {
				fmt.Printf("Predecessor: %s\n", chord.predecessor().IpAddrString())
			}
			chord.PredMutex.Unlock()

		case "join":
			targetAddr := strings.SplitN(args[1], ":", 2)
			fmt.Printf("Joining Address: %s\n", targetAddr)
			newEntry := NewNode(targetAddr[0], targetAddr[1])
			chord.SuccessorList[1] = newEntry
			chord.SuccessorList = append(chord.SuccessorList, chord.SuccessorList[1])
			fmt.Printf("Successfully joined Address: %s\n", chord.SuccessorList[1].IpAddr)

		default:
			fmt.Println("Invalid command.")
		}
		fmt.Printf("[balancer] CMD-> ")
	}
}
