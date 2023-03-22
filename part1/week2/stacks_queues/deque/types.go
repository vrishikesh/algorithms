package deque

type Iterator[T comparable] interface {
	HasNext() bool
	GetNext() T
}

type Iterable[T comparable] interface {
	Iterator() Iterator[T]
}

type IDeque[T comparable] interface {
	// is the deque empty?
	IsEmpty() bool

	// return the number of items on the deque
	Size() int

	// add the item to the front
	AddFirst(Item T)

	// add the item to the back
	AddLast(Item T)

	// remove and return the item from the front
	RemoveFirst() T

	// remove and return the item from the back
	RemoveLast() T

	// return an iterator over items in order from front to back
	Iterator() Iterator[T]
}
