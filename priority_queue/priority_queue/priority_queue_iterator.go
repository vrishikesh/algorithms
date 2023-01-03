package priority_queue

type PriorityQueueIterator struct {
	heap  []*Elem
	index int
}

func NewPriorityQueueIterator(heap []*Elem) *PriorityQueueIterator {
	return &PriorityQueueIterator{
		heap:  heap,
		index: 0,
	}
}

func (i *PriorityQueueIterator) HasNext() bool {
	return i.index < len(i.heap)
}

func (i *PriorityQueueIterator) GetNext() *Elem {
	current := i.heap[i.index]
	i.index += 1
	return current
}
