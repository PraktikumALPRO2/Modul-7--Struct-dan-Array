package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	// Input nama klub
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	const maxPertandingan = 9 // Batas maksimal pertandingan
	i := 1

	for i <= maxPertandingan {
		// Input skor pertandingan
		fmt.Printf("Pertandingan %d: ", i)
		_, err := fmt.Scan(&skorA, &skorB)

		// Validasi input
		if err != nil {
			fmt.Println("Input tidak valid. Silakan masukkan skor dalam format angka.")
			continue
		}

		// Kondisi untuk menghentikan jika skor negatif
		if skorA < 0 || skorB < 0 {
			fmt.Println("Skor negatif tidak diperbolehkan. Pertandingan dihentikan.")
			break
		}

		// Menentukan pemenang atau seri
		if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorA < skorB {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}

		i++
	}

	// Menampilkan hasil pertandingan
	fmt.Println("Pertandingan selesai")
	for j := 0; j < len(hasil); j++ {
		fmt.Printf("Hasil %d = %s\n", j+1, hasil[j])
	}
}
