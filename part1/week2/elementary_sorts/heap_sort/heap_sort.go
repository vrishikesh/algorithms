package heap_sort

import (
	"priority_queue/priority_queue"
)

type HeapSort[T int] struct {
	pq *priority_queue.PriorityQueue
}

func New[T int](items ...int) *HeapSort[T] {
	return &HeapSort[T]{
		pq: priority_queue.NewPriorityQueue(items...),
	}
}

func (s *HeapSort[T]) Sort(items []int) {
	for i := 0; !s.pq.IsEmpty(); i++ {
		items[i] = s.pq.Dequeue().Score
	}
}

func (s *HeapSort[T]) Name() string {
	return "Heap"
}
