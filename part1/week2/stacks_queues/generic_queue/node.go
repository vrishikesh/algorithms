package generic_queue

type node[T comparable] struct {
	Val  T
	Next *node[T]
}
