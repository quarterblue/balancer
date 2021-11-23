package chord

import (
	"fmt"
	"log"
	"math/big"

	pb "github.com/quarterblue/balancer/proto"
)

type Node struct {
	// IP Address of the entry
	IpAddr string

	// Port of the entry
	Port string

	// Identifer of the entry constructing by using a base hash function SHA-1 of IP Addr
	Identifier *big.Int
}

func NewNode(ipAddr, port string) *Node {
	addr := AddrToIpPort(ipAddr, port)
	hash := hashString(addr)
	node := &Node{
		IpAddr:     ipAddr,
		Port:       port,
		Identifier: hash,
	}
	return node
}

func (n *Node) IpAddrString() string {
	return n.IpAddr + ":" + n.Port
}

func (n *Node) Notify(c *Chord) {
	targetAddr := AddrToIpPort(n.IpAddr, n.Port)

	request := &pb.NodeRequest{}
	_, err := gRpcNode(targetAddr, "notify", request)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (n *Node) FindSuccessor(ipAddr, port string) *Node {
	targetAddr := AddrToIpPort(n.IpAddr, n.Port)
	fmt.Printf("target addr: %s\n", targetAddr)

	request := &pb.NodeRequest{
		Ipaddr: ipAddr,
		Port:   port,
	}

	response, err := gRpcNode(targetAddr, "findsuccessor", request)

	if err != nil {
		log.Println("Something wrong with findSucc")
		log.Println(err)
		return nil
	}
	return &Node{
		IpAddr:     response.GetIpaddr(),
		Port:       response.GetPort(),
		Identifier: hashString(AddrToIpPort(response.GetIpaddr(), response.GetPort())),
	}
}
