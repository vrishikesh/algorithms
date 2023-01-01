package shell_sort

type ShellSort[T int] struct{}

func New[T int]() *ShellSort[T] {
	return &ShellSort[T]{}
}

func (s *ShellSort[T]) Sort(items []T) {
	N := len(items)

	h := 1
	for h < N/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && items[j] < items[j-h]; j -= h {
				items[j-h], items[j] = items[j], items[j-h]
			}
		}
		h /= 3
	}
}

func (s *ShellSort[T]) Name() string {
	return "Shell"
}
