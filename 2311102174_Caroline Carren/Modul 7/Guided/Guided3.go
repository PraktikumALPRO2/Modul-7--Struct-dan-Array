// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

func main() {
	// Membuat map dengan nama buah sebagai kunci dan harga sebagai nilai
	hargaBuah := map[string]int{
		"Apel":   5000, // Kunci "Apel", Nilai 5000
		"Pisang": 3000, // Kunci "Pisang", Nilai 3000
		"Mangga": 7000, // Kunci "Mangga", Nilai 7000
	}

	// Menampilkan harga dari setiap buah
	fmt.Println("Harga Buah:")
	// Melakukan iterasi untuk menampilkan kunci dan nilai dari map
	for buah, harga := range hargaBuah {
		// Menampilkan nama buah dan harganya
		fmt.Printf("%s: Rp%d\n", buah, harga)
	}

	// Menampilkan harga dari buah Mangga secara langsung menggunakan kunci
	fmt.Print("Harga buah Mangga = ", hargaBuah["Mangga"])
}
