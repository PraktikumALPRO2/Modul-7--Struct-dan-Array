package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var winners []string

	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	for {
		fmt.Printf("Pertandingan %d: ", len(winners)+1)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			winners = append(winners, klubA)
			fmt.Printf("Hasil %d: %s\n", len(winners), klubA)
		} else if skorB > skorA {
			winners = append(winners, klubB)
			fmt.Printf("Hasil %d: %s\n", len(winners), klubB)
		} else {
			winners = append(winners, "Draw")
			fmt.Printf("Hasil %d: Draw\n", len(winners))
		}
	}

	fmt.Println("Pertandingan selesai")
}