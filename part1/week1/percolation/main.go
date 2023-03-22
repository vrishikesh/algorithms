package main

import (
	"fmt"
	"os"
	"strconv"

	"percolation/stats"
)

func main() {
	if len(os.Args) < 3 {
		panic("args size and trails required")
	}

	args := os.Args[1:]
	fmt.Println(args)

	size, err := strconv.Atoi(args[0])
	if err != nil {
		panic("size is not number")
	}

	trials, err := strconv.Atoi(args[1])
	if err != nil {
		panic("trails is not number")
	}

	var s stats.IPercolationStats = stats.New(size, trials)

	fmt.Printf("mean                      = %f\n", s.Mean())
	fmt.Printf("stddev                    = %f\n", s.Stddev())
	fmt.Printf("95%% confidence interval   = [%f, %f]\n\n", s.ConfidenceLo(), s.ConfidenceHi())
}
