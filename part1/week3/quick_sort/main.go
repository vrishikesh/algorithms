package main

import (
	"fmt"

	"elementary_sorts/shuffling"
	"quick_sort/quick_sort"
)

func main() {
	slice := shuffling.Generate(10)
	fmt.Println(slice)

	s := quick_sort.NewRecursive()
	s.Sort(slice)
	fmt.Println(slice)
}
