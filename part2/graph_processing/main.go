package main

import "fmt"

func main() {
	// g1 := NewAdjGraph(5)
	// g1.AddEdge(0, 1)
	// g1.AddEdge(0, 2)
	// g1.AddEdge(0, 3)
	// g1.AddEdge(1, 4)
	// g1.AddEdge(2, 4)
	// g1.AddEdge(3, 4)
	// fmt.Println(g1)

	// g2 := NewSparseGraph(5)
	// g2.AddEdge(0, 1)
	// g2.AddEdge(0, 2)
	// g2.AddEdge(0, 3)
	// g2.AddEdge(1, 4)
	// g2.AddEdge(2, 4)
	// g2.AddEdge(3, 4)
	// fmt.Println(g2)

	g3 := NewMyGraph()
	g3.AddEdge(0, 1)
	g3.AddEdge(0, 2)
	g3.AddEdge(0, 3)
	g3.AddEdge(1, 4)
	g3.AddEdge(2, 4)
	g3.AddEdge(3, 4)
	fmt.Println(g3)

	dfp := NewDepthFirstPaths(g3, 0)
	fmt.Println(dfp)
	fmt.Println(dfp.PathTo(4))

	bfp := NewBreadthFirstPaths(g3, 0)
	fmt.Println(bfp)
	// fmt.Println(bfp.PathTo(4))
}
