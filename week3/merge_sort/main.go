package main

import (
	"fmt"

	"elementary_sorts/shuffling"
	"merge_sort/merge_sort"
)

func main() {
	slice := shuffling.Generate(10)
	fmt.Println(slice)

	s := merge_sort.New()
	s.Sort(slice)
	fmt.Println(slice)
}
