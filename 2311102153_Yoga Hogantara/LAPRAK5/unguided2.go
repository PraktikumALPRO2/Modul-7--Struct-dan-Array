package main

import (
	"fmt"
	"math"
)

func cetakArray(array []int) {
	fmt.Print("Array: ")
	for _, v := range array {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
// INDEKS GANJIL
func cetakganjil(array []int) {
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < len(array); i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}
// INDEKS GENAP
func cetakgenap(array []int) {
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(array); i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

func hapusindeks(array []int, index int) []int {
	if index >= 0 && index < len(array) {
		return append(array[:index], array[index+1:]...)
	}
	fmt.Println("Indeks tidak valid.")
	return array
}

func rata2(array []int) float64 {
	if len(array) == 0 {
		fmt.Println("array kosong.")
		return 0
	}
	sum := 0
	for _, v := range array {
		sum += v
	}
	return float64(sum) / float64(len(array))
}
func standarDeviasi(array []int) float64 {
	if len(array) == 0 {
		fmt.Println("Array kosong.")
		return 0
	}
	mean := rata2(array)
	var simpanganbaku float64
	for _, v := range array {
		simpanganbaku += math.Pow(float64(v)-mean, 2)
	}
	return math.Sqrt(simpanganbaku / float64(len(array)))
}

func frekuensi(array []int, num int) int {
	count := 0
	for _, f := range array {
		if f == num {
			count++
		}
	}
	return count
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen dalam array: ")
	fmt.Scan(&n)

	array := make([]int, n)
	fmt.Print("Masukkan elemen-elemen array: ")
	for i := range array {
		fmt.Scan(&array[i])
	}

	for {
		fmt.Println("\nPilih operasi:")
		fmt.Println("1. elemen array")
		fmt.Println("2. elemen dengan indeks ganjil")
		fmt.Println("3. elemen dengan indeks genap")
		fmt.Println("4. Hapus elemen pada indeks ")
		fmt.Println("5. rata-rata ")
		fmt.Println("6. standar deviasi")
		fmt.Println("7. frekuensi")
		fmt.Println("8. keluar")
		fmt.Print("Pilihan Anda: ")
		var pilih153 int
		fmt.Scan(&pilih153)

		switch pilih153 {
		case 1:
			cetakArray(array)
		case 2:
			cetakganjil(array)
		case 3:
			cetakgenap(array)
		case 4:
			var index int
			fmt.Print("Masukkan indeks elemen yang ingin dihapus: ")
			fmt.Scan(&index)
			array = hapusindeks(array, index)
			cetakArray(array)
		case 5:
			rata := rata2(array)
			if rata != 0 {
				fmt.Printf("Rata-rata: %.2f\n", rata)
			}
		case 6:
			sd := standarDeviasi(array)
			if sd != 0 {
				fmt.Printf("Standar deviasi: %.2f\n", sd)
			}
		case 7:
			var num int
			fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
			fmt.Scan(&num)
			frek := frekuensi(array, num)
			fmt.Printf("Frekuensi bilangan %d adalah %d kali\n", num, frek)
		case 8:
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
