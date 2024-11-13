// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import (
	"fmt"
	"math"
)

// Mendefinisikan tipe data struct 'Point' untuk merepresentasikan koordinat (X, Y) dari suatu titik
type Point struct {
	X int
	Y int
}

// Mendefinisikan tipe data struct 'Circle' untuk merepresentasikan lingkaran dengan pusat (Center) dan radius
type Circle struct {
	Center Point
	Radius int
}

// Fungsi untuk menghitung jarak antara dua titik (p1 dan p2) menggunakan rumus jarak Euclidean
func distance(p1, p2 Point) float64 {
	return math.Sqrt(float64((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y)))
}

// Fungsi untuk menentukan posisi titik relatif terhadap dua lingkaran (circle1 dan circle2)
func checkPointPosition(circle1, circle2 Circle, point Point) string {
	// Menghitung jarak antara titik dan pusat lingkaran pertama
	dist1 := distance(circle1.Center, point)

	// Menghitung jarak antara titik dan pusat lingkaran kedua
	dist2 := distance(circle2.Center, point)

	// Memeriksa apakah titik berada di dalam kedua lingkaran, atau salah satu lingkaran, atau di luar keduanya
	if dist1 <= float64(circle1.Radius) && dist2 <= float64(circle2.Radius) {
		return "Titik di dalam lingkaran 1 dan 2" // Jika titik berada di dalam lingkaran 1 dan 2
	} else if dist1 <= float64(circle1.Radius) {
		return "Titik di dalam lingkaran 1" // Jika titik hanya di dalam lingkaran 1
	} else if dist2 <= float64(circle2.Radius) {
		return "Titik di dalam lingkaran 2" // Jika titik hanya di dalam lingkaran 2
	} else {
		return "Titik di luar lingkaran 1 dan 2" // Jika titik berada di luar kedua lingkaran
	}
}

func main() {
	// Deklarasi dua variabel lingkaran (circle1 dan circle2) dan satu variabel titik (point)
	var circle1, circle2 Circle
	var point Point

	// Mengambil input untuk koordinat pusat dan radius lingkaran 1 dari pengguna
	fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 1:")
	fmt.Scanln(&circle1.Center.X, &circle1.Center.Y, &circle1.Radius)

	// Mengambil input untuk koordinat pusat dan radius lingkaran 2 dari pengguna
	fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 2:")
	fmt.Scanln(&circle2.Center.X, &circle2.Center.Y, &circle2.Radius)

	// Mengambil input untuk koordinat titik sembarang dari pengguna
	fmt.Println("Masukkan koordinat titik sembarang:")
	fmt.Scanln(&point.X, &point.Y)

	// Memeriksa posisi titik relatif terhadap kedua lingkaran dan menampilkan hasilnya
	position := checkPointPosition(circle1, circle2, point)
	fmt.Println("Posisi titik:", position)
}
