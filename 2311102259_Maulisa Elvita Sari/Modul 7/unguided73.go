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

	fmt.Println("\nMasukkan skor pertandingan (masukkan skor negatif untuk menghentikan):")
	for i := 1; ; i++ {
		fmt.Printf("Pertandingan %d - Skor %s: ", i, klubA)
		fmt.Scan(&skorA)
		fmt.Printf("Pertandingan %d - Skor %s: ", i, klubB)
		fmt.Scan(&skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}
	}

	fmt.Println("\nHasil pertandingan:")
	for i, hasil := range pemenang {
		fmt.Printf("Pertandingan %d: %s\n", i+1, hasil)
	}
}
