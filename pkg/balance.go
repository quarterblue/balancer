package balancer

import (
	"sort"
	"strconv"
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
	bound       int
}

type Balance struct {
	// Configuration to set Hash algorithm, replication factor and bounds
	config Config

	// Mutex to protect against concurrent access to hashMap and nodeList
	mutex sync.RWMutex

	// Node list
	hashMap map[uint64]interface{}

	// Sorted list of keys
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
	b.mutex.Lock()
	defer b.mutex.Unlock()
	hash := b.config.Hasher.Hash([]byte(key))
	// val, ok := b.hashMap[hash]
	return hash
}

func (b *Balance) AddNode(node string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	for i := 0; i < b.config.replication; i++ {
		hash := b.config.Hasher.Hash([]byte(node + strconv.Itoa(i)))
		b.nodeList = append(b.nodeList, hash)
		b.hashMap[hash] = node
	}

	// Keep the slice sorted
	sort.Slice(b.nodeList, func(i, j int) bool { return b.nodeList[i] < b.nodeList[j] })
}

func (b *Balance) RemoveNode(n string) {

}
