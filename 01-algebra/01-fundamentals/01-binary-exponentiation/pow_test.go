package exp

import (
	"testing"
)

func TestPow(t *testing.T) {
	dataList := []struct {
		ans PowableInt
		b   PowableInt
		n   int
	}{
		{ans: 1, b: 1, n: 10},
		{ans: 1024, b: 2, n: 10},
		{ans: 59049, b: 3, n: 10},
		{ans: 1162261467, b: 3, n: 19},
		{ans: 1073741824, b: 2, n: 30},
	}
	for _, data := range dataList {
		got := Pow(data.b, data.n).(PowableInt)
		if got != data.ans {
			t.Fatalf("error: pow, got = %v, data = %+v\n", got, data)
		} else {
			t.Logf("   ok: pow, data = %+v\n", data)
		}
	}
}
