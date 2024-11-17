package main

import (
	"fmt"
	"strings"
)

type Pertandingan_145 struct {
	clubA  string
	clubB  string
	scoreA int
	scoreB int
}

func main() {
	var clubA, clubB string
	var matches []Pertandingan_145
	var winners []string

	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scanln(&clubA)
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scanln(&clubB)

	round := 1
	for {
		var scoreA, scoreB int
		fmt.Printf("Pertandingan %d -  skor: ", round)
		fmt.Scan(&scoreA, &scoreB)

		if scoreA < 0 || scoreB < 0 {
			fmt.Println("Input skor selesai.")
			break
		}

		matches = append(matches, Pertandingan_145{clubA, clubB, scoreA, scoreB})

		if scoreA > scoreB {
			winners = append(winners, clubA)
		} else if scoreB > scoreA {
			winners = append(winners, clubB)
		}

		round++
	}

	fmt.Println("\nHasil Pertandingan:")
	for i, match := range matches {
		fmt.Printf("Pertandingan %d: %s %d - %d %s\n", i+1, match.clubA, match.scoreA, match.scoreB, match.clubB)
	}

	fmt.Println("\nDaftar Klub yang Menang:")
	for _, winner := range winners {
		fmt.Println(winner)
	}

	countA, countB := 0, 0
	for _, winner := range winners {
		if strings.EqualFold(winner, clubA) {
			countA++
		} else if strings.EqualFold(winner, clubB) {
			countB++
		}
	}

	fmt.Printf("\nJumlah kemenangan:\n%s: %d\n%s: %d\n", clubA, countA, clubB, countB)
}
