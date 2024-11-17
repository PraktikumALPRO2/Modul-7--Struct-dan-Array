/* Liya Khoirunnisa - 2311102124 */
package main

import (
	"fmt"
	"math" // Untuk operasi matematika
)

// Fungsi untuk menampilkan isi array
func tampilkanArray(data []int) {
	fmt.Println("Isi array:", data)
}

// Fungsi untuk menampilkan elemen dengan indeks ganjil
func tampilkanIndeksGanjil(data []int) {
	fmt.Print("Elemen berindeks ganjil: ")
	for i := 1; i < len(data); i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen dengan indeks genap
func tampilkanIndeksGenap(data []int) {
	fmt.Print("Elemen berindeks genap: ")
	for i := 0; i < len(data); i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen dengan indeks kelipatan x
func tampilkanKelipatan(data []int, x int) {
	fmt.Print("Elemen dengan indeks kelipatan ", x, ": ")
	for i := 0; i < len(data); i++ {
		if i%x == 0 {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()
}

// Fungsi untuk menghapus elemen pada indeks tertentu
func hapusIndeks(data *[]int, indeks int) {
	if indeks < 0 || indeks >= len(*data) {
		fmt.Println("Indeks tidak valid.")
		return
	}
	*data = append((*data)[:indeks], (*data)[indeks+1:]...)
	fmt.Println("Elemen pada indeks ke-", indeks, "telah dihapus.")
	tampilkanArray(*data)
}

// Fungsi untuk menghitung rata-rata
func hitungRataRata(data []int) float64 {
	if len(data) == 0 {
		return 0
	}
	jumlah := 0
	for _, nilai := range data {
		jumlah += nilai
	}
	return float64(jumlah) / float64(len(data))
}

// Fungsi untuk menghitung simpangan baku
func hitungSimpanganBaku(data []int) float64 {
	if len(data) == 0 {
		return 0
	}
	rataRata := hitungRataRata(data)
	jumlah := 0.0
	for _, nilai := range data {
		jumlah += math.Pow(float64(nilai)-rataRata, 2)
	}
	return math.Sqrt(jumlah / float64(len(data)))
}

// Fungsi untuk menghitung frekuensi dari suatu bilangan
func frekuensi(data []int, nilai int) int {
	hitung := 0
	for _, v := range data {
		if v == nilai {
			hitung++
		}
	}
	return hitung
}

func main() {
	// Deklarasi variabel
	var n, pilihan int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n) // Meminta input jumlah elemen

	// Inisialisai array dengan ukuran n
	data := make([]int, n)
	fmt.Println("Masukkan elemen array:")

	// Perulangan meminta input elemen array
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&data[i])
	}

	// Menu
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tampilkan isi array")
		fmt.Println("2. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("3. Tampilkan elemen dengan indeks genap")
		fmt.Println("4. Tampilkan elemen dengan indeks kelipatan x")
		fmt.Println("5. Hapus elemen pada indeks tertentu")
		fmt.Println("6. Tampilkan rata-rata")
		fmt.Println("7. Tampilkan simpangan baku")
		fmt.Println("8. Tampilkan frekuensi dari suatu bilangan")
		fmt.Println("9. Keluar")

		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan) // Input pilihan menu

		// Switch case pilihan menu
		switch pilihan {
		case 1:
			// Menampilkan isi array
			tampilkanArray(data)
		case 2:
			// Menampilkan elemen berindeks ganjil
			tampilkanIndeksGanjil(data)
		case 3:
			// Menampilkan elemen berindeks genap
			tampilkanIndeksGenap(data)
		case 4:
			var x int
			fmt.Print("Masukkan bilangan x: ")
			fmt.Scan(&x) // Input kelipatan x

			// Menampilkan elemen dengan indeks kelipatan
			tampilkanKelipatan(data, x)
		case 5:
			var indeks int
			fmt.Print("Masukkan indeks yang ingin dihapus: ")
			fmt.Scan(&indeks) // Input indeks yang ingin di hapus
			
			// Menghapus elemen pada indeks yang diinputkan
			hapusIndeks(&data, indeks) 
		case 6:
			// Menghitung rata-rata
			rataRata := hitungRataRata(data)
			fmt.Printf("Rata-rata: %.2f\n", rataRata)
		case 7:
			// Menghitung simpangan baku
			simpanganBaku := hitungSimpanganBaku(data)
			fmt.Printf("Simpangan baku: %.2f\n", simpanganBaku)
		case 8:
			// Menghitung frekuensi nilai
			var nilai int
			fmt.Print("Masukkan bilangan untuk frekuensi: ")
			fmt.Scan(&nilai)
			frekuensiVal := frekuensi(data, nilai)
			fmt.Printf("Frekuensi dari %d: %d\n", nilai, frekuensiVal)
		case 9:
			// Keluar 
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}
