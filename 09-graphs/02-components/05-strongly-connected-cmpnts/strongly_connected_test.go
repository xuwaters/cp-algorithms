package strongly_connected

import (
	"testing"
)

func TestBridges(t *testing.T) {
	dataList := []struct {
		n     int
		scc   []int // strongly connected components
		edges [][2]int
	}{
		{
			n:     5,
			scc:   []int{0, 1, 2},
			edges: [][2]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}, {4, 2}},
		},
		{
			n:     5,
			scc:   []int{0, 2},
			edges: [][2]int{{0, 1}, {0, 2}, {1, 0}, {2, 3}, {3, 4}, {4, 2}},
		},
		{
			n:     5,
			scc:   []int{0},
			edges: [][2]int{{0, 1}, {0, 2}, {1, 0}, {2, 0}, {2, 3}, {3, 4}, {4, 2}},
		},
		{
			n:     5,
			scc:   []int{0},
			edges: [][2]int{{0, 1}, {0, 2}, {1, 0}, {2, 3}, {3, 1}, {3, 4}, {4, 2}},
		},
	}

	for _, data := range dataList {
		t.Logf(">> data = %+v", data)
		g := NewGraph(data.n)
		g.AddEdges(data.edges)
		got := g.FindStronglyConnectedComponents()
		if equals(got, data.scc) {
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
		if v != b[i] {
			return false
		}
	}
	return true
}
