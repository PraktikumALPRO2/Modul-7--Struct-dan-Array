package main

import (
	"fmt"
)

func main() {
	var KlubA, KlubB string
	fmt.Print("Masukkan nama Tim A: ")
	fmt.Scanln(&KlubA)
	fmt.Print("Masukkan nama Tim B: ")
	fmt.Scanln(&KlubB)

	rekapHasilPertandingan := []string{}
	jumlahPertandingan := 1

	for {
		var skorKlubA, skorKlubB int
		fmt.Printf("\nPertandingan %d - Masukkan skor %s: ", jumlahPertandingan, KlubA)
		fmt.Scan(&skorKlubA)
		fmt.Printf("Pertandingan %d - Masukkan skor %s: ", jumlahPertandingan, KlubB)
		fmt.Scan(&skorKlubB)

		// Cek apakah skor valid (tidak negatif)
		if skorKlubA < 0 || skorKlubB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		// Menentukan hasil pertandingan
		if skorKlubA > skorKlubB {
			fmt.Printf("\nHasil %d: %s menang\n", jumlahPertandingan, KlubA)
			rekapHasilPertandingan = append(rekapHasilPertandingan, KlubA)
		} else if skorKlubB > skorKlubA {
			fmt.Printf("\nHasil %d: %s menang\n", jumlahPertandingan, KlubB)
			rekapHasilPertandingan = append(rekapHasilPertandingan, KlubB)
		} else {
			fmt.Printf("\nHasil %d: Draw\n", jumlahPertandingan)
			rekapHasilPertandingan = append(rekapHasilPertandingan, "Draw")
		}

		jumlahPertandingan++
	}

	fmt.Println("\nDaftar tim yang memenangkan pertandingan:")
	for i, hasil := range rekapHasilPertandingan {
    fmt.Printf("Hasil %d: %s\n", i+1, hasil)
	}

}
