package balancer

type FingerTable map[int]Finger

type Finger struct {
	// (n + 2^(i-1)) mod 2^m (1 <= i <= m)
	start int

	// (n + 2^i - 1) mod 2^m
	end int

	// interval
	interval [2]int

	// succesor node
	successor *Entry
}

func InitFingerTable() (FingerTable, error) {
	newFT := make(map[int]Finger)
	return newFT, nil
}
