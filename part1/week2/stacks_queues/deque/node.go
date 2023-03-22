package deque

type node[T comparable] struct {
	Item T
	Next *node[T]
	Prev *node[T]
}
