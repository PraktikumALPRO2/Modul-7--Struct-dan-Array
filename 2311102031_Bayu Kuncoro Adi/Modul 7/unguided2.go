package main

import (
	"fmt"
	"math"
)

func tampilkanArray(arr []int) {
	fmt.Println("Isi Array:")
	for _, elemen := range arr {
		fmt.Print(elemen, " ")
	}
	fmt.Println()
}

func tampilkanIndeksGanjil(arr []int) {
	fmt.Println("Elemen dengan indeks ganjil:")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func tampilkanIndeksGenap(arr []int) {
	fmt.Println("Elemen dengan indeks genap:")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func tampilkanIndeksKelipatanX(arr []int, x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

func hapusElemen(arr []int, indeks int) []int {
	arr = append(arr[:indeks], arr[indeks+1:]...)
	return arr
}

func hitungRataRata(arr []int) float64 {
	var sum int
	for _, elemen := range arr {
		sum += elemen
	}
	return float64(sum) / float64(len(arr))
}

func hitungStandarDeviasi(arr []int) float64 {
	mean := hitungRataRata(arr)
	var sumSquares float64
	for _, elemen := range arr {
		sumSquares += math.Pow(float64(elemen)-mean, 2)
	}
	variance := sumSquares / float64(len(arr))
	return math.Sqrt(variance)
}

func hitungFrekuensi(arr []int, bilangan int) int {
	var count int
	for _, elemen := range arr {
		if elemen == bilangan {
			count++
		}
	}
	return count
}

func main() {
	var N, x, indeksHapus, bilanganCari int

	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&N)

	arr := make([]int, N)

	fmt.Println("Masukkan nilai untuk setiap elemen array:")
	for i := 0; i < N; i++ {
		fmt.Printf("Elemen ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	tampilkanArray(arr)

	tampilkanIndeksGanjil(arr)

	tampilkanIndeksGenap(arr)

	fmt.Print("Masukkan nilai x untuk indeks kelipatan x: ")
	fmt.Scan(&x)
	tampilkanIndeksKelipatanX(arr, x)

	fmt.Print("Masukkan indeks elemen yang akan dihapus: ")
	fmt.Scan(&indeksHapus)
	arr = hapusElemen(arr, indeksHapus)
	fmt.Println("Array setelah elemen dihapus:")
	tampilkanArray(arr)

	rataRata := hitungRataRata(arr)
	fmt.Printf("Rata-rata array: %.2f\n", rataRata)

	standarDeviasi := hitungStandarDeviasi(arr)
	fmt.Printf("Standar deviasi array: %.2f\n", standarDeviasi)

	fmt.Print("Masukkan bilangan untuk mencari frekuensinya: ")
	fmt.Scan(&bilanganCari)
	frekuensi := hitungFrekuensi(arr, bilanganCari)
	fmt.Printf("Frekuensi bilangan %d: %d kali\n", bilanganCari, frekuensi)
}
