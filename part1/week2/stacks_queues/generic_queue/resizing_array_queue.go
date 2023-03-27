package generic_queue

import (
	"fmt"
	"strings"
)

type ResizingArrayQueue[T comparable] struct {
	s          []T
	head, tail int
}

func NewResizingArrayQueue[T comparable]() *ResizingArrayQueue[T] {
	return &ResizingArrayQueue[T]{
		s:    make([]T, 1),
		head: 0,
		tail: 0,
	}
}

func (q *ResizingArrayQueue[T]) IsEmpty() bool {
	return q.head-q.tail == 0
}

func (q *ResizingArrayQueue[T]) resize(n int) {
	arr := make([]T, n)
	for i, j := q.head, 0; i < q.tail; i, j = i+1, j+1 {
		arr[j] = q.s[i]
	}
	q.s = arr
}

func (q *ResizingArrayQueue[T]) Enqueue(val T) {
	if q.tail == len(q.s) {
		q.resize(2 * q.tail)
	}

	q.s[q.tail] = val
	q.tail += 1
}

func (q *ResizingArrayQueue[T]) Dequeue() T {
	var zero T
	item := q.s[q.head]
	q.s[q.head] = zero
	q.head += 1

	if q.tail > 0 && (q.tail-q.head) == len(q.s)/4 {
		q.resize(len(q.s) / 2)
	}

	return item
}

func (q *ResizingArrayQueue[T]) String() string {
	var sb strings.Builder

	for i := q.head; i < q.tail; i++ {
		sb.WriteString(fmt.Sprint(q.s[i]) + " ")
	}

	return sb.String()
}
