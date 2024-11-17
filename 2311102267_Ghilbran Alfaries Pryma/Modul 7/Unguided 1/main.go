package main

import (
	"fmt"
	"math"
)

// Tipe bentukan Titik untuk menyimpan koordinat x dan y
type Titik struct {
	x, y int
}

// Tipe bentukan Lingkaran untuk menyimpan titik pusat (cx, cy) dan radius r
type Lingkaran struct {
	cx, cy, r int
}

// Fungsi untuk menghitung jarak antara dua titik (x1, y1) dan (x2, y2)
func jarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

// Fungsi untuk menentukan apakah titik berada di dalam lingkaran
func titikDiDalamLingkaran(t Titik, l Lingkaran) bool {
	// Menghitung jarak antara titik dan pusat lingkaran
	jarakTitikKePusat := jarak(t.x, t.y, l.cx, l.cy)
	// Jika jarak titik ke pusat lingkaran lebih kecil atau sama dengan radius, maka titik berada di dalam lingkaran
	return jarakTitikKePusat <= float64(l.r)
}

func main() {
	var l1, l2 Lingkaran
	var t Titik

	// Membaca input untuk lingkaran 1 (cx, cy, r)
	fmt.Scan(&l1.cx, &l1.cy, &l1.r)

	// Membaca input untuk lingkaran 2 (cx, cy, r)
	fmt.Scan(&l2.cx, &l2.cy, &l2.r)

	// Membaca input untuk titik sembarang (x, y)
	fmt.Scan(&t.x, &t.y)

	// Mengecek apakah titik di dalam lingkaran 1 dan 2
	inLingkaran1 := titikDiDalamLingkaran(t, l1)
	inLingkaran2 := titikDiDalamLingkaran(t, l2)

	// Menentukan hasil berdasarkan posisi titik terhadap lingkaran
	if inLingkaran1 && inLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
