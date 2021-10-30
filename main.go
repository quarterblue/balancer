package main

import (
	"flag"
	"os"

	"github.com/quarterblue/chord-dht/cmd"
)

func main() {
	port := flag.String("port", "3410", "Port that this node should listen on.")
	ring := flag.Bool("ring", false, "To create a ring")
	join := flag.String("join", "172.38.18.10", "To join address")
	address := flag.String("address", "127.0.0.1", "Your address")
	flag.Parse()

	if *join == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	s := cmd.Settings{
		Port:    *port,
		Ring:    *ring,
		Join:    *join,
		Address: *address,
	}

	cmd.Looper(s)
}
