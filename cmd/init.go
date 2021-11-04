package cmd

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"
)

func Initialize() {
	fmt.Println("Init")
}

type Storer interface {
	Ping(msg string, reply *string) error
	Get(key string, reply *string) error
	Put(key string, value string, reply *string) error
	Delete(key string, reply *string) error
}

// Implementation of the Storer interface, currently supports only string to string KV
type Store struct {
	KeyValue map[string]string
	mutex    *sync.RWMutex
}

func (s *Store) Ping(r Request, reply *string) error {
	*reply = "You pinged me!"
	return nil
}

func (s *Store) Get(r Request, reply *string) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	value, ok := s.KeyValue[r.Key]
	if ok {
		*reply = value
	}
	return nil
}

func (s *Store) Put(r Request, reply *string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.KeyValue[r.Key] = r.Value
	*reply = "success"
	return nil
}

func (s *Store) Delete(key string, reply *string) error {
	return nil
}

type RPCServer struct {
	Settings Settings
}

func (r *RPCServer) init(address string, store *Store) {
	rpc.Register(store)
	rpc.HandleHTTP()
	addr := r.Settings.Address + ":" + r.Settings.Port
	fmt.Println("Listening on: ", addr)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Listen Error:", err)
	}
	if err := http.Serve(l, nil); err != nil {
		log.Fatalf("http.Serve: %v", err)
	}
}
