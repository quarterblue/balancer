package balancer

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/rpc"
	"strings"

	"github.com/quarterblue/balancer/proto"
	"google.golang.org/grpc"
)

func AddrToIpPort(ipAddr, port string) string {
	return ipAddr + ":" + port
}

func IpPortToAddr(ipPort string) (ipAddr, port string) {
	s := strings.SplitN(ipPort, ":", 2)
	return s[0], s[1]
}

// Returns the hashed value of the string, converted into big integer
func hashString(elt string) *big.Int {
	hasher := sha1.New()
	hasher.Write([]byte(elt))
	return new(big.Int).SetBytes(hasher.Sum(nil))
}

// Returns the IP address of the current machine
func GetLocalAddress() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func between(start, elt, end *big.Int, inclusive bool) bool {
	if end.Cmp(start) > 0 {
		return (start.Cmp(elt) < 0 && elt.Cmp(end) < 0) || (inclusive && elt.Cmp(end) == 0)
	} else {
		return start.Cmp(elt) < 0 || elt.Cmp(end) < 0 || (inclusive && elt.Cmp(end) == 0)
	}
}

// func call(address string, method string, request Request, response interface{}) error {
// 	client, err := rpc.DialHTTP("tcp", address)
// 	if err != nil {
// 		log.Printf("Error connecting: %v", err)
// 		return err
// 	}
// 	defer client.Close()

// 	if err = client.Call(method, request, response); err != nil {
// 		log.Printf("Client call: %s, %v", method, err)
// 		return err
// 	}

// 	return nil
// }

func calltwo(address string, method string, request SRequest, response interface{}) error {
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Printf("Error connecting: %v", err)
		return err
	}
	defer client.Close()

	if err = client.Call(method, request, response); err != nil {
		log.Printf("Client call: %s, %v", method, err)
		return err
	}

	return nil
}

func gRpcCall(addr, method string, request *proto.KVRequest) (*proto.KVResponse, error) {
	var response *proto.KVResponse
	var err error

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewAddServiceClient(conn)

	ctx := context.Background()

	switch method {
	case "ping":
		fmt.Println("Ping")
		response, err = client.Ping(ctx, request)
	case "get":
		fmt.Println("Get")
		response, err = client.Get(ctx, request)
	case "put":
		fmt.Println("Put")
		response, err = client.Put(ctx, request)
	case "delete":
		fmt.Println("Delete")
		response, err = client.Delete(ctx, request)
	default:
		return nil, errors.New("unrecognized method request")
	}

	if err != nil {
		return nil, errors.New("gRPC Error")
	}

	return response, nil
}
