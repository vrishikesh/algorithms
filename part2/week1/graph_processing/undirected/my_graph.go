package undirected

import (
	"fmt"
	"strings"
)

type MyGraph struct {
	edges        map[int][]int
	vertices     map[int]struct{}
	noOfVertices int
	noOfEdges    int
}

func NewMyGraph() *MyGraph {
	return &MyGraph{
		edges:        make(map[int][]int, 0),
		vertices:     make(map[int]struct{}, 0),
		noOfVertices: 0,
		noOfEdges:    0,
	}
}

func (g *MyGraph) AddEdge(v, w int) {
	g.edges[v] = append(g.edges[v], w)
	g.edges[w] = append(g.edges[w], v)
	if _, ok := g.vertices[v]; !ok {
		g.noOfVertices++
	}
	if _, ok := g.vertices[w]; !ok {
		g.noOfVertices++
	}
	g.vertices[v] = struct{}{}
	g.vertices[w] = struct{}{}
	g.noOfEdges++
}

func (g *MyGraph) Edges(v int) []int {
	return g.edges[v]
}

func (g *MyGraph) V() int {
	return g.noOfVertices
}

func (g *MyGraph) E() int {
	return g.noOfEdges
}

func (g *MyGraph) String() string {
	var s strings.Builder

	s.WriteString(fmt.Sprintf("Vertices: %v\n", g.noOfVertices))
	s.WriteString(fmt.Sprintf("Edges: %v\n", g.noOfEdges))

	for v := range g.vertices {
		s.WriteString(fmt.Sprintf("%v: ", v))
		for _, e := range g.Edges(v) {
			s.WriteString(fmt.Sprintf("%v ", e))
		}
		s.WriteString("\n")
	}

	return s.String()
}
