package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// Configuration settings
type Settings struct {
	// Represents the port this node will listen on
	Port string

	// Represents the port this node will listen on
	Ring bool

	// IP address of the node to contact to join the ring initially
	Join string

	// IP address of the current node
	Address string
}

// Finds out the IP address of the current machine
func GetLocalAddress() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func Looper(s Settings) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Listening on: %s\n", s.Port)
	fmt.Printf("Ring Join: %t\n", s.Ring)
	fmt.Printf("Join address: %s\n", s.Join)
	fmt.Printf("Node address: %s\n", s.Address)

	for {
		fmt.Print("Command: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
	}
}
