package gcd

import "testing"

func TestGreatestCommonDivisor(t *testing.T) {
	dataList := []struct {
		ans int
		a   int
		b   int
	}{
		{ans: 8, a: 24, b: 32},
		{ans: 1, a: 31, b: 32},
		{ans: 31, a: 31, b: 62},
		{ans: 31, a: 62 * 71, b: 31 * 31},
		{ans: 831038, a: 0, b: 831038},
		{ans: 831038, a: 831038, b: 0},
		{ans: 1, a: 832040, b: 514229},
	}

	for _, data := range dataList {
		got := GreatestCommonDivisor(data.a, data.b)
		if got != data.ans {
			t.Fatalf("ERR: got = %+v, data = %+v", got, data)
		} else {
			t.Logf(" OK: got = %+v, data = %+v", got, data)
		}
	}

	for _, data := range dataList {
		got := GreatestCommonDivisorLoop(data.a, data.b)
		if got != data.ans {
			t.Fatalf("ERR: got = %+v, data = %+v", got, data)
		} else {
			t.Logf(" OK: got = %+v, data = %+v", got, data)
		}
	}
}

func TestLeastCommonMultiple(t *testing.T) {
	dataList := []struct {
		ans int
		a   int
		b   int
	}{
		{ans: 35, a: 5, b: 7},
		{ans: 31 * 2 * 71 * 31, a: 62 * 71, b: 31 * 31},
		{ans: 514229 * 832040, a: 832040, b: 514229},
	}

	for _, data := range dataList {
		got := LeastCommonMultiple(data.a, data.b)
		if got != data.ans {
			t.Fatalf("ERR: got = %+v, data = %+v", got, data)
		} else {
			t.Logf(" OK: got = %+v, data = %+v", got, data)
		}
	}
}
