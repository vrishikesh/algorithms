package main

import (
	"fmt"
	"strings"
)

type Graph struct {
	adj          [][]int
	noOfVertices int
}

func NewGraph(v int) *Graph {
	adj := make([][]int, v)
	for i := 0; i < v; i++ {
		adj[i] = make([]int, 0)
	}

	return &Graph{
		noOfVertices: v,
		adj:          adj,
	}
}

func (g *Graph) AddEdge(from, to int) {
	g.adj[from] = append(g.adj[from], to)
	g.adj[to] = append(g.adj[to], from)
}

func (g *Graph) Adj(v int) []int {
	return g.adj[v]
}

func (g *Graph) V() int {
	return g.noOfVertices
}

func (g *Graph) String() string {
	var s strings.Builder

	for i := 0; i < g.V(); i++ {
		s.WriteString(fmt.Sprintf("%v: ", i))
		for _, v := range g.Adj(i) {
			s.WriteString(fmt.Sprintf("%v ", v))
		}
		s.WriteString("\n")
	}

	return s.String()
}
