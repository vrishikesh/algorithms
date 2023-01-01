package main

type LinkedListIterator[T comparable] struct {
	current *node[T]
}

func NewLinkedListIterator[T comparable](c *node[T]) *LinkedListIterator[T] {
	return &LinkedListIterator[T]{
		current: c,
	}
}

func (i *LinkedListIterator[T]) HasNext() bool {
	return i.current != nil
}

func (i *LinkedListIterator[T]) GetNext() T {
	item := i.current.Item
	i.current = i.current.Next
	return item
}
