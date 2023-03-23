package main

import "fmt"

func main() {
	g := NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)

	for i := 0; i < g.V(); i++ {
		for _, v := range g.Adj(i) {
			fmt.Println(i, v)
		}
	}

	fmt.Println(g)
}
