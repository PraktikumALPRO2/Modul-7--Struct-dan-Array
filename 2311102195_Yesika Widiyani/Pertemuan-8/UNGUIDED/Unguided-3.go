package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string

	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scanln(&klubB)

	fmt.Println("\nMasukkan skor pertandingan:")
	for i := 1; ; i++ {
		fmt.Printf("Pertandingan %d (Skor %s - %s): ", i, klubA, klubB)
		fmt.Scan(&skorA, &skorB)

		// Berhenti jika skor tidak valid
		if skorA < 0 || skorB < 0 {
			break
		}

		// Menentukan pemenang
		if skorA > skorB {
			fmt.Println("Hasil", i, ":", klubA)
			pemenang = append(pemenang, klubA)
		} else if skorA < skorB {
			fmt.Println("Hasil", i, ":", klubB)
			pemenang = append(pemenang, klubB)
		} else {
			fmt.Println("Hasil", i, ": Draw")
			pemenang = append(pemenang, "Draw")
		}
	}

	// Menampilkan daftar pemenang
	fmt.Println("\nPertandingan selesai.")
	fmt.Println("Daftar hasil pertandingan:")
	for i, hasil := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, hasil)
	}
}