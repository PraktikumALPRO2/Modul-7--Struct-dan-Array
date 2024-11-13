// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import (
	"fmt"
	"math"
)

func main() {
	// Deklarasi array dengan kapasitas 10 untuk menampung elemen-elemen array
	array := make([]int, 10)

	// Meminta user memasukkan jumlah elemen yang ingin dimasukkan ke dalam array
	var n int
	fmt.Print("Masukkan jumlah elemen (maksimal 10): ")
	fmt.Scanln(&n)

	// Memastikan jumlah elemen tidak melebihi kapasitas array
	if n > 10 {
		n = 10
		fmt.Println("Jumlah elemen dibatasi maksimal 10. Menggunakan 10 elemen.")
	}

	// Meminta user untuk memasukkan nilai-nilai elemen ke dalam array
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		var elemen int
		fmt.Scanln(&elemen)
		array[i] = elemen
	}

	for {
		// Menampilkan pilihan menu operasi yang bisa dilakukan pada array
		fmt.Println("\nPilihan menu:")
		fmt.Println("1. Menampilkan keseluruhan isi dari array")
		fmt.Println("2. Menampilkan elemen-elemen array dengan indeks ganjil saja")
		fmt.Println("3. Menampilkan elemen-elemen array dengan indeks genap saja")
		fmt.Println("4. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x")
		fmt.Println("5. Menghapus elemen array pada indeks tertentu")
		fmt.Println("6. Menampilkan rata-rata dari bilangan yang ada di dalam array")
		fmt.Println("7. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array")
		fmt.Println("8. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array")
		fmt.Println("9. Keluar")

		// Meminta user untuk memilih salah satu menu
		var pilihan int
		fmt.Print("\nMasukkan pilihan: ")
		fmt.Scanln(&pilihan)

		// Memproses pilihan yang dimasukkan oleh user
		switch pilihan {
		case 1:
			tampilkanKeseluruhanIsiArray(array, n)
		case 2:
			tampilkanElemenIndeksGanjil(array, n)
		case 3:
			tampilkanElemenIndeksGenap(array, n)
		case 4:
			tampilkanElemenKelipatan(array, n)
		case 5:
			n = hapusElemenArray(array, n) // Update n after deletion
		case 6:
			tampilkanRataRata(array, n)
		case 7:
			tampilkanStandarDeviasi(array, n)
		case 8:
			tampilkanFrekuensi(array, n)
		case 9:
			fmt.Println("Terima kasih! Program selesai.")
			return // Exit the program
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// Fungsi untuk menampilkan seluruh isi array
func tampilkanKeseluruhanIsiArray(array []int, n int) {
	fmt.Println("Isi array:")
	for i := 0; i < n; i++ {
		fmt.Println(array[i])
	}
}

// Fungsi untuk menampilkan elemen array dengan indeks ganjil
func tampilkanElemenIndeksGanjil(array []int, n int) {
	fmt.Println("Elemen dengan indeks ganjil:")
	for i := 1; i < n; i += 2 {
		fmt.Println(array[i])
	}
}

// Fungsi untuk menampilkan elemen array dengan indeks genap
func tampilkanElemenIndeksGenap(array []int, n int) {
	fmt.Println("Elemen dengan indeks genap:")
	for i := 0; i < n; i += 2 {
		fmt.Println(array[i])
	}
}

// Fungsi untuk menampilkan elemen array dengan indeks kelipatan bilangan x
func tampilkanElemenKelipatan(array []int, n int) {
	var x int
	fmt.Print("Masukkan bilangan kelipatan: ")
	fmt.Scanln(&x)
	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Println(array[i])
		}
	}
}

// Fungsi untuk menghapus elemen array pada indeks tertentu
func hapusElemenArray(array []int, n int) int {
	var indeks int
	fmt.Print("Masukkan indeks elemen yang ingin dihapus: ")
	fmt.Scanln(&indeks)
	if indeks < 0 || indeks >= n {
		fmt.Println("Indeks tidak valid")
		return n // Return n unchanged
	}
	// Menggeser elemen-elemen setelah elemen yang dihapus untuk menghapusnya
	for i := indeks; i < n-1; i++ {
		array[i] = array[i+1]
	}
	n-- // Mengurangi jumlah elemen
	fmt.Println("Elemen pada indeks", indeks, "telah dihapus")
	return n // Return updated n
}

// Fungsi untuk menampilkan rata-rata dari elemen array
func tampilkanRataRata(array []int, n int) {
	var total int
	// Menghitung total dari seluruh elemen dalam array
	for i := 0; i < n; i++ {
		total += array[i]
	}
	// Menghitung rata-rata
	rataRata := float64(total) / float64(n)
	fmt.Printf("Rata-rata: %.2f\n", rataRata)
}

// Fungsi untuk menghitung dan menampilkan standar deviasi (simpangan baku) dari elemen array
func tampilkanStandarDeviasi(array []int, n int) {
	var total, rataRata float64
	// Menghitung total nilai dalam array
	for i := 0; i < n; i++ {
		total += float64(array[i])
	}
	// Menghitung rata-rata
	rataRata = total / float64(n)

	var totalDeviasi float64
	// Menghitung deviasi dari setiap elemen terhadap rata-rata
	for i := 0; i < n; i++ {
		totalDeviasi += math.Pow(float64(array[i])-rataRata, 2)
	}
	// Menghitung standar deviasi
	standarDeviasi := math.Sqrt(totalDeviasi / float64(n))
	fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi)
}

// Fungsi untuk menampilkan frekuensi kemunculan suatu bilangan di dalam array
func tampilkanFrekuensi(array []int, n int) {
	var bilangan int
	// Meminta user untuk memasukkan bilangan yang ingin dihitung frekuensinya
	fmt.Print("Masukkan bilangan untuk mencari frekuensi: ")
	fmt.Scanln(&bilangan)
	frekuensi := 0
	// Menghitung frekuensi kemunculan bilangan di dalam array
	for i := 0; i < n; i++ {
		if array[i] == bilangan {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi dari bilangan %d: %d\n", bilangan, frekuensi)
}
