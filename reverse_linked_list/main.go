package main

import "fmt"

func main() {
	l := NewLinkedList[int]()
	l.AddNode(1)
	l.AddNode(2)
	l.AddNode(3)

	fmt.Println("Linked List:")
	iter := l.Iterator()
	for iter.HasNext() {
		fmt.Println(iter.GetNext())
	}

	fmt.Println("Linked List After Reverse:")
	l.Reverse()
	iter = l.Iterator()
	for iter.HasNext() {
		fmt.Println(iter.GetNext())
	}
}
