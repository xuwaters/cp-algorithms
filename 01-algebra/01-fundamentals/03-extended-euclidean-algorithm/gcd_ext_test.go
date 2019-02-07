package gcd_ext

import (
	"testing"
)

func TestGreatestCommonDivisorExt(t *testing.T) {
	dataList := []struct {
		g int // gcd (a, b)
		a int
		b int
	}{
		{g: 8, a: 24, b: 32},
		{g: 1, a: 31, b: 32},
		{g: 31, a: 31, b: 62},
		{g: 31, a: 62 * 71, b: 31 * 31},
		{g: 831038, a: 0, b: 831038},
		{g: 831038, a: 831038, b: 0},
		{g: 1, a: 832040, b: 514229},
	}

	for _, data := range dataList {
		g, x, y := GreatestCommonDivisorExt(data.a, data.b)
		if g != data.g {
			t.Fatalf("ERR: got = %+v, data = %+v", g, data)
		} else {
			t.Logf(" OK: got = %+v, data = %+v", g, data)
		}
		if x*data.a+y*data.b != g {
			t.Fatalf("ERR: x = %v, y = %v, data = %v", x, y, data)
		} else {
			t.Logf(" OK: x = %v, y = %v, data = %v", x, y, data)
		}
	}
}
