package min_stack_queue

type Entry struct {
	Val int
	Min int
}

type Stack struct {
	entries []Entry
}

func NewStack() *Stack {
	return &Stack{
		entries: nil,
	}
}

func (s *Stack) Len() int {
	return len(s.entries)
}

func (s *Stack) top() Entry {
	return s.entries[s.Len()-1]
}

func (s *Stack) Push(val int) {
	min := val
	if s.Len() > 0 {
		top := s.top()
		if top.Min < min {
			min = top.Min
		}
	}
	s.entries = append(s.entries, Entry{
		Val: val,
		Min: min,
	})
}

func (s *Stack) Pop() int {
	if s.Len() == 0 {
		panic("invalid stack state")
	}
	val := s.top().Val
	s.entries = s.entries[0 : s.Len()-1]
	return val
}

func (s *Stack) Min() int {
	if s.Len() == 0 {
		panic("invalid stack state")
	}
	return s.top().Min
}
