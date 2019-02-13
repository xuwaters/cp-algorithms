package kruskal

import (
	"sort"
)

// Edge is bidirectional
type Edge struct {
	From   int // from
	To     int // to
	Weight int // weight
}

func Kruskal(edges []Edge) (int, []Edge) {
	// sort by increasing order of weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	// prepare sets
	set := NewDisjointSet()
	for _, e := range edges {
		set.MakeSet(e.From)
		set.MakeSet(e.To)
	}

	cost := 0
	mst := make([]Edge, 0)

	for _, e := range edges {
		if set.FindSet(e.From) != set.FindSet(e.To) {
			mst = append(mst, e)
			cost += e.Weight
			set.UnionSets(e.From, e.To)
		}
	}

	return cost, mst
}
