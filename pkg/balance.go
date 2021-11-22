package balancer

import (
	"sync"

	"github.com/cespare/xxhash"
)

type Entity struct {
	Name string
}

type Balance struct {
	sync.RWMutex
	hashMap map[uint64]interface{}
}

func (b *Balance) AddNode(n string) uint64 {
	return 2
}

func (b *Balance) RemoveNode(n string) {

}

func HashFunc(element string) uint64 {
	hasher := xxhash.New()
	hasher.Write([]byte(element))
	return hasher.Sum64()
}
