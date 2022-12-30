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

func (s *LinkedStack[T]) String() string {
	curr := s.first
	var sb string

	for curr != nil {
		val := fmt.Sprint(curr.Val)
		sb = val + " " + sb
		curr = curr.Next
	}

	return sb
}
