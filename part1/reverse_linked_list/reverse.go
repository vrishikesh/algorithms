package main

type LinkedList[T comparable] struct {
	first *node[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.first == nil
}

func (l *LinkedList[T]) AddNode(item T) {
	second := l.first
	n := &node[T]{Item: item}
	if l.IsEmpty() {
		l.first = n
	} else {
		l.first = n
		l.first.Next = second
	}
}

func (l *LinkedList[T]) Iterator() *LinkedListIterator[T] {
	return NewLinkedListIterator(l.first)
}

func (l *LinkedList[T]) Reverse() {
	if l.IsEmpty() {
		panic("cannot reverse empty list")
	}

	current := l.first

	for current.Next != nil {
		next := current.Next
		current.Next = current.Next.Next
		next.Next = l.first
		l.first = next
	}
}
