package main

import (
	"fmt"
	"math"
)

// Definisikan struktur untuk titik dengan koordinat x dan y
type Titik struct {
	x, y float64
}

// Definisikan struktur untuk lingkaran dengan titik pusat dan radius
type Lingkaran struct {
	pusat  Titik
	radius float64
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(t1, t2 Titik) float64 {
	return math.Sqrt((t1.x-t2.x)*(t1.x-t2.x) + (t1.y-t2.y)*(t1.y-t2.y))
}

// Fungsi untuk mengecek apakah titik berada di dalam lingkaran
func diDalamLingkaran(t Titik, l Lingkaran) bool {
	return jarak(t, l.pusat) <= l.radius
}

// Fungsi untuk menentukan posisi titik terhadap dua lingkaran
func cekPosisiTitik(t Titik, l1, l2 Lingkaran) string {
	diLingkaran1 := diDalamLingkaran(t, l1)
	diLingkaran2 := diDalamLingkaran(t, l2)

	switch {
	case diLingkaran1 && diLingkaran2:
		return "Titik di dalam lingkaran 1 dan 2"
	case diLingkaran1:
		return "Titik di dalam lingkaran 1"
	case diLingkaran2:
		return "Titik di dalam lingkaran 2"
	default:
		return "Titik di luar lingkaran 1 dan 2"
	}
}

func main() {
	var lingkaran1, lingkaran2 Lingkaran
	var titik Titik

	// Input baris pertama untuk lingkaran 1
	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 1 (x y radius):")
	fmt.Scan(&lingkaran1.pusat.x, &lingkaran1.pusat.y, &lingkaran1.radius)

	// Input baris kedua untuk lingkaran 2
	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 2 (x y radius):")
	fmt.Scan(&lingkaran2.pusat.x, &lingkaran2.pusat.y, &lingkaran2.radius)

	// Input baris ketiga untuk titik yang akan dicek
	fmt.Println("Masukkan koordinat titik yang akan dicek (x y):")
	fmt.Scan(&titik.x, &titik.y)

	// Menentukan posisi titik
	hasil := cekPosisiTitik(titik, lingkaran1, lingkaran2)
	fmt.Println(hasil)
}
