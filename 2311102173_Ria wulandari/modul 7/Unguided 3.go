package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var menangA, menangB, seri int

	// Meminta input nama klub
	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	// Meminta input skor dari pengguna hingga 9 kali atau skor negatif diberikan
	for i := 1; i <= 9; i++ {
		fmt.Printf("Pertandingan %d (masukkan skorA skorB): ", i)
		_, err := fmt.Scan(&skorA, &skorB)

		// Menangani error input
		if err != nil {
			fmt.Println("Input tidak valid, coba lagi.")
			i-- // Mengulangi pertandingan yang sama jika input tidak valid
			continue
		}

		// Cek jika ada skor negatif untuk berhenti
		if skorA < 0 || skorB < 0 {
			fmt.Println("Input negatif terdeteksi, pertandingan selesai.")
			break
		}

		// Menentukan pemenang dari pertandingan
		if skorA > skorB {
			fmt.Printf("Hasil %d: %s menang\n", i, klubA)
			menangA++
		} else if skorB > skorA {
			fmt.Printf("Hasil %d: %s menang\n", i, klubB)
			menangB++
		} else {
			fmt.Printf("Hasil %d: Draw\n", i)
			seri++
		}
	}

	// Menampilkan hasil akhir
	fmt.Println("\nHasil Akhir:")
	fmt.Printf("%s menang: %d kali\n", klubA, menangA)
	fmt.Printf("%s menang: %d kali\n", klubB, menangB)
	fmt.Printf("Seri: %d kali\n", seri)
}
