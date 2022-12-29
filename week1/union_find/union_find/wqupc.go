package union_find

import "fmt"

type WQUPC struct {
	ID []int
	SZ []int
}

func NewWQUPC(count int) *WQUPC {
	var wqupc = WQUPC{
		ID: make([]int, count),
		SZ: make([]int, count),
	}

	for i := 0; i < count; i++ {
		wqupc.ID[i] = i
		wqupc.SZ[i] = 1
	}

	return &wqupc
}

func (t *WQUPC) Union(p int, q int) {
	i := t.root(p)
	j := t.root(q)
	if i == j {
		return
	}

	if t.SZ[i] < t.SZ[j] {
		t.ID[i] = j
		t.SZ[j] += t.SZ[i]
	} else {
		t.ID[j] = i
		t.SZ[i] += t.SZ[j]
	}
}

func (t *WQUPC) Connected(p int, q int) bool {
	return t.root(p) == t.root(q)
}

func (t *WQUPC) Find(int) int {
	return 0
}

func (t *WQUPC) Count() int {
	return len(t.ID)
}

func (t *WQUPC) root(i int) int {
	for i != t.ID[i] {
		t.ID[i] = t.ID[t.ID[i]]
		i = t.ID[i]
	}

	return i
}

func (t *WQUPC) Print() {
	fmt.Printf("WQUPC.ID: %#v\n", t.ID)
	fmt.Printf("WQUPC.SZ: %#v\n", t.SZ)
}
