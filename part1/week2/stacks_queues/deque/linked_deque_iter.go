package deque

type LinkedDequeIterator[T comparable] struct {
	current *node[T]
}

func (i *LinkedDequeIterator[T]) HasNext() bool {
	return i.current != nil
}

func (i *LinkedDequeIterator[T]) GetNext() T {
	item := i.current.Item
	i.current = i.current.Next
	return item
}
