package percolation

type IPercolation interface {
	// opens the site (row, col) if it is not open already
	Open(int, int)

	// is the site (row, col) open?
	IsOpen(int, int) bool

	// is the site (row, col) full?
	IsFull(int, int) bool

	// returns the number of open sites
	NumberOfOpenSites() int

	// does the system percolate?
	Percolates() bool

	// print
	Print()
}
