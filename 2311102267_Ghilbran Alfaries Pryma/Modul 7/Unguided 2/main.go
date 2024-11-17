package main

import (
	"fmt"
	"math"
)

// Menampilkan seluruh isi array
func tampilkanArray(arr []int) {
	fmt.Println("Isi array:", arr)
}

// Menampilkan elemen dengan indeks genap
func tampilkanIndeksGenap(arr []int) {
	fmt.Println("Elemen dengan indeks genap:")
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 { // Indeks genap
			fmt.Printf("Indeks %d: %d\n", i, arr[i])
		}
	}
	fmt.Println()
}

// Menampilkan elemen dengan indeks ganjil
func tampilkanIndeksGanjil(arr []int) {
	fmt.Println("Elemen dengan indeks ganjil:")
	for i := 0; i < len(arr); i++ {
		if i%2 != 0 { // Indeks ganjil
			fmt.Printf("Indeks %d: %d\n", i, arr[i])
		}
	}
	fmt.Println()
}

// Menampilkan elemen dengan indeks kelipatan x
func tampilkanIndeksKelipatanX(arr []int, x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	for i := 0; i < len(arr); i++ {
		if i%x == 0 { // Indeks kelipatan x
			fmt.Printf("Indeks %d: %d\n", i, arr[i])
		}
	}
	fmt.Println()
}

// Menghapus elemen array pada indeks tertentu
func hapusElemen(arr []int, indeks int) []int {
	return append(arr[:indeks], arr[indeks+1:]...)
}

// Menghitung rata-rata
func hitungRataRata(arr []int) float64 {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return float64(sum) / float64(len(arr))
}

// Menghitung standar deviasi
func hitungStandarDeviasi(arr []int) float64 {
	if len(arr) <= 1 {
		return 0
	}

	mean := hitungRataRata(arr)
	var variance float64
	for _, v := range arr {
		variance += math.Pow(float64(v)-mean, 2)
	}
	variance /= float64(len(arr))
	return math.Sqrt(variance)
}

// Menghitung frekuensi dari bilangan tertentu
func hitungFrekuensi(arr []int, bilangan int) int {
	count := 0
	for _, v := range arr {
		if v == bilangan {
			count++
		}
	}
	return count
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	var arr []int
	for i := 0; i < n; i++ {
		var elemen int
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&elemen)
		arr = append(arr, elemen)
	}

	// a. Menampilkan seluruh isi array
	tampilkanArray(arr)

	// b. Menampilkan elemen dengan indeks genap
	tampilkanIndeksGenap(arr)

	// c. Menampilkan elemen dengan indeks ganjil
	tampilkanIndeksGanjil(arr)

	// d. Menampilkan elemen dengan indeks kelipatan x
	var x int
	fmt.Print("Masukkan nilai x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	tampilkanIndeksKelipatanX(arr, x)

	// e. Menghapus elemen pada indeks tertentu
	var indeksHapus int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&indeksHapus)
	arr = hapusElemen(arr, indeksHapus)
	fmt.Println("Array setelah penghapusan elemen:", arr)

	// f. Menampilkan rata-rata
	rataRata := hitungRataRata(arr)
	fmt.Printf("Rata-rata elemen dalam array: %.2f\n", rataRata)

	// g. Menampilkan standar deviasi
	standarDeviasi := hitungStandarDeviasi(arr)
	fmt.Printf("Standar deviasi elemen dalam array: %.2f\n", standarDeviasi)

	// h. Menampilkan frekuensi suatu bilangan tertentu
	var bilangan int
	fmt.Print("Masukkan bilangan untuk menghitung frekuensinya: ")
	fmt.Scan(&bilangan)
	frekuensi := hitungFrekuensi(arr, bilangan)
	fmt.Printf("Frekuensi bilangan %d: %d\n", bilangan, frekuensi)
}
