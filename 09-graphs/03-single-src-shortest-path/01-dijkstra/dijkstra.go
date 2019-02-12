package dijkstra

type WeightedGraph struct {
	Adj []map[int]int // from -> { to: weight }, where weight >= 0
}

func NewWeightedGraph(n int) *WeightedGraph {
	g := &WeightedGraph{
		Adj: make([]map[int]int, n),
	}
	g.init()
	return g
}

type Edge struct {
	From   int
	To     int
	Weight int
}

func (g *WeightedGraph) init() {
	for i := 0; i < len(g.Adj); i++ {
		g.Adj[i] = make(map[int]int)
	}
}

func (g *WeightedGraph) AddEdges(edges []Edge) {
	for _, e := range edges {
		g.Adj[e.From][e.To] = e.Weight
	}
}

type PathEntry struct {
	Distance int // shortest path, -1 if not reachable
	Parent   int // parent
}

func (g *WeightedGraph) ShortestPaths(s int) []PathEntry {
	n := len(g.Adj)
	visited := make([]bool, n)
	dist := make([]PathEntry, n)
	for i := 0; i < n; i++ {
		dist[i].Parent = -1
		dist[i].Distance = -1
	}

	dist[s].Distance = 0

	for i := 0; i < n; i++ {
		v := -1
		for j := 0; j < n; j++ {
			if visited[j] {
				continue
			}
			if dist[j].Distance < 0 {
				continue
			}
			if v < 0 || dist[j].Distance < dist[v].Distance {
				v = j
			}
		}
		if v < 0 {
			break
		}
		visited[v] = true
		for to, weight := range g.Adj[v] {
			currDist := dist[v].Distance + weight
			if dist[to].Distance < 0 || currDist < dist[to].Distance {
				dist[to].Distance = currDist
				dist[to].Parent = v
			}
		}
	}

	return dist
}

func ToDistances(entries []PathEntry) []int {
	dist := make([]int, len(entries))
	for i, e := range entries {
		dist[i] = e.Distance
	}
	return dist
}
