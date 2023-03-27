package graph_processing

import (
	"fmt"
	"strings"
)

type DepthFirstPaths struct {
	marked []bool
	edgeTo []int
	s      int
}

func NewDepthFirstPaths(g *MyGraph, s int) *DepthFirstPaths {
	dfp := &DepthFirstPaths{
		marked: make([]bool, g.V()),
		edgeTo: make([]int, g.V()),
		s:      s,
	}
	dfp.dfs(g, s)
	return dfp
}

func (dfp *DepthFirstPaths) dfs(g *MyGraph, v int) {
	dfp.marked[v] = true
	edges, ok := g.Edges(v)
	if ok {
		for _, e := range edges {
			if !dfp.marked[e] {
				dfp.edgeTo[e] = v
				dfp.dfs(g, e)
			}
		}
	}
}

func (dfp *DepthFirstPaths) HasPathTo(v int) bool {
	return dfp.marked[v]
}

func (dfp *DepthFirstPaths) PathTo(v int) []int {
	if !dfp.HasPathTo(v) {
		return nil
	}
	path := []int{}
	for x := v; x != dfp.s; x = dfp.edgeTo[x] {
		path = append(path, x)
	}
	path = append(path, dfp.s)
	return path
}

func (dfp *DepthFirstPaths) String() string {
	var s strings.Builder
	for v := range dfp.marked {
		s.WriteString(fmt.Sprintf("%v: ", v))
		path := dfp.PathTo(v)
		for _, p := range path {
			s.WriteString(fmt.Sprintf("%v ", p))
		}
		s.WriteString("\n")
	}
	return s.String()
}
