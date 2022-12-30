package generic_stack

type IStack[T any] interface {
	IsEmpty() bool
	Push(T)
	Pop() T
}
