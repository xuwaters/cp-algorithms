package binomial

import (
	"fmt"
	"testing"
)

func TestBinomialCoefficient(t *testing.T) {
	for n := 0; n <= 10; n++ {
		for k := 0; k <= n; k++ {
			fmt.Printf(" %4d", BinomialCoefficients(n, k))
		}
		fmt.Println()
	}
}
