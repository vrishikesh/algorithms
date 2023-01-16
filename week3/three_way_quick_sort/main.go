package main

import (
	"elementary_sorts/shuffling"
	"fmt"
	"three_way_quick_sort/three_way_quick_sort"
)

func main() {
	slice := shuffling.Generate(10)
	fmt.Println(slice)

	qsort := three_way_quick_sort.NewRecursive()
	qsort.Sort(slice)
	fmt.Println(slice)
}
