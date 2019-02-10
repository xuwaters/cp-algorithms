package parquet

//
// given a grid of size NxM, find number of ways to fill the grid with figures of size 2x1
//
// D[i][mask], where i = [1..N], mask = [0 .. 2^M-1]
// mask = 0 means filled
//
//
//
//
//

type ParquetCount struct {
	N int     // row
	M int     // col
	D [][]int // count, [i][mask]
}

func NewParquetCount(n, m int) *ParquetCount {
	p := &ParquetCount{
		N: n,
		M: m,
		D: make([][]int, n+1),
	}
	p.init()
	return p
}
func (p *ParquetCount) init() {
	mlen := 1 << uint(p.M)
	for r := 0; r < p.N; r++ {
		p.D[r] = make([]int, mlen)
	}
}

// calculate the number of ways to achieve mask state in row r
func (p *ParquetCount) calculate(r int, c int, mask int, nextMask int) {
	if r == p.N {
		return
	}
	if c >= p.M {
		p.D[r+1][nextMask] += p.D[r][mask]
		return
	}
	currMask := 1 << uint(c)
	if (mask & currMask) != 0 {
		// skip current column
		p.calculate(r, c+1, mask, nextMask)
	} else {
		// fill vertically
		p.calculate(r, c+1, mask, nextMask|currMask)
		// fill horizontally (c and c+1 are both 0, need to fill)
		if c+1 < p.M && (mask&currMask) == 0 && (mask&(currMask<<1) == 0) {
			p.calculate(r, c+2, mask, nextMask)
		}
	}
}

func (p *ParquetCount) Solve() int {
	p.D[0][0] = 1
	for r := 0; r < p.N; r++ {
		maskLimit := 1 << uint(p.M)
		for mask := 0; mask < maskLimit; mask++ {
			p.calculate(r, 0, mask, 0)
		}
	}
	return p.D[p.N][0]
}
