package balancer

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sync"
)

// Successor list size
const sListSize = 5

// m-bit hash of the node IP Addr or the key (Default uses SHA-1)
// current implementation stores a big int casted from the m-bit hash bytes
type Identifier *big.Int

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
	// IP Address of the node
	IpAddr string

	// Port of the node
	Port string

	// The m-bit identifier of the node
	Identifier Identifier

	// The array of n first successors; this is to replicate the successors to deal with failures in nodes
	SuccessorList []*Entry

	// The finger route table described in the paper; maintains up to m entries (nodes)
	Finger FingerTable

	// The previous node on the identifier circle
	Predecessor Entry

	// Mutex for protecting
	Mutex sync.RWMutex
}

func NewChordNode(ipAddr, port string, successor *Entry) *Node {
	addr := ipAddr + ":" + port
	iden := hashString(addr)

	sList := []*Entry{successor}
	node := &Node{
		IpAddr:        ipAddr,
		Port:          port,
		Identifier:    iden,
		SuccessorList: sList,
	}

	return node
}

func (n *Node) GetSuccessor(req Request, response *Entry) error {
	*response = *n.SuccessorList[1]
	return nil
}

func (n *Node) GetSuccessorList(req Request, response *[]*Entry) error {
	*response = n.SuccessorList
	return nil
}

func (n *Node) stabilize() {
	successor := n.successor()

	request := SRequest{}
	response := &SResponse{}

	if err := calltwo(successor.IpAddrString(), "Node.GetPredecessor", request, &response); err != nil {
		log.Fatalf("Calling Node.GetPredecessor: %v", err)
	}

	// if (x E (n, successor))
	if between(n.Identifier, response.successor.Identifier, n.SuccessorList[1].Identifier, true) {
		fmt.Println("its true")
	}

	fmt.Println(response.successor)
	fmt.Println("Stabilizing")
}

func (n *Node) predecessor() Entry {
	return n.Predecessor
}

func (n *Node) GetPredecessor(req SRequest, response *SResponse) error {
	fmt.Println("Getting Predecessor")
	s := SResponse{}
	s.predecessor = n.predecessor()
	*response = s
	return nil
}

func (n *Node) successor() *Entry {
	return n.SuccessorList[1]
	// return n.Finger[1].successor
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
