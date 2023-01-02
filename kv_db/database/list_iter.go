package database

import (
	"fmt"
)

type LinkedListIterator[T comparable] struct {
	current *node[T]
}

func NewLinkedListIterator[T comparable](n *node[T]) *LinkedListIterator[T] {
	return &LinkedListIterator[T]{
		current: n,
	}
}

func (l *LinkedListIterator[T]) HasNext() bool {
	return l.current != nil
}

func (l *LinkedListIterator[T]) GetNext() string {
	n := l.current
	l.current = l.current.Next
	return fmt.Sprintf("Key: %v, Value: %v", n.Key, n.Value)
}
