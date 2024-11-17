package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y float64
}

type Lingkaran struct {
	pusat Titik
	radius float64
}

func jarak(p1, p2 Titik) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func diDalamLingkaran(l Lingkaran, t Titik) bool {
	return jarak(l.pusat, t) <= l.radius
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64

	fmt.Println("Masukkan data lingkaran 1 (pusat x, pusat y, radius):")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Println("Masukkan data lingkaran 2 (pusat x, pusat y, radius):")
	fmt.Scan(&cx2, &cy2, &r2)

	lingkaran1 := Lingkaran{pusat: Titik{x: cx1, y: cy1}, radius: r1}
	lingkaran2 := Lingkaran{pusat: Titik{x: cx2, y: cy2}, radius: r2}

	var tx, ty float64
	fmt.Println("Masukkan koordinat titik (x, y):")
	fmt.Scan(&tx, &ty)
	titik := Titik{x: tx, y: ty}

	diLingkaran1 := diDalamLingkaran(lingkaran1, titik)
	diLingkaran2 := diDalamLingkaran(lingkaran2, titik)

	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
