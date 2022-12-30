package main

import (
	"fmt"
	"stacks_queues/stack"
)

func main()  {
	var s stack.IStack = stack.NewResizingArrayStackOfStrings()

	s.Push("How")
	s.Push("are")
	s.Push("the")
	s.Pop()
	s.Push("you?")
	fmt.Println(s)
}
