package stats

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"percolation/percolation"
)

type PercolationStats struct {
	count        float64
	sum          float64
	mean         float64
	stddev       float64
	confidenceLo float64
	confidenceHi float64
}

// perform independent trials on an n-by-n grid
func New(n int, trials int) *PercolationStats {
	if n <= 0 || trials <= 0 {
		panic("number or trials should not be less than or equal to zero")
	}

	ratios := make([]float64, 0, trials)

	for i := 0; i < trials; i++ {
		var p percolation.IPercolation = percolation.New(n)

		rand.Seed(time.Now().UnixNano())
		percolates := false

		for !percolates {
			x, y := rand.Intn(n)+1, rand.Intn(n)+1
			p.Open(x, y)
			percolates = p.Percolates()
			// fmt.Printf("M: %02d, N: %02d, Percolates?: %t\n", x, y, percolates)
		}

		numberOfOpenSites := float64(p.NumberOfOpenSites())
		totalSize := float64(n * n)
		percolationRatio := numberOfOpenSites / totalSize
		ratios = append(ratios, percolationRatio)

		fmt.Printf("NumberOfOpenSites: %.2f, Size: %.2f, Percolation ratio: %.2f\n", numberOfOpenSites, totalSize, numberOfOpenSites/totalSize)
	}

	s := &PercolationStats{}
	s.count = float64(trials)
	s.sum = s.CalculateSum(ratios)
	s.mean = s.CalculateMean(s.sum, s.count)
	s.stddev = s.CalculateStddev(ratios, s.count, s.mean)
	s.confidenceLo, s.confidenceHi = s.CalculateConfidenceInterval(s.stddev, s.mean, s.count)
	return s
}

func (s *PercolationStats) CalculateSum(ratios []float64) float64 {
	var sum float64
	for _, r := range ratios {
		sum += r
	}
	return sum
}

func (s *PercolationStats) Sum() float64 {
	return s.sum
}

func (s *PercolationStats) CalculateMean(sum float64, count float64) float64 {
	return sum / count
}

func (s *PercolationStats) Mean() float64 {
	return s.mean
}

func (s *PercolationStats) CalculateStddev(
	ratios []float64,
	count float64,
	mean float64,
) float64 {
	var numerator float64
	for _, r := range ratios {
		numerator += math.Pow(r-mean, 2)
	}
	denominator := count - 1
	variance := numerator / denominator
	return math.Sqrt(variance)
}

func (s *PercolationStats) Stddev() float64 {
	return s.stddev
}

func (s *PercolationStats) CalculateConfidenceInterval(
	stddev, mean, count float64) (float64, float64) {
	// The sampling mean most likely follows a normal distribution. In this case, the standard error of the mean (SEM) can be calculated using the following equation:
	sem := stddev / math.Sqrt(count)
	ci := 1.960 * sem // 1.960 is for 95% CI
	confidenceLo := mean - ci
	confidenceHi := mean + ci

	return confidenceLo, confidenceHi
}

func (s *PercolationStats) ConfidenceLo() float64 {
	return s.confidenceLo
}

func (s *PercolationStats) ConfidenceHi() float64 {
	return s.confidenceHi
}
