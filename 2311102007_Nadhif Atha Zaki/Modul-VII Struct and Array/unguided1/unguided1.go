package main

import (
	"fmt"
	"math"
)

// Struktur data untuk merepresentasikan titik dengan koordinat X dan Y
type Titik struct {
	X int
	Y int
}

// Struktur data untuk merepresentasikan lingkaran dengan pusat (Tengah) dan radius
type Lingkaran struct {
	Tengah Titik
	Radius int
}

// Fungsi untuk menghitung jarak antara dua titik menggunakan math
func hitungJarak(t1, t2 Titik) float64 {
	return math.Sqrt(math.Pow(float64(t1.X-t2.X), 2) + math.Pow(float64(t1.Y-t2.Y), 2))
}

// Fungsi untuk memeriksa posisi titik relatif terhadap dua lingkaran
func tentukanPosisiTitik(lingkaran1, lingkaran2 Lingkaran, titik Titik) string {
	// Menghitung jarak titik dari pusat lingkaran pertama
	jarakKeLingkaran1 := hitungJarak(lingkaran1.Tengah, titik)

	// Menghitung jarak titik dari pusat lingkaran kedua
	jarakKeLingkaran2 := hitungJarak(lingkaran2.Tengah, titik)

	// Menentukan apakah titik berada di dalam kedua lingkaran, salah satu lingkaran, atau di luar keduanya
	if jarakKeLingkaran1 <= float64(lingkaran1.Radius) && jarakKeLingkaran2 <= float64(lingkaran2.Radius) {
		return "Titik berada di dalam lingkaran 1 dan 2"
	} else if jarakKeLingkaran1 <= float64(lingkaran1.Radius) {
		return "Titik berada di dalam lingkaran 1"
	} else if jarakKeLingkaran2 <= float64(lingkaran2.Radius) {
		return "Titik berada di dalam lingkaran 2"
	} else {
		return "Titik berada di luar lingkaran 1 dan 2"
	}
}

func main() {
	// Deklarasi variabel untuk dua lingkaran dan satu titik
	var lingkaran1, lingkaran2 Lingkaran
	var titik Titik

	// Meminta input untuk pusat dan radius lingkaran pertama
	fmt.Println("Masukkan pusat dan radius lingkaran 1:")
	fmt.Scanln(&lingkaran1.Tengah.X, &lingkaran1.Tengah.Y, &lingkaran1.Radius)

	// Meminta input untuk pusat dan radius lingkaran kedua
	fmt.Println("Masukkan pusat dan radius lingkaran 2:")
	fmt.Scanln(&lingkaran2.Tengah.X, &lingkaran2.Tengah.Y, &lingkaran2.Radius)

	// Meminta input untuk koordinat titik
	fmt.Println("Masukkan koordinat titik sembarang:")
	fmt.Scanln(&titik.X, &titik.Y)

	// Menentukan posisi titik terhadap kedua lingkaran dan menampilkan hasilnya
	hasil := tentukanPosisiTitik(lingkaran1, lingkaran2, titik)
	fmt.Println("Posisi titik:", hasil)
}
