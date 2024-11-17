package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	fmt.Print("Klub A : ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scanln(&klubB)

	for i := 1; i <= 9; i++ {
		fmt.Printf("Pertandingan %d : ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}
		var result string
		if skorA > skorB {
			result = klubA
		} else if skorB > skorA {
			result = klubB
		} else {
			result = "Draw"
		}

		hasil = append(hasil, fmt.Sprintf("Pertandingan %d: %s", i, result))
	}

	fmt.Println("\nHasil Pertandingan:")
	for _, result := range hasil {
		fmt.Println(result)
	}
	fmt.Println("Pertandingan selesai")
}
