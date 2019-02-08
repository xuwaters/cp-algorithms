package modular

import (
	"testing"
)

func TestInverse(t *testing.T) {
	primes := []int{
		7, 19, 31,
	}
	for _, p := range primes {
		inv := CalcInverse(p)
		t.Logf("inverse of %d: %+v", p, inv)
		for i := 1; i < len(inv); i++ {
			if inv[i]*i%p != 1 {
				t.Fatalf("CalcInverse of %d (mod %d) invalid: %+v", i, p, inv[i])
			}
			ii := ModInverse(i, p)
			if ii*i%p != 1 {
				t.Fatalf("ModInverse of %d (mod %d) invalid: %+v", i, p, ii)
			}
		}
	}
}
