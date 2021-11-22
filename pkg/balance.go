package balancer

import (
	"sync"
)

type Entity struct {
	Name string
}

// Hashing algorithm interface, implement this interface to use as a hasher
type Hasher interface {
	Hash([]byte) uint64
}

type Config struct {
	Hasher      Hasher
	replication int
	vNum        int
}

type Balance struct {
	config   Config
	mutex    sync.RWMutex
	hashMap  map[uint64]interface{}
	nodeList []uint64
}

// Creates and returns a new balancer with the given config
func NewBalance(config Config) *Balance {
	// Default hashing algorithm is xxhash implementation by cespare
	if config.Hasher == nil {
		hasher := &XxHash{}
		config.Hasher = hasher
	}

	return &Balance{
		config:   config,
		mutex:    sync.RWMutex{},
		hashMap:  make(map[uint64]interface{}),
		nodeList: []uint64{},
	}
}

func (b *Balance) FindNode(key string) interface{} {
	hash := b.config.Hasher.Hash([]byte(key))
	b.mutex.Lock()
	defer b.mutex.Unlock()
	// val, ok := b.hashMap[hash]
	return hash
}

func (b *Balance) AddNode(node string) uint64 {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	for i := 0; i < b.config.replication; i++ {
		hash := b.config.Hasher.Hash([]byte(node))
		b.nodeList = append(b.nodeList, hash)
	}
	return 2
}

func (b *Balance) RemoveNode(n string) {

}
