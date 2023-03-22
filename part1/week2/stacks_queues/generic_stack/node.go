package generic_stack

type node[T any] struct {
	Val  T
	Next *node[T]
}
