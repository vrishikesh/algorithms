package merge_sort

type MergeSortRecursive[T int] struct{}

func NewRecursive[T int]() *MergeSortRecursive[T] {
	return &MergeSortRecursive[T]{}
}

func (s *MergeSortRecursive[T]) Sort(items []T) {
	aux := make([]T, len(items))

	high := len(items) - 1
	s.dive(items, aux, 0, high)
}

func (s *MergeSortRecursive[T]) dive(slice, aux []T, low, high int) {
	if low >= high {
		return
	}

	mid := (high + low) / 2

	s.dive(slice, aux, low, mid)
	s.dive(slice, aux, mid+1, high)
	if slice[mid] < slice[mid+1] {
		return
	}

	s.merge(
		slice,
		aux,
		low,
		mid,
		high,
	)
}

func (s *MergeSortRecursive[T]) merge(slice, aux []T, low, mid, high int) {
	for i := low; i <= high; i++ {
		aux[i] = slice[i]
	}

	i, j, k := low, mid+1, low
	for i <= mid && j <= high {
		if aux[i] <= aux[j] {
			slice[k] = aux[i]
			i += 1
		} else {
			slice[k] = aux[j]
			j += 1
		}

		k += 1
	}

	for i <= mid {
		slice[k] = aux[i]
		i += 1
		k += 1
	}
}

func (s *MergeSortRecursive[T]) Name() string {
	return "Recursive Merge"
}
