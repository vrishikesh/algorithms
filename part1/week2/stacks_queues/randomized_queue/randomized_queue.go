package randomized_queue

import (
	"math/rand"
	"time"
)

type RandomizedQueue[T comparable] struct {
	Items []T
	first int
	last  int
}

func NewRandomizedQueue[T comparable]() *RandomizedQueue[T] {
	return &RandomizedQueue[T]{
		Items: make([]T, 20),
		first: 0,
		last:  0,
	}
}

func (q *RandomizedQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *RandomizedQueue[T]) Size() int {
	return q.last - q.first
}

func (q *RandomizedQueue[T]) Enqueue(item T) {
	var zero T
	if item == zero {
		panic("item cannot be zero valued")
	}

	q.Items[q.last] = item
	q.last += 1
}

func (q *RandomizedQueue[T]) Dequeue() T {
	if q.IsEmpty() {
		panic("cannot dequeue from empty queue")
	}

	n := GetRandInt(q.first, q.last)
	q.Items[q.first], q.Items[n] = q.Items[n], q.Items[q.first]
	item := q.Items[q.first]
	q.first += 1
	return item
}

func GetRandInt(first, last int) int {
	rand.Seed(time.Now().UnixMicro())
	n := rand.Intn(last-first) + first
	return n
}

func (q *RandomizedQueue[T]) Sample() T {
	n := GetRandInt(q.first, q.last)
	return q.Items[n]
}

func (q *RandomizedQueue[T]) Iterator() Iterator[T] {
	size := q.last - q.first + 1
	items := make([]T, 0, size)
	for i := q.first; i < q.last; i++ {
		items = append(items, q.Items[i])
	}

	return &RandomizedQueueIterator[T]{
		Items: items,
	}
}
