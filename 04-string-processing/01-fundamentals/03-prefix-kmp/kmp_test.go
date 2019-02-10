package kmp

import "testing"

func TestPrefixFunction(t *testing.T) {

	dataList := []struct {
		prefix  []int
		pattern string
	}{
		{
			prefix:  []int{0, 1, 0, 1, 2, 3, 0, 0, 0},
			pattern: "ssissippi",
		},
		{
			prefix:  []int{0, 0, 0, 1, 2, 3, 0},
			pattern: "abcabcd",
		},
	}

	for _, data := range dataList {
		got := PrefixFunction(data.pattern)
		if !ArrayEquals(got, data.prefix) {
			t.Fatalf("ERR: got = %+v, data = %+v", got, data)
		} else {
			t.Logf(" OK: got = %+v, data = %+v", got, data)
		}
	}
}

func ArrayEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
