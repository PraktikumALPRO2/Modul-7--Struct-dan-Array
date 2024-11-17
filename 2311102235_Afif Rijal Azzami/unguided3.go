package main

import "fmt"

func main() {
	// Deklarasi variabel
	var klubA235, klubB235 string
	var skorA235, skorB235 int
	var pemenang235 []string
	var pertandingan235 int = 1

	// Input nama klub
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA235)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB235)

	fmt.Println("\nMasukkan skor untuk setiap pertandingan. Masukkan skor negatif untuk menghentikan.")

	// Loop untuk menerima skor dan menentukan pemenang
	for {
		fmt.Printf("Pertandingan %d - Skor %s: ", pertandingan235, klubA235)
		fmt.Scan(&skorA235)
		fmt.Printf("Pertandingan %d - Skor %s: ", pertandingan235, klubB235)
		fmt.Scan(&skorB235)

		// Kondisi untuk menghentikan program jika skor negatif
		if skorA235 < 0 || skorB235 < 0 {
			fmt.Println("Pertandingan selesai.")
			break
		}

		// Menentukan hasil pertandingan
		if skorA235 > skorB235 {
			pemenang235 = append(pemenang235, klubA235)
		} else if skorB235 > skorA235 {
			pemenang235 = append(pemenang235, klubB235)
		} else {
			pemenang235 = append(pemenang235, "Draw")
		}

		pertandingan235++
	}

	// Menampilkan hasil pertandingan
	fmt.Println("\nHasil pertandingan:")
	for i235, hasil235 := range pemenang235 {
		fmt.Printf("Hasil %d: %s\n", i235+1, hasil235)
	}
}
