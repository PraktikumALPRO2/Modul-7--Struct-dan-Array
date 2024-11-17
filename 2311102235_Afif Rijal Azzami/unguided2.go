package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	array := make([]int, n)

	// Input array
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	for {
		// Menu
		fmt.Println("\n=== MENU ===")
		fmt.Println("a. Tampilkan keseluruhan isi array")
		fmt.Println("b. Tampilkan elemen array dengan indeks ganjil")
		fmt.Println("c. Tampilkan elemen array dengan indeks genap")
		fmt.Println("d. Tampilkan elemen array dengan indeks kelipatan x")
		fmt.Println("e. Hapus elemen array pada indeks tertentu")
		fmt.Println("f. Tampilkan rata-rata dari bilangan dalam array")
		fmt.Println("g. Tampilkan standar deviasi dari bilangan dalam array")
		fmt.Println("h. Tampilkan frekuensi dari bilangan tertentu")
		fmt.Println("i. Keluar")
		fmt.Print("Pilih opsi: ")
		var opsi string
		fmt.Scan(&opsi)

		switch opsi {
		case "a":
			// Tampilkan seluruh isi array
			fmt.Println("Isi array:", array)

		case "b":
			// Tampilkan elemen array dengan indeks ganjil
			fmt.Println("Elemen dengan indeks ganjil:")
			for i := 1; i < n; i += 2 {
				fmt.Printf("Index %d: %d\n", i, array[i])
			}

		case "c":
			// Tampilkan elemen array dengan indeks genap
			fmt.Println("Elemen dengan indeks genap:")
			for i := 0; i < n; i += 2 {
				fmt.Printf("Index %d: %d\n", i, array[i])
			}

		case "d":
			// Tampilkan elemen array dengan indeks kelipatan x
			var x int
			fmt.Print("Masukkan nilai x: ")
			fmt.Scan(&x)
			fmt.Println("Elemen dengan indeks kelipatan", x, ":")
			for i := 0; i < n; i++ {
				if i%x == 0 {
					fmt.Printf("Index %d: %d\n", i, array[i])
				}
			}

		case "e":
			// Hapus elemen array pada indeks tertentu
			var index int
			fmt.Print("Masukkan indeks yang ingin dihapus: ")
			fmt.Scan(&index)
			if index >= 0 && index < n {
				array = append(array[:index], array[index+1:]...)
				n-- // Kurangi jumlah elemen
				fmt.Println("Array setelah penghapusan:", array)
			} else {
				fmt.Println("Indeks tidak valid!")
			}

		case "f":
			// Tampilkan rata-rata
			sum := 0
			for _, val := range array {
				sum += val
			}
			avg := float64(sum) / float64(len(array))
			fmt.Printf("Rata-rata: %.2f\n", avg)

		case "g":
			// Tampilkan standar deviasi
			sum := 0
			for _, val := range array {
				sum += val
			}
			avg := float64(sum) / float64(len(array))

			var variance float64
			for _, val := range array {
				variance += math.Pow(float64(val)-avg, 2)
			}
			variance /= float64(len(array))
			stdDev := math.Sqrt(variance)
			fmt.Printf("Standar deviasi: %.2f\n", stdDev)

		case "h":
			// Tampilkan frekuensi dari bilangan tertentu
			var num int
			fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
			fmt.Scan(&num)
			count := 0
			for _, val := range array {
				if val == num {
					count++
				}
			}
			fmt.Printf("Frekuensi bilangan %d: %d kali\n", num, count)

		case "i":
			// Keluar
			fmt.Println("Program selesai.")
			return

		default:
			fmt.Println("Opsi tidak valid!")
		}
	}
}