package main

import (
	"fmt"
	"strings"
)

type SparseGraph struct {
	edges [][]int
	v     int
}

func NewSparseGraph(v int) *SparseGraph {
	edges := make([][]int, v)
	for i := 0; i < v; i++ {
		edges[i] = make([]int, 0)
	}

	return &SparseGraph{
		v:     v,
		edges: edges,
	}
}

func (g *SparseGraph) AddEdge(from, to int) {
	g.edges[from] = append(g.edges[from], to)
	g.edges[to] = append(g.edges[to], from)
}

func (g *SparseGraph) Edges(v int) []int {
	return g.edges[v]
}

func (g *SparseGraph) V() int {
	return g.v
}

func (g *SparseGraph) String() string {
	var s strings.Builder

	for i := 0; i < g.V(); i++ {
		s.WriteString(fmt.Sprintf("%v: ", i))
		for _, v := range g.Edges(i) {
			s.WriteString(fmt.Sprintf("%v ", v))
		}
		s.WriteString("\n")
	}

	return s.String()
}
