package main

import "fmt"

func main() {
	q := NewPriorityQueue[string]()

	q.Enqueue(3, "")
	q.Enqueue(2, "")
	q.Enqueue(4, "")
	q.Enqueue(7, "")

	q.Dequeue()
	q.Enqueue(0, "")

	iter := q.Iterator()
	for iter.HasNext() {
		fmt.Println(iter.GetNext())
	}
}
