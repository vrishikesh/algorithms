package quick_sort

type QuickSortRecursive[T int] struct{}

func NewRecursive[T int]() *QuickSortRecursive[T] {
	return &QuickSortRecursive[T]{}
}

func (s *QuickSortRecursive[T]) Sort(items []T) {
	high := len(items) - 1
	s.dive(items, 0, high)
}

func (s *QuickSortRecursive[T]) dive(slice []T, low, high int) {
	if low >= high {
		return
	}

	k := s.partition(slice, low, high)
	s.dive(slice, low, k-1)
	s.dive(slice, k+1, high)
}

func (s *QuickSortRecursive[T]) partition(slice []T, low, high int) int {
	var i, j int = low, high
	for {
		for i = low + 1; i < high; i++ {
			if !s.less(slice, i, low) {
				break
			}
		}

		for j = high; j > low; j-- {
			if !s.less(slice, low, j) {
				break
			}
		}

		if i >= j {
			break
		}
		s.swap(slice, i, j)
	}

	s.swap(slice, low, j)
	return j
}

func (s *QuickSortRecursive[T]) less(slice []T, x, y int) bool {
	return slice[x] < slice[y]
}

func (s *QuickSortRecursive[T]) swap(slice []T, x, y int) {
	slice[x], slice[y] = slice[y], slice[x]
}

func (s *QuickSortRecursive[T]) Name() string {
	return "Recursive Quick"
}
