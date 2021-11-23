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

// Interface of the Chord Node RPC, requires implementation of all node RPC functionalities
type NodeRPC interface {
	FindSuccessor(ctx context.Context, request *pb.NodeRequest) (*pb.Node, error)
	GetSuccessor(ctx context.Context, request *pb.NodeRequest) (*pb.Node, error)
	GetPredecessor(ctx context.Context, request *pb.NodeRequest) (*pb.Node, error)
	GetSuccessorList(ctx context.Context, request *pb.NodeRequest) (*pb.SuccessorResponse, error)
}

// Interface of the Chord Node, requires implementation of all node functionalities and Chord Node RPC functionalities
type ChordNode interface {
	NodeRPC
	Stabilize()
	ClosestPrecedingNode()
	Join(node *Node)
	FixFingers()
	Notify()
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

	// Mutex for protecting Predecessor
	PredMutex sync.Mutex

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
		PredMutex:     sync.Mutex{},
		FingerMutex:   sync.RWMutex{},
		SuccMutex:     sync.RWMutex{},
		next:          0,
	}

	return chord
}

func (c *Chord) Join(node *Node) {
	c.Predecessor = nil
	successor := node.FindSuccessor(c.IpAddr, c.Port)
	fmt.Println(successor)
	c.SuccMutex.Lock()
	c.SuccessorList[0] = successor
	c.SuccMutex.Unlock()
}

func (c *Chord) Stabilize() {
	fmt.Println("Stabilizing")
	successor := c.successor()

	// Same identifier, we are the only node in the ring
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

	c.SuccessorList[0].Notify(c)
}

func (c *Chord) FixFingers() {
	// Iterate up to m-bit, account for 0 index subtracting 1
	if c.next > (m - 1) {
		c.next = 0
	}

	ctx := context.Background()

	response, err := c.FindSuccessor(ctx, &pb.NodeRequest{})
	if err != nil {
		log.Fatal(err)
	}

	n := &Node{
		IpAddr: response.GetIpaddr(),
		Port:   response.GetPort(),
	}

	f := &Finger{
		start: fingerStart(c.IpAddr, c.Port, c.next),
		end:   fingerEnd(c.IpAddr, c.Port, c.next),
		node:  n,
	}

	c.FingerMutex.Lock()
	c.FingerTable[c.next] = f
	c.FingerMutex.Unlock()
}

func (c *Chord) InitFingerTable(join *Node) {
	c.FingerMutex.Lock()
	defer c.FingerMutex.Unlock()

	// c.FingerTable[1].node = join.FindSuccessor()
}

func (c *Chord) ClosestPrecedingNode(id *big.Int) *Node {

	// For i = m down to 1
	for i := m; m > 1; i-- {
		// Finger[i] E (n, id)
	}

	return nil
}

func (c *Chord) predecessor() *Node {
	return c.Predecessor
}

func (c *Chord) successor() *Node {
	return c.SuccessorList[0]
}

// These are the gRPC methods exposed to other peer nodes in the ring

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

func (c *Chord) Notify(ctx context.Context, request *pb.NodeRequest) (*pb.Node, error) {
	c.PredMutex.Lock()
	defer c.PredMutex.Unlock()
	addr := AddrToIpPort(request.Ipaddr, request.Port)
	if c.Predecessor == nil || checkBetween(c.Predecessor.Identifier, hashString(addr), c.Identifier, true) {
		c.Predecessor = NewNode(request.Ipaddr, request.Port)
	}

	return &pb.Node{}, nil
}

func (c *Chord) FindSuccessor(ctx context.Context, request *pb.NodeRequest) (*pb.Node, error) {
	pSucc := AddrToIpPort(request.Ipaddr, request.Port)
	pSuccHash := hashString(pSucc)

	return &pb.Node{
		Ipaddr: c.IpAddr,
		Port:   c.Port,
	}, nil

	if checkBetween(c.Identifier, pSuccHash, c.successor().Identifier, true) {
		// ID E (n, succcessor)
		c.SuccMutex.RLock()
		successor := &pb.Node{
			Ipaddr: c.SuccessorList[0].IpAddr,
			Port:   c.SuccessorList[0].IpAddr,
		}
		c.SuccMutex.RUnlock()
		return successor, nil
	}

	nPrime := c.ClosestPrecedingNode(pSuccHash)

	node := nPrime.FindSuccessor(request.Ipaddr, request.Port)

	successor := &pb.Node{
		Ipaddr: node.IpAddr,
		Port:   node.Port,
	}

	return successor, nil
}
