package treap

import (
	"math/rand"
)

// build treap from sorted list, O(n)

func BuildTreapFromSortedArray(arr []int) *TreapNode {
	// arr are keys, priority is randomly generated
	n := len(arr)
	if n == 0 {
		return nil
	}
	mid := n / 2
	root := &TreapNode{
		Key:   arr[mid],
		Prior: rand.Intn(10000),
	}
	root.Left = BuildTreapFromSortedArray(arr[0:mid])
	root.Right = BuildTreapFromSortedArray(arr[mid:])
	heapify(root)
	return root
}

// adjust priority
func heapify(root *TreapNode) {
	if root == nil {
		return
	}
	maxNode := root
	if root.Left != nil && root.Left.Prior > maxNode.Prior {
		maxNode = root.Left
	}
	if root.Right != nil && root.Right.Prior > maxNode.Prior {
		maxNode = root.Right
	}
	if maxNode != root {
		maxNode.Prior, root.Prior = root.Prior, maxNode.Prior
		heapify(maxNode)
	}
}
