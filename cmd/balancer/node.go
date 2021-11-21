package balancer

import (
	"fmt"
	"math/big"
)

type Node struct {
	// IP Address of the entry
	IpAddr string

	// Port of the entry
	Port string

	// Identifer of the entry constructing by using a base hash function SHA-1 of IP Addr
	Identifier *big.Int
}

func NewEntry(ipAddr, port string) *Node {
	addr := ipAddr + ":" + port
	hash := hashString(addr)
	entry := &Node{
		IpAddr:     ipAddr,
		Port:       port,
		Identifier: hash,
	}
	return entry
}

func (n *Node) IpAddrString() string {
	return n.IpAddr + ":" + n.Port
}

func (n *Node) notify(c *Chord) {
	fmt.Println("Hello")
}

func (n *Node) FindSuccessor(c *Chord) *Node {
	targetAddr := AddrToIpPort(n.IpAddr, n.Port)
	fmt.Println(targetAddr)

	// response, err := gRpcCall(targetAddr, "ping", request)
	// if err != nil {
	// 	log.Println("Calling Store.Ping: %v", err)
	// }
	return nil
}
