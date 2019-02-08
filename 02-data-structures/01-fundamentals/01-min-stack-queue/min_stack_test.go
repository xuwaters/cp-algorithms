package min_stack_queue

import "testing"

func TestStack(t *testing.T) {
	dataList := []struct {
		arr []int
	}{
		{arr: []int{3, 24, 4, 3, 9, 5, 8, 433, 4, 234, 234, 234, 8}},
		{arr: []int{1, 3, 4, 54, 56, 67, 78, 8, 9}},
	}

	for _, data := range dataList {
		s := NewStack()

		idx := -1
		for i := 0; i < len(data.arr); i++ {
			s.Push(data.arr[i])
			if idx < 0 || data.arr[i] < data.arr[idx] {
				idx = i
			}
			if s.Min() != data.arr[idx] {
				t.Fatalf("invalid min value from stack")
			}
		}
		for i := len(data.arr) - 1; i >= 0; i-- {
			if s.Pop() != data.arr[i] {
				t.Fatalf("invalid pop sequence")
			}
		}
	}
}
