package stats

type IPercolationStats interface {
	// sample mean of percolation threshold
	Mean() float64

	// sample standard deviation of percolation threshold
	Stddev() float64

	// low endpoint of 95% confidence interval
	ConfidenceLo() float64

	// high endpoint of 95% confidence interval
	ConfidenceHi() float64
}
