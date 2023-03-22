package generic_stack

import (
	"fmt"
)

type LinkedStack[T any] struct {
	first *node[T]
}

func NewLinkedStack[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{}
}

func (s *LinkedStack[T]) IsEmpty() bool {
	return s.first == nil
}

func (s *LinkedStack[T]) Push(val T) {
	second := s.first
	s.first = &node[T]{Val: val, Next: second}
}

func (s *LinkedStack[T]) Pop() T {
	val := s.first.Val
	s.first = s.first.Next
	return val
}

func (s *LinkedStack[T]) Peek() T {
	return s.first.Val
}

func (s *LinkedStack[T]) String() string {
	current := s.first
	var sb string

	for current != nil {
		val := fmt.Sprint(current.Val)
		sb = val + " " + sb
		current = current.Next
	}

	return sb
}
