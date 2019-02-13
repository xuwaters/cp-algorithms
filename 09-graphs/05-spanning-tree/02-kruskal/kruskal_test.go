package kruskal

import "testing"

func TestKruskal(t *testing.T) {

	dataList := []struct {
		n      int
		edges  []Edge
		weight int
	}{
		{
			n:      2,
			weight: 2,
			edges: []Edge{
				{From: 0, To: 1, Weight: 2},
			},
		},
		{
			n:      8,
			weight: 15,
			edges: []Edge{
				{From: 0, To: 1, Weight: 2},
				{From: 0, To: 2, Weight: 4},
				{From: 0, To: 3, Weight: 6},
				{From: 0, To: 4, Weight: 1},
				{From: 1, To: 2, Weight: 3},
				{From: 2, To: 3, Weight: 1},
				{From: 2, To: 6, Weight: 1},
				{From: 3, To: 4, Weight: 7},
				{From: 3, To: 6, Weight: 1},
				{From: 4, To: 6, Weight: 9},
				{From: 5, To: 6, Weight: 3},
				{From: 5, To: 7, Weight: 4},
				{From: 6, To: 7, Weight: 6},
			},
		},
		{
			n:      8,
			weight: 13,
			edges: []Edge{
				{From: 0, To: 1, Weight: 2},
				{From: 0, To: 2, Weight: 4},
				{From: 0, To: 3, Weight: 6},
				{From: 0, To: 4, Weight: 1},
				{From: 1, To: 2, Weight: 3},
				{From: 2, To: 3, Weight: 1},
				{From: 2, To: 6, Weight: 1},
				{From: 2, To: 4, Weight: 1},
				{From: 3, To: 4, Weight: 7},
				{From: 3, To: 6, Weight: 1},
				{From: 4, To: 6, Weight: 9},
				{From: 5, To: 6, Weight: 3},
				{From: 5, To: 7, Weight: 4},
				{From: 6, To: 7, Weight: 6},
			},
		},
	}

	for _, data := range dataList {
		w, mst := Kruskal(data.edges)
		if w != data.weight {
			t.Fatalf("ERR: weight = %d, mst = %+v, data = %+v", w, mst, data)
		} else {
			t.Logf(" OK: weight = %d, mst = %+v, data = %+v", w, mst, data)
		}
	}
}
