package main

import (
	"fmt"
	"math"
)

type titik153 struct {
	x, y int
}
type o153 struct {
	center titik153
	radius int
}

func jarak(p1, p2 titik153) float64 {
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)))
}
func dalamo(o o153, point titik153) bool {
	return jarak(o.center, point) <= float64(o.radius)
}

func main() {
	var o1, o2 o153
	var point titik153

	fmt.Print("INPUT KOORDINAT PUSAT DAN RADIUS O1 (cx cy r): ")
	fmt.Scan(&o1.center.x, &o1.center.y, &o1.radius)

	fmt.Print("INPUT KOORDINAT PUSAT DAN RADIUS O2 (cx cy r): ")
	fmt.Scan(&o2.center.x, &o2.center.y, &o2.radius)

	fmt.Print("INPUT KOORDINAT SEMBARANG (x y): ")
	fmt.Scan(&point.x, &point.y)

	diO1 := dalamo(o1, point)
	diO2 := dalamo(o2, point)

	if diO1 && diO2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diO1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diO2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
