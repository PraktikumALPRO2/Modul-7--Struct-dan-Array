//Haifa Zahra Azzimmi
//2311102163

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var winningClubs []string // Slice untuk menyimpan nama klub yang menang
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Rekap Skor Pertandingan Bola")
	for {
		fmt.Print("Masukkan nama klub pertama: ")
		club1, _ := reader.ReadString('\n')
		club1 = strings.TrimSpace(club1)

		fmt.Print("Masukkan nama klub kedua: ")
		club2, _ := reader.ReadString('\n')
		club2 = strings.TrimSpace(club2)

		fmt.Print("Masukkan skor klub pertama: ")
		score1Str, _ := reader.ReadString('\n')
		score1Str = strings.TrimSpace(score1Str)
		score1, err1 := strconv.Atoi(score1Str)
		if err1 != nil || score1 < 0 {
			fmt.Println("Skor tidak valid! Program berhenti.")
			break
		}

		fmt.Print("Masukkan skor klub kedua: ")
		score2Str, _ := reader.ReadString('\n')
		score2Str = strings.TrimSpace(score2Str)
		score2, err2 := strconv.Atoi(score2Str)
		if err2 != nil || score2 < 0 {
			fmt.Println("Skor tidak valid! Program berhenti.")
			break
		}

		if score1 > score2 {
			winningClubs = append(winningClubs, club1)
			fmt.Printf("Pemenang: %s\n", club1)
		} else if score2 > score1 {
			winningClubs = append(winningClubs, club2)
			fmt.Printf("Pemenang: %s\n", club2)
		} else {
			fmt.Println("Pertandingan seri, tidak ada pemenang.")
		}
		fmt.Println()
	}

	fmt.Println("\nDaftar Klub yang Menang:")
	for i, club := range winningClubs {
		fmt.Printf("%d. %s\n", i+1, club)
	}
}
