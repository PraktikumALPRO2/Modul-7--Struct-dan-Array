package main

import (
	"fmt"
	"math"
)

type ArrayOperations struct {
	data []float64
}

// Membuat array baru
func NewArrayOperations(capacity int) *ArrayOperations {
	return &ArrayOperations{data: make([]float64, 0, capacity)}
}

// Menambah elemen ke array
func (a *ArrayOperations) Add(value float64) {
	if len(a.data) < cap(a.data) {
		a.data = append(a.data, value)
	}
}

// Menampilkan seluruh isi array
func (a *ArrayOperations) DisplayAll() {
	fmt.Println("Seluruh isi array:")
	for i, v := range a.data {
		fmt.Printf("Index %d: %.2f\n", i, v)
	}
}

// Menampilkan elemen dengan indeks ganjil atau genap
func (a *ArrayOperations) DisplayIndices(isEven bool) {
	if isEven {
		fmt.Println("Elemen dengan indeks genap:")
	} else {
		fmt.Println("Elemen dengan indeks ganjil:")
	}
	for i, v := range a.data {
		if (isEven && i%2 == 0) || (!isEven && i%2 != 0) {
			fmt.Printf("Index %d: %.2f\n", i, v)
		}
	}
}

// Menampilkan elemen dengan indeks kelipatan x
func (a *ArrayOperations) DisplayMultipleIndices(x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	for i, v := range a.data {
		if i%x == 0 {
			fmt.Printf("Index %d: %.2f\n", i, v)
		}
	}
}

// Menghapus elemen pada indeks tertentu
func (a *ArrayOperations) DeleteAtIndex(index int) {
	if index >= 0 && index < len(a.data) {
		a.data = append(a.data[:index], a.data[index+1:]...)
	}
}

// Menghitung rata-rata
func (a *ArrayOperations) CalculateAverage() float64 {
	sum := 0.0
	for _, v := range a.data {
		sum += v
	}
	return sum / float64(len(a.data))
}

// Menghitung standar deviasi
func (a *ArrayOperations) CalculateStandardDeviation() float64 {
	if len(a.data) < 2 {
		return 0
	}
	mean := a.CalculateAverage()
	sumSquaredDiff := 0.0
	for _, v := range a.data {
		diff := v - mean
		sumSquaredDiff += diff * diff
	}
	return math.Sqrt(sumSquaredDiff / float64(len(a.data)-1))
}

// Menghitung frekuensi suatu bilangan
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
	arr := NewArrayOperations(10)

	values := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.0, 3.0}
	for _, value := range values {
		arr.Add(value)
	}

	fmt.Println("\n=== Demonstrasi Operasi Array ===")
	arr.DisplayAll()
	fmt.Println()
	arr.DisplayIndices(false) // Indeks ganjil
	fmt.Println()
	arr.DisplayIndices(true) // Indeks genap
	fmt.Println()
	arr.DisplayMultipleIndices(2)
	fmt.Println()

	fmt.Println("Menghapus elemen index 2:")
	arr.DeleteAtIndex(2)
	arr.DisplayAll()
	fmt.Println()

	fmt.Printf("Rata-rata: %.2f\n", arr.CalculateAverage())
	fmt.Printf("Standar deviasi: %.2f\n", arr.CalculateStandardDeviation())
	fmt.Printf("Frekuensi angka 2.0: %d\n", arr.CalculateFrequency(2.0))
}