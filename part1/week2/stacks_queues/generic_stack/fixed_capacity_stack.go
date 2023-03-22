package generic_stack

import (
	"fmt"
	"strings"
)

type FixedCapacityStack[T any] struct {
	s []T
	n int
}

func NewFixedCapacityStack[T any](capacity int) *FixedCapacityStack[T] {
	return &FixedCapacityStack[T]{
		s: make([]T, capacity),
		n: 0,
	}
}

func (s *FixedCapacityStack[T]) IsEmpty() bool {
	return s.n == 0
}

func (s *FixedCapacityStack[T]) Push(val T) {
	s.s[s.n] = val
	s.n += 1
}

func (s *FixedCapacityStack[T]) Pop() T {
	var zero T
	s.n -= 1
	item := s.s[s.n]
	s.s[s.n] = zero
	return item
}

func (s *FixedCapacityStack[T]) String() string {
	var sb strings.Builder

	for i := 0; i < s.n; i++ {
		val := fmt.Sprint(s.s[i])
		sb.WriteString(val + " ")
	}

	return sb.String()
}
