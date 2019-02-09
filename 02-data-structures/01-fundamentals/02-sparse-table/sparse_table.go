package sparse_table

import (
	"math"
)

// Sparse Table is a data structure, that allows answering range queries.
// It can answer most range queries in O(logn), but its true power is answering range
// minimum queries (or equivalent range maximum queries). For those queries it can
//  compute the answer in O(1) time.

type MergeFunc func(a, b int) int

type SparseTable struct {
	Content [][]int // [logn][n]int
	Merge   MergeFunc
}

func NewSparseTable(arr []int, f MergeFunc) *SparseTable {
	table := &SparseTable{
		Content: nil,
		Merge:   f,
	}
	table.init(arr)
	return table
}

func (table *SparseTable) init(arr []int) {
	n := len(arr)
	layers := int(math.Ceil(math.Log2(float64(n))))
	table.Content = make([][]int, layers+1)
	for i := 0; i < len(table.Content); i++ {
		table.Content[i] = make([]int, n)
	}
	copy(table.Content[0], arr)
	for j := 1; j < layers; j++ {
		currOffset := 1 << uint(j)
		for i := 0; i+currOffset <= n; i++ {
			prevOffset := 1 << uint(j-1)
			table.Content[j][i] = table.Merge(table.Content[j-1][i], table.Content[j-1][i+prevOffset])
		}
	}
}

type RangeSumQuery struct {
	table *SparseTable
}

func NewRangeSumQuery(arr []int) *RangeSumQuery {
	return &RangeSumQuery{
		table: NewSparseTable(arr, func(a, b int) int {
			return a + b
		}),
	}
}

func (r *RangeSumQuery) Sum(start, end int) int {
	// [start, end)
	layers := len(r.table.Content)
	sum := 0
	for j := layers - 1; j >= 0; j-- {
		length := 1 << uint(j)
		if start+length <= end {
			sum += r.table.Content[j][start]
			start += length
		}
	}
	return sum
}

type RangeMinQuery struct {
	table *SparseTable
	log   []int
}

func NewRangeMinQuery(arr []int) *RangeMinQuery {
	r := &RangeMinQuery{
		table: NewSparseTable(arr, func(a, b int) int {
			if a < b {
				return a
			}
			return b
		}),
	}
	r.init()
	return r
}

func (r *RangeMinQuery) init() {
	logLen := len(r.table.Content[0]) + 1
	r.log = make([]int, logLen)
	r.log[1] = 0
	for i := 2; i < logLen; i++ {
		r.log[i] = r.log[i/2] + 1
	}
}

func (r *RangeMinQuery) Min(start, end int) int {
	// [start, end)
	layer := r.log[end-start]
	layerLen := (1 << uint(layer))
	return r.table.Merge(
		r.table.Content[layer][start],
		r.table.Content[layer][end-layerLen],
	)
}
