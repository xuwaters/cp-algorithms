package diophantine

import (
	"testing"
)

func TestDiophantineEquation(t *testing.T) {
	dataList := []struct {
		a     int
		b     int
		c     int
		found bool
	}{
		{c: 3, a: 832040, b: 514229, found: true},
		{c: 62, a: 62 * 71, b: 31 * 31, found: true},
		{c: 52, a: 62 * 71, b: 31 * 31, found: false},
	}

	for _, data := range dataList {
		found, g, x, y := FindAnyDiophantineSolution(data.a, data.b, data.c)
		if found != data.found {
			t.Fatalf("ERR: data = %+v", data)
		} else if found {
			if data.a*x+data.b*y != data.c {
				t.Fatalf("ERR: g = %v, x = %v, y = %v, data = %+v", g, x, y, data)
			} else {
				t.Logf(" OK: found, g = %v, x = %v, y = %v, data = %+v", g, x, y, data)
			}
		} else {
			t.Logf(" OK: not_found, data = %+v", data)
		}
	}
}
