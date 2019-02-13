package pape

import (
	"container/list"
)

type Edge struct {
	From   int
	To     int
	Weight int
}

type PathEntry struct {
	Distance int
	Parent   int
}

// The algorithm performs usually quite fast.
// In most cases even faster than Dijkstra's algorithm.
// However there exist cases for which the algorithm takes exponential time.

func ShortestPath(start int, n int, adj map[int][]Edge) []PathEntry {
	const INF = 1e9

	m := make([]int, n) // 0 - already calculated, 1 - currently calculated, 2 - not calculated
	dist := make([]PathEntry, n)
	for i := 0; i < n; i++ {
		dist[i].Distance = INF
		dist[i].Parent = -1
		m[i] = 2
	}
	q := list.New()
	q.PushBack(start)
	for q.Len() > 0 {
		u := q.Remove(q.Front()).(int)
		m[u] = 0
		for _, e := range adj[u] {
			if dist[e.To].Distance > dist[u].Distance+e.Weight {
				dist[e.To].Distance = dist[u].Distance + e.Weight
				dist[e.To].Parent = u
				if m[e.To] == 2 {
					m[e.To] = 1
					q.PushBack(e.To)
				} else if m[e.To] == 0 {
					m[e.To] = 1
					q.PushFront(e.To)
				}
			}
		}
	}

	return dist
}
