package dijkstra

import "testing"

func TestDijkstra(t *testing.T) {
	dataList := []struct {
		n        int
		s        int
		shortest []int
		edges    []Edge
	}{
		{
			n:        8,
			s:        0,
			shortest: []int{0, 2, 4, 5, 1, -1, -1, -1},
			edges: []Edge{
				{0, 1, 2}, {1, 0, 2},
				{0, 2, 4}, {2, 0, 4},
				{0, 3, 6}, {3, 0, 6},
				{0, 4, 1}, {4, 0, 1},
				{1, 2, 3}, {2, 1, 3},
				{2, 3, 1}, {3, 2, 1},
				{3, 4, 7}, {4, 3, 7},
				{5, 6, 3}, {6, 5, 3},
				{5, 7, 4}, {7, 5, 4},
				{6, 7, 6}, {7, 6, 6},
			},
		},
	}

	for _, data := range dataList {
		g := NewWeightedGraph(data.n)
		g.AddEdges(data.edges)
		got := ToDistances(g.ShortestPaths(data.s))
		if equals(got, data.shortest) {
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
