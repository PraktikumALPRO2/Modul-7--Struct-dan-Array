package main

import (
	"fmt"
)

func main() {
	// Membuat array dengan kapasitas 10 untuk menampung elemen-elemen
	elemenArray := make([]int, 10)

	// Meminta input dari pengguna untuk jumlah elemen yang akan dimasukkan ke dalam array
	var jumlahElemen int
	fmt.Print("Masukkan jumlah elemen (maksimal 10): ")
	fmt.Scanln(&jumlahElemen)

	// Memastikan jumlah elemen tidak melebihi kapasitas array
	if jumlahElemen > 10 {
		jumlahElemen = 10
		fmt.Println("Jumlah elemen dibatasi maksimal 10. Menggunakan 10 elemen.")
	}

	// Meminta pengguna untuk memasukkan nilai elemen-elemen array
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < jumlahElemen; i++ {
		var elemen int
		fmt.Scanln(&elemen)
		elemenArray[i] = elemen
	}

	for {
		// Menampilkan menu pilihan operasi yang dapat dilakukan pada array
		fmt.Println("\nMenu pilihan:")
		fmt.Println("1. Menampilkan seluruh isi array")
		fmt.Println("2. Menampilkan elemen array dengan indeks ganjil")
		fmt.Println("3. Menampilkan elemen array dengan indeks genap")
		fmt.Println("4. Menampilkan elemen array dengan indeks kelipatan bilangan x")
		fmt.Println("5. Menghapus elemen array pada indeks tertentu")
		fmt.Println("6. Menampilkan rata-rata nilai dalam array")
		fmt.Println("7. Menampilkan standar deviasi dari nilai dalam array")
		fmt.Println("8. Menampilkan frekuensi kemunculan suatu bilangan dalam array")
		fmt.Println("9. Keluar")

		// Meminta input dari pengguna untuk memilih salah satu menu
		var pilihan int
		fmt.Print("\nMasukkan pilihan: ")
		fmt.Scanln(&pilihan)

		// Menangani pilihan yang dimasukkan pengguna
		switch pilihan {
		case 1:
			tampilkanSeluruhIsiArray(elemenArray, jumlahElemen)
		case 2:
			tampilkanElemenIndeksGanjil(elemenArray, jumlahElemen)
		case 3:
			tampilkanElemenIndeksGenap(elemenArray, jumlahElemen)
		case 4:
			tampilkanElemenKelipatan(elemenArray, jumlahElemen)
		case 5:
			jumlahElemen = hapusElemenArray(elemenArray, jumlahElemen) // Perbarui jumlah elemen setelah penghapusan
		case 6:
			tampilkanRataRata(elemenArray, jumlahElemen)
		case 7:
			tampilkanStandarDeviasi(elemenArray, jumlahElemen)
		case 8:
			tampilkanFrekuensi(elemenArray, jumlahElemen)
		case 9:
			fmt.Println("Terima kasih! Program selesai.")
			return // Keluar dari program
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// Fungsi untuk menampilkan seluruh isi array
func tampilkanSeluruhIsiArray(array []int, jumlahElemen int) {
	fmt.Println("Isi array:")
	for i := 0; i < jumlahElemen; i++ {
		fmt.Println(array[i])
	}
}

// Fungsi untuk menampilkan elemen array dengan indeks ganjil
func tampilkanElemenIndeksGanjil(array []int, jumlahElemen int) {
	fmt.Println("Elemen dengan indeks ganjil:")
	for i := 1; i < jumlahElemen; i += 2 {
		fmt.Println(array[i])
	}
}

// Fungsi untuk menampilkan elemen array dengan indeks genap
func tampilkanElemenIndeksGenap(array []int, jumlahElemen int) {
	fmt.Println("Elemen dengan indeks genap:")
	for i := 0; i < jumlahElemen; i += 2 {
		fmt.Println(array[i])
	}
}

// Fungsi untuk menampilkan elemen array dengan indeks kelipatan bilangan x
func tampilkanElemenKelipatan(array []int, jumlahElemen int) {
	var x int
	fmt.Print("Masukkan bilangan kelipatan: ")
	fmt.Scanln(&x)
	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	for i := 0; i < jumlahElemen; i++ {
		if i%x == 0 {
			fmt.Println(array[i])
		}
	}
}

// Fungsi untuk menghapus elemen array pada indeks tertentu
func hapusElemenArray(array []int, jumlahElemen int) int {
	var indeks int
	fmt.Print("Masukkan indeks elemen yang ingin dihapus: ")
	fmt.Scanln(&indeks)
	if indeks < 0 || indeks >= jumlahElemen {
		fmt.Println("Indeks tidak valid")
		return jumlahElemen // Kembalikan jumlahElemen tanpa perubahan
	}
	// Menggeser elemen-elemen setelah elemen yang dihapus untuk menghapusnya
	for i := indeks; i < jumlahElemen-1; i++ {
		array[i] = array[i+1]
	}
	jumlahElemen-- // Kurangi jumlah elemen
	fmt.Println("Elemen pada indeks", indeks, "telah dihapus")
	return jumlahElemen // Kembalikan jumlah elemen yang baru
}

// Fungsi untuk menampilkan rata-rata nilai dalam array
func tampilkanRataRata(array []int, jumlahElemen int) {
	var total int
	// Menghitung total dari semua elemen dalam array
	for i := 0; i < jumlahElemen; i++ {
		total += array[i]
	}
	// Menghitung rata-rata
	rataRata := float64(total) / float64(jumlahElemen)
	fmt.Printf("Rata-rata: %.2f\n", rataRata)
}

// Fungsi untuk menghitung dan menampilkan standar deviasi dari elemen array
func tampilkanStandarDeviasi(array []int, jumlahElemen int) {
	var total, rataRata float64
	// Menghitung total nilai dalam array
	for i := 0; i < jumlahElemen; i++ {
		total += float64(array[i])
	}
	// Menghitung rata-rata
	rataRata = total / float64(jumlahElemen)

	var totalDeviasi float64
	// Menghitung deviasi setiap elemen terhadap rata-rata secara manual
	for i := 0; i < jumlahElemen; i++ {
		deviasi := float64(array[i]) - rataRata
		totalDeviasi += deviasi * deviasi
	}
	// Menghitung standar deviasi secara manual
	standarDeviasi := totalDeviasi / float64(jumlahElemen)
	standarDeviasi = hitungAkarKuadrat(standarDeviasi)
	fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi)
}

// Fungsi untuk menghitung akar kuadrat secara manual (seperti math.Sqrt)
func hitungAkarKuadrat(x float64) float64 {
	// Metode estimasi akar kuadrat menggunakan metode Newton-Raphson
	if x < 0 {
		return 0
	}
	z := x / 2.0
	for i := 0; i < 10; i++ {
		z = (z + x/z) / 2.0
	}
	return z
}

// Fungsi untuk menampilkan frekuensi kemunculan suatu bilangan dalam array
func tampilkanFrekuensi(array []int, jumlahElemen int) {
	var bilangan int
	// Meminta input untuk bilangan yang frekuensinya ingin dihitung
	fmt.Print("Masukkan bilangan untuk mencari frekuensi: ")
	fmt.Scanln(&bilangan)
	frekuensi := 0
	// Menghitung frekuensi kemunculan bilangan dalam array
	for i := 0; i < jumlahElemen; i++ {
		if array[i] == bilangan {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi dari bilangan %d: %d\n", bilangan, frekuensi)
}
