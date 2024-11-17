package main

import (
	"fmt"
	"math"
)

func main() {
	var n_158, index, x_158, target int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n_158)

	array := make([]int, n_158)

	// Mengisi array
	for i := 0; i < n_158; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	for {
		fmt.Println("\nPilih operasi:")
		fmt.Println("1. Tampilkan seluruh isi array")
		fmt.Println("2. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("3. Tampilkan elemen dengan indeks genap")
		fmt.Println("4. Tampilkan elemen dengan indeks kelipatan bilangan tertentu")
		fmt.Println("5. Hapus elemen pada indeks tertentu")
		fmt.Println("6. Hitung rata-rata")
		fmt.Println("7. Hitung standar deviasi")
		fmt.Println("8. Hitung frekuensi bilangan tertentu")
		fmt.Println("9. Keluar")
		fmt.Print("Pilihan: ")
		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("Isi array:", array)
		case 2:
			fmt.Println("Elemen dengan indeks ganjil:")
			for i := 1; i < len(array); i += 2 {
				fmt.Printf("Indeks %d: %d\n_158", i, array[i])
			}
		case 3:
			fmt.Println("Elemen dengan indeks genap:")
			for i := 0; i < len(array); i += 2 {
				fmt.Printf("Indeks %d: %d\n_158", i, array[i])
			}
		case 4:
			fmt.Print("Masukkan bilangan untuk kelipatan indeks (x): ")
			fmt.Scan(&x_158)
			fmt.Println("Elemen dengan indeks kelipatan", x_158, ":")
			for i := 0; i < len(array); i++ {
				if i%x_158 == 0 {
					fmt.Printf("Indeks %d: %d\n_158", i, array[i])
				}
			}
		case 5:
			fmt.Print("Masukkan indeks yang ingin dihapus: ")
			fmt.Scan(&index)
			if index >= 0 && index < len(array) {
				array = append(array[:index], array[index+1:]...)
				fmt.Println("Isi array setelah dihapus:", array)
			} else {
				fmt.Println("Indeks tidak valid!")
			}
		case 6:
			sum := 0
			for _, val := range array {
				sum += val
			}
			average := float64(sum) / float64(len(array))
			fmt.Printf("Rata-rata: %.2f\n_158", average)
		case 7:
			sum := 0
			for _, val := range array {
				sum += val
			}
			average := float64(sum) / float64(len(array))

			var varianceSum float64
			for _, val := range array {
				varianceSum += math.Pow(float64(val)-average, 2)
			}
			standardDeviation := math.Sqrt(varianceSum / float64(len(array)))
			fmt.Printf("Standar deviasi: %.2f\n_158", standardDeviation)
		case 8:
			fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
			fmt.Scan(&target)
			count := 0
			for _, val := range array {
				if val == target {
					count++
				}
			}
			fmt.Printf("Frekuensi %d: %d kali\n_158", target, count)
		case 9:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}


