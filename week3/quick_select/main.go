package main

import (
	"fmt"
	"os"
	"strconv"

	"merge_sort/merge_sort"
	"quick_select/quick_select"
)

func main() {
	var N int = 1
	if len(os.Args) > 1 {
		N, _ = strconv.Atoi(os.Args[1])
	}
	if N < 1 {
		panic("number should be greater than or equal to 1")
	}

	// slice := shuffling.Generate(20)
	slice := []int{999, 1, 6, 324, 545, 234, 543, 54, 76, 878, 64, 453, 87, 565}
	fmt.Println(slice)

	s := quick_select.NewQuickSelect()
	k := s.Select(slice, len(slice)-N)
	fmt.Printf("top %d element is %d\n", N, k)

	m := merge_sort.NewIterative()
	m.Sort(slice)
	fmt.Println(slice)
}
