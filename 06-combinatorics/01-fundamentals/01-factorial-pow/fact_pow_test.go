package fact_pow

import "testing"

func TestFactors(t *testing.T) {
	dataList := []struct {
		k   int
		ans [][2]int
	}{
		{k: 1, ans: [][2]int{}},
		{k: 2, ans: [][2]int{{2, 1}}},
		{k: 27648, ans: [][2]int{{2, 10}, {3, 3}}},
		{k: 180238, ans: [][2]int{{2, 1}, {227, 1}, {397, 1}}},
		{k: 64971473288, ans: [][2]int{{2, 3}, {227, 2}, {397, 2}}},
	}
	for _, data := range dataList {
		got := Factors(data.k)
		if !arrayEquals(got, data.ans) {
			t.Fatalf("ERR: got = %+v, data = %+v", got, data)
		} else {
			t.Logf(" OK: got = %+v, data = %+v", got, data)
		}
	}
}

func arrayEquals(a, b [][2]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func TestFactPow(t *testing.T) {
	dataList := []struct {
		n   int
		k   int
		ans int
	}{
		{n: 2, k: 5, ans: 0},
		{n: 10000, k: 769, ans: 13},
		{n: 10000, k: 2, ans: 9995},
		{n: 10000, k: 27648, ans: 999},
		{n: 10000, k: 180238, ans: 25},
		{n: 10000, k: 64971473288, ans: 12},
	}
	for _, data := range dataList {
		got := FactPow(data.n, data.k)
		if got != data.ans {
			t.Fatalf("ERR: got = %+v, data = %+v", got, data)
		} else {
			t.Logf(" OK: got = %+v, data = %+v", got, data)
		}
	}
}
