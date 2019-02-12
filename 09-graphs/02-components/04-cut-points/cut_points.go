package cut_points

import "sort"

// find bridges offline in O(|V| + |E|)

type Graph struct {
	Adj [][]int
	//
	visited   []bool
	tin       []int
	low       []int
	timer     int
	cutPoints []int
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
	// undirected graph
	for _, e := range edges {
		g.Adj[e[0]] = append(g.Adj[e[0]], e[1])
		g.Adj[e[1]] = append(g.Adj[e[1]], e[0])
	}
}

func (g *Graph) dfs(v int, p int) {
	// p is parent of v
	g.visited[v] = true

	g.tin[v] = g.timer // v's first visited timestamp
	g.low[v] = g.timer // min timestamp of backedges of v and v's descendants
	g.timer++

	children := 0
	for _, to := range g.Adj[v] {
		if to == p {
			continue
		}
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
			if g.tin[v] <= g.low[to] && p != -1 {
				g.cutPoints = append(g.cutPoints, v)
			}
			children++
		}
	}

	if p == -1 && children > 1 {
		g.cutPoints = append(g.cutPoints, v)
	}
}

func (g *Graph) V() int {
	return len(g.Adj)
}

func (g *Graph) FindCutPoints() []int {
	n := g.V()
	g.visited = make([]bool, n)
	g.tin = make([]int, n)
	g.low = make([]int, n)
	g.cutPoints = make([]int, 0)
	g.timer = 1 // timer start from 1
	for i := 0; i < n; i++ {
		if !g.visited[i] {
			g.dfs(i, -1)
		}
	}
	g.normalizeCutPoints()
	return g.cutPoints
}

func (g *Graph) normalizeCutPoints() {
	sort.Ints(g.cutPoints)
}
