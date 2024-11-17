package main

import (
	"fmt"
	"math"
)

type titik struct { 
	x, y int
}

type lingkaran struct {
	titikPusat titik
	radius     int
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(p, q titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

// Fungsi untuk memeriksa apakah titik p berada di dalam lingkaran c
func dalamL(c lingkaran, p titik) bool {
	return jarak(c.titikPusat, p) <= float64(c.radius)
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y int

	// Input lingkaran 1 (koordinat dan radius pada satu baris)
	fmt.Println("Masukkan koordinat x1, y1 radius  1 untuk lingkaran 1:")
	fmt.Scan(&cx1, &cy1, &r1)

	// Input lingkaran 2 (koordinat dan radius pada satu baris)
	fmt.Println("Masukkan koordinat x2, y2 radius 2 untuk lingkaran 2:")
	fmt.Scan(&cx2, &cy2, &r2)

	// Input titik sembarang (koordinat pada satu baris)
	fmt.Println("Masukkan koordinat x dan y untuk titik sembarang:")
	fmt.Scan(&x, &y)

	// Deklarasi lingkaran dan titik sembarang
	lingkaran1 := lingkaran{titik{cx1, cy1}, r1}
	lingkaran2 := lingkaran{titik{cx2, cy2}, r2}
	titikSembarang := titik{x, y}

	// Mengecek posisi titik terhadap lingkaran 1 dan lingkaran 2
	dalamL1 := dalamL(lingkaran1, titikSembarang)
	dalamL2 := dalamL(lingkaran2, titikSembarang)

	// Menampilkan hasil sesuai kondisi
	if dalamL1 && dalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
