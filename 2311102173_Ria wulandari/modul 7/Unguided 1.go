package main

import (
	"fmt"
	"math"
)

// Struct untuk menyimpan informasi lingkaran
type Lingkaran struct {
	x, y, radius float64
}

// Fungsi untuk menghitung jarak titik dari pusat lingkaran
func hitungJarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// Fungsi untuk mengecek apakah titik berada di dalam lingkaran
func diDalamLingkaran(l Lingkaran, x, y float64) bool {
	return hitungJarak(l.x, l.y, x, y) <= l.radius
}

func main() {
	// Input koordinat pusat dan radius dua lingkaran
	var lingkaran1, lingkaran2 Lingkaran
	fmt.Println("Masukkan titik pusat dan radius lingkaran 1 (cx cy r):")
	fmt.Scan(&lingkaran1.x, &lingkaran1.y, &lingkaran1.radius)

	fmt.Println("Masukkan titik pusat dan radius lingkaran 2 (cx cy r):")
	fmt.Scan(&lingkaran2.x, &lingkaran2.y, &lingkaran2.radius)

	// Input titik yang akan dicek
	var x, y float64
	fmt.Println("Masukkan koordinat titik (x y):")
	fmt.Scan(&x, &y)

	// Cek posisi titik terhadap lingkaran
	diLingkaran1 := diDalamLingkaran(lingkaran1, x, y)
	diLingkaran2 := diDalamLingkaran(lingkaran2, x, y)

	// Output hasil
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
