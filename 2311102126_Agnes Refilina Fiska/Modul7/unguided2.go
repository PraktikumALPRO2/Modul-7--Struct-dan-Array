package main

import (
	"fmt"
	"math"
)

// ArrayOps struct to hold the array and its operations
type ArrayOps struct {
	arr    []int
	length int
}

// NewArrayOps creates a new ArrayOps instance
func NewArrayOps(capacity int) *ArrayOps {
	return &ArrayOps{
		arr:    make([]int, 0, capacity),
		length: 0,
	}
}

// inputArray fills the array with user input
func (a *ArrayOps) inputArray() {
	fmt.Print("Masukkan jumlah elemen (N): ")
	var n int
	fmt.Scan(&n)
	a.length = n

	fmt.Println("Masukkan", n, "bilangan:")
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		a.arr = append(a.arr, num)
	}
}

// displayAll shows all elements
func (a *ArrayOps) displayAll() {
	fmt.Print("Seluruh elemen array: ")
	for _, v := range a.arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

// displayOddIndices shows elements at odd indices
func (a *ArrayOps) displayOddIndices() {
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < len(a.arr); i += 2 {
		fmt.Print(a.arr[i], " ")
	}
	fmt.Println()
}

// displayEvenIndices shows elements at even indices
func (a *ArrayOps) displayEvenIndices() {
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(a.arr); i += 2 {
		fmt.Print(a.arr[i], " ")
	}
	fmt.Println()
}

// displayMultipleIndices shows elements at indices multiple of x
func (a *ArrayOps) displayMultipleIndices() {
	fmt.Print("Masukkan nilai x untuk menampilkan kelipatan: ")
	var x int
	fmt.Scan(&x)

	fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
	for i := 0; i < len(a.arr); i++ {
		if i%x == 0 {
			fmt.Print(a.arr[i], " ")
		}
	}
	fmt.Println()
}

// deleteElement removes element at specified index
func (a *ArrayOps) deleteElement() {
	fmt.Print("Masukkan indeks yang akan dihapus: ")
	var idx int
	fmt.Scan(&idx)

	a.arr = append(a.arr[:idx], a.arr[idx+1:]...)
	a.length--

	fmt.Print("Array setelah penghapusan: ")
	a.displayAll()
}

// calculateAverage computes mean of array elements
func (a *ArrayOps) calculateAverage() float64 {
	sum := 0
	for _, v := range a.arr {
		sum += v
	}
	avg := float64(sum) / float64(len(a.arr))
	fmt.Printf("Rata-rata: %.2f\n", avg)
	return avg
}

// calculateStdDev computes standard deviation
func (a *ArrayOps) calculateStdDev() {
	avg := a.calculateAverage()
	sumSquaredDiff := 0.0

	for _, v := range a.arr {
		diff := float64(v) - avg
		sumSquaredDiff += diff * diff
	}

	stdDev := math.Sqrt(sumSquaredDiff / float64(len(a.arr)))
	fmt.Printf("Standar deviasi: %.2f\n", stdDev)
}

// calculateFrequency counts occurrences of a number
func (a *ArrayOps) calculateFrequency() {
	fmt.Print("Masukkan bilangan yang ingin dihitung frekuensinya: ")
	var num int
	fmt.Scan(&num)

	freq := 0
	for _, v := range a.arr {
		if v == num {
			freq++
		}
	}

	fmt.Printf("Frekuensi bilangan %d: %d\n", num, freq)
}

func main() {
	// Initialize array with capacity 100
	arrOps := NewArrayOps(100)

	// Input array elements
	arrOps.inputArray()

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tampilkan semua elemen")
		fmt.Println("2. Tampilkan elemen indeks ganjil")
		fmt.Println("3. Tampilkan elemen indeks genap")
		fmt.Println("4. Tampilkan elemen indeks kelipatan x")
		fmt.Println("5. Hapus elemen")
		fmt.Println("6. Hitung rata-rata")
		fmt.Println("7. Hitung standar deviasi")
		fmt.Println("8. Hitung frekuensi")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih menu (0-8): ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 0:
			return
		case 1:
			arrOps.displayAll()
		case 2:
			arrOps.displayOddIndices()
		case 3:
			arrOps.displayEvenIndices()
		case 4:
			arrOps.displayMultipleIndices()
		case 5:
			arrOps.deleteElement()
		case 6:
			arrOps.calculateAverage()
		case 7:
			arrOps.calculateStdDev()
		case 8:
			arrOps.calculateFrequency()
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
