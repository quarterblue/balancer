package internal

import (
	"math/big"
)

type Finger struct {
	// (n + 2^(i-1)) mod 2^m (1 <= i <= m)
	start *big.Int

	// (n + 2^i - 1) mod 2^m
	end *big.Int

	// interval
	// interval [2]int

	// succesor node
	successor *Node
}

// func fingerStart(ipAddr, port string, i int) *big.Int {
// 	start := hashString(AddrToIpPort(ipAddr, port))
// 	start.Add(start, big.NewInt(int64(math.Pow(2, float64(i)))))
// 	// startMod := *start % *big.NewInt(int64(math.Pow(2, m)))
// 	return startMod

// }

// func fingerEnd(IpAddr, port string, i int) *big.Int {

// }
