package fixed

import (
	"fmt"
	"strings"
)

//
// directed, unweighted graph G
// for each pair of vertices (i, j), find the number of paths of length k from i to j
//
//
// Ck+1 = Sum(Ck[i][p]*G[p][j], p = 1..n)
// So: Ck = G^k
// use Binary Exponentiation to calculate
//

func FixedLengthPaths(g [][]int, k int) [][]int {
	n := len(g)
	p := matCopy(matNew(n), g)
	res := matOne(matNew(n))
	temp := matNew(n)
	for k > 0 {
		if (k & 1) != 0 {
			matProduct(temp, res, p)
			temp, res = res, temp
		}
		matProduct(temp, p, p)
		temp, p = p, temp
		k >>= 1
	}
	return res
}

func matNew(n int) [][]int {
	m := make([][]int, n)
	for r := 0; r < n; r++ {
		m[r] = make([]int, n)
	}
	return m
}

func matProduct(dst, a, b [][]int) [][]int {
	n := len(a)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			sum := 0
			for i := 0; i < n; i++ {
				sum += a[r][i] * b[i][c]
			}
			dst[r][c] = sum
		}
	}
	return dst
}

func matOne(dst [][]int) [][]int {
	n := len(dst)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if r == c {
				dst[r][c] = 1
			} else {
				dst[r][c] = 0
			}
		}
	}
	return dst
}

func matCopy(dst, src [][]int) [][]int {
	for r := 0; r < len(dst); r++ {
		copy(dst[r], src[r])
	}
	return dst
}

func matPrint(m [][]int) string {
	lines := []string{}
	for r := 0; r < len(m); r++ {
		lines = append(lines, fmt.Sprintf(" %2d: %+3v", r, m[r]))
	}
	return strings.Join(lines, "\n")
}

func matEquals(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for r := 0; r < len(a); r++ {
		if len(a[r]) != len(b[r]) {
			return false
		}
		for c := 0; c < len(a[r]); c++ {
			if a[r][c] != b[r][c] {
				return false
			}
		}
	}
	return true
}
