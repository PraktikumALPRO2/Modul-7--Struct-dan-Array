package main
import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	i := 1
	for {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorA < skorB {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Seri") 
		}
		i++
	}

	fmt.Println("Daftar klub yang memenangkan pertandingan:")
	for j := 0; j < len(hasil); j++ {
		fmt.Printf("Pertandingan %d: %s\n", j+1, hasil[j])
	}

	fmt.Println("Pertandingan selesai")
}
