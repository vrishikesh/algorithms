package database

type Database[T comparable] struct {
	dict   map[T]*node[T]
	list   *LinkedList[T]
	buffer int
	size   int
}

func NewDatabase[T comparable](buffer int) *Database[T] {
	return &Database[T]{
		dict:   make(map[T]*node[T]),
		list:   NewLinkedList[T](),
		buffer: buffer,
		size:   0,
	}
}

func (d *Database[T]) Set(key, value T) {
	if n, ok := d.dict[key]; ok {
		d.list.RemoveNode(n)
	} else if d.size < d.buffer {
		d.size += 1
	} else {
		old := d.list.RemoveFirst()
		delete(d.dict, old.Key)
	}

	d.add(key, value)
}

func (d *Database[T]) Get(key T) T {
	var zero T

	if n, ok := d.dict[key]; ok {
		d.list.RemoveNode(n)
		d.add(key, n.Value)
		return n.Value
	}

	return zero
}

func (d *Database[T]) add(key, value T) {
	n := d.list.AddLast(key, value)
	d.dict[key] = n
}

func (d *Database[T]) Iterator() *LinkedListIterator[T] {
	return &LinkedListIterator[T]{
		current: d.list.first,
	}
}
