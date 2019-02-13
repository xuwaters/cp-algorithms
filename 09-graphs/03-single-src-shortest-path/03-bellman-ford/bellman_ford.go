package ford

// single source shortest path with negative weighted edges

type Edge struct {
	From   int
	To     int
	Weight int
}

func BellmanFord(n int, start int, edges []Edge) []int {
	const INF = 1000000000
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	dist[start] = 0

	elen := len(edges)

	for i := 0; i < n-1; i++ {
		relaxed := false
		for j := 0; j < elen; j++ {
			e := edges[j]
			if dist[e.From] < INF {
				if dist[e.To] > dist[e.From]+e.Weight {
					dist[e.To] = dist[e.From] + e.Weight
					relaxed = true
				}
			}
		}
		if !relaxed {
			break
		}
	}

	return dist
}
