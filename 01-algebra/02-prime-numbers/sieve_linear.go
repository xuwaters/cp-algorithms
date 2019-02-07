package prime

// SievePrimesLinearly runs in O(n) complexity
func SievePrimesLinearly(n int) []int {
	primes := make([]int, 0)
	lp := make([]int, n+1) // lp[i] is the least prime factor of i
	for i := 2; i <= n; i++ {
		if lp[i] == 0 {
			lp[i] = i
			primes = append(primes, i)
		}
		for j := 0; j < len(primes) && primes[j] <= lp[i] && i*primes[j] <= n; j++ {
			lp[i*primes[j]] = primes[j]
		}
	}
	return primes
}
