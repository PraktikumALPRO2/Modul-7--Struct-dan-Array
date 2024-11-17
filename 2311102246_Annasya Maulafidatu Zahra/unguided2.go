package main

import (
	"fmt"
	"math"
)

func displayArray(arr []int) {
	fmt.Println("Isi array:", arr)
}

func displayOddIndices(arr []int) {
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func displayEvenIndices(arr []int) {
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func displayMultiples(arr []int, x int) {
	fmt.Print("Elemen dengan indeks kelipatan ", x, ": ")
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

func removeElement(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func calculateAverage(arr []int) float64 {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return float64(sum) / float64(len(arr))
}

func calculateStandardDeviation(arr []int, average float64) float64 {
	sum := 0.0
	for _, value := range arr {
		sum += math.Pow(float64(value)-average, 2)
	}
	return math.Sqrt(sum / float64(len(arr)))
}

func frequency(arr []int, number int) int {
	count := 0
	for _, value := range arr {
		if value == number {
			count++
		}
	}
	return count
}

func main() {
	var n, x, index, number int

	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	displayArray(arr)
	displayOddIndices(arr)
	displayEvenIndices(arr)
	fmt.Print("Masukkan bilangan x untuk kelipatan: ")
	fmt.Scan(&x)
	displayMultiples(arr, x)

	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&index)
	arr = removeElement(arr, index)
	displayArray(arr)

	average := calculateAverage(arr)
	fmt.Printf("Rata-rata: %.2f\n", average)

	stdDev := calculateStandardDeviation(arr, average)
	fmt.Printf("Standar deviasi: %.2f\n", stdDev)

	fmt.Print("Masukkan bilangan untuk menghitung frekuensi: ")
	fmt.Scan(&number)
	freq := frequency(arr, number)
	fmt.Printf("Frekuensi bilangan %d: %d\n", number, freq)
}
