package balancer

import (
	"context"
	"errors"
	"fmt"

	"github.com/quarterblue/balancer/proto"
	"google.golang.org/grpc"
)

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
