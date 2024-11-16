package main

import (
	"fmt"
	"math"
)

func main() {
	var jumlahElemen, kelipatan, indeksHapus int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&jumlahElemen)

	daftarBilangan := make([]int, jumlahElemen)
	for i := 0; i < jumlahElemen; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&daftarBilangan[i])
	}

	fmt.Println("\nPilihan operasi:")
	fmt.Println("a. Tampilkan keseluruhan isi array")
	fmt.Println("b. Tampilkan elemen dengan indeks ganjil")
	fmt.Println("c. Tampilkan elemen dengan indeks genap")
	fmt.Println("d. Tampilkan elemen dengan indeks kelipatan tertentu")
	fmt.Println("e. Hapus elemen pada indeks tertentu")
	fmt.Println("f. Tampilkan rata-rata elemen")
	fmt.Println("g. Tampilkan standar deviasi elemen")
	fmt.Println("h. Tampilkan frekuensi elemen tertentu")

	var pilihan string
	fmt.Print("\nPilih operasi: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case "a":
		fmt.Println("Isi array:", daftarBilangan)
	case "b":
		fmt.Println("Elemen dengan nilai ganjil:")
		for i := 0; i < jumlahElemen; i++ {
			if daftarBilangan[i]%2 != 0 { // Cek jika elemen adalah bilangan ganjil
				fmt.Print(daftarBilangan[i], " ")
			}
		}
		fmt.Println()
	case "c":
		fmt.Println("Elemen dengan nilai genap:")
		for i := 1; i < jumlahElemen; i += 2 { 
			fmt.Print(daftarBilangan[i], " ")
		}
		fmt.Println()
	case "d":
		fmt.Print("Masukkan nilai kelipatan: ")
		fmt.Scan(&kelipatan)
		fmt.Printf("Elemen dengan kelipatan indeks %d:\n", kelipatan)
		for i := 0; i < jumlahElemen; i++ {
			if daftarBilangan[i] % kelipatan == 0 { // Memeriksa apakah nilai elemen adalah kelipatan dari angka yang dimasukkan
				fmt.Print(daftarBilangan[i], " ")
			}
		}
		fmt.Println()	
	case "e":
		fmt.Print("Masukkan indeks yang akan dihapus: ")
		fmt.Scan(&indeksHapus)
		if indeksHapus >= 0 && indeksHapus < jumlahElemen {
			daftarBilangan = append(daftarBilangan[:indeksHapus], daftarBilangan[indeksHapus+1:]...)
			fmt.Println("Array setelah penghapusan:", daftarBilangan)
		} else {
			fmt.Println("Indeks tidak valid.")
		}
	case "f":
		total := 0
		for _, bilangan := range daftarBilangan {
			total += bilangan
		}
		rataRata := float64(total) / float64(jumlahElemen)
		fmt.Println("Rata-rata:", rataRata)
	case "g":
		total := 0
		for _, bilangan := range daftarBilangan {
			total += bilangan
		}
		rataRata := float64(total) / float64(jumlahElemen)
		var jumlahVarian float64
		for _, bilangan := range daftarBilangan {
			jumlahVarian += math.Pow(float64(bilangan)-rataRata, 2)
		}
		standarDeviasi := math.Sqrt(jumlahVarian / float64(jumlahElemen))
		fmt.Println("Standar deviasi:", standarDeviasi)
	case "h":
		fmt.Print("Masukkan elemen yang ingin dicari frekuensinya: ")
		fmt.Scan(&kelipatan)
		frekuensi := 0
		for _, bilangan := range daftarBilangan {
			if bilangan == kelipatan {
				frekuensi++
			}
		}
		fmt.Printf("Frekuensi %d: %d\n", kelipatan, frekuensi)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}
