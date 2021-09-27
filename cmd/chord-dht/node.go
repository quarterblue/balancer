package main

// m-bit hash of the node IP Addr or the key (Default uses SHA-1)
type Identifier []byte

type FingerTable map[string]Finger

type Node struct {
	finger      FingerTable
	successor   Entry
	predecessor Entry
	identifier  Identifier
}

type Finger struct {
	// (n + 2^(i-1)) mod 2^m (1 <= i <= m)
	start int
	// (n + 2^i - 1) mod 2^m
	end int
	// interval
	interval [2]int
	// succesor node
	successor int
}

type Entry struct {
	// IP Address of the entry
	ipAddr string

	// Port of the entry
	port string

	// Identifer of the entry constructing by using a base hash function SHA-1 of IP Addr
	identifier Identifier
}

func (n *Node) findSuccessor(id string) (*Entry, error) {
	return nil, nil
}

func NewEntry(ipAddr, port string, identifier []byte) *Entry {
	newEntry := Entry{
		ipAddr,
		port,
		identifier,
	}
	return &newEntry
}

func NewNode() {

}

func InitFingerTable(nPrime Entry) error {

	return nil
}
