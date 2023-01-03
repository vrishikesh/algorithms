package main

type PriorityQueue[T comparable] struct {
	heap []*Elem[T]
}

func NewPriorityQueue[T comparable]() *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}

func (q *PriorityQueue[T]) IsEmpty() bool {
	return len(q.heap) == 0
}

func (q *PriorityQueue[T]) Enqueue(score int) {
	var data T
	e := &Elem[T]{
		Score: score,
		Data:  data,
	}
	q.heap = append(q.heap, e)
	q.bubbleUp(len(q.heap) - 1)
}

func (q *PriorityQueue[T]) Dequeue() *Elem[T] {
	if q.IsEmpty() {
		panic("heap empty cannot dequeue")
	}

	e := q.heap[0]
	lastIdx := len(q.heap) - 1
	q.heap[0] = q.heap[lastIdx]
	q.heap[lastIdx] = nil
	q.heap = q.heap[:lastIdx]
	q.bubbleDown(0)

	return e
}

func (q *PriorityQueue[T]) bubbleUp(idx int) {
	if q.checkBigParent(idx) {
		q.bubbleUp(q.swapBigParent(idx))
	}
}

func (q *PriorityQueue[T]) bubbleDown(idx int) {
	if q.checkSmallChild(idx) {
		q.bubbleUp(q.swapSmallestChild(idx))
	}
}

func (q *PriorityQueue[T]) getParentIdx(idx int) int {
	return (idx - 1) / 2
}

func (q *PriorityQueue[T]) checkBigParent(idx int) bool {
	pIdx := q.getParentIdx(idx)
	h := q.heap
	return h[idx].Score < h[pIdx].Score
}

func (q *PriorityQueue[T]) checkSmallChild(idx int) bool {
	idx1 := idx*2 + 1
	idx2 := idx*2 + 2

	var indexes []int

	if idx1 < len(q.heap) {
		indexes = append(indexes, idx1)
	}

	if idx2 < len(q.heap) {
		indexes = append(indexes, idx2)
	}

	for _, index := range indexes {
		if q.heap[index].Score < q.heap[idx].Score {
			return true
		}
	}

	return false
}

func (q *PriorityQueue[T]) swapBigParent(idx int) int {
	pIdx := q.getParentIdx(idx)
	h := q.heap
	h[idx], h[pIdx] = h[pIdx], h[idx]
	return pIdx
}

func (q *PriorityQueue[T]) swapSmallestChild(idx int) int {
	idx1 := idx*2 + 1
	idx2 := idx*2 + 2

	var indexes []int

	if idx1 < len(q.heap) {
		indexes = append(indexes, idx1)
	}

	if idx2 < len(q.heap) {
		indexes = append(indexes, idx2)
	}

	smallest := idx

	for _, index := range indexes {
		if q.heap[index].Score < q.heap[smallest].Score {
			smallest = index
		}
	}

	h := q.heap
	h[idx], h[smallest] = h[smallest], h[idx]

	return smallest
}

func (q *PriorityQueue[T]) Iterator() *PriorityQueueIterator[T] {
	return NewPriorityQueueIterator(q.heap)
}
