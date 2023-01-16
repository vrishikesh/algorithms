package quick_select

type QuickSelect[T int] struct{}

func NewQuickSelect[T int]() *QuickSelect[T] {
	return &QuickSelect[T]{}
}

func (s *QuickSelect[T]) Select(slice []T, k int) T {
	var low, high = 0, len(slice) - 1

	for low < high {
		j := s.partition(slice, low, high)
		if k < j {
			high = j - 1
		} else if k > j {
			low = j + 1
		} else {
			return slice[k]
		}
	}

	return slice[k]
}

func (s *QuickSelect[T]) partition(slice []T, low, high int) int {
	var i, j = low, high + 1
	for {
		for i = i + 1; i < high; i++ {
			if !s.less(slice, i, low) {
				break
			}
		}

		for j = j - 1; j > low; j-- {
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

func (s *QuickSelect[T]) less(slice []T, x, y int) bool {
	return slice[x] < slice[y]
}

func (s *QuickSelect[T]) swap(slice []T, x, y int) {
	slice[x], slice[y] = slice[y], slice[x]
}
