package database

type LinkedList[T comparable] struct {
	first *node[T]
	last  *node[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.first == nil
}

func (l *LinkedList[T]) AddLast(key, value T) *node[T] {
	n := &node[T]{Key: key, Value: value}
	if l.IsEmpty() {
		l.first = n
	} else {
		l.last.Next = n
		n.Prev = l.last
	}
	l.last = n
	return n
}

func (l *LinkedList[T]) RemoveFirst() *node[T] {
	n := l.first
	l.first = l.first.Next
	n.Next = nil
	if l.IsEmpty() {
		l.last = nil
	} else {
		l.first.Prev = nil
	}
	return n
}

func (l *LinkedList[T]) RemoveNode(n *node[T]) {
	prev := n.Prev
	next := n.Next

	if prev != nil {
		prev.Next = next
	}

	if next != nil {
		next.Prev = prev
	}

	if prev == nil && next == nil {
		l.first = nil
		l.last = nil
	} else if prev == nil {
		l.first = next
	} else if next == nil {
		l.last = prev
	}

	n.Next = nil
	n.Prev = nil
}
