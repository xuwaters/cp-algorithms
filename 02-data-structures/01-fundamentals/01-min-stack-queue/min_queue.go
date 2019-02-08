package min_stack_queue

type Queue struct {
	s1 *Stack // push stack
	s2 *Stack // pop stack
}

func NewQueue() *Queue {
	return &Queue{
		s1: NewStack(),
		s2: NewStack(),
	}
}

func (q *Queue) Len() int {
	return q.s1.Len() + q.s2.Len()
}

func (q *Queue) Min() int {
	if q.Len() == 0 {
		panic("invalid queue state")
	}
	if q.s1.Len() == 0 {
		return q.s2.Min()
	}
	if q.s2.Len() == 0 {
		return q.s1.Min()
	}
	min1 := q.s1.Min()
	min2 := q.s2.Min()
	if min1 < min2 {
		return min1
	}
	return min2
}

func (q *Queue) Enqueue(val int) {
	q.s1.Push(val)
}

func (q *Queue) Dequeue() int {
	if q.Len() == 0 {
		panic("invalid queue state")
	}
	if q.s2.Len() == 0 {
		for q.s1.Len() > 0 {
			q.s2.Push(q.s1.Pop())
		}
	}
	return q.s2.Pop()
}
