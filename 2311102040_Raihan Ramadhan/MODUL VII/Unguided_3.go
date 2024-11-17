package main

import (
	"fmt"
)

func main() {
	// Input nama klub
	var clubA, clubB string
	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scanln(&clubA)

	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scanln(&clubB)

	var winners []string // Array untuk menyimpan pemenang

	// Input skor pertandingan
	for i := 1; ; i++ {
		var scoreA, scoreB int
		fmt.Printf("Masukkan skor pertandingan %d (%s vs %s):\n", i, clubA, clubB)

		fmt.Printf("Skor %s: ", clubA)
		fmt.Scanln(&scoreA)

		fmt.Printf("Skor %s: ", clubB)
		fmt.Scanln(&scoreB)

		// Cek jika skor tidak valid (negatif)
		if scoreA < 0 || scoreB < 0 {
			fmt.Println("Skor tidak valid! Program selesai.")
			break
		}

		// Tentukan pemenang
		if scoreA > scoreB {
			winners = append(winners, clubA)
			fmt.Printf("Hasil pertandingan %d: %s menang\n", i, clubA)
		} else if scoreA < scoreB {
			winners = append(winners, clubB)
			fmt.Printf("Hasil pertandingan %d: %s menang\n", i, clubB)
		} else {
			winners = append(winners, "Draw")
			fmt.Printf("Hasil pertandingan %d: Seri (Draw)\n", i)
		}
	}

	// Tampilkan daftar pemenang
	fmt.Println("\nDaftar pemenang pertandingan:")
	for _, winner := range winners {
		if winner != "Draw" {
			fmt.Println(winner)
		}
	}
}
