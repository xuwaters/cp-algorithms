package sqrt_decomp

import (
	"math"
)

//
// split input array of length n into blocks of length sqrt(n)
//

type RangeSumQuery struct {
	data   []int
	blocks []int
}

func NewRangeSumQuery(arr []int) *RangeSumQuery {
	n := len(arr)
	block := int(math.Sqrt(float64(n))) + 1
	// '+1' to make sure block * block >= n
	r := &RangeSumQuery{
		data:   arr,
		blocks: make([]int, block),
	}
	r.init()
	return r
}

func (r *RangeSumQuery) init() {
	n := len(r.data)
	blk := len(r.blocks)
	for i := 0; i < n; i++ {
		r.blocks[i/blk] += r.data[i]
	}
}

func (r *RangeSumQuery) Sum(start, end int) int {
	// range is [start, end)
	sum := 0
	blk := len(r.blocks)
	startIdx := start / blk
	endIdx := end / blk
	nextBlock := (startIdx + 1) * blk

	if end <= nextBlock {
		for i := start; i < end; i++ {
			sum += r.data[i]
		}
	} else {
		for i := start; i < nextBlock; i++ {
			sum += r.data[i]
		}
		for i := startIdx + 1; i < endIdx; i++ {
			sum += r.blocks[i]
		}
		for i := endIdx * blk; i < end; i++ {
			sum += r.data[i]
		}
	}

	return sum
}
