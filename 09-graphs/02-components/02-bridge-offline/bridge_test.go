package bridge

import (
	"testing"
)

func TestBridges(t *testing.T) {
	dataList := []struct {
		n       int
		bridges [][2]int
		edges   [][2]int
	}{
		{
			n:       5,
			bridges: [][2]int{{0, 1}, {0, 2}},
			edges:   [][2]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}, {2, 4}},
		},
	}

	for _, data := range dataList {
		g := NewGraph(data.n)
		g.AddEdges(data.edges)
		bridges := g.FindBridges()
		if equals(bridges, data.bridges) {
			t.Logf(" OK: bridges = %+v, data = %+v", bridges, data)
		} else {
			t.Fatalf("ERR: bridges = %+v, data = %+v", bridges, data)
		}
	}
}

func equals(a, b [][2]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i][0] != v[0] || b[i][1] != v[1] {
			return false
		}
	}
	return true
}
