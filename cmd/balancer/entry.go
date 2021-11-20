package balancer

import (
	"fmt"
	"math/big"
)

type Entry struct {
	// IP Address of the entry
	IpAddr string

	// Port of the entry
	Port string

	// Identifer of the entry constructing by using a base hash function SHA-1 of IP Addr
	Identifier *big.Int
}

func NewEntry(ipAddr, port string) *Entry {
	addr := ipAddr + ":" + port
	hash := hashString(addr)
	entry := &Entry{
		IpAddr:     ipAddr,
		Port:       port,
		Identifier: hash,
	}
	return entry
}

func (e *Entry) IpAddrString() string {
	return e.IpAddr + ":" + e.Port
}

func (e *Entry) notify(n *Node) {
	fmt.Println("Hello")
}
