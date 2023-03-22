package main

import (
	"fmt"

	"convex_hull/point"
	"elementary_sorts/shuffling"
	"stacks_queues/generic_stack"
)

func main() {
	slice := shuffling.Generate(30)
	var points []*point.Point2D

	for i := 1; i < len(slice); i += 2 {
		x, y := float64(slice[i-1]), float64(slice[i])
		newPoint := point.NewPoint2D(x, y)
		points = append(points, newPoint)
	}

	sortPointsByY(points)
	fmt.Printf("sortPointsByY: %v\n\n", points)

	sortPointsByPolarOrder(points)
	fmt.Printf("sortPointsByPolarOrder: %v\n\n", points)

	hull := generic_stack.NewLinkedStack[*point.Point2D]()
	// hull.Push(points[0])
	// hull.Push(points[1])

	// for i := 2; i < len(points); i++ {
	// 	top := hull.Pop()
	// 	for point.CCW(hull.Peek(), top, points[i]) <= 0 {
	// 		top = hull.Pop()
	// 	}
	// 	hull.Push(top)
	// 	hull.Push(points[i])
	// }

	fmt.Println(points)
	fmt.Println()
	fmt.Println(hull.String())
}

func sortPointsByPolarOrder(points []*point.Point2D) {
	for i := 1; i < len(points); i++ {
		for j := 2; j < len(points); j++ {
			if point.CCW(points[0], points[j-1], points[j]) > 0 {
				points[j], points[j-1] = points[j-1], points[j]
			}
		}
	}
}

func sortPointsByY(points []*point.Point2D) {
	for i := 0; i < len(points); i++ {
		for j := 1; j < len(points); j++ {
			if points[j-1].Y > points[j].Y {
				points[j], points[j-1] = points[j-1], points[j]
			}
		}
	}
}
