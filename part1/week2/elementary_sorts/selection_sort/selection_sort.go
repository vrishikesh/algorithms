package selection_sort

type SelectionSort[T int] struct{}

func New[T int]() *SelectionSort[T] {
	return &SelectionSort[T]{}
}

func (s *SelectionSort[T]) Sort(items []T) {
	for i := 0; i < len(items); i++ {
		min := items[i]
		pos := i
		for j := i; j < len(items); j++ {
			if items[j] < min {
				min = items[j]
				pos = j
			}
		}

		items[i], items[pos] = items[pos], items[i]
	}
}

func (s *SelectionSort[T]) Name() string {
	return "Selection"
}
