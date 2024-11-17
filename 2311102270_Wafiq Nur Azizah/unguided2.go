package main

import (
	"fmt"
	"math"
)

func main() {
	var jumlahElemen int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&jumlahElemen)
	dataArray := make([]int, jumlahElemen)

	for i := 0; i < jumlahElemen; i++ {
		fmt.Print("Masukkan nilai untuk elemen array ke-", i, ": ")
		fmt.Scan(&dataArray[i])
	}

	menuPilihan(dataArray)
}

func menuPilihan(array []int) {
	var pilihanMenu int
	n := len(array)

	fmt.Println("-----MENU PENGOLAHAN DATA ARRAY-----")
	fmt.Println("1. Tampilkan seluruh isi array")
	fmt.Println("2. Tampilkan elemen array dengan indeks ganjil")
	fmt.Println("3. Tampilkan elemen array dengan indeks genap")
	fmt.Println("4. Tampilkan elemen array yang indeksnya kelipatan x")
	fmt.Println("5. Hapus elemen array pada indeks tertentu")
	fmt.Println("6. Tampilkan rata-rata elemen array")
	fmt.Println("7. Tampilkan simpangan baku array")
	fmt.Println("8. Tampilkan frekuensi suatu bilangan dalam array")
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scan(&pilihanMenu)

	switch pilihanMenu {
	case 1:
		for i := 0; i < n; i++ {
			fmt.Print(array[i], " ")
		}
		fmt.Println()
	case 2:
		for i := 1; i < n; i += 2 {
			fmt.Print(array[i], " ")
		}
		fmt.Println()
	case 3:
		for i := 0; i < n; i += 2 {
			fmt.Print(array[i], " ")
		}
		fmt.Println()
	case 4:
		var kelipatanX int
		fmt.Print("Masukkan nilai kelipatan (x): ")
		fmt.Scan(&kelipatanX)
		for i := 0; i < n; i++ {
			if i%kelipatanX == 0 {
				fmt.Print(array[i], " ")
			}
		}
		fmt.Println()
	case 5:
		var indeksHapus int
		fmt.Print("Masukkan indeks yang ingin dihapus: ")
		fmt.Scan(&indeksHapus)
		if indeksHapus < 0 || indeksHapus >= n {
			fmt.Println("Indeks tidak valid!")
			return
		}
		arrayBaru := make([]int, 0, n-1)
		for i := 0; i < n; i++ {
			if i != indeksHapus {
				arrayBaru = append(arrayBaru, array[i])
			}
		}
		fmt.Println("Array setelah penghapusan:", arrayBaru)
	case 6:
		var jumlah int
		for i := 0; i < n; i++ {
			jumlah += array[i]
		}
		rataRata := float64(jumlah) / float64(n)
		fmt.Println("Rata-rata elemen array adalah:", rataRata)
	case 7:
		var selisih, totalSelisih float64
		var jumlah int
		for i := 0; i < n; i++ {
			jumlah += array[i]
		}
		rataRata := float64(jumlah) / float64(n)
		for i := 0; i < n; i++ {
			selisih = float64(array[i]) - rataRata
			totalSelisih += selisih * selisih
		}
		simpanganBaku := math.Sqrt(totalSelisih / float64(n))
		fmt.Println("Simpangan baku dari array adalah:", simpanganBaku)
	case 8:
		var frekuensi, angka int
		fmt.Print("Masukkan angka yang frekuensinya ingin dihitung: ")
		fmt.Scan(&angka)
		for i := 0; i < n; i++ {
			if array[i] == angka {
				frekuensi++
			}
		}
		fmt.Printf("Frekuensi bilangan %d dalam array adalah %d\n", angka, frekuensi)
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}
