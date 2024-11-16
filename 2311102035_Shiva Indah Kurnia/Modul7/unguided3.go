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
	fmt.Print("Klub A: ")
	scanner.Scan()
	clubA := scanner.Text()

	fmt.Print("Klub B: ")
	scanner.Scan()
	clubB := scanner.Text()

	// Proses input skor pertandingan
	for matchNumber := 1; ; matchNumber++ {
		fmt.Printf("Pertandingan %d: ", matchNumber)
		scanner.Scan()
		input := scanner.Text()

		scores := strings.Fields(input)
		if len(scores) != 2 {
			break
		}

		var scoreA, scoreB int
		if _, err1 := fmt.Sscanf(scores[0], "%d", &scoreA); err1 != nil || scoreA < 0 {
			break
		}
		if _, err2 := fmt.Sscanf(scores[1], "%d", &scoreB); err2 != nil || scoreB < 0 {
			break
		}

		var winner string
		if scoreA > scoreB {
			winner = clubA
		} else if scoreB > scoreA {
			winner = clubB
		} else {
			winner = "Draw"
		}

		winners = append(winners, winner)
	}

	fmt.Println("\nHasil pertandingan:")
	for i, winner := range winners {
		fmt.Printf("Hasil %d: %s\n", i+1, winner)
	}
	fmt.Println("Pertandingan selesai")
}