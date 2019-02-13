package prim

//
// Weighted undirected graph
// minimum spanning tree
//

type Edge struct {
	From   int
	To     int
	Weight int
}

// Dense, Weighted, Undirected Graph
type Graph struct {
	Adj []map[int]int
}

func NewGraph(n int, edges []Edge) *Graph {
	g := &Graph{
		Adj: make([]map[int]int, n),
	}
	g.init()
	g.AddEdges(edges)
	return g
}

func (g *Graph) init() {
	n := g.V()
	for i := 0; i < n; i++ {
		g.Adj[i] = make(map[int]int)
	}
}

func (g *Graph) V() int {
	return len(g.Adj)
}

func (g *Graph) AddEdges(edges []Edge) {
	for _, e := range edges {
		g.Adj[e.From][e.To] = e.Weight
		g.Adj[e.To][e.From] = e.Weight
	}
}

func (g *Graph) Prim() []Edge {
	n := g.V()

	selected := make(map[int]bool)

	// the minimum edge from n to mst
	minEdges := NewPriorityQueue()
	minEdges.Enqueue(QueueItem{
		Key:      0, // edge.From
		Priority: 0, // edge.Weight
		Payload:  Edge{From: 0, To: -1, Weight: 0},
	})

	mst := make([]Edge, 0)

	for i := 0; i < n; i++ {
		// select min edge from minEdges
		if minEdges.Len() == 0 {
			// no mst
			return nil
		}
		item := minEdges.Dequeue()
		ve := item.Payload.(Edge)
		selected[ve.From] = true
		if ve.To >= 0 {
			mst = append(mst, ve)
		}
		// expand
		for to, weight := range g.Adj[ve.From] {
			if selected[to] {
				continue
			}
			toItem := minEdges.GetByKey(to)
			if toItem == nil || weight < toItem.Priority {
				// update by key or insert
				minEdges.Enqueue(QueueItem{
					Key:      to,
					Priority: weight,
					Payload:  Edge{From: to, To: ve.From, Weight: weight},
				})
			}
		}
	}

	return mst
}

func MSTWeight(mst []Edge) int {
	w := 0
	for _, e := range mst {
		w += e.Weight
	}
	return w
}
