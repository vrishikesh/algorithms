package main

import (
	"fmt"
	"stacks_queues/deque"
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

	var d deque.IDeque[string] = deque.NewLinkedDeque[string]()

	d.AddFirst("first")
	d.AddLast("second")
	d.AddLast("third")
	d.AddLast("last")
	d.AddFirst("zero")

	d.RemoveLast()
	d.RemoveFirst()
	d.RemoveFirst()
	d.RemoveLast()

	iter := d.Iterator()
	for iter.HasNext() {
		fmt.Println(iter.GetNext())
	}
}
