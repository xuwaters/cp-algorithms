package bridge

import (
	"sort"
)

// find bridges offline in O(|V| + |E|)

type Graph struct {
	Adj [][]int
	//
	visited []bool
	tin     []int
	low     []int
	bridges [][2]int
	timer   int
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
			// here: g.low[v] <= g.low[to]
			if g.low[v] < g.low[to] {
				g.bridges = append(g.bridges, [2]int{v, to})
			}
		}
	}
}

func (g *Graph) V() int {
	return len(g.Adj)
}

func (g *Graph) FindBridges() [][2]int {
	n := g.V()
	g.visited = make([]bool, n)
	g.tin = make([]int, n)
	g.low = make([]int, n)
	g.bridges = make([][2]int, 0)
	g.timer = 1 // timer start from 1
	for i := 0; i < n; i++ {
		if !g.visited[i] {
			g.dfs(i, -1)
		}
	}
	g.normalizeBridges()
	return g.bridges
}

func (g *Graph) normalizeBridges() {
	for i := 0; i < len(g.bridges); i++ {
		b := g.bridges[i]
		if b[1] < b[0] {
			b[0], b[1] = b[1], b[0]
		}
		g.bridges[i] = b
	}
	sort.Slice(g.bridges, func(i, j int) bool {
		if g.bridges[i][0] == g.bridges[j][0] {
			return g.bridges[i][1] < g.bridges[j][1]
		}
		return g.bridges[i][0] < g.bridges[j][0]
	})
}
