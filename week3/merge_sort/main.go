package main

import (
	"fmt"

	"elementary_sorts/shuffling"
	"merge_sort/merge_sort"
)

func main() {
	slice := shuffling.Generate(10)
	fmt.Println(slice)

	s := merge_sort.NewRecursive()
	s.Sort(slice)
	fmt.Println(slice)

	// s := merge_sort.NewIterative()
	// s.Sort(slice)
	// fmt.Println(slice)
}
