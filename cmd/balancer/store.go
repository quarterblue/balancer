package balancer

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/quarterblue/balancer/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Initialize() {
	fmt.Println("Init")
}

type Storer interface {
	Ping(ctx context.Context, request *proto.KVRequest) (*proto.KVResponse, error)
	Get(ctx context.Context, request *proto.KVRequest) (*proto.KVResponse, error)
	Put(ctx context.Context, request *proto.KVRequest) (*proto.KVResponse, error)
	Delete(ctx context.Context, request *proto.KVRequest) (*proto.KVResponse, error)
}

// Implementation of the Storer interface, currently supports only string to string KV
type Store struct {
	KeyValue map[string]string
	mutex    *sync.RWMutex
}

func (s *Store) Ping(ctx context.Context, request *proto.KVRequest) (*proto.KVResponse, error) {
	return &proto.KVResponse{Value: "", Msg: "Ping Back!", Ping: true}, nil
}

func (s *Store) Get(ctx context.Context, request *proto.KVRequest) (*proto.KVResponse, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	value, ok := s.KeyValue[request.GetKey()]
	reply := &proto.KVResponse{
		Value: "",
		Msg:   "",
		Ping:  false,
	}

	if ok {
		reply.Value = value
		reply.Ping = true
		return reply, nil
	} else {
		return reply, nil
	}

}
func (s *Store) Put(ctx context.Context, request *proto.KVRequest) (*proto.KVResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.KeyValue[request.GetKey()] = request.GetValue()
	return &proto.KVResponse{Value: request.GetValue(), Msg: "Success", Ping: true}, nil
}

func (s *Store) Delete(ctx context.Context, request *proto.KVRequest) (*proto.KVResponse, error) {
	return &proto.KVResponse{Value: "", Msg: "Not Implemented", Ping: false}, nil
}

// type RPCServer struct {
// 	Settings Settings
// }

// func (r *RPCServer) init(address string, store *Store, node *Node) {
// 	rpc.Register(store)
// 	rpc.Register(node)
// 	rpc.HandleHTTP()
// 	addr := r.Settings.Address + ":" + r.Settings.Port
// 	fmt.Println("Listening on: ", addr)
// 	l, err := net.Listen("tcp", addr)
// 	if err != nil {
// 		log.Fatal("Listen Error:", err)
// 	}
// 	if err := http.Serve(l, nil); err != nil {
// 		log.Fatalf("http.Serve: %v", err)
// 	}
// }

type RPCServer struct {
	Settings Settings
}

func (r *RPCServer) init(address string, store *Store, node *Chord) {
	addr := r.Settings.Address + ":" + r.Settings.Port
	listener, err := net.Listen("tcp", addr)
	fmt.Println("Listening on: ", addr)
	if err != nil {
		log.Fatal("Listen Error:", err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, store)
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
