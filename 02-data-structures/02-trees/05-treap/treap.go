package treap

//
// Treap (Cartesian Tree)
//   Tree + Heap => Treap
//
// At each node, data = (Key, Priority)
// Binary Search Tree is sorted by 'Key', and Heap is by 'Priority'
//
//   LeftNode.Key < root.Key <= RightNode.Key
//   LeftNode.Priority <= root.Priority && RigthNode.Priority <= root.Priority (MaxHeap)
//
//
//
// Implementation Operations:
//  Split: split tree into two parts by given key. O(log(N))
//  Insert: insert a new node into tree. O(log(N))
//  Merge:
//
//

type TreapNode struct {
	Key   int
	Prior int
	Left  *TreapNode
	Right *TreapNode
}

type Treap struct {
	root *TreapNode
}

func NewTreap() *Treap {
	t := &Treap{
		root: nil,
	}
	return t
}

func (root *TreapNode) Split(key int) (left, right *TreapNode) {
	if root == nil {
		return
	}
	if key < root.Key {
		left, root.Left = root.Left.Split(key)
		right = root
	} else {
		root.Right, right = root.Right.Split(key)
		left = root
	}
	return
}

// Insert returns new tree node
func (root *TreapNode) Insert(it *TreapNode) (tree *TreapNode) {
	if root == nil {
		tree = it
		return
	}
	if it.Prior > root.Prior {
		it.Left, it.Right = root.Split(it.Key)
		tree = it
	} else {
		if it.Key < root.Key {
			root.Left = root.Left.Insert(it)
		} else {
			root.Right = root.Right.Insert(it)
		}
		tree = root
	}
	return
}

// Merge is used by erase to merge two children nodes
func Merge(left *TreapNode, right *TreapNode) (tree *TreapNode) {
	if left == nil || right == nil {
		if left == nil {
			tree = right
		} else {
			tree = left
		}
		return
	}

	if left.Prior > right.Prior {
		left.Right = Merge(left.Right, right)
		tree = left
	} else {
		right.Left = Merge(left, right.Left)
		tree = right
	}

	return
}

func (root *TreapNode) Erase(key int) (tree *TreapNode) {
	if key == root.Key {
		tree = Merge(root.Left, root.Right)
	} else {
		if key < root.Key {
			root.Left = root.Left.Erase(key)
		} else {
			root.Right = root.Right.Erase(key)
		}
		tree = root
	}
	return
}

// Unite two trees
func Unite(left *TreapNode, right *TreapNode) (tree *TreapNode) {
	if left == nil || right == nil {
		if left == nil {
			tree = right
		} else {
			tree = left
		}
		return
	}

	if left.Prior < right.Prior {
		// to make sure left is the new root
		left, right = right, left
	}
	pLeft, pRight := right.Split(left.Key)
	left.Left = Unite(left.Left, pLeft)
	left.Right = Unite(left.Right, pRight)
	tree = left
	return
}
