package main

import "context"

// m-bit hash of the node IP Addr or the key (Default uses SHA-1)
type Identifier []byte

type FingerTable map[string]Finger

type Node struct {
	// The route table
	finger FingerTable

	// The next node on the identifier cirle; finger[1].node
	successor Entry

	// The previous node on the identifier circle
	predecessor Entry

	// The m-bit identifier
	identifier Identifier
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
	ipAddr string

	// Port of the entry
	port string

	// Identifer of the entry constructing by using a base hash function SHA-1 of IP Addr
	identifier Identifier
}

func (n *Node) findSuccessor(ctx context.Context, id Identifier) (*Entry, error) {
	if withinFingerRange(n.identifier, id) {
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

func NewEntry(ipAddr, port string, identifier []byte) *Entry {
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
