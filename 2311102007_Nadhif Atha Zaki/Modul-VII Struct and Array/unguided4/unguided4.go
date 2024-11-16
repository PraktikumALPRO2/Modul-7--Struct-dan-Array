package main

import "fmt"

const MaksKarakter int = 127 // Konstanta batas maksimum jumlah karakter dalam array

// Tipe data 'arrayKarakter' adalah array yang bisa menyimpan hingga 127 karakter bertipe rune (Unicode)
type arrayKarakter [MaksKarakter]rune

// Fungsi untuk mengisi array dengan karakter-karakter yang dimasukkan oleh pengguna
func inputArrayKarakter(array *arrayKarakter, jumlah int) {
	/* I.S. Data karakter tersedia dalam input
	   F.S. Array array terisi sejumlah jumlah karakter yang dimasukkan pengguna,
	   Proses berhenti jika karakter titik ('.') dimasukkan atau jumlah <= MaksKarakter */
	var indeks int
	var karakter rune
	fmt.Println("Masukkan karakter: ")
	for indeks = 0; indeks < jumlah; indeks++ {
		// Memasukkan karakter satu per satu
		fmt.Scanf("%c", &karakter)
		// Jika karakter yang dimasukkan adalah titik, berhenti
		if karakter == '.' {
			break
		}
		// Menyimpan karakter ke dalam array
		array[indeks] = karakter
	}
}

// Fungsi untuk menampilkan seluruh isi array ke layar
func tampilkanKarakter(array arrayKarakter, jumlah int) {
	/* I.S. Terdefinisi array yang berisi sejumlah jumlah karakter
	   F.S. n karakter dalam array ditampilkan di layar */
	for i := 0; i < jumlah; i++ {
		// Mencetak karakter satu per satu
		fmt.Printf("%c", array[i])
	}
	// Menambahkan baris baru setelah mencetak karakter
	fmt.Println()
}

// Fungsi untuk membalikkan urutan karakter dalam array
func balikkanUrutan(array *arrayKarakter, jumlah int) {
	/* I.S. Terdefinisi array yang berisi sejumlah jumlah karakter
	   F.S. Urutan isi array dibalik */
	for i := 0; i < jumlah/2; i++ {
		// Menyimpan nilai sementara untuk pertukaran
		sementara := array[i]
		// Melakukan pertukaran posisi karakter
		array[i] = array[jumlah-i-1]
		array[jumlah-i-1] = sementara
	}
}

// Fungsi untuk memeriksa apakah array membentuk palindrom
func cekApakahPalindrom(array arrayKarakter, jumlah int) bool {
	/* Mengembalikan true jika susunan karakter di dalam array membentuk palindrom,
	   dan false jika tidak. Petunjuk: Gunakan fungsi balikkanUrutan */
	var arraySalinan arrayKarakter
	// Menyalin array ke arraySalinan
	copy(arraySalinan[:], array[:])
	// Membalikkan arraySalinan
	balikkanUrutan(&arraySalinan, jumlah)

	// Memeriksa apakah array dan arraySalinan sama
	for i := 0; i < jumlah; i++ {
		if arraySalinan[i] != array[i] {
			// Jika ada perbedaan, berarti bukan palindrom
			return false
		}
	}
	// Jika semua elemen sama, berarti palindrom
	return true
}

func main() {
	// Deklarasi array untuk menyimpan karakter dan jumlah karakter yang dimasukkan
	var array arrayKarakter
	var jumlahKarakter int
	// Meminta pengguna untuk memasukkan jumlah karakter
	fmt.Println("Masukkan jumlah karakter: ")
	fmt.Scanln(&jumlahKarakter)
	// Mengisi array dengan karakter-karakter yang dimasukkan
	inputArrayKarakter(&array, jumlahKarakter)
	// Menampilkan karakter-karakter yang dimasukkan
	fmt.Println("Karakter yang dimasukkan adalah:")
	tampilkanKarakter(array, jumlahKarakter)

	// Membalikkan urutan array dan menampilkan reverse text
	balikkanUrutan(&array, jumlahKarakter)
	fmt.Println("Teks terbalik (reverse) adalah:")
	tampilkanKarakter(array, jumlahKarakter)

	// Memeriksa apakah array membentuk palindrom
	if cekApakahPalindrom(array, jumlahKarakter) {
		// Jika ya, tampilkan pesan palindrom
		fmt.Println("Array ini adalah palindrom.")
	} else {
		// Jika tidak, tampilkan pesan bukan palindrom
		fmt.Println("Array ini bukan palindrom.")
	}
}
