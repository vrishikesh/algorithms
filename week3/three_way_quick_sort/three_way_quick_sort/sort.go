package three_way_quick_sort

type ThreeWayQuickSort[T int] struct{}

func NewRecursive[T int]() *ThreeWayQuickSort[T] {
	return &ThreeWayQuickSort[T]{}
}

func (s *ThreeWayQuickSort[T]) Sort(items []T) {
	high := len(items) - 1
	s.dive(items, 0, high)
}

func (s *ThreeWayQuickSort[T]) dive(slice []T, low, high int) {
	if low >= high {
		return
	}

	lt, gt := s.partition(slice, low, high)
	s.dive(slice, low, lt-1)
	s.dive(slice, gt+1, high)
}

func (s *ThreeWayQuickSort[T]) partition(slice []T, low, high int) (int, int) {
	var lt, i, gt = low, low, high

	for i <= gt {
		if s.less(slice, i, lt) {
			s.swap(slice, i, lt)
			lt += 1
			i += 1
		} else if s.less(slice, lt, i) {
			s.swap(slice, i, gt)
			gt -= 1
		} else {
			i += 1
		}
	}

	return lt, gt
}

func (s *ThreeWayQuickSort[T]) less(slice []T, x, y int) bool {
	return slice[x] < slice[y]
}

func (s *ThreeWayQuickSort[T]) swap(slice []T, x, y int) {
	slice[x], slice[y] = slice[y], slice[x]
}

func (s *ThreeWayQuickSort[T]) Name() string {
	return "Three Way Quick"
}
