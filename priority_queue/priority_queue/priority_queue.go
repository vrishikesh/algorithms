package priority_queue

type PriorityQueue struct {
	heap []*Elem
}

func NewPriorityQueue(items ...int) *PriorityQueue {
	var heap []*Elem
	for _, item := range items {
		e := &Elem{
			Score: item,
			Data:  nil,
		}
		heap = append(heap, e)
	}

	pq := &PriorityQueue{}
	pq.heap = heap

	for i := pq.lastNonLeafIdx(); i >= 0; i-- {
		pq.bubbleDown(i)
	}

	return pq
}

func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.heap) == 0
}

func (pq *PriorityQueue) Enqueue(score int) {
	e := &Elem{
		Score: score,
		Data:  nil,
	}
	pq.heap = append(pq.heap, e)
	pq.bubbleUp(len(pq.heap) - 1)
}

func (pq *PriorityQueue) Dequeue() *Elem {
	if pq.IsEmpty() {
		panic("heap empty cannot dequeue")
	}

	e := pq.heap[0]
	pq.replaceRootWithLastChild()
	pq.bubbleDown(0)

	return e
}

func (pq *PriorityQueue) bubbleUp(idx int) {
	if pq.hasBiggerParent(idx) {
		pq.bubbleUp(pq.switchWithParent(idx))
	}
}

func (pq *PriorityQueue) bubbleDown(idx int) {
	if pq.hasSmallerChild(idx) {
		pq.bubbleDown(pq.switchWithSmallestChild(idx))
	}
}

func (pq *PriorityQueue) parentIdx(idx int) int {
	return (idx - 1) / 2
}

func (pq *PriorityQueue) hasBiggerParent(idx int) bool {
	pIdx := pq.parentIdx(idx)
	h := pq.heap
	return h[idx].Score < h[pIdx].Score
}

func (pq *PriorityQueue) childrenIdxs(idx int) []int {
	idxs := make([]int, 0, 2)

	idx1 := idx*2 + 1
	idx2 := idx*2 + 2

	if idx1 < len(pq.heap) {
		idxs = append(idxs, idx1)
	}

	if idx2 < len(pq.heap) {
		idxs = append(idxs, idx2)
	}

	return idxs
}

func (pq *PriorityQueue) hasSmallerChild(idx int) bool {
	if idx >= len(pq.heap) {
		return false
	}

	score := pq.heap[idx].Score

	for _, i := range pq.childrenIdxs(idx) {
		childScore := pq.heap[i].Score
		if childScore < score {
			return true
		}
	}

	return false
}

func (pq *PriorityQueue) switchWithParent(idx int) int {
	pIdx := pq.parentIdx(idx)
	h := pq.heap

	h[idx], h[pIdx] = h[pIdx], h[idx]
	return pIdx
}

func (pq *PriorityQueue) switchWithSmallestChild(idx int) int {
	h := pq.heap
	smallest := idx

	for _, i := range pq.childrenIdxs(idx) {
		if pq.heap[i].Score < pq.heap[smallest].Score {
			smallest = i
		}
	}

	h[idx], h[smallest] = h[smallest], h[idx]

	return smallest
}

func (pq *PriorityQueue) lastNonLeafIdx() int {
	lastIdx := len(pq.heap) - 1

	return pq.parentIdx(lastIdx)
}

func (pq *PriorityQueue) replaceRootWithLastChild() {
	lastIdx := len(pq.heap) - 1
	pq.heap[0] = pq.heap[lastIdx]
	pq.heap[lastIdx] = nil
	pq.heap = pq.heap[:lastIdx]
}

func (pq *PriorityQueue) Iterator() *PriorityQueueIterator {
	return NewPriorityQueueIterator(pq.heap)
}
