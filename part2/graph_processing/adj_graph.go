package main

import (
	"fmt"
	"strings"
)

type AdjGraph struct {
	edges [][]int
	v     int
}

func NewAdjGraph(v int) *AdjGraph {
	edges := make([][]int, v)
	for i := 0; i < v; i++ {
		edges[i] = make([]int, v)
	}

	return &AdjGraph{
		v:     v,
		edges: edges,
	}
}

func (g *AdjGraph) AddEdge(from, to int) {
	g.edges[from][to] = 1
	g.edges[to][from] = 1
}

func (g *AdjGraph) Edges(v int) []int {
	return g.edges[v]
}

func (g *AdjGraph) V() int {
	return g.v
}

func (g *AdjGraph) String() string {
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
