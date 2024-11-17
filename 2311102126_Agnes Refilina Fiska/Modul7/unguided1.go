package main

import (
	"fmt"
	"math"
)

// Point represents a 2D point with x,y coordinates
type Point struct {
	x, y float64
}

// Circle represents a circle with center point and radius
type Circle struct {
	center Point
	radius float64
}

// distance calculates Euclidean distance between two points
func distance(p1, p2 Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

// isPointInside checks if a point is inside a circle
func (c Circle) isPointInside(p Point) bool {
	return distance(c.center, p) <= c.radius
}

// checkPointPosition determines the position of a point relative to two circles
func checkPointPosition(c1, c2 Circle, p Point) string {
	in1 := c1.isPointInside(p)
	in2 := c2.isPointInside(p)

	if in1 && in2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if in1 {
		return "Titik di dalam lingkaran 1"
	} else if in2 {
		return "Titik di dalam lingkaran 2"
	}
	return "Titik di luar lingkaran 1 dan 2"
}

func main() {
	var T int
	fmt.Println("Masukkan jumlah test case:")
	fmt.Scan(&T)

	for t := 1; t <= T; t++ {
		// Read first circle coordinates and radius
		var x1, y1, r1 float64
		fmt.Scan(&x1, &y1, &r1)
		circle1 := Circle{center: Point{x1, y1}, radius: r1}

		// Read second circle coordinates and radius
		var x2, y2, r2 float64
		fmt.Scan(&x2, &y2, &r2)
		circle2 := Circle{center: Point{x2, y2}, radius: r2}

		// Read test point coordinates
		var px, py float64
		fmt.Scan(&px, &py)
		testPoint := Point{px, py}

		// Check and print result
		result := checkPointPosition(circle1, circle2, testPoint)
		fmt.Printf("Case #%d: %s\n", t, result)
	}
}
