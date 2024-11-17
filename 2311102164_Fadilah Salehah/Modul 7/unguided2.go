// Fadilah Salehah
// 2311102164

package main

import (
	"fmt"
	"math"
)

func main() {
	var jumlah int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&jumlah)

	dataArray := make([]int, jumlah)

	// Input elemen array
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&dataArray[i])
	}

	for {
		// Menu opsi
		fmt.Println("\n=== MENU ===")
		fmt.Println("a. Tampilkan seluruh elemen array")
		fmt.Println("b. Tampilkan elemen array pada indeks ganjil")
		fmt.Println("c. Tampilkan elemen array pada indeks genap")
		fmt.Println("d. Tampilkan elemen array pada indeks kelipatan tertentu")
		fmt.Println("e. Hapus elemen array berdasarkan indeks")
		fmt.Println("f. Hitung rata-rata elemen array")
		fmt.Println("g. Hitung standar deviasi elemen array")
		fmt.Println("h. Hitung frekuensi elemen tertentu")
		fmt.Println("i. Keluar")
		fmt.Print("Pilih menu: ")
		var pilihan string
		fmt.Scan(&pilihan)

		switch pilihan {
		case "a":
			// Tampilkan seluruh elemen array
			fmt.Println("Isi array:", dataArray)

		case "b":
			// Tampilkan elemen dengan indeks ganjil
			fmt.Println("Elemen pada indeks ganjil:")
			for i := 1; i < jumlah; i += 2 {
				fmt.Printf("Indeks %d: %d\n", i, dataArray[i])
			}

		case "c":
			// Tampilkan elemen dengan indeks genap
			fmt.Println("Elemen pada indeks genap:")
			for i := 0; i < jumlah; i += 2 {
				fmt.Printf("Indeks %d: %d\n", i, dataArray[i])
			}

		case "d":
			// Tampilkan elemen pada indeks kelipatan tertentu
			var kelipatan int
			fmt.Print("Masukkan nilai kelipatan: ")
			fmt.Scan(&kelipatan)
			fmt.Printf("Elemen pada indeks kelipatan %d:\n", kelipatan)
			for i := 0; i < jumlah; i++ {
				if i%kelipatan == 0 {
					fmt.Printf("Indeks %d: %d\n", i, dataArray[i])
				}
			}

		case "e":
			// Hapus elemen berdasarkan indeks
			var indeks int
			fmt.Print("Masukkan indeks yang ingin dihapus: ")
			fmt.Scan(&indeks)
			if indeks >= 0 && indeks < jumlah {
				dataArray = append(dataArray[:indeks], dataArray[indeks+1:]...)
				jumlah-- // Kurangi jumlah elemen
				fmt.Println("Array setelah penghapusan:", dataArray)
			} else {
				fmt.Println("Indeks tidak valid!")
			}

		case "f":
			// Hitung rata-rata elemen array
			total := 0
			for _, nilai := range dataArray {
				total += nilai
			}
			rataRata := float64(total) / float64(len(dataArray))
			fmt.Printf("Rata-rata: %.2f\n", rataRata)

		case "g":
			// Hitung standar deviasi elemen array
			total := 0
			for _, nilai := range dataArray {
				total += nilai
			}
			rataRata := float64(total) / float64(len(dataArray))

			var variansi float64
			for _, nilai := range dataArray {
				variansi += math.Pow(float64(nilai)-rataRata, 2)
			}
			variansi /= float64(len(dataArray))
			stdDeviasi := math.Sqrt(variansi)
			fmt.Printf("Standar deviasi: %.2f\n", stdDeviasi)

		case "h":
			// Hitung frekuensi elemen tertentu
			var elemen int
			fmt.Print("Masukkan elemen yang ingin dihitung frekuensinya: ")
			fmt.Scan(&elemen)
			frekuensi := 0
			for _, nilai := range dataArray {
				if nilai == elemen {
					frekuensi++
				}
			}
			fmt.Printf("Frekuensi elemen %d: %d kali\n", elemen, frekuensi)

		case "i":
			// Keluar dari program
			fmt.Println("Program selesai. Terima kasih!")
			return

		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi!")
		}
	}
}