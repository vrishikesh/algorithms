package main

import "fmt"

type UFer interface {
	union(int, int)
	connected(int, int) bool
	find(int) int
	count() int
}

type UF struct {
	id []int
}

func NewUF(count int) *UF {
	var uf = UF{
		id: make([]int, count),
	}

	for i := 0; i < count; i++ {
		uf.id[i] = i
	}

	return &uf
}

func (t *UF) union(p int, q int) {
	pid, qid := t.id[p], t.id[q]
	for i := 0; i < len(t.id); i++ {
		if t.id[i] == pid {
			t.id[i] = qid
		}
	}
}

func (t *UF) connected(p int, q int) bool {
	return t.id[p] == t.id[q]
}

func (t *UF) find(int) int {
	return 0
}

func (t *UF) count() int {
	return len(t.id)
}

func main() {
	var count = 10
	var list = [...][2]int{
		{4, 3},
		{3, 8},
		{6, 5},
		{9, 4},
		{2, 1},
		{8, 9},
		{5, 0},
		{7, 2},
		{6, 1},
		{1, 0},
		{6, 7},
	}

	uf := NewUF(count)

	for _, v := range list {
		p, q := v[0], v[1]
		if !uf.connected(p, q) {
			uf.union(p, q)
			fmt.Printf("p: %d, q: %d\n", p, q)
		} else {
			fmt.Printf("%d and %d are connected\n", p, q)
		}
	}

	fmt.Printf("%#v\n", uf.id)
}
