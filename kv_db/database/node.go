package database

type node[T comparable] struct {
	Key   T
	Value T
	Next  *node[T]
	Prev  *node[T]
}
