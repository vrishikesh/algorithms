package main

import (
	"fmt"

	UF "union_find/weighted_quick_union"
)

func main() {
	var count = 10
	var list = [...][2]int{
		{4, 3},
		{3, 8},
		{6, 5},
		{9, 4},
		{2, 1},
		{8, 9},
		{5, 0},
		{7, 2},
		{6, 1},
		{1, 0},
		{6, 7},
	}

	var uf UFer = UF.New(count)

	for _, v := range list {
		p, q := v[0], v[1]
		if !uf.Connected(p, q) {
			uf.Union(p, q)
			fmt.Printf("p: %d, q: %d\n", p, q)
		} else {
			fmt.Printf("%d and %d are connected\n", p, q)
		}
	}

	uf.Print()
}
