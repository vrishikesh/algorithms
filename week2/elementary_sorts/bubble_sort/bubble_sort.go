package bubble_sort

type BubbleSort[T int] struct{}

func New[T int]() *BubbleSort[T] {
	return &BubbleSort[T]{}
}

func (s *BubbleSort[T]) Sort(items []T) {
	for i := 0; i < len(items); i++ {
		for j := len(items) - 1; j > i; j-- {
			if items[j] < items[j-1] {
				items[j-1], items[j] = items[j], items[j-1]
			}
		}
	}
}

func (s *BubbleSort[T]) Name() string {
	return "Bubble"
}
