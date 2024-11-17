package main

import (
	"fmt"
	"math"
)

func main() {
	// Input jumlah elemen array
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	// Inisialisasi array
	arr := make([]int, n)
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	// a. Menampilkan keseluruhan isi array
	fmt.Println("Keseluruhan isi array:", arr)

	// b. Menampilkan elemen-elemen array dengan indeks ganjil
	fmt.Print("Elemen array dengan indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// c. Menampilkan elemen-elemen array dengan indeks genap
	fmt.Print("Elemen array dengan indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x
	var x int
	fmt.Print("Masukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen array dengan indeks kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// e. Menghapus elemen array pada indeks tertentu
	var index int
	fmt.Print("Masukkan indeks elemen yang akan dihapus: ")
	fmt.Scan(&index)
	if index >= 0 && index < n {
		arr = append(arr[:index], arr[index+1:]...)
		fmt.Println("Array setelah elemen dihapus:", arr)
	} else {
		fmt.Println("Indeks tidak valid!")
	}

	// f. Menampilkan rata-rata dari bilangan dalam array
	sum := 0
	for _, val := range arr {
		sum += val
	}
	mean := float64(sum) / float64(len(arr))
	fmt.Printf("Rata-rata bilangan dalam array: %.2f\n", mean)

	// g. Menampilkan standar deviasi
	varianceSum := 0.0
	for _, val := range arr {
		varianceSum += math.Pow(float64(val)-mean, 2)
	}
	stdDev := math.Sqrt(varianceSum / float64(len(arr)))
	fmt.Printf("Standar deviasi bilangan dalam array: %.2f\n", stdDev)

	// h. Menampilkan frekuensi dari suatu bilangan tertentu
	var target int
	fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&target)
	count := 0
	for _, val := range arr {
		if val == target {
			count++
		}
	}
	fmt.Printf("Frekuensi bilangan %d dalam array: %d\n", target, count)
}
