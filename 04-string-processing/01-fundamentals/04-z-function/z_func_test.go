package z_func

import "testing"

func TestZFunction(t *testing.T) {

	dataList := []struct {
		z   []int
		str string
	}{
		{str: "aaaaa", z: []int{0, 4, 3, 2, 1}},
		{str: "aaabaab", z: []int{0, 2, 1, 0, 2, 1, 0}},
		{str: "abacaba", z: []int{0, 0, 1, 0, 3, 0, 1}},
	}

	for _, data := range dataList {
		got := ZFunction(data.str)
		if !ArrayEquals(got, data.z) {
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
		if b[i] != v {
			return false
		}
	}
	return true
}
