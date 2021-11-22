package chord

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/quarterblue/balancer/proto"
	"google.golang.org/grpc"
)

func gRpcCall(addr, method string, request *pb.KVRequest) (*pb.KVResponse, error) {
	var response *pb.KVResponse
	var err error

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewAddServiceClient(conn)

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

func gRpcNode(addr, method string, request *pb.NodeRequest) (*pb.Node, error) {
	var response *pb.Node
	var err error

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewNodeServiceClient(conn)

	ctx := context.Background()

	switch method {
	case "predecessor":
		fmt.Println("Predecessor")
		response, err = client.Predecessor(ctx, request)
	case "findsuccessor":
		fmt.Println("FindSuccesor")
		response, err = client.FindSuccessor(ctx, request)
	case "notify":
		fmt.Println("Notify")
		response, err = client.Notify(ctx, request)
	default:
		return nil, errors.New("unrecognized method request")
	}

	if err != nil {
		return nil, errors.New("gRPC Error")
	}

	return response, nil
}
