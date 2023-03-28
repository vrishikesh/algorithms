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

func (g *MyGraph) AddEdge(v1, v2 int) {
	g.edges[v1] = append(g.edges[v1], v2)
	g.edges[v2] = append(g.edges[v2], v1)
	if _, ok := g.vertices[v1]; !ok {
		g.noOfVertices++
	}
	if _, ok := g.vertices[v2]; !ok {
		g.noOfVertices++
	}
	g.vertices[v1] = struct{}{}
	g.vertices[v2] = struct{}{}
	g.noOfEdges++
}

func (g *MyGraph) Edges(v int) ([]int, bool) {
	edges, ok := g.edges[v]
	return edges, ok
}

func (g *MyGraph) V() int {
	return g.noOfVertices
}

func (g *MyGraph) String() string {
	var s strings.Builder

	s.WriteString(fmt.Sprintf("Vertices: %v\n", g.noOfVertices))
	s.WriteString(fmt.Sprintf("Edges: %v\n", g.noOfEdges))

	for v := range g.vertices {
		s.WriteString(fmt.Sprintf("%v: ", v))
		edges, ok := g.Edges(v)
		if ok {
			for _, e := range edges {
				s.WriteString(fmt.Sprintf("%v ", e))
			}
		}
		s.WriteString("\n")
	}

	return s.String()
}
