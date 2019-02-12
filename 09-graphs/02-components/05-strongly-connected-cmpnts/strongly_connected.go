package strongly_connected

import (
	"sort"
)

type Graph struct {
	Adj [][]int
	//
	visited []bool
	tin     []int
	low     []int
	timer   int
	scc     []int // strongly connected components count
}

func NewGraph(n int) *Graph {
	g := &Graph{
		Adj: make([][]int, n),
	}
	g.init()
	return g
}

func (g *Graph) init() {
	// do nothing
}

func (g *Graph) AddEdges(edges [][2]int) {
	// directed graph
	for _, e := range edges {
		g.Adj[e[0]] = append(g.Adj[e[0]], e[1])
	}
}

func (g *Graph) dfs(v int, p int) {
	// p is parent of v
	g.visited[v] = true

	g.tin[v] = g.timer // v's first visited timestamp
	g.low[v] = g.timer // min timestamp of backedges of v and v's descendants
	g.timer++

	for _, to := range g.Adj[v] {
		// this is a directed graph, don't need to check to == parent
		if g.visited[to] {
			if g.low[v] > g.tin[to] {
				g.low[v] = g.tin[to]
			}
		} else {
			g.dfs(to, v)
			// this line should be put before 'if' check
			if g.low[v] > g.low[to] {
				g.low[v] = g.low[to]
			}
			// here: g.low[v] <= g.low[to]
		}
	}

	if g.tin[v] == g.low[v] {
		g.scc = append(g.scc, v)
	}
}

func (g *Graph) V() int {
	return len(g.Adj)
}

func (g *Graph) FindStronglyConnectedComponents() []int {
	n := g.V()
	g.visited = make([]bool, n)
	g.tin = make([]int, n)
	g.low = make([]int, n)
	g.timer = 1 // timer start from 1
	g.scc = make([]int, 0)
	for i := 0; i < n; i++ {
		if !g.visited[i] {
			g.dfs(i, -1)
		}
	}
	sort.Ints(g.scc)
	return g.scc
}
