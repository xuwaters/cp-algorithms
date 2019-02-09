package sqrt_decomp

import (
	"testing"
)

func TestRangeMinQuery(t *testing.T) {
	type entry struct {
		ans   int
		start int
		end   int
	}
	dataList := []struct {
		arr     []int
		entries []entry
	}{
		{
			arr: []int{41, 35, 51, 3, 47, 91, 17, 34, 7, 46, 90, 76, 88, 82, 51, 99, 22, 25, 53, 59},
			entries: []entry{
				{ans: 0, start: 0, end: 0},
				{ans: 41, start: 0, end: 1},
				{ans: 76, start: 0, end: 2},
				{ans: 177, start: 0, end: 5},
				{ans: 268, start: 0, end: 6},
				{ans: 319, start: 0, end: 8},
				{ans: 227, start: 1, end: 6},
				{ans: 285, start: 1, end: 9},
				{ans: 698, start: 8, end: 20},
				{ans: 645, start: 10, end: 20},
				{ans: 1017, start: 0, end: 20},
			},
		},
	}

	for _, data := range dataList {
		r := NewRangeSumQuery(data.arr)
		for _, e := range data.entries {
			got := r.Sum(e.start, e.end)
			if got != e.ans {
				t.Fatalf("ERR: got = %+v, e = %+v, arr = %+v", got, e, data.arr[e.start:e.end])
			} else {
				t.Logf(" OK: got = %+v, e = %+v, arr = %+v", got, e, data.arr[e.start:e.end])
			}
		}
	}
}
