package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik
func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// Fungsi untuk mengecek apakah sebuah titik berada di dalam lingkaran
func diDalamLingkaran(cx, cy, radius, x, y float64) bool {
	return jarak(cx, cy, x, y) <= radius
}

func main() {
	// Koordinat pusat dan radius lingkaran 1
	cx1, cy1, r1 := 8.0, 8.0, 4.0
	// Koordinat pusat dan radius lingkaran 2
	cx2, cy2, r2 := 10.0, 15.0, 10.0

	// Daftar titik yang akan diuji
	titik := [][]float64{
		{8, 4},
		{10, 15},
		{-15, 4},
		{0, 0},
	}

	// Proses setiap titik
	for _, t := range titik {
		x, y := t[0], t[1]
		diLingkaran1 := diDalamLingkaran(cx1, cy1, r1, x, y)
		diLingkaran2 := diDalamLingkaran(cx2, cy2, r2, x, y)

		if diLingkaran1 && diLingkaran2 {
			fmt.Printf("Titik (%.0f, %.0f) di dalam lingkaran 1 dan 2\n", x, y)
		} else if diLingkaran1 {
			fmt.Printf("Titik (%.0f, %.0f) di dalam lingkaran 1\n", x, y)
		} else if diLingkaran2 {
			fmt.Printf("Titik (%.0f, %.0f) di dalam lingkaran 2\n", x, y)
		} else {
			fmt.Printf("Titik (%.0f, %.0f) di luar lingkaran 1 dan 2\n", x, y)
		}
	}
}
