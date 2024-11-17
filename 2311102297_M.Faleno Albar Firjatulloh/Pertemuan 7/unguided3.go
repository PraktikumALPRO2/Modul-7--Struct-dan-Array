package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var winners []string

	// Input nama klub
	fmt.Print("Klub A : ")
	scanner.Scan()
	clubA := scanner.Text()

	fmt.Print("Klub B : ")
	scanner.Scan()
	clubB := scanner.Text()

	// Proses input skor pertandingan
	matchNumber := 1
	for {
		fmt.Printf("Pertandingan %d : ", matchNumber)
		scanner.Scan()
		input := scanner.Text()

		// Split input menjadi dua skor
		scores := strings.Split(input, " ")
		if len(scores) != 2 {
			continue
		}

		// Convert string ke integer
		var scoreA, scoreB int
		_, err1 := fmt.Sscanf(scores[0], "%d", &scoreA)
		_, err2 := fmt.Sscanf(scores[1], "%d", &scoreB)

		// Cek apakah skor valid
		if err1 != nil || err2 != nil || scoreA < 0 || scoreB < 0 {
			fmt.Printf("\nHasil pertandingan:\n")
			for i, winner := range winners {
				fmt.Printf("Hasil %d : %s\n", i+1, winner)
			}
			fmt.Println("Pertandingan selesai")
			break
		}

		// Tentukan pemenang
		var winner string
		if scoreA > scoreB {
			winner = clubA
		} else if scoreB > scoreA {
			winner = clubB
		} else {
			winner = "Draw"
		}

		// Tambahkan pemenang ke array
		winners = append(winners, winner)
		matchNumber++
	}
}
