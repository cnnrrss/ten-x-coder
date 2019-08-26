package main

import hp "container/heap"

type path struct {
    value int
    nodes []string
}

type minPath []*path

func (h minPath) Len() int           { return len(h) }
func (h minPath) Less(i, j int) bool { return h[i].value < h[j].value }
func (h minPath) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minPath) Push(x interface{}) {
	*h = append(*h, x.(path))
}

func (h *minPath) Pop() interface {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

type edge struct {
	node string
	weight int
}

type graph struct {
	// adjacency list a:[b, e, d], b:[e, f], etc..
	nodes map[string][]edge
}

func newGraph() *graph {
	return &graph{nodes: make(map[string][]edge)}
}

func (g *graph) addEdge(src, dest string, weight int) {
	g.nodes[src] = append(g.nodes[src], edge{node:dest, weight:weight})
	// g.nodes[dest] = append(g.nodes[dest], edge{node:src, weight:weight})
}

func (g *graph) getEdges(node string) []edge {
	return g.nodes[node]
}

func (g *graph) getPath(src, dest string) (int, []string) {
	start := &path{value: 0, nodes: []string{src}}
	h := &heap{values: &minPath{start}}
	visited := make(map[string]bool) // track visited nodes
	for len(*h.values) > 0 {
		p := h.Pop(h.values).(*path)
		node := p.nodes[len(p.nodes)-1]
		
		if visited[node] {
			continue
		}

		if node == dest {
			return p.value, p.nodes
		}

		for _, e := range g.getEdges(node) {
			if !visited[e.node] {
				newPath := &path{value: p.value + e.weight, nodes: append()}
				h.Push()
			}
		}
	}
}