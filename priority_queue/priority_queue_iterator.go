package main

type PriorityQueueIterator[T comparable] struct {
	heap  []*Elem[T]
	index int
}

func NewPriorityQueueIterator[T comparable](heap []*Elem[T]) *PriorityQueueIterator[T] {
	return &PriorityQueueIterator[T]{
		heap:  heap,
		index: 0,
	}
}

func (i *PriorityQueueIterator[T]) HasNext() bool {
	return i.index < len(i.heap)
}

func (i *PriorityQueueIterator[T]) GetNext() *Elem[T] {
	current := i.heap[i.index]
	i.index += 1
	return current
}
