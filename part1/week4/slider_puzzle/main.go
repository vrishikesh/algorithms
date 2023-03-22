package main

import (
	"fmt"

	"slider_puzzle/slider_puzzle"
)

func main() {
	// n := 3
	// size := n * n
	// slice := shuffling.Generate(size)
	matrix := slider_puzzle.Matrix{
		{0, 1, 3},
		{4, 2, 5},
		{7, 8, 6},
	}

	b := slider_puzzle.New(matrix)
	fmt.Println(b)
	b.Start()
}
