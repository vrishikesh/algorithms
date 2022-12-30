package main

import (
	"fmt"
	"stacks_queues/randomized_queue"
)

func main() {
	// var s stack.IStack = stack.NewResizingArrayStackOfStrings()
	// s.Push("How")
	// s.Push("are")
	// s.Push("the")
	// s.Pop()
	// s.Push("you?")
	// fmt.Println(s)

	// var q queue.IQueue = queue.NewResizingArrayQueueOfStrings()
	// q.Enqueue("the")
	// q.Enqueue("How")
	// q.Enqueue("are")
	// q.Dequeue()
	// q.Enqueue("you?")
	// fmt.Println(q)

	// var s generic_stack.IStack[string] = generic_stack.NewFixedCapacityStack[string](10)
	// s.Push("How")
	// s.Push("are")
	// s.Push("the")
	// s.Pop()
	// s.Push("you?")
	// fmt.Println(s)

	// var d deque.IDeque[string] = deque.NewLinkedDeque[string]()

	// d.AddFirst("first")
	// d.AddLast("second")
	// d.AddLast("third")
	// d.AddLast("last")
	// d.AddFirst("zero")

	// d.RemoveLast()
	// d.RemoveFirst()
	// d.RemoveFirst()
	// d.RemoveLast()

	// iter := d.Iterator()
	// for iter.HasNext() {
	// 	fmt.Println(iter.GetNext())
	// }

	var q randomized_queue.IRandomizedQueue[string] = randomized_queue.NewRandomizedQueue[string]()

	q.Enqueue("first")
	q.Enqueue("second")
	q.Enqueue("third")
	q.Enqueue("forth")
	q.Enqueue("fifth")
	q.Enqueue("sixth")
	q.Enqueue("seventh")
	q.Enqueue("eighth")
	q.Enqueue("ninth")
	q.Enqueue("tenth")

	fmt.Printf("Sample: %s\n", q.Sample())
	fmt.Printf("Sample: %s\n", q.Sample())
	fmt.Printf("Sample: %s\n", q.Sample())
	fmt.Printf("Sample: %s\n", q.Sample())

	fmt.Printf("Dequeue: %s\n", q.Dequeue())
	fmt.Printf("Dequeue: %s\n\n", q.Dequeue())

	iter := q.Iterator()
	for iter.HasNext() {
		fmt.Println(iter.GetNext())
	}
}
