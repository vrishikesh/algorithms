package main

import (
	"fmt"

	"elementary_sorts/shuffling"
)

func main() {
	q := NewPriorityQueue[string]()

	randomNumbers := shuffling.Generate(10)
	fmt.Println(randomNumbers)

	for _, n := range randomNumbers {
		q.Enqueue(n)
	}

	q.Dequeue()
	q.Dequeue()

	iter := q.Iterator()
	for iter.HasNext() {
		fmt.Println(iter.GetNext())
	}
}
