package queue

type IQueue interface {
	IsEmpty() bool
	Enqueue(string)
	Dequeue() string
}
