package fenwick

//
// To support both range udpate and range query
//   we use two Fenwich Tree
//     B1 for range add: Add(B1, L, val), Add(B1, R+1, -val)
//     B2 for range fix: Add(B2, L, val*(L-1)), Add(B2, R+1, -val*R)
//
//   PrefixSum(idx) = sum(B1, idx) * idx - sum(B2, idx)
//   RangeSum(L, R) = PrefixSum(R) - PrefixSum(L-1)
//

type RangeQuery struct {
	b1 *FenwickTree
	b2 *FenwickTree
}

func NewRangeQuery(n int) *RangeQuery {
	return &RangeQuery{
		b1: NewFenwickTree(n),
		b2: NewFenwickTree(n),
	}
}

func (r *RangeQuery) RangeAdd(start, end int, val int) {
	r.b1.Add(start, val)
	r.b1.Add(end+1, -val)
	r.b2.Add(start, val*(start-1))
	r.b2.Add(end+1, -val*end)
}

func (r *RangeQuery) PrefixSum(idx int) int {
	return r.b1.Sum(idx)*idx - r.b2.Sum(idx)
}

func (r *RangeQuery) RangeSum(start, end int) int {
	return r.PrefixSum(end) - r.PrefixSum(start-1)
}
