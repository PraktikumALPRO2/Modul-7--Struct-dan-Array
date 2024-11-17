// Fadilah Salehah
// 2311102164

package main

import "fmt"

func main() {
	// Deklarasi variabel
	var timA, timB string
	var skorTimA, skorTimB int
	var hasilPertandingan []string
	var nomorPertandingan int = 1

	// Input nama tim
	fmt.Print("Nama Tim A: ")
	fmt.Scan(&timA)
	fmt.Print("Nama Tim B: ")
	fmt.Scan(&timB)

	fmt.Println("\nMasukkan skor untuk setiap pertandingan. Masukkan skor negatif untuk menghentikan.")

	// Perulangan untuk menerima skor dan menentukan pemenang
	for {
		fmt.Printf("Pertandingan %d - Skor %s: ", nomorPertandingan, timA)
		fmt.Scan(&skorTimA)
		fmt.Printf("Pertandingan %d - Skor %s: ", nomorPertandingan, timB)
		fmt.Scan(&skorTimB)

		// Kondisi untuk menghentikan jika skor negatif
		if skorTimA < 0 || skorTimB < 0 {
			fmt.Println("Input skor dihentikan.")
			break
		}

		// Menentukan pemenang atau hasil seri
		if skorTimA > skorTimB {
			hasilPertandingan = append(hasilPertandingan, timA)
		} else if skorTimB > skorTimA {
			hasilPertandingan = append(hasilPertandingan, timB)
		} else {
			hasilPertandingan = append(hasilPertandingan, "Draw")
		}

		nomorPertandingan++
	}

	// Menampilkan hasil akhir pertandingan
	fmt.Println("\nRekap Hasil Pertandingan:")
	for i, hasil := range hasilPertandingan {
		fmt.Printf("Pertandingan %d: %s\n", i+1, hasil)
	}
}