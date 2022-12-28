package weighted_quick_union

import "fmt"

type WQU struct {
	ID []int
	SZ []int
}

func New(count int) *WQU {
	var wqu = WQU{
		ID: make([]int, count),
		SZ: make([]int, count),
	}

	for i := 0; i < count; i++ {
		wqu.ID[i] = i
		wqu.SZ[i] = 1
	}

	return &wqu
}

func (t *WQU) Union(p int, q int) {
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

func (t *WQU) Connected(p int, q int) bool {
	return t.root(p) == t.root(q)
}

func (t *WQU) Find(int) int {
	return 0
}

func (t *WQU) Count() int {
	return len(t.ID)
}

func (t *WQU) root(i int) int {
	for i != t.ID[i] {
		i = t.ID[i]
	}

	return i
}

func (t *WQU) Print() {
	fmt.Printf("%#v\n", t.ID)
	fmt.Printf("%#v\n", t.SZ)
}
