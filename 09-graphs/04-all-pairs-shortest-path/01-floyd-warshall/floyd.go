package floyd

// find all pairs shortest path
const INF = int(1e9)

func Floyd(d [][]int) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := len(d)
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if d[i][k] < INF && d[k][j] < INF {
					d[i][j] = min(d[i][j], d[i][k]+d[k][j])
				}
			}
		}
	}
}
