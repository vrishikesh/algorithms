package deque

type LinkedDeque[T comparable] struct {
	first *node[T]
	last  *node[T]
	size  int
}

func NewLinkedDeque[T comparable]() *LinkedDeque[T] {
	return &LinkedDeque[T]{}
}

func (d *LinkedDeque[T]) IsEmpty() bool {
	return d.first == nil
}

func (d *LinkedDeque[T]) Size() int {
	return d.size
}

func (d *LinkedDeque[T]) AddFirst(item T) {
	var zero T
	if zero == item {
		panic("cannot add empty value in deque")
	}

	second := d.first
	d.first = &node[T]{Item: item, Next: second, Prev: nil}
	if d.last == nil {
		d.last = d.first
	}

	d.size += 1
}

func (d *LinkedDeque[T]) AddLast(item T) {
	var zero T
	if zero == item {
		panic("cannot add empty value in deque")
	}

	secondLast := d.last
	d.last = &node[T]{Item: item, Next: nil, Prev: secondLast}
	if d.IsEmpty() {
		d.first = d.last
	} else {
		secondLast.Next = d.last
	}

	d.size += 1
}

func (d *LinkedDeque[T]) RemoveFirst() T {
	if d.IsEmpty() {
		panic("cannot remove item from empty deque")
	}

	item := d.first.Item
	d.first = d.first.Next
	if d.first == nil {
		d.last = d.first
	} else {
		d.first.Prev = nil
	}

	d.size -= 1
	return item
}

func (d *LinkedDeque[T]) RemoveLast() T {
	if d.IsEmpty() {
		panic("cannot remove item from empty deque")
	}

	item := d.last.Item
	d.last = d.last.Prev
	if d.last == nil {
		d.first = d.last
	} else {
		d.last.Next = nil
	}

	d.size -= 1
	return item
}

func (d *LinkedDeque[T]) Iterator() Iterator[T] {
	return &LinkedDequeIterator[T]{
		current: d.first,
	}
}
