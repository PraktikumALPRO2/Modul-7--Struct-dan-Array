package main

import "fmt"

func main() {
	var klubA, klubB string
	var menang []string

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	matchNumber := 1

	for {
		var skorA, skorB int

		fmt.Printf("Pertandingan %d : ", matchNumber)
		_, err := fmt.Scan(&skorA, &skorB)

		if err != nil {
			fmt.Println("GAVALID PISAHKAN 2 ANGKA DENGAN SPACE!! ")
			continue
		}
		if skorA < 0 || skorB < 0 {
			break
		}
		if skorA > skorB {
			menang = append(menang, klubA)
		} else if skorB > skorA {
			menang = append(menang, klubB)
		} else {
			menang = append(menang, "Draw")
		}
		matchNumber++
	}

	for i, winner := range menang {
		fmt.Printf("Pertandingan %d: %s\n", i+1, winner)
	}
	fmt.Println("Pertandingan selesai")
}