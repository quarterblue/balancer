package balancer

import (
	"github.com/cespare/xxhash"
)

type XxHash struct{}

func (x *XxHash) Hash(str []byte) uint64 {
	hasher := xxhash.New()
	hasher.Write(str)
	return hasher.Sum64()
}
