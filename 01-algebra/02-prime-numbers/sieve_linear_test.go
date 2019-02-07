package prime

import "testing"

func TestSievePrimesLinearly(t *testing.T) {
	primes := SievePrimesLinearly(100)
	t.Logf("linearly sieve primes: %+v, %+v", len(primes), primes)

	n := 100000000
	largePrimes := SievePrimesLinearly(n)
	t.Logf("linearly sieve primes of large number: %+v, %v", len(largePrimes), n)
	if 5761455 != len(largePrimes) {
		t.Fatalf("linearly sieve primes failure")
	}
}
