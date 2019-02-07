package exp

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	dataList := []struct {
		ans int
		n   int
	}{
		{n: 0, ans: 0},
		{n: 1, ans: 1},
		{n: 2, ans: 1},
		{n: 3, ans: 2},
		{n: 4, ans: 3},
		{n: 5, ans: 5},
		{n: 6, ans: 8},
		{n: 7, ans: 13},
		{n: 8, ans: 21},
		{n: 9, ans: 34},
		{n: 10, ans: 55},
		{n: 11, ans: 89},
		{n: 12, ans: 144},
		{n: 29, ans: 514229},
		{n: 30, ans: 832040},
	}

	for _, data := range dataList {
		got := fibonacci(data.n)
		if got != data.ans {
			t.Fatalf("ERR: got = %+v, data = %+v\n", got, data)
		} else {
			t.Logf(" OK: got = %+v, data = %+v\n", got%3, data)
		}
	}
}
