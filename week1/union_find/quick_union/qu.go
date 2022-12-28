package quick_union

import "fmt"

type QU struct {
	ID []int
}

func New(count int) *QU {
	var qu = QU{
		ID: make([]int, count),
	}

	for i := 0; i < count; i++ {
		qu.ID[i] = i
	}

	return &qu
}

func (t *QU) Union(p int, q int) {
	i := t.root(p)
	j := t.root(q)
	t.ID[i] = j
}

func (t *QU) Connected(p int, q int) bool {
	return t.root(p) == t.root(q)
}

func (t *QU) Find(int) int {
	return 0
}

func (t *QU) Count() int {
	return len(t.ID)
}

func (t *QU) root(i int) int {
	for i != t.ID[i] {
		i = t.ID[i]
	}

	return i
}

func (t *QU) Print() {
	fmt.Printf("%#v\n", t.ID)
}
