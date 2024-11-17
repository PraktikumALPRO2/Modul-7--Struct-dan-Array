package main

import (
	"fmt"
	"math"
)

type ArrayOperations struct {
	data []float64
	size int
}

// Membuat array baru
func NewArrayOperations(capacity int) *ArrayOperations {
	return &ArrayOperations{
		data: make([]float64, 0, capacity),
		size: 0,
	}
}

// Menambah elemen ke array
func (a *ArrayOperations) Add(value float64) {
	if a.size < cap(a.data) {
		a.data = append(a.data, value)
		a.size++
	}
}

// a. Menampilkan seluruh isi array
func (a *ArrayOperations) DisplayAll() {
	fmt.Println("Seluruh isi array:")
	for i, v := range a.data {
		fmt.Printf("Index %d: %.2f\n", i, v)
	}
}

// b. Menampilkan elemen dengan indeks ganjil
func (a *ArrayOperations) DisplayOddIndices() {
	fmt.Println("Elemen dengan indeks ganjil:")
	for i, v := range a.data {
		if i%2 != 0 {
			fmt.Printf("Index %d: %.2f\n", i, v)
		}
	}
}

// c. Menampilkan elemen dengan indeks genap
func (a *ArrayOperations) DisplayEvenIndices() {
	fmt.Println("Elemen dengan indeks genap:")
	for i, v := range a.data {
		if i%2 == 0 {
			fmt.Printf("Index %d: %.2f\n", i, v)
		}
	}
}

// d. Menampilkan elemen dengan indeks kelipatan x
func (a *ArrayOperations) DisplayMultipleIndices(x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	for i, v := range a.data {
		if i%x == 0 {
			fmt.Printf("Index %d: %.2f\n", i, v)
		}
	}
}

// e. Menghapus elemen pada indeks tertentu
func (a *ArrayOperations) DeleteAtIndex(index int) {
	if index >= 0 && index < a.size {
		a.data = append(a.data[:index], a.data[index+1:]...)
		a.size--
	}
}

// f. Menghitung rata-rata
func (a *ArrayOperations) CalculateAverage() float64 {
	if a.size == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range a.data {
		sum += v
	}
	return sum / float64(a.size)
}

// g. Menghitung standar deviasi
func (a *ArrayOperations) CalculateStandardDeviation() float64 {
	if a.size < 2 {
		return 0
	}
	mean := a.CalculateAverage()
	sumSquaredDiff := 0.0
	for _, v := range a.data {
		diff := v - mean
		sumSquaredDiff += diff * diff
	}
	variance := sumSquaredDiff / float64(a.size-1)
	return math.Sqrt(variance)
}

// h. Menghitung frekuensi suatu bilangan
func (a *ArrayOperations) CalculateFrequency(number float64) int {
	frequency := 0
	for _, v := range a.data {
		if v == number {
			frequency++
		}
	}
	return frequency
}

func main() {
	// Membuat array dengan kapasitas 10
	arr := NewArrayOperations(10)

	// Menambahkan beberapa nilai
	arr.Add(1.0)
	arr.Add(2.0)
	arr.Add(3.0)
	arr.Add(4.0)
	arr.Add(5.0)
	arr.Add(2.0)
	arr.Add(3.0)

	// Demonstrasi semua operasi
	fmt.Println("\n=== Demonstrasi Operasi Array ===")

	// a. Menampilkan semua elemen
	arr.DisplayAll()

	fmt.Println()
	// b. Menampilkan indeks ganjil
	arr.DisplayOddIndices()

	fmt.Println()
	// c. Menampilkan indeks genap
	arr.DisplayEvenIndices()

	fmt.Println()
	// d. Menampilkan kelipatan 2
	arr.DisplayMultipleIndices(2)

	fmt.Println()
	// e. Menghapus elemen index 2
	fmt.Println("Menghapus elemen index 2:")
	arr.DeleteAtIndex(2)
	arr.DisplayAll()

	fmt.Println()
	// f. Menampilkan rata-rata
	fmt.Printf("Rata-rata: %.2f\n", arr.CalculateAverage())

	// g. Menampilkan standar deviasi
	fmt.Printf("Standar deviasi: %.2f\n", arr.CalculateStandardDeviation())

	// h. Menampilkan frekuensi angka 2.0
	fmt.Printf("Frekuensi angka 2.0: %d\n", arr.CalculateFrequency(2.0))
}
