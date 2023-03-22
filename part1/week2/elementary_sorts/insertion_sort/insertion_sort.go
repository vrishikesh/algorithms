package insertion_sort

type InsertionSort[T int] struct{}

func New[T int]() *InsertionSort[T] {
	return &InsertionSort[T]{}
}

func (s *InsertionSort[T]) Sort(items []T) {
	for i := 1; i < len(items); i++ {
		for j := i; j > 0; j-- {
			if items[j] < items[j-1] {
				items[j-1], items[j] = items[j], items[j-1]
			} else {
				break
			}
		}
	}
}

func (s *InsertionSort[T]) Name() string {
	return "Insertion"
}
