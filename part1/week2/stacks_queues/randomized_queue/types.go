package randomized_queue

type Iterator[T comparable] interface {
	HasNext() bool
	GetNext() T
}

type Iterable[T comparable] interface {
	Iterator() Iterator[T]
}

type IRandomizedQueue[T comparable] interface {
	// is the randomized queue empty?
	IsEmpty() bool

	// return the number of items on the randomized queue
	Size() int

	// add the item
	Enqueue(item T)

	// remove and return a random item
	Dequeue() T

	// return a random item (but do not remove it)
	Sample() T

	// return an independent iterator over items in random order
	Iterator() Iterator[T]
}
