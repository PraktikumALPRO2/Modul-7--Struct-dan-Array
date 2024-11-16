package main

import (
	"fmt"
	"math"
)

func tampilkanSeluruhArray(arr []int) {
	fmt.Println("Seluruh elemen array:", arr)
}

func tampilkanIndeksGanjil(arr []int) {
	fmt.Println("Elemen dengan indeks ganjil:")
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}
}

func tampilkanIndeksGenap(arr []int) {
	fmt.Println("Elemen dengan indeks genap:")
	for i := 0; i < len(arr); i += 2 {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}
}

func tampilkanIndeksKelipatan(arr []int, x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	for i := x; i < len(arr); i += x {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}
}

func hapusElemenPadaIndeks(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		fmt.Println("Indeks tidak valid")
		return arr
	}
	fmt.Printf("Menghapus elemen pada indeks %d\n", index)
	arr = append(arr[:index], arr[index+1:]...)
	return arr
}

func hitungRataRata(arr []int) float64 {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return float64(sum) / float64(len(arr))
}

func hitungStandarDeviasi(arr []int) float64 {
	rataRata := hitungRataRata(arr)
	sum := 0.0
	for _, val := range arr {
		sum += math.Pow(float64(val)-rataRata, 2)
	}
	return math.Sqrt(sum / float64(len(arr)))
}

func hitungFrekuensi(arr []int, bilangan int) int {
	count := 0
	for _, val := range arr {
		if val == bilangan {
			count++
		}
	}
	return count
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	tampilkanSeluruhArray(arr)
	tampilkanIndeksGanjil(arr)
	tampilkanIndeksGenap(arr)

	var x int
	fmt.Print("Masukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	tampilkanIndeksKelipatan(arr, x)

	var index int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&index)
	arr = hapusElemenPadaIndeks(arr, index)
	tampilkanSeluruhArray(arr)

	rataRata := hitungRataRata(arr)
	fmt.Printf("Rata-rata elemen array: %.2f\n", rataRata)

	standarDeviasi := hitungStandarDeviasi(arr)
	fmt.Printf("Standar deviasi elemen array: %.2f\n", standarDeviasi)

	var bilangan int
	fmt.Print("Masukkan bilangan untuk menghitung frekuensinya: ")
	fmt.Scan(&bilangan)
	frekuensi := hitungFrekuensi(arr, bilangan)
	fmt.Printf("Frekuensi %d di dalam array: %d kali\n", bilangan, frekuensi)
}
