package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	titikPusat Titik
	radius int
}

// Fungsi untuk menghitung jarak antara dua titik
func jarakAntarTitik(a, b Titik) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2))
}

// Fungsi untuk mengecek apakah titik berada di dalam lingkaran
func cekTitikDiDalam(lingkaran Lingkaran, titik Titik) bool {
	return jarakAntarTitik(lingkaran.titikPusat, titik) <= float64(lingkaran.radius)
}

func main() {
	var lingkaran1, lingkaran2 Lingkaran
	var titikSembarang Titik

	// Input titik pusat dan radius lingkaran 1
	fmt.Print("Masukkan titik pusat dan radius lingkaran 1: ")
	fmt.Scan(&lingkaran1.titikPusat.x, &lingkaran1.titikPusat.y, &lingkaran1.radius)


	// Input titik pusat dan radius lingkaran 2
	fmt.Print("Masukkan titik pusat dan radius lingkaran 2: ")
	fmt.Scan(&lingkaran2.titikPusat.x, &lingkaran2.titikPusat.y, &lingkaran2.radius)

	// Input titik sembarang
	fmt.Print("Masukkan titik sembarang: ")
	fmt.Scan(&titikSembarang.x, &titikSembarang.y)

	// Mengecek posisi titik terhadap lingkaran 1 dan 2
	diDalamLingkaran1 := cekTitikDiDalam(lingkaran1, titikSembarang)
	diDalamLingkaran2 := cekTitikDiDalam(lingkaran2, titikSembarang)

	// Menentukan output berdasarkan posisi titik
	if diDalamLingkaran1 && diDalamLingkaran2 {
		fmt.Println("\nTitik di dalam lingkaran 1 dan 2")
	} else if diDalamLingkaran1 {
		fmt.Println("\nTitik di dalam lingkaran 1")
	} else if diDalamLingkaran2 {
		fmt.Println("\nTitik di dalam lingkaran 2")
	} else {
		fmt.Println("\nTitik di luar lingkaran 1 dan 2")
	}
}
