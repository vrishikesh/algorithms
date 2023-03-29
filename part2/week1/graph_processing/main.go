package main

import (
	"fmt"

	"graph_processing/directed"
	"graph_processing/paths"
	"graph_processing/sort"
	"graph_processing/undirected"
)

func main() {
	g3 := undirected.NewMyGraph()
	g3.AddEdge(0, 1)
	g3.AddEdge(0, 2)
	g3.AddEdge(0, 3)
	g3.AddEdge(1, 4)
	g3.AddEdge(2, 4)
	g3.AddEdge(3, 4)
	fmt.Println(g3)

	dfp := paths.NewDepthFirstPaths(g3, 0)
	fmt.Println(dfp)
	fmt.Println(dfp.PathTo(4))

	// bfp := undirected.NewBreadthFirstPaths(g3, 0)
	// fmt.Println(bfp)
	// fmt.Println(bfp.PathTo(4))

	g4 := directed.NewDigraph(5)
	g4.AddEdge(0, 1)
	g4.AddEdge(0, 2)
	g4.AddEdge(0, 3)
	g4.AddEdge(1, 4)
	g4.AddEdge(2, 4)
	g4.AddEdge(3, 4)
	fmt.Println(g4)

	dfp = paths.NewDepthFirstPaths(g4, 0)
	fmt.Println(dfp)
	fmt.Println(dfp.PathTo(4))

	ts := sort.NewTopologicalSort(g4)
	fmt.Println(ts)
}
