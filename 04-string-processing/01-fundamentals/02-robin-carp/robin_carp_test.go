package robin_carp

import "testing"

func TestRobinCarp(t *testing.T) {
	dataList := []struct {
		ans     []int
		pattern string
		text    string
	}{
		{
			ans:     []int{1, 2},
			pattern: "aa",
			text:    "baaac",
		},
		{
			ans:     []int{5, 7, 9},
			pattern: "ab",
			text:    "helloababab",
		},
		{
			ans:     []int{7, 17, 38, 49},
			pattern: "abelmak",
			text:    "helloababelmakabaabelmakbafliefjlafjelabelmakjlahabelmakalemfei",
		},
	}

	for _, data := range dataList {
		got := RobinCarp(data.pattern, data.text)
		if !ArrayEquals(got, data.ans) {
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
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
