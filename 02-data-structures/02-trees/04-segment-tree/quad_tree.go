package segment_tree

import "fmt"

type Rect struct {
	x int
	y int
	w int
	h int
}

func (r Rect) Covers(o Rect) bool {
	if o.x < r.x || o.y < r.y {
		return false
	}
	if o.x+o.w > r.x+r.w || o.y+o.h > r.y+r.h {
		return false
	}
	return true
}

func (r Rect) Intersects(o Rect) bool {
	if o.x >= r.x+r.w || o.x+o.w <= r.x {
		return false
	}
	if o.y >= r.y+r.h || o.y+o.h <= r.y {
		return false
	}
	return true
}

func (r Rect) Contains(x, y int) bool {
	if x < r.x || x >= r.x+r.w {
		return false
	}
	if y < r.y || y >= r.y+r.h {
		return false
	}
	return true
}

type QuadTree struct {
	region   Rect // 2^k
	children []*QuadTree
	ids      []int
}

func NewQuadTree(region Rect) *QuadTree {
	// find region w, h
	t := &QuadTree{
		region:   region,
		children: nil,
	}
	t.init()
	return t
}

func (t *QuadTree) init() {
	smallestPowerOfTwo := func(n int) int {
		v := 1
		for v < n {
			v <<= 1
		}
		return v
	}
	t.region.w = smallestPowerOfTwo(t.region.w)
	t.region.h = smallestPowerOfTwo(t.region.h)
}

func (t *QuadTree) AddRect(r Rect, id int) {
	if !r.Intersects(t.region) {
		return
	}
	if r.Covers(t.region) {
		t.ids = append(t.ids, id)
		return
	}
	if len(t.children) == 0 {
		t.splitChildren()
	}
	for _, child := range t.children {
		child.AddRect(r, id)
	}
}

func (t *QuadTree) splitChildren() {
	if len(t.children) == 0 {
		t.children = make([]*QuadTree, 4)
		w := t.region.w / 2
		h := t.region.h / 2
		for i := 0; i < len(t.children); i++ {
			x := t.region.x + w*(i%2)
			y := t.region.y + h*(i/2)
			t.children[i] = &QuadTree{
				region: Rect{x: x, y: y, w: w, h: h},
			}
		}
	}
}

func (t *QuadTree) QueryRectsByPoints(x, y int) []int {
	ret := make([]int, 0)

	var dfs func(*QuadTree, int, int)
	dfs = func(root *QuadTree, xx, yy int) {
		if root.region.Contains(xx, yy) {
			ret = append(ret, root.ids...)
			for _, child := range root.children {
				dfs(child, xx, yy)
			}
		}
	}

	dfs(t, x, y)
	return ret
}

func (t *QuadTree) QueryRectsByRegion(r Rect) []int {
	idmap := make(map[int]bool, 0)
	var dfs func(root *QuadTree)
	dfs = func(root *QuadTree) {
		if !root.region.Intersects(r) {
			return
		}
		for _, id := range root.ids {
			idmap[id] = true
		}
		for _, child := range root.children {
			dfs(child)
		}
	}
	dfs(t)

	ret := make([]int, 0)
	for k := range idmap {
		ret = append(ret, k)
	}
	return ret
}

func (t *QuadTree) Print(depth int) {
	prefix := fmt.Sprintf(fmt.Sprintf("%c%ds", '%', depth*2), "")
	fmt.Printf("%s[\n", prefix)
	fmt.Printf("%s  region = %+v,\n", prefix, t.region)
	fmt.Printf("%s  ids = %+v,\n", prefix, t.ids)
	for _, child := range t.children {
		child.Print(depth + 1)
	}
	fmt.Printf("%s],\n", prefix)
}
