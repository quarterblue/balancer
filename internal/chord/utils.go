package chord

import (
	"crypto/sha1"
	"log"
	"math/big"
	"net"
	"strings"
)

func AddrToIpPort(ipAddr, port string) string {
	return ipAddr + ":" + port
}

func IpPortToAddr(ipPort string) (ipAddr, port string) {
	s := strings.SplitN(ipPort, ":", 2)
	return s[0], s[1]
}

// Returns the hashed value of the string, converted into big integer
func hashString(elt string) *big.Int {
	hasher := sha1.New()
	hasher.Write([]byte(elt))
	return new(big.Int).SetBytes(hasher.Sum(nil))
}

// Returns the IP address of the current machine
func GetLocalAddress() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func between(start, elt, end *big.Int, inclusive bool) bool {
	if end.Cmp(start) > 0 {
		return (start.Cmp(elt) < 0 && elt.Cmp(end) < 0) || (inclusive && elt.Cmp(end) == 0)
	} else {
		return start.Cmp(elt) < 0 || elt.Cmp(end) < 0 || (inclusive && elt.Cmp(end) == 0)
	}
}
