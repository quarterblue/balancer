package main

import (
	"flag"
	"log"
	"os"

	chord "github.com/quarterblue/balancer/internal/chord"
)

func main() {
	port := flag.String("port", "3410", "Port that this node should listen on.")
	ring := flag.Bool("ring", false, "Boolean to start a ring.")
	join := flag.String("join", "", "IP Address to join initially.")
	address := flag.String("address", "", "This machines IP address.")
	cli := flag.Bool("cli", false, "Boolean to start interactive CLI mode.")
	flag.Parse()

	if *join == "" && !*ring {
		log.Printf("You must either create a ring or join a ring.\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	settings := chord.Settings{
		Port:    *port,
		Ring:    *ring,
		Join:    *join,
		Address: *address,
		Cli:     *cli,
	}

	chord.ExecutionLoop(settings)
	// fmt.Println(s)

	// hasher := balancer.XxHash{}
	// hashed := hasher.Hash([]byte("172.38.2.40:3006"))
	// fmt.Printf("Hashed: %d\n", hashed)
	// hashed2 := hasher.Hash([]byte("172.38.2.40:3007"))
	// fmt.Printf("Hashed: %d\n", hashed2)
}
