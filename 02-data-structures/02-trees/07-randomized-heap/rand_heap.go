package rand_heap

import (
	"math/rand"
)

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

// Unite two min heaps into one
func Unite(t1, t2 *Tree) *Tree {
	if t1 == nil || t2 == nil {
		if t1 == nil {
			return t2
		} else {
			return t1
		}
	}
	// make sure t1 is the new root
	if t2.Val < t1.Val {
		t1, t2 = t2, t1
	}
	// randomize
	if rand.Intn(2) == 1 {
		t1.Left, t1.Right = t1.Right, t1.Left
	}
	t1.Left = Unite(t1.Left, t2)
	return t1
}
