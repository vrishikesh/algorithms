package merge_sort

type MergeSort[T int] struct{}

func New[T int]() *MergeSort[T] {
	return &MergeSort[T]{}
}

func (s *MergeSort[T]) Sort(items []T) {
	aux := make([]T, len(items))

	high := len(items) - 1
	s.dive(items, aux, 0, high)
}

func (s *MergeSort[T]) dive(slice, aux []T, low, high int) (int, int) {
	mid := (high + low) / 2

	if low < high {
		lo1, hi1 := s.dive(slice, aux, low, mid)
		lo2, hi2 := s.dive(slice, aux, mid+1, high)
		if slice[hi1] < slice[lo2] {
			return low, high
		}

		s.merge(
			slice,
			aux,
			lo1,
			hi1,
			lo2,
			hi2,
		)
	}

	return low, high
}

func (s *MergeSort[T]) merge(slice, aux []T, lo1, hi1, lo2, hi2 int) {
	for i := lo1; i <= hi2; i++ {
		aux[i] = slice[i]
	}

	i, j, k := lo1, lo2, lo1
	for i <= hi1 && j <= hi2 {
		if aux[i] <= aux[j] {
			slice[k] = aux[i]
			i += 1
		} else {
			slice[k] = aux[j]
			j += 1
		}

		k += 1
	}

	for i <= hi1 {
		slice[k] = aux[i]
		i += 1
		k += 1
	}

	for j <= hi2 {
		slice[k] = aux[j]
		j += 1
		k += 1
	}
}

func (s *MergeSort[T]) Name() string {
	return "Merge"
}
