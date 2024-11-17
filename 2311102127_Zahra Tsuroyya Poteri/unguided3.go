package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasilPertandingan []string

	// Meminta input nama klub
	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	// Proses input pertandingan
	for i := 1; i <= 9; i++ {
		// Input skor pertandingan
		fmt.Printf("Pertandingan %d: ", i)
		_, err := fmt.Scanln(&skorA, &skorB)

		// Cek jika input tidak valid atau skor negatif
		if err != nil || skorA < 0 || skorB < 0 {
			// Hanya berhenti jika input tidak valid
			break
		}

		// Menentukan pemenang
		if skorA > skorB {
			hasilPertandingan = append(hasilPertandingan, klubA)
		} else if skorB > skorA {
			hasilPertandingan = append(hasilPertandingan, klubB)
		} else {
			hasilPertandingan = append(hasilPertandingan, "Draw")
		}
	}

	// Menampilkan hasil pertandingan yang valid
	for i, hasil := range hasilPertandingan {
		fmt.Printf("Hasil %d: %s\n", i+1, hasil)
	}

	// Menampilkan pesan bahwa pertandingan telah selesai
	fmt.Println("Pertandingan selesai")
}
