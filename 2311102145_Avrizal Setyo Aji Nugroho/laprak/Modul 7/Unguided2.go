package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	array_145 := make([]int, n)
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&array_145[i])
	}

	//Menampilkan seluruh isi array
	fmt.Println("Seluruh isi array:", array_145)

	// Menampilkan  indeks ganjil
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < len(array_145); i += 2 {
		fmt.Print(array_145[i], " ")
	}
	fmt.Println()

	// Menampilkan indeks genap
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(array_145); i += 2 {
		fmt.Print(array_145[i], " ")
	}
	fmt.Println()

	// Menampilkan elemen indeks kelipatan x
	var x int
	fmt.Print("Masukkan bilangan x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Print("Elemen dengan indeks kelipatan ", x, ": ")
	for i := 0; i < len(array_145); i++ {
		if i%x == 0 {
			fmt.Print(array_145[i], " ")
		}
	}
	fmt.Println()

	// Menghapus array pada indeks tertentu
	var index int
	fmt.Print("Masukkan indeks elemen yang akan dihapus: ")
	fmt.Scan(&index)
	if index >= 0 && index < len(array_145) {
		array_145 = append(array_145[:index], array_145[index+1:]...)
		fmt.Println("Isi array setelah dihapus:", array_145)
	} else {
		fmt.Println("Indeks tidak valid!")
	}

	// Menampilkan rata-rata array
	var sum int
	for _, v := range array_145 {
		sum += v
	}
	rataRata := float64(sum) / float64(len(array_145))
	fmt.Printf("f. Rata-rata elemen array: %.2f\n", rataRata)

	// Menampilkan standar deviasi elemen array
	var variance float64
	for _, v := range array_145 {
		variance += math.Pow(float64(v)-rataRata, 2)
	}
	stdDev := math.Sqrt(variance / float64(len(array_145)))
	fmt.Printf("Standar deviasi elemen array: %.2f\n", stdDev)

	// Menampilkan frekuensi suatu bilangan dalam array
	var target int
	fmt.Print("Masukkan bilangan untuk menghitung frekuensinya: ")
	fmt.Scan(&target)
	frekuensi := 0
	for _, v := range array_145 {
		if v == target {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi bilangan %d: %d kali\n", target, frekuensi)
}
