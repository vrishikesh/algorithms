package generic_queue

import (
	"fmt"
	"strings"
)

type LinkedQueue[T comparable] struct {
	first *node[T]
	last  *node[T]
}

func NewLinkedQueue[T comparable]() *LinkedQueue[T] {
	return &LinkedQueue[T]{}
}

func (q *LinkedQueue[T]) IsEmpty() bool {
	return q.first == nil
}

func (q *LinkedQueue[T]) Enqueue(val T) {
	secondLast := q.last
	q.last = &node[T]{Val: val}
	if q.IsEmpty() {
		q.first = q.last
	} else {
		secondLast.Next = q.last
	}
}

func (q *LinkedQueue[T]) Dequeue() T {
	val := q.first.Val
	q.first = q.first.Next
	if q.IsEmpty() {
		q.last = nil
	}
	return val
}

func (s *LinkedQueue[T]) String() string {
	curr := s.first
	var sb strings.Builder

	for curr != nil {
		sb.WriteString(fmt.Sprint(curr.Val) + " ")
		curr = curr.Next
	}

	return sb.String()
}
