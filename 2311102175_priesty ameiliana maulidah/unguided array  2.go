package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("Isi keseluruhan array:", arr)

	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < N; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < N; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan bilangan x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	fmt.Print("Elemen dengan indeks kelipatan ", x, ": ")
	for i := 0; i < N; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var indexToDelete int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&indexToDelete)
	arr = append(arr[:indexToDelete], arr[indexToDelete+1:]...)
	fmt.Println("Isi array setelah penghapusan:", arr)

	var sum, count int
	for _, value := range arr {
		sum += value
		count++
	}
	average := float64(sum) / float64(count)
	fmt.Println("Rata-rata:", average)

	var varianceSum float64
	for _, value := range arr {
		varianceSum += math.Pow(float64(value)-average, 2)
	}
	stdDev := math.Sqrt(varianceSum / float64(count))
	fmt.Println("Standar deviasi:", stdDev)

	var target int
	fmt.Print("Masukkan bilangan untuk frekuensi: ")
	fmt.Scan(&target)
	frequency := 0
	for _, value := range arr {
		if value == target {
			frequency++
		}
	}
	fmt.Printf("Frekuensi dari bilangan %d: %d\n", target, frequency)
}