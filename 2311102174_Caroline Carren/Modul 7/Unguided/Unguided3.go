// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	// Input nama klub
	fmt.Print("Masukkan nama klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Masukkan nama klub B: ")
	fmt.Scanln(&klubB)

	// Loop untuk input skor
	for {
		fmt.Printf("Masukkan skor untuk %s dan %s : ", klubA, klubB)
		_, err := fmt.Scanf("%d %d\n", &skorA, &skorB)

		// Cek jika input tidak valid
		if err != nil {
			fmt.Println("Input tidak valid. Pastikan untuk memasukkan dua angka.")
			// Mengosongkan input buffer
			var dummy string
			fmt.Scanln(&dummy)
			continue
		}

		// Break loop jika salah satu skor negatif
		if skorA < 0 || skorB < 0 {
			break
		}

		// Tentukan hasil pertandingan
		if skorA > skorB {
			hasil = append(hasil, klubA) // klubA menang
		} else if skorA < skorB {
			hasil = append(hasil, klubB) // klubB menang
		} else {
			hasil = append(hasil, "Draw") // seri
		}
	}

	// Tampilkan hasil
	fmt.Println("Hasil pertandingan:")
	for i, h := range hasil {
		fmt.Printf("Hasil %d: %s\n", i+1, h)
	}
	fmt.Println("Pertandingan selesai")
}
