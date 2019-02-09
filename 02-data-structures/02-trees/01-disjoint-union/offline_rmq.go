package disjoint_set

type Query struct {
	L   int
	R   int
	Idx int
}

type OfflineRMQ struct {
	arr       []int
	container [][]Query // [i] contains []Query whose R == i
	parent    []int     // index of smaller value to the right of index i
	ans       []int
}

func NewOfflineRMQ(arr []int, queries [][2]int) *OfflineRMQ {
	n := len(arr)
	o := &OfflineRMQ{
		arr:       arr,
		container: make([][]Query, n),
		parent:    make([]int, n),
		ans:       make([]int, len(queries)),
	}
	o.init(queries)
	return o
}

func (o *OfflineRMQ) init(queries [][2]int) {
	for i := 0; i < len(queries); i++ {
		q := Query{
			L:   queries[i][0],
			R:   queries[i][1],
			Idx: i,
		}
		o.container[q.R] = append(o.container[q.R], q)
	}
	for i := 0; i < len(o.parent); i++ {
		o.parent[i] = i
	}
}

func (o *OfflineRMQ) Solve() []int {
	s := NewStack()
	for i := 0; i < len(o.arr); i++ {
		// use a non-decreasing stack
		for s.Len() > 0 && o.arr[s.Top()] > o.arr[i] {
			o.parent[s.Top()] = i
			s.Pop()
		}
		s.Push(i)
		// answer queries
		for _, q := range o.container[i] {
			// path compression
			o.ans[q.Idx] = o.arr[o.findSet(q.L)]
		}
	}
	return o.ans
}

func (o *OfflineRMQ) findSet(idx int) int {
	p := o.parent[idx]
	if p == idx {
		return idx
	}
	p = o.findSet(p)
	o.parent[idx] = p
	return p
}

type Stack struct {
	payload []int
}

func NewStack() *Stack {
	return &Stack{
		payload: nil,
	}
}

func (s *Stack) Push(val int) {
	s.payload = append(s.payload, val)
}

func (s *Stack) Pop() int {
	ret := s.Top()
	s.payload = s.payload[0 : s.Len()-1]
	return ret
}

func (s *Stack) Len() int {
	return len(s.payload)
}

func (s *Stack) Top() int {
	return s.payload[s.Len()-1]
}
