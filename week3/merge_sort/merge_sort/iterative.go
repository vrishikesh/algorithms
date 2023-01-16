package merge_sort

type MergeSortIterative[T int] struct{}

func NewIterative[T int]() *MergeSortIterative[T] {
	return &MergeSortIterative[T]{}
}

func (s *MergeSortIterative[T]) Sort(items []T) {
	s.iterate(items)
}

func (s *MergeSortIterative[T]) iterate(slice []T) {
	var N = len(slice)
	var aux = make([]T, N)

	for sz := 1; sz < N; sz *= 2 {
		for low := 0; low+sz < N; low += sz + sz {
			high := s.min((low + sz + sz - 1), N-1)
			mid := (low + sz - 1)
			if slice[mid] <= slice[mid+1] {
				continue
			}

			s.merge(slice, aux, low, mid, high)
		}
	}
}

func (s *MergeSortIterative[T]) min(m, n int) int {
	if m <= n {
		return m
	}
	return n
}

func (s *MergeSortIterative[T]) merge(slice, aux []T, low, mid, high int) {
	for i := low; i <= high; i++ {
		aux[i] = slice[i]
	}

	var i, j, k = low, mid + 1, low
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

func (s *MergeSortIterative[T]) Name() string {
	return "Iterative Merge"
}
