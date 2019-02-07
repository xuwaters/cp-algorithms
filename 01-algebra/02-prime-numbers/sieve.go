package prime

import (
	"math"
)

// CountPrimesBySieve count prime numbers <= n
func CountPrimesBySieve(n int) int {
	// prepare small primes
	nsqrt := int(math.Sqrt(float64(n)))
	primes := GeneratePrimes(nsqrt)
	// Sieve by blocks
	cnt := 0
	BS := nsqrt // block size
	isPrime := make([]bool, BS)

	resetIsPrime := func() {
		for i := 0; i < len(isPrime); i++ {
			isPrime[i] = true
		}
	}

	for k := 0; k*BS <= n; k++ {
		resetIsPrime()
		start := k * BS
		for _, p := range primes {
			startIndex := (start + p - 1) / p
			// scan from max(p*p, startIndex*p)
			if startIndex < p {
				startIndex = p
			}
			for j := startIndex*p - start; j < BS; j += p {
				isPrime[j] = false
			}
		}
		if k == 0 {
			isPrime[0] = false
			isPrime[1] = false
		}
		for i := 0; i < BS && i+start <= n; i++ {
			if isPrime[i] {
				cnt++
			}
		}
	}

	return cnt
}

func GeneratePrimes(n int) []int {
	primes := make([]int, 0)
	isNotPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if !isNotPrime[i] {
			primes = append(primes, i)
			for j := i * i; j <= n; j += i {
				isNotPrime[j] = true
			}
		}
	}
	return primes
}
