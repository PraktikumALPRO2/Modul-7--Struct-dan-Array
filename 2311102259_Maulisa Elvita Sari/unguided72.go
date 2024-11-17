package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&n)

	array := make([]int, n)
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	fmt.Println("\na. Isi array:")
	fmt.Println(array)

	fmt.Println("\nb. Elemen dengan indeks ganjil:")
	for i := 1; i < len(array); i += 2 {
		fmt.Printf("array[%d] = %d\n", i, array[i])
	}

	fmt.Println("\nc. Elemen dengan indeks genap:")
	for i := 0; i < len(array); i += 2 {
		fmt.Printf("array[%d] = %d\n", i, array[i])
	}

	var x int
	fmt.Print("\nd. Masukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	for i := 0; i < len(array); i++ {
		if i%x == 0 {
			fmt.Printf("array[%d] = %d\n", i, array[i])
		}
	}

	var idx int
	fmt.Print("\ne. Masukkan indeks elemen yang akan dihapus: ")
	fmt.Scan(&idx)
	if idx >= 0 && idx < len(array) {
		array = append(array[:idx], array[idx+1:]...)
		fmt.Println("Isi array setelah elemen dihapus:")
		fmt.Println(array)
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	sum := 0
	for _, v := range array {
		sum += v
	}
	rataRata := float64(sum) / float64(len(array))
	fmt.Printf("\nf. Rata-rata elemen array: %.2f\n", rataRata)

	var deviasiSum float64
	for _, v := range array {
		deviasiSum += math.Pow(float64(v)-rataRata, 2)
	}
	standarDeviasi := math.Sqrt(deviasiSum / float64(len(array)))
	fmt.Printf("g. Standar deviasi elemen array: %.2f\n", standarDeviasi)

	var target int
	fmt.Print("\nh. Masukkan bilangan yang ingin dihitung frekuensinya: ")
	fmt.Scan(&target)
	frekuensi := 0
	for _, v := range array {
		if v == target {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi bilangan %d dalam array: %d\n", target, frekuensi)
}
