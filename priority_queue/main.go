package main

import (
	"fmt"

	"elementary_sorts/shuffling"
	"priority_queue/priority_queue"
)

func main() {
	capacity := 10
	slice := shuffling.Generate(capacity)
	fmt.Printf("slice: %v\n\n", slice)

	q := priority_queue.NewPriorityQueue(slice...)

	for i := 0; i < capacity; i++ {
		DequeueAndPrint(q)
	}
}

func DequeueAndPrint(q *priority_queue.PriorityQueue) {
	fmt.Println("Dequeue:", q.Dequeue())
	iter := q.Iterator()
	for iter.HasNext() {
		fmt.Println(iter.GetNext())
	}
	fmt.Println()
}
