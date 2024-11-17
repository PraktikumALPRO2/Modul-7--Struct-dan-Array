/* Liya Khoirunnisa - 2311102124 */
package main

import (
	"fmt"
)

// Struct untuk menyimpan informasi pertandingan
type pertandingan struct {
	klubA  string // Nama klub A
	klubB  string // Nama klub B
	skorA  int    // Skor klub A
	skorB  int    // Skor klub B
	hasil  string // Hasil pertandingan
}

func main() {
	// Deklarasi variabel
	var klubA, klubB string
	var listPertandingan []pertandingan // Array untuk menyimpan hasil pertandingan

	// Input nama klub A
	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)

	// Input nama klub B
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	fmt.Println()

	// Proses input skor hingga skor negatif
	pertandinganKe := 1 // Variabel untuk menghitung pertandingan ke berapa
	for {
		var skorA, skorB int
		fmt.Printf("Pertandingan %d: ", pertandinganKe)
		fmt.Scan(&skorA, &skorB)

		// Cek validitas skor
		if skorA < 0 || skorB < 0 {
			break 
		}

		// Menentukan pemenang
		var hasil string
		if skorA > skorB {
			hasil = klubA
		} else if skorB > skorA {
			hasil = klubB
		} else {
			hasil = "Draw"
		}

		// Simpan hasil pertandingan dalam array
		listPertandingan = append(listPertandingan, pertandingan{
			klubA: klubA,
			klubB: klubB,
			skorA: skorA,
			skorB: skorB,
			hasil: hasil,
		})

		pertandinganKe++ // Increment untuk pertandingan berikutnya
	}

	// Menampilkan hasil pertandingan
	fmt.Println("\nHasil Pertandingan:")
	for i, p := range listPertandingan {
		fmt.Printf("Hasil %d: %s\n", i+1, p.hasil)
	}

	// Menampilkan pesan bahwa pertandingan selesai
	fmt.Println("Pertandingan selesai")
}