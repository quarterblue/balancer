package chord

import (
	"math"
	"math/big"
)

type Finger struct {
	// (n + 2^(i-1)) mod 2^m (1 <= i <= m)
	start *big.Int

	// (n + 2^i - 1) mod 2^m
	end *big.Int

	// succesor node
	node *Node
}

func fingerStart(ipAddr, port string, i int) *big.Int {
	start := hashString(AddrToIpPort(ipAddr, port))
	start.Add(start, big.NewInt(int64(math.Pow(2, float64(i)))))
	startMod := new(big.Int)
	startMod = startMod.Mod(start, big.NewInt(int64(math.Pow(2, m))))

	return startMod
}

func fingerEnd(ipAddr, port string, i int) *big.Int {
	end := hashString(AddrToIpPort(ipAddr, port))
	end.Add(end, big.NewInt(int64(math.Pow(2, float64(i)+1))))
	endMod := new(big.Int)
	endMod = endMod.Mod(end, big.NewInt(int64(math.Pow(2, m))))

	return endMod
}
