package fenwick

//
// fenwick tree is also called Binary Indexed Tree (BIT)
//
// Fenwick tree:
//   calculate f([l,r]) in O(log(n)) time
//   update an element of A in O(log(n)) time
//   require O(N) memory
//   is easy to use and code
//

type FenwickTree struct {
	// one based indexing approach
	data []int
}

func NewFenwickTree(n int) *FenwickTree {
	return &FenwickTree{
		data: make([]int, n+1),
	}
}

func (t *FenwickTree) delta(idx int) int {
	return idx & (-idx)
}

func (t *FenwickTree) Len() int {
	return len(t.data)
}

func (t *FenwickTree) Add(idx int, val int) {
	for idx++; idx < t.Len(); idx += t.delta(idx) {
		t.data[idx] += val
	}
}

func (t *FenwickTree) Sum(idx int) int {
	sum := 0
	for idx++; idx > 0; idx -= t.delta(idx) {
		sum += t.data[idx]
	}
	return sum
}

func (t *FenwickTree) SumOfRange(start, end int) int {
	// sum of range [start, end]
	return t.Sum(end) - t.Sum(start-1)
}
