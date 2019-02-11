package suffix_array

import (
	"testing"
)

func TestSuffixArray(t *testing.T) {
	data := []rune("mississippimississippi")
	n := len(data)

	s := NewSuffixArray(data)
	s.Build()

	arr := s.Arr()
	t.Logf("  rank array = %+v\n", s.rank[len(s.rank)-1])
	t.Logf("suffix array = %+v\n", arr)
	for i := 0; i < n; i++ {
		suffix := string(data[arr[i]:])

		lcp := 0
		if i > 0 {
			lcp = s.LongestCommonPrefix(arr[i], arr[i-1])
		}

		t.Logf("  %2d: (lcp = %2d) %s\n", i, lcp, suffix)
	}
}
