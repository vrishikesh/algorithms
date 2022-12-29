package percolation

import (
	"fmt"

	UF "union_find/union_find"
)

type Percolation struct {
	ID                []int
	numberOfOpenSites int
	size              int
	uf                UF.UFer
}

// creates n-by-n grid, with all sites initially blocked
func New(n int) *Percolation {
	if n <= 0 {
		panic("number should not be less than or equal to zero")
	}

	count := (n * n) + 2
	uf := UF.NewWQUPC(count)

	for i := 1; i <= n; i++ {
		uf.Union(0, i)
	}

	for i := n*(n-1) + 1; i <= n*n; i++ {
		uf.Union(n*n+1, i)
	}

	return &Percolation{
		ID:   make([]int, count),
		size: n,
		uf:   uf,
	}
}

func (p *Percolation) Encode(m, n int) int {
	return (m-1)*p.size + n
}

func (p *Percolation) IsInRange(n int) bool {
	return n > 0 && n <= p.size
}

func (p *Percolation) Open(m int, n int) {
	if !p.IsInRange(m) || !p.IsInRange(n) {
		panic("number out of range")
	}

	if p.IsOpen(m, n) {
		return
	}

	index := p.Encode(m, n)
	p.ID[index] = 1
	p.numberOfOpenSites += 1

	for _, arr := range [][2]int{{m + 1, n}, {m - 1, n}, {m, n + 1}, {m, n - 1}} {
		if !p.IsInRange(arr[0]) || !p.IsInRange(arr[1]) {
			continue
		}

		if !p.IsOpen(arr[0], arr[1]) {
			continue
		}

		sibling := p.Encode(arr[0], arr[1])
		if p.uf.Connected(index, sibling) {
			continue
		}

		p.uf.Union(index, sibling)
	}
}

func (p *Percolation) IsOpen(m int, n int) bool {
	index := p.Encode(m, n)
	return p.ID[index] == 1
}

func (p *Percolation) IsFull(int, int) bool {
	return false
}

func (p *Percolation) NumberOfOpenSites() int {
	return p.numberOfOpenSites
}

func (p *Percolation) Percolates() bool {
	return p.uf.Connected(0, p.size*p.size+1)
}

func (p *Percolation) Print() {
	p.uf.Print()
	fmt.Printf("Percolation.ID: %#v\n", p.ID)
}
