package directed

import (
	"fmt"
	"strings"
)

type Digraph struct {
	edges [][]int
	v     int
	e     int
}

func NewDigraph(v int) *Digraph {
	return &Digraph{
		edges: make([][]int, v),
		v:     v,
		e:     0,
	}
}

func (g *Digraph) AddEdge(v, w int) {
	g.edges[v] = append(g.edges[v], w)
	g.e++
}

func (g *Digraph) V() int {
	return g.v
}

func (g *Digraph) E() int {
	return g.e
}

func (g *Digraph) Edges(v int) []int {
	return g.edges[v]
}

func (g *Digraph) Reverse() *Digraph {
	r := NewDigraph(g.v)
	for v := 0; v < g.v; v++ {
		for _, w := range g.edges[v] {
			r.AddEdge(w, v)
		}
	}
	return r
}

func (g *Digraph) String() string {
	var s strings.Builder
	for v := 0; v < g.v; v++ {
		s.WriteString(fmt.Sprintf("%d ", v))
		for _, w := range g.edges[v] {
			s.WriteString(fmt.Sprintf("-> %d ", w))
		}
		s.WriteString("\n")
	}
	return s.String()
}
