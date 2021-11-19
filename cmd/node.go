package cmd

import (
	"context"
	"fmt"
	"math/big"
	"sync"
)

// Successor list size
const sListSize = 5

// m-bit hash of the node IP Addr or the key (Default uses SHA-1)
// current implementation stores a big int casted from the m-bit hash bytes
type Identifier *big.Int

type FingerTable map[string]Finger

// Interface of the Chord Node RPC, requires implementation of all node functionalities described in the paper
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
	// The finger route table described in the paper; maintains up to m entries (nodes)
	Finger FingerTable

	// The array of n first successors; this is to replicate the successors to deal with failures in nodes
	SuccessorList []*Entry

	// The next node on the identifier circle; finger[1].node
	Successor Entry

	// The previous node on the identifier circle
	Predecessor Entry

	// The m-bit identifier of the node
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
	sList := []*Entry{}
	node := &Node{
		IpAddr:        ipAddr,
		Port:          port,
		Identifier:    iden,
		SuccessorList: sList,
	}

	return node
}

func (n *Node) GetSuccessor(req Request, response *Entry) error {
	*response = n.Successor
	return nil
}

func (n *Node) GetSuccessorList(req Request, response *[]*Entry) error {
	*response = n.SuccessorList
	return nil
}

func (n *Node) stabilize() {
	successor := n.Successor
	fmt.Println(successor)
	fmt.Println("Stabilizing")
}

func (n *Node) GetPredecessor(req Request, response *Entry) error {
	fmt.Println("Getting Predecessor")
	*response = n.Predecessor
	return nil
}

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
	IpAddr string

	// Port of the entry
	Port string

	// Identifer of the entry constructing by using a base hash function SHA-1 of IP Addr
	Identifier Identifier
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

func NewNode() {

}

func InitFingerTable() (FingerTable, error) {
	newFT := make(map[string]Finger)
	return newFT, nil
}
