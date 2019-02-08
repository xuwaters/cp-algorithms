package fft

import (
	"testing"
)

func TestFastFourierTransform(t *testing.T) {
	dataList := []struct {
		ans []int
		a   []int
		b   []int
	}{
		{
			ans: []int{5, 16, 35, 34, 24},
			a:   []int{1, 2, 3},
			b:   []int{5, 6, 8},
		},
	}
	for _, data := range dataList {
		got := MultiplyArray(data.a, data.b)
		if !ArrayEqual(got, data.ans) {
			t.Fatalf("ERR: got = %+v, data = %+v", got, data)
		} else {
			t.Logf(" OK: got = %+v, data = %+v", got, data)
		}
	}
}

func ArrayEqual(a, b []int) bool {
	for i := 0; i < len(a) || i < len(b); i++ {
		if i < len(a) && i < len(b) {
			if a[i] != b[i] {
				return false
			}
		} else if i < len(a) {
			if a[i] != 0 {
				return false
			}
		} else {
			if b[i] != 0 {
				return false
			}
		}
	}
	return true
}
