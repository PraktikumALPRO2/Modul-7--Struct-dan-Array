package main

import (
	"fmt"
)

type Klub struct {
	nama string
}

type Pertandingan struct {
	skorA int
	skorB int
}

func main() {
	var klubA, klubB Klub
	var pertandingan []Pertandingan
	var pemenang []string

	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scanln(&klubA.nama)
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scanln(&klubB.nama)

	fmt.Println("\nMasukkan skor pertandingan (format: skorA skorB)")

	for {
		var skorA, skorB int
		fmt.Print("Skor: ")
		fmt.Scanf("%d %d", &skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		pertandingan = append(pertandingan, Pertandingan{skorA, skorB})

		if skorA > skorB {
			pemenang = append(pemenang, klubA.nama)
		} else if skorA < skorB {
			pemenang = append(pemenang, klubB.nama)
		} else {
			pemenang = append(pemenang, "Draw")
		}
	}

	fmt.Println("\nHasil pertandingan:")
	for i, p := range pertandingan {
		fmt.Printf("Pertandingan %d: %s %d - %d %s\n", i+1, klubA.nama, p.skorA, p.skorB, klubB.nama)
	}

	fmt.Println("\nDaftar pemenang:")
	for i, p := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, p)
	}
}
