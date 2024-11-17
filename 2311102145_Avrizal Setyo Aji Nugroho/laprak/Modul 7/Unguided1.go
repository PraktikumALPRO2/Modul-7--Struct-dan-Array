package main

import (
	"fmt"
	"math"
)

type Circle_145 struct {
	x, y, radius float64
}

type Poin_145 struct {
	x, y float64
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

// Fungsi untuk mengecek apakah suatu titik berada di dalam lingkaran
func didalam(c Circle_145, p Poin_145) bool {
	return jarak(c.x, c.y, p.x, p.y) <= c.radius
}

func main() {
	var circle1, circle2 Circle_145
	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 1 (cx1, cy1, r1):")
	fmt.Scan(&circle1.x, &circle1.y, &circle1.radius)

	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 2 (cx2, cy2, r2):")
	fmt.Scan(&circle2.x, &circle2.y, &circle2.radius)

	var point Poin_145
	fmt.Println("Masukkan koordinat titik (x, y):")
	fmt.Scan(&point.x, &point.y)

	inCircle1 := didalam(circle1, point)
	inCircle2 := didalam(circle2, point)

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
