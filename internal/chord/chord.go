package chord

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"sync"

	pb "github.com/quarterblue/balancer/proto"
)

const (
	replicationFactor = 3   // Replication factor for redundant storage
	sListSize         = 5   // Successor list size
	m                 = 160 // bit size (for SHA-1)
)

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
	FingerTable map[int]*Finger

	// The previous node on the identifier circle
	Predecessor *Node

	// Mutex for protecting finger table
	FingerMutex sync.RWMutex

	// Mutex for protecting successor list
	SuccMutex sync.RWMutex

	// Next finger index to fix
	next int
}

func NewChordNode(ipAddr, port string, successor *Node) *Chord {
	addr := ipAddr + ":" + port
	iden := hashString(addr)

	chord := &Chord{
		IpAddr:        ipAddr,
		Port:          port,
		Identifier:    iden,
		SuccessorList: []*Node{successor},
		FingerTable:   map[int]*Finger{},
		Predecessor:   nil,
		FingerMutex:   sync.RWMutex{},
		SuccMutex:     sync.RWMutex{},
		next:          0,
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
	var sList = []*pb.Node{}

	for _, n := range c.SuccessorList {
		sList = append(sList, &pb.Node{
			Ipaddr: n.IpAddr,
			Port:   n.Port,
		})
	}

	return &pb.SuccessorResponse{
		Successorlist: sList,
	}, nil
}

func (c *Chord) Stabilize() {
	fmt.Println("Stabilizing")
	successor := c.successor()

	// Same identifier
	if successor.Identifier.Cmp(c.Identifier) == 0 {
		return
	}

	request := &pb.NodeRequest{}
	successorAddr := AddrToIpPort(successor.IpAddr, successor.Port)

	response, err := gRpcNode(successorAddr, "predecessor", request)
	if err != nil {
		log.Println(err)
		return
	}

	pSucc := AddrToIpPort(response.GetIpaddr(), response.GetPort())
	pSuccHash := hashString(pSucc)

	// if (x E (n, successor))
	if checkBetween(c.Identifier, pSuccHash, c.successor().Identifier, true) {
		// successor = x;
		newSuccessor := &Node{
			IpAddr:     response.GetIpaddr(),
			Port:       response.GetPort(),
			Identifier: pSuccHash,
		}

		c.SuccessorList = append([]*Node{newSuccessor}, c.SuccessorList...)
	}

	c.SuccessorList[0].notify(c)
}

func (c *Chord) predecessor() *Node {
	return c.Predecessor
}

func (c *Chord) successor() *Node {
	return c.SuccessorList[0]
}

func (c *Chord) FindSuccessor(ctx context.Context, request *pb.NodeRequest) (*pb.Node, error) {

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

func (c *Chord) FixFingers() {
	// Iterate up to m-bit, account for 0 index subtracting 1
	if c.next > (m - 1) {
		c.next = 0
	}

	ctx := context.Background()

	response, err := c.FindSuccessor(ctx, &pb.NodeRequest{Node: ""})
	if err != nil {
		log.Fatal(err)
	}

	n := &Node{
		IpAddr: response.GetIpaddr(),
		Port:   response.GetPort(),
	}

	f := &Finger{
		// start:     hashString(AddrToIpPort(c.IpAddr, c.Port)) + ,
		// end:       hashString(AddrToIpPort(c.IpAddr, c.Port)),
		successor: n,
	}

	c.FingerMutex.Lock()
	c.FingerTable[c.next] = f
	c.FingerMutex.Unlock()
}
