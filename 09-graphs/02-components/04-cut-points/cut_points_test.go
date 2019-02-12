package cut_points

import (
	"testing"
)

func TestCutPoints(t *testing.T) {
	dataList := []struct {
		n         int
		cutPoints []int
		edges     [][2]int
	}{
		{
			n:         5,
			cutPoints: []int{0, 2},
			edges:     [][2]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}, {2, 4}},
		},
	}

	for _, data := range dataList {
		g := NewGraph(data.n)
		g.AddEdges(data.edges)
		got := g.FindCutPoints()
		if equals(got, data.cutPoints) {
			t.Logf(" OK: got = %+v, data = %+v", got, data)
		} else {
			t.Fatalf("ERR: got = %+v, data = %+v", got, data)
		}
	}
}

func equals(a, b []int) bool {
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
