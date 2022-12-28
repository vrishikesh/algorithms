package quick_find

import "fmt"

type QF struct {
	ID []int
}

func New(count int) *QF {
	var uf = QF{
		ID: make([]int, count),
	}

	for i := 0; i < count; i++ {
		uf.ID[i] = i
	}

	return &uf
}

func (t *QF) Union(p int, q int) {
	pid, qid := t.ID[p], t.ID[q]
	for i := 0; i < len(t.ID); i++ {
		if t.ID[i] == pid {
			t.ID[i] = qid
		}
	}
}

func (t *QF) Connected(p int, q int) bool {
	return t.ID[p] == t.ID[q]
}

func (t *QF) Find(int) int {
	return 0
}

func (t *QF) Count() int {
	return len(t.ID)
}

func (t *QF) Print() {
	fmt.Printf("%#v\n", t.ID)
}
