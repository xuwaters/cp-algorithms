package fixed

//
// find the shortest path from i to j in k steps.
//
// Lk+1[i][j]=Min(Lk[i][p]+G[p][j], p=1..n)
//
// suppose G[i][j] = -1 means from i to j is unreachable
//

func ShortestFixedLength(g [][]int, k int) [][]int {
	n := len(g)

	if k <= 0 {
		// unreachable
		return matFill(matNew(n), -1)
	}

	p := matCopy(matNew(n), g)
	res := matCopy(matNew(n), g)
	temp := matNew(n)
	k--
	for k > 0 {
		if (k & 1) != 0 {
			matMin(temp, res, p)
			temp, res = res, temp
		}
		matMin(temp, p, p)
		temp, p = p, temp
		k >>= 1
	}
	return res
}

func matMin(dst, a, b [][]int) [][]int {
	n := len(a)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			min := -1
			first := true
			for i := 0; i < n; i++ {
				if a[r][i] < 0 || b[i][c] < 0 {
					// unreachable
					continue
				}
				curr := a[r][i] + b[i][c]
				if first || curr < min {
					min = curr
					first = false
				}
			}
			dst[r][c] = min
		}
	}
	return dst
}

func matFill(dst [][]int, v int) [][]int {
	n := len(dst)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			dst[r][c] = v
		}
	}
	return dst
}
