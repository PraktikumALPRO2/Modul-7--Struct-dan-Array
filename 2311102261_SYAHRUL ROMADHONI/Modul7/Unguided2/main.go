package main

import (
	"fmt"
	"math"
)

func tampilkanArray(arr []int) {
	fmt.Println("Isi array:", arr)
}

func tampilkanIndeksGanjil(arr []int) {
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func tampilkanIndeksGenap(arr []int) {
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func tampilkanKelipatanIndeks(arr []int, x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
	for i := x; i < len(arr); i += x {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func hapusElemen(arr []int, idx int) []int {
	if idx < 0 || idx >= len(arr) {
		fmt.Println("Indeks tidak valid!")
		return arr
	}
	fmt.Printf("Menghapus elemen di indeks %d: %d\n", idx, arr[idx])
	return append(arr[:idx], arr[idx+1:]...)
}

func rataRata(arr []int) float64 {
	total := 0
	for _, val := range arr {
		total += val
	}
	return float64(total) / float64(len(arr))
}

func standarDeviasi(arr []int) float64 {
	rata := rataRata(arr)
	var total float64
	for _, val := range arr {
		total += math.Pow(float64(val)-rata, 2)
	}
	return math.Sqrt(total / float64(len(arr)))
}

func hitungFrekuensi(arr []int, target int) int {
	count := 0
	for _, val := range arr {
		if val == target {
			count++
		}
	}
	return count
}

func main() {
	arr := []int{5, 10, 15, 20, 25, 30, 35, 40}

	tampilkanArray(arr)

	tampilkanIndeksGanjil(arr)

	tampilkanIndeksGenap(arr)

	var x int
	fmt.Print("Masukkan kelipatan indeks yang diinginkan: ")
	fmt.Scan(&x)
	tampilkanKelipatanIndeks(arr, x)

	var idx int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idx)
	arr = hapusElemen(arr, idx)
	tampilkanArray(arr)

	fmt.Printf("Rata-rata elemen array: %.2f\n", rataRata(arr))

	fmt.Printf("Standar deviasi elemen array: %.2f\n", standarDeviasi(arr))

	var target int
	fmt.Print("Masukkan bilangan untuk menghitung frekuensinya: ")
	fmt.Scan(&target)
	fmt.Printf("Frekuensi bilangan %d: %d\n", target, hitungFrekuensi(arr, target))
}
