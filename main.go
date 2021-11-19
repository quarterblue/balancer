package main

import (
	"flag"
	"os"

	"github.com/quarterblue/balancer/cmd"
)

func main() {
	port := flag.String("port", "3410", "Port that this node should listen on.")
	ring := flag.Bool("ring", false, "Boolean to start a ring")
	join := flag.String("join", "172.38.18.10", "IP Address to join initially")
	address := flag.String("address", "", "This machines IP address")
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
