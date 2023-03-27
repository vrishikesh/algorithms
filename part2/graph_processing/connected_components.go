package main

type ConnectedComponents struct {
	marked []bool
	id     []int
	count  int
}

func NewConnectedComponents(g *MyGraph) *ConnectedComponents {
	cc := &ConnectedComponents{
		marked: make([]bool, g.V()),
		id:     make([]int, g.V()),
	}
	for v := 0; v < g.V(); v++ {
		if !cc.marked[v] {
			cc.dfs(g, v)
			cc.count++
		}
	}
	return cc
}

func (cc *ConnectedComponents) dfs(g *MyGraph, v int) {
	cc.marked[v] = true
	cc.id[v] = cc.count
	edges, ok := g.Edges(v)
	if ok {
		for _, e := range edges {
			if !cc.marked[e] {
				cc.dfs(g, e)
			}
		}
	}
}

func (cc *ConnectedComponents) Connected(v, w int) bool {
	return cc.id[v] == cc.id[w]
}

func (cc *ConnectedComponents) Id(v int) int {
	return cc.id[v]
}

func (cc *ConnectedComponents) Count() int {
	return cc.count
}
