package queue

import "strings"

type LinkedQueueOfStrings struct {
	first *node
	last  *node
}

func NewLinkedQueueOfStrings() *LinkedQueueOfStrings {
	return &LinkedQueueOfStrings{}
}

func (q *LinkedQueueOfStrings) IsEmpty() bool {
	return q.first == nil
}

func (q *LinkedQueueOfStrings) Enqueue(val string) {
	secondLast := q.last
	q.last = &node{Val: val}
	if q.IsEmpty() {
		q.first = q.last
	} else {
		secondLast.Next = q.last
	}
}

func (q *LinkedQueueOfStrings) Dequeue() string {
	val := q.first.Val
	q.first = q.first.Next
	if q.IsEmpty() {
		q.last = nil
	}
	return val
}

func (s *LinkedQueueOfStrings) String() string {
	curr := s.first
	var sb strings.Builder

	for curr != nil {
		sb.WriteString(curr.Val + " ")
		curr = curr.Next
	}

	return sb.String()
}
