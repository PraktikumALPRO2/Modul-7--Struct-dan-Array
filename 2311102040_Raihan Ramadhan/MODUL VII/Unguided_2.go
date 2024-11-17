package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen dalam array: ")
	fmt.Scan(&n)

	array := make([]int, n)
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	for {
		var choice int
		fmt.Println("\nMenu:")
		fmt.Println("1. Tampilkan semua elemen")
		fmt.Println("2. Tampilkan elemen pada indeks ganjil")
		fmt.Println("3. Tampilkan elemen pada indeks genap")
		fmt.Println("4. Tampilkan elemen pada indeks kelipatan x")
		fmt.Println("5. Hapus elemen pada indeks tertentu")
		fmt.Println("6. Tampilkan rata-rata elemen")
		fmt.Println("7. Tampilkan deviasi standar elemen")
		fmt.Println("8. Tampilkan frekuensi angka tertentu")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			tampilkanSemua(array)
		case 2:
			tampilkanIndeksGanjil(array)
		case 3:
			tampilkanIndeksGenap(array)
		case 4:
			var x int
			fmt.Print("Masukkan nilai x: ")
			fmt.Scan(&x)
			tampilkanKelipatanX(array, x)
		case 5:
			var index int
			fmt.Print("Masukkan indeks untuk dihapus: ")
			fmt.Scan(&index)
			array = hapusPadaIndeks(array, index)
			tampilkanSemua(array)
		case 6:
			fmt.Printf("Rata-rata: %.2f\n", rataRata(array))
		case 7:
			fmt.Printf("Deviasi Standar: %.2f\n", deviasiStandar(array))
		case 8:
			var number int
			fmt.Print("Masukkan angka untuk mencari frekuensinya: ")
			fmt.Scan(&number)
			fmt.Printf("Frekuensi angka %d: %d\n", number, frekuensi(array, number))
		case 9:
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func tampilkanSemua(array []int) {
	fmt.Println("Elemen array:", array)
}

func tampilkanIndeksGanjil(array []int) {
	fmt.Print("Elemen pada indeks ganjil: ")
	for i := 1; i < len(array); i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

func tampilkanIndeksGenap(array []int) {
	fmt.Print("Elemen pada indeks genap: ")
	for i := 0; i < len(array); i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

func tampilkanKelipatanX(array []int, x int) {
	fmt.Printf("Elemen pada indeks yang merupakan kelipatan dari %d: ", x)
	for i := x; i < len(array); i += x {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

func hapusPadaIndeks(array []int, index int) []int {
	if index < 0 || index >= len(array) {
		fmt.Println("Indeks tidak valid.")
		return array
	}
	return append(array[:index], array[index+1:]...)
}

func rataRata(array []int) float64 {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return float64(sum) / float64(len(array))
}

func deviasiStandar(array []int) float64 {
	avg := rataRata(array)
	var varianceSum float64
	for _, value := range array {
		varianceSum += math.Pow(float64(value)-avg, 2)
	}
	return math.Sqrt(varianceSum / float64(len(array)))
}

func frekuensi(array []int, number int) int {
	count := 0
	for _, value := range array {
		if value == number {
			count++
		}
	}
	return count
}
