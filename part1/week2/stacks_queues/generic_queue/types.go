package generic_queue

type IQueue[T comparable] interface {
	IsEmpty() bool
	Enqueue(T)
	Dequeue() T
}
