package cmd

import (
	"context"
	"math/big"
	"sync"
)

// m-bit hash of the node IP Addr or the key (Default uses SHA-1)
type Identifier *big.Int

type FingerTable map[string]Finger

// Interface of the Chord Node
type ChordNode interface {
	FindSuccessor(Identifier) Node
	ClosestPrecedingNode(Identifier) Node
	Create()
	Join(Identifier)
	Stabilize()
	Notify(Identifier)
	FixFinger()
	CheckPredecessor() bool
}

// Implementation of the Chord Node
type Node struct {
	// The route table
	Finger FingerTable

	// The next node on the identifier circle; finger[1].node
	Successor Entry

	// The previous node on the identifier circle
	Predecessor Entry

	// The m-bit identifier
	Identifier Identifier

	// IP Address of the node
	IpAddr string

	// Port of the node
	Port string

	// Mutex for protecting
	Mutex sync.RWMutex
}

func NewChordNode(ipAddr, port string) *Node {
	addr := ipAddr + ":" + port
	iden := hashString(addr)
	node := &Node{
		IpAddr:     ipAddr,
		Port:       port,
		Identifier: iden,
	}

	return node
}

// func (n *Node) create() error {
// 	n.successor = Entry{
// 		ipAddr:     n.ipAddr,
// 		port:       n.port,
// 		identifier: n.identifier,
// 	}

// 	return nil
// }

type Finger struct {
	// (n + 2^(i-1)) mod 2^m (1 <= i <= m)
	start int

	// (n + 2^i - 1) mod 2^m
	end int

	// interval
	interval [2]int

	// succesor node
	successor int
}

type Entry struct {
	// IP Address of the entry
	ipAddr string

	// Port of the entry
	port string

	// Identifer of the entry constructing by using a base hash function SHA-1 of IP Addr
	identifier Identifier
}

func (n *Node) findSuccessor(ctx context.Context, id Identifier) (*Entry, error) {
	if withinFingerRange(n.Identifier, id) {
		// return n.finger[id], nil
		return nil, nil
	} else {
		// nPrime := closestPrecedingNode(id)
		// return nPrime.findSuccessor(ctx, id)

		return nil, nil
	}
}

func closestPrecedingNode(id Identifier) *Entry {
	return nil
}

func withinFingerRange(n, successor Identifier) bool {
	return true
}

func NewEntry(ipAddr, port string, identifier Identifier) *Entry {
	newEntry := Entry{
		ipAddr,
		port,
		identifier,
	}
	return &newEntry
}

func NewNode() {

}

func InitFingerTable() (FingerTable, error) {
	newFT := make(map[string]Finger)
	return newFT, nil
}
