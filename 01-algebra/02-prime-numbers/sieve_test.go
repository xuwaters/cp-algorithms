package prime

import "testing"

func TestGeneratePrimes(t *testing.T) {
	primes := GeneratePrimes(100)
	t.Logf("primes: %+v, %+v", len(primes), primes)
}

func TestCountPrimesBySieve(t *testing.T) {
	dataList := []struct {
		n   int
		cnt int
	}{
		{n: 10, cnt: 4},
		{n: 100, cnt: 25},
		{n: 10000, cnt: 1229},
		{n: 10000000, cnt: 664579},
		{n: 100000000, cnt: 5761455},
	}
	for _, data := range dataList {
		got := CountPrimesBySieve(data.n)
		if got != data.cnt {
			t.Fatalf("ERR: got = %d, data = %+v", got, data)
		} else {
			t.Logf(" OK: got = %d, data = %+v", got, data)
		}
	}
}
