package types

type Sorter[T comparable] interface {
	Name() string
	Sort([]T)
}
