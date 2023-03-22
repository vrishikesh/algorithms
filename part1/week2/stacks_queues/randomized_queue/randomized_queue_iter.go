package randomized_queue

type RandomizedQueueIterator[T comparable] struct {
	Items []T
	index int
}

func (i *RandomizedQueueIterator[T]) HasNext() bool {
	return i.index < len(i.Items)
}

func (i *RandomizedQueueIterator[T]) GetNext() T {
	n := GetRandInt(i.index, len(i.Items))
	i.Items[i.index], i.Items[n] = i.Items[n], i.Items[i.index]

	item := i.Items[i.index]
	i.index += 1
	return item
}
