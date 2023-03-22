package queue

import (
	"strings"
)

type ResizingArrayQueueOfStrings struct {
	s          []string
	head, tail int
}

func NewResizingArrayQueueOfStrings() *ResizingArrayQueueOfStrings {
	return &ResizingArrayQueueOfStrings{
		s:    make([]string, 1),
		head: 0,
		tail: 0,
	}
}

func (q *ResizingArrayQueueOfStrings) IsEmpty() bool {
	return q.head-q.tail == 0
}

func (q *ResizingArrayQueueOfStrings) resize(n int) {
	arr := make([]string, n)
	for i, j := q.head, 0; i < q.tail; i, j = i+1, j+1 {
		arr[j] = q.s[i]
	}
	q.s = arr
}

func (q *ResizingArrayQueueOfStrings) Enqueue(val string) {
	if q.tail == len(q.s) {
		q.resize(2 * q.tail)
	}

	q.s[q.tail] = val
	q.tail += 1
}

func (q *ResizingArrayQueueOfStrings) Dequeue() string {
	item := q.s[q.head]
	q.s[q.head] = ""
	q.head += 1

	if q.tail > 0 && (q.tail-q.head) == len(q.s)/4 {
		q.resize(len(q.s) / 2)
	}

	return item
}

func (q *ResizingArrayQueueOfStrings) String() string {
	var sb strings.Builder

	for i := q.head; i < q.tail; i++ {
		sb.WriteString(q.s[i] + " ")
	}

	return sb.String()
}
