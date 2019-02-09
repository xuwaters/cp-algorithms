package sparse_table

import "testing"

func TestRangeSumQuery(t *testing.T) {

	type Query struct {
		start int
		end   int
		ans   int
	}

	dataList := []struct {
		arr     []int
		queries []Query
	}{
		{
			arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
			queries: []Query{
				{start: 1, end: 15, ans: 105},
				{start: 2, end: 15, ans: 104},
				{start: 3, end: 15, ans: 102},
				{start: 4, end: 8, ans: 22},
			},
		},
		{
			arr: []int{0, 12, 23, 38, 49, 5, 61, 78, 845, 91, 10, 11, 12, 13, 14},
			queries: []Query{
				{start: 1, end: 15, ans: 1262},
				{start: 2, end: 15, ans: 1250},
				{start: 3, end: 15, ans: 1227},
				{start: 4, end: 8, ans: 193},
			},
		},
	}

	for _, data := range dataList {
		r := NewRangeSumQuery(data.arr)
		t.Logf(">> range sum arr = %+v\n", data.arr)
		for _, q := range data.queries {
			got := r.Sum(q.start, q.end)
			if got != q.ans {
				t.Fatalf("ERR: range sum: got = %+v, query = %+v, arr = %+v", got, q, data.arr[q.start:q.end])
			} else {
				t.Logf(" OK: range sum: got = %+v, query = %+v, arr = %+v", got, q, data.arr[q.start:q.end])
			}
		}
	}
}

func TestRangeMinQuery(t *testing.T) {

	type Query struct {
		start int
		end   int
		ans   int
	}

	dataList := []struct {
		arr     []int
		queries []Query
	}{
		{
			arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
			queries: []Query{
				{start: 1, end: 15, ans: 1},
				{start: 2, end: 15, ans: 2},
				{start: 3, end: 15, ans: 3},
				{start: 4, end: 8, ans: 4},
			},
		},
		{
			arr: []int{0, 1, 2, 3, 4, 2, 6, 7, 1, 9, 10, 11, 12, 4, 14},
			queries: []Query{
				{start: 1, end: 15, ans: 1},
				{start: 2, end: 15, ans: 1},
				{start: 3, end: 15, ans: 1},
				{start: 4, end: 8, ans: 2},
				{start: 7, end: 13, ans: 1},
				{start: 10, end: 14, ans: 4},
			},
		},
	}

	for _, data := range dataList {
		r := NewRangeMinQuery(data.arr)
		t.Logf(">> range min arr = %+v\n", data.arr)
		for _, q := range data.queries {
			got := r.Min(q.start, q.end)
			if got != q.ans {
				t.Fatalf("ERR: range min: got = %+v, query = %+v, arr = %+v", got, q, data.arr[q.start:q.end])
			} else {
				t.Logf(" OK: range min: got = %+v, query = %+v, arr = %+v", got, q, data.arr[q.start:q.end])
			}
		}
	}
}
