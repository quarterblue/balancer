package main

import (
	"flag"
	"fmt"
	"os"

	chord "github.com/quarterblue/balancer/internal/chord"
	balancer "github.com/quarterblue/balancer/pkg"
)

func main() {
	port := flag.String("port", "3410", "Port that this node should listen on.")
	ring := flag.Bool("ring", false, "Boolean to start a ring")
	join := flag.String("join", "127.0.0.1", "IP Address to join initially")
	address := flag.String("address", "", "This machines IP address")
	flag.Parse()

	if *join == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	s := chord.Settings{
		Port:    *port,
		Ring:    *ring,
		Join:    *join,
		Address: *address,
	}

	// balancer.Looper(s)
	fmt.Println(s)

	hasher := balancer.XxHash{}
	hashed := hasher.Hash([]byte("172.38.2.40:3006"))
	fmt.Printf("Hashed: %d\n", hashed)
	hashed2 := hasher.Hash([]byte("172.38.2.40:3007"))
	fmt.Printf("Hashed: %d\n", hashed2)
}
