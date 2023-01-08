package point

import "fmt"

type Point2D struct {
	X, Y float64
}

func NewPoint2D(x, y float64) *Point2D {
	return &Point2D{
		X: x,
		Y: y,
	}
}

func (p *Point2D) String() string {
	return fmt.Sprintf("X: %.2f, Y: %.2f\n", p.X, p.Y)
}

func CCW(a, b, c *Point2D) float64 {
	area := ((b.X - a.X) * (c.Y - a.Y)) - ((b.Y - a.Y) * (c.X - a.X))
	if area < 0 {
		return -1 // clockwise
	} else if area > 0 {
		return 1 // counter clockwise
	} else {
		return 0 // collinear
	}
}
