/* Liya Khoirunnisa - 2311102124 */
package main

import (
	"fmt"
	"math" // Untuk operasi matematika
)

// Struct untuk Lingkaran
type Lingkaran struct {
	cx, cy, radius int
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

// Fungsi untuk menentukan apakah titik berada di dalam lingkaran
func cekDalamLingkaran(circle Lingkaran, x, y int) bool {
	return jarak(circle.cx, circle.cy, x, y) <= float64(circle.radius)
}

func main() {
	// Deklarasi variabel untuk lingkaran 1 dan 2
	var lingkaran1, lingkaran2 Lingkaran
	var x, y int

	// Input lingkaran 1
	fmt.Println("Lingkaran 1")
	fmt.Print("Masukkan titik pusat (cx1, cy1) dan radius r1: ")
	fmt.Scan(&lingkaran1.cx, &lingkaran1.cy, &lingkaran1.radius)

	// Input lingkaran 2
	fmt.Println("\nLingkaran 2")
	fmt.Print("Masukkan titik pusat (cx2, cy2) dan radius r2: ")
	fmt.Scan(&lingkaran2.cx, &lingkaran2.cy, &lingkaran2.radius)

	// Input titik sembarang
	fmt.Print("\nMasukkan titik sembarang (x, y): ")
	fmt.Scan(&x, &y)

	// Mengecek letak titik di dalam lingkaran
	titikDalam1 := cekDalamLingkaran(lingkaran1, x, y)
	titikDalam2 := cekDalamLingkaran(lingkaran2, x, y)

	// Menentukan output berdasarkan posisi titik
	fmt.Print("\nPosisi titik: ")
	if titikDalam1 && titikDalam2 {
		// Jika titik berada di dalam kedua lingkaran
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if titikDalam1 {
		// Jika titik hanya berada di dalam lingkaran 1
		fmt.Println("Titik di dalam lingkaran 1")
	} else if titikDalam2 {
		// Jika titik hanya berada di dalam lingkaran 2
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		// Jika titik berada di luar kedua lingkaran
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}