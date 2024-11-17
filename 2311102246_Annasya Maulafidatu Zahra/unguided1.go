package main

import (
	"fmt"
	"math"
)

// Struct untuk menyimpan koordinat titik
type Point struct {
	x float64
	y float64
}

// Struct untuk menyimpan informasi lingkaran
type Circle struct {
	center Point
	radius float64
}

// Fungsi untuk menentukan apakah titik berada di dalam lingkaran
func (c Circle) contains(p Point) bool {
	distance := math.Sqrt(math.Pow(p.x-c.center.x, 2) + math.Pow(p.y-c.center.y, 2))
	return distance < c.radius
}

func main() {
	var circle1, circle2 Circle
	var point Point

	// Input untuk lingkaran 1
	fmt.Scanf("%f %f %f", &circle1.center.x, &circle1.center.y, &circle1.radius)

	// Input untuk lingkaran 2
	fmt.Scanf("%f %f %f", &circle2.center.x, &circle2.center.y, &circle2.radius)

	// Input untuk titik sembarang
	fmt.Scanf("%f %f", &point.x, &point.y)

	// Menentukan posisi titik
	inCircle1 := circle1.contains(point)
	inCircle2 := circle2.contains(point)

	if inCircle1 && inCircle2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inCircle1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inCircle2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
