package balancer

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/quarterblue/balancer/proto"
	pb "github.com/quarterblue/balancer/proto"
)

// Successor list size
const sListSize = 5

// m-bit hash of the node IP Addr or the key (Default uses SHA-1)
// current implementation stores a big int casted from the m-bit hash bytes
// type Identifier *big.Int

// Interface of the Chord Node RPC, requires implementation of all node functionalities described in the paper
type NodeRPC interface {
	FindSuccessor()
	ClosestPrecedingNode()
	GetPredecessor()
	GetSuccessorList()
	Create()
	Join()
	Stabilize()
	Notify()
	FixFinger()
	CheckPredecessor() bool
}

// Implementation of the Chord Node
type Chord struct {
	// IP Address of the chord node
	IpAddr string

	// Port of the chord node
	Port string

	// The m-bit identifier of the chord node stored as big.Int
	Identifier *big.Int

	// Slice of n first successors; this is to replicate the successors to deal with failures in nodes
	SuccessorList []*Node

	// The finger route table described in the paper; maintains up to m entries (nodes)
	Finger *FingerTable

	// The previous node on the identifier circle
	Predecessor *Node

	// Mutex for protecting finger table
	FingerMutex sync.RWMutex

	// Mutex for protecting successor list
	SuccMutex sync.RWMutex
}

func NewChordNode(ipAddr, port string, successor *Node) *Chord {
	addr := ipAddr + ":" + port
	iden := hashString(addr)

	chord := &Chord{
		IpAddr:        ipAddr,
		Port:          port,
		Identifier:    iden,
		SuccessorList: []*Node{successor},
		Finger:        new(FingerTable),
		Predecessor:   nil,
		FingerMutex:   sync.RWMutex{},
		SuccMutex:     sync.RWMutex{},
	}

	return chord
}

func (c *Chord) Join(node *Node) {
	c.Predecessor = nil
	succesor := node.FindSuccessor(c)
	fmt.Println(succesor)
}

func (c *Chord) GetPredecessor(ctx context.Context, request *pb.NodeRequest) (*pb.Node, error) {
	if c.Predecessor == nil {
		return nil, errors.New("predecessor is nil")
	}

	return &pb.Node{
		Ipaddr: c.Predecessor.IpAddr,
		Port:   c.Predecessor.Port,
	}, nil
}

func (c *Chord) GetSuccessor(ctx context.Context, request *pb.NodeRequest) (*pb.Node, error) {
	return &pb.Node{
		Ipaddr: c.SuccessorList[0].IpAddr,
		Port:   c.SuccessorList[0].Port,
	}, nil
}

func (c *Chord) GetSuccessorList(ctx context.Context, request *pb.NodeRequest) (*pb.SuccessorResponse, error) {
	return nil, nil
}

func (c *Chord) stabilize() {
	// successor := c.successor()
	// fmt.Println("Stabilizing")
	// // Same identifier
	// if successor.Identifier.Cmp(c.Identifier) == 0 {
	// 	return
	// }

	// if err := calltwo(successor.IpAddrString(), "Node.GetPredecessor", request, &response); err != nil {
	// 	log.Fatalf("Calling Node.GetPredecessor: %v", err)
	// }

	// // if (x E (n, successor))
	// if between(c.Identifier, response.Successor.Identifier, c.SuccessorList[0].Identifier, true) {
	// 	fmt.Println("its true")

	// 	newSuccessor := &Node{
	// 		IpAddr:     response.Predecessor.IpAddr,
	// 		Port:       response.Predecessor.Port,
	// 		Identifier: response.Predecessor.Identifier,
	// 	}

	// 	c.SuccessorList = append([]*Node{newSuccessor}, c.SuccessorList...)
	// }

	// c.SuccessorList[0].notify(c)

	// fmt.Println(response.Successor)
}

func (c *Chord) predecessor() *Node {
	return c.Predecessor
}

func (c *Chord) successor() *Node {
	return c.SuccessorList[0]
	// return n.Finger[1].successor
}

func (c *Chord) FindSuccessor(context.Context, *proto.NodeRequest) (*proto.Node, error) {

	// if between(c.Identifier, response.Successor.Identifier, c.SuccessorList[0].Identifier, true) {
	// 	fmt.Println("True")
	// }
	return nil, nil
}

// func (n *Node) findSuccessor(ctx context.Context, id *big.Int) (*Entry, error) {
// 	if withinFingerRange(n.Identifier, id) {
// 		// return n.finger[id], nil
// 		return nil, nil
// 	} else {
// 		// nPrime := closestPrecedingNode(id)
// 		// return nPrime.findSuccessor(ctx, id)

// 		return nil, nil
// 	}
// }

func closestPrecedingNode(id *big.Int) *Node {
	return nil
}

func withinFingerRange(n, successor *big.Int) bool {
	return true
}
