package main

import (
	"fmt"
)

type Match struct {
	clubAScore int
	clubBScore int
}

func determineWinner(clubA, clubB string, match Match) string {
	if match.clubAScore < 0 || match.clubBScore < 0 {
		return "Invalid"
	}
	if match.clubAScore > match.clubBScore {
		return clubA
	} else if match.clubBScore > match.clubAScore {
		return clubB
	}
	return "Draw"
}

func main() {
	var clubA, clubB string

	// Input klub names
	fmt.Print("Klub A : ")
	fmt.Scan(&clubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&clubB)

	matches := make([]Match, 9)  // Array untuk 9 pertandingan
	winners := make([]string, 9) // Array untuk menyimpan pemenang

	// Input skor pertandingan
	for i := 0; i < 9; i++ {
		fmt.Printf("Pertandingan %d : ", i+1)
		var scoreA, scoreB int
		fmt.Scan(&scoreA, &scoreB)

		// Jika skor negatif, hentikan input
		if scoreA < 0 || scoreB < 0 {
			fmt.Println("Skor tidak valid (negatif). Program berhenti.")
			break
		}

		matches[i] = Match{clubAScore: scoreA, clubBScore: scoreB}
		winners[i] = determineWinner(clubA, clubB, matches[i])

		// Tampilkan komentar untuk pertandingan pertama
		if i == 0 {
			fmt.Printf("// %s = %d sedangkan %s = %d\n",
				clubA, scoreA, clubB, scoreB)
		}
	}

	// Tampilkan hasil pertandingan
	fmt.Println("\nHasil pertandingan:")
	for i := 0; i < len(winners); i++ {
		if winners[i] == "" {
			continue
		}
		fmt.Printf("Hasil %d : %s\n", i+1, winners[i])
	}

	fmt.Println("Pertandingan selesai")
}
