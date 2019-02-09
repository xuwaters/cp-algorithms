package fenwick

//
// add value to range [L, R]
// query value of point A[i]
// by using FenwickTree
//  Add(L, x)
//  Add(R+1, x)
//  then A[i] = Sum(i)
//
type RangeUpdate struct {
	fenwick *FenwickTree
}

func NewRangeUpdate(n int) *RangeUpdate {
	return &RangeUpdate{
		fenwick: NewFenwickTree(n),
	}
}

func (r *RangeUpdate) AddRange(start, end int, val int) {
	// add val to range [start, end]
	r.fenwick.Add(start, val)
	r.fenwick.Add(end+1, -val)
}

func (r *RangeUpdate) Query(idx int) int {
	return r.fenwick.Sum(idx)
}
