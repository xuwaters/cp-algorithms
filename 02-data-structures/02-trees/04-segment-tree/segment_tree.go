package segment_tree

//
// SegmentTree works same as QuadTree
//

type Segment struct {
	Start  int
	Length int
}

// Covers returns whether current segment covers o segment
func (s Segment) Covers(o Segment) bool {
	if o.Start < s.Start {
		return false
	}
	if o.Start+o.Length > s.Start+s.Length {
		return false
	}
	return true
}

// Intersects returns whether s intersects with o
func (s Segment) Intersects(o Segment) bool {
	if o.Start+o.Length <= s.Start {
		return false
	}
	if o.Start >= s.Start+s.Length {
		return false
	}
	return true
}

func (s Segment) Contains(x int) bool {
	return x >= s.Start && x < s.Start+s.Length
}

type SegmentTree struct {
	Segment
	Children []*SegmentTree
	Payload  []int
}

func NewSegmentTree(start int, length int) *SegmentTree {
	root := &SegmentTree{
		Segment: Segment{
			Start:  start,
			Length: findBestLegnth(length),
		},
	}
	return root
}

func findBestLegnth(length int) int {
	f := uint(0)
	t := uint(32)
	for f < t {
		m := (f + t) / 2
		currLen := 1 << m
		if currLen < length {
			f = m + 1
		} else {
			t = m
		}
	}
	return 1 << f
}

func (tree *SegmentTree) AddSegment(seg Segment, id int) {
	if !tree.Intersects(seg) {
		return
	}
	if seg.Covers(tree.Segment) {
		tree.Payload = append(tree.Payload, id)
		return
	}
	if len(tree.Children) == 0 {
		tree.splitChildren()
	}
	for _, child := range tree.Children {
		child.AddSegment(seg, id)
	}
}

func (tree *SegmentTree) splitChildren() {
	if len(tree.Children) > 0 {
		return
	}
	tree.Children = make([]*SegmentTree, 2)
	childLength := tree.Segment.Length >> 1
	for i := 0; i < len(tree.Children); i++ {
		tree.Children[i] = &SegmentTree{
			Segment: Segment{
				Start:  tree.Start + (i * childLength),
				Length: childLength,
			},
		}
	}
}

func (tree *SegmentTree) QuerySegmentsByPoint(pos int) []int {
	ret := make([]int, 0)

	var dfs func(*SegmentTree)
	dfs = func(root *SegmentTree) {
		if !root.Contains(pos) {
			return
		}
		ret = append(ret, root.Payload...)
		for _, child := range root.Children {
			dfs(child)
		}
	}

	dfs(tree)
	return ret
}

func (tree *SegmentTree) QuerySegmentsBySegment(seg Segment) []int {
	idmap := make(map[int]bool, 0)

	var dfs func(root *SegmentTree)
	dfs = func(root *SegmentTree) {
		if !root.Intersects(seg) {
			return
		}
		for _, id := range root.Payload {
			idmap[id] = true
		}
		for _, child := range root.Children {
			dfs(child)
		}
	}

	dfs(tree)

	ret := make([]int, 0)
	for k := range idmap {
		ret = append(ret, k)
	}

	return ret
}
