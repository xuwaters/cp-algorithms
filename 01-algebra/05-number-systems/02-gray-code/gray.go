package gray

func gray(n int) int {
	return n ^ (n >> 1)
}

func inverseGray(g int) int {
	n := 0
	for g > 0 {
		n ^= g
		g >>= 1
	}
	return n
}
