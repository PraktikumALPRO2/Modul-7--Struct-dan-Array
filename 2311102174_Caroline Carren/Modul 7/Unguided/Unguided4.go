// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

const NMAX int = 127 // Konstanta untuk batas maksimum jumlah karakter dalam array

// Tipe data tabel adalah array yang menampung hingga 127 karakter bertipe rune (karakter Unicode)
type tabel [NMAX]rune

// Fungsi untuk mengisi array dengan karakter-karakter yang dimasukkan oleh pengguna
func isiArray(t *tabel, n int) {
	/* I.S. Data tersedia dalam piranti masukan
	   F.S. Array t berisi sejumlah n karakter yang dimasukkan user,
	   Proses input berhenti jika karakter yang dimasukkan adalah TITIK ('.') dan n <= NMAX */
	var i int
	var karakter rune
	fmt.Println("Masukkan karakter : ")
	for i = 0; i < n; i++ {
		// Memasukkan karakter satu per satu
		fmt.Scanf("%c", &karakter)
		// Jika karakter yang dimasukkan adalah titik, berhenti memasukkan data
		if karakter == '.' {
			break
		}
		// Menyimpan karakter ke dalam array
		t[i] = karakter
	}
}

// Fungsi untuk mencetak seluruh isi array ke layar
func cetakArray(t tabel, n int) {
	/* I.S. Terdefinisi array t yang berisi sejumlah n karakter
	   F.S. n karakter dalam array muncul di layar */
	for i := 0; i < n; i++ {
		// Mencetak karakter satu per satu
		fmt.Printf("%c", t[i])
	}
	// Menambahkan baris baru setelah mencetak karakter
	fmt.Println()
}

// Fungsi untuk membalikkan urutan karakter dalam array
func balikanArray(t *tabel, n int) {
	/* I.S. Terdefinisi array t yang berisi sejumlah n karakter
	   F.S. Urutan isi array t terbalik */
	for i := 0; i < n/2; i++ {
		// Menyimpan nilai sementara untuk pertukaran
		temp := t[i]
		// Melakukan pertukaran posisi karakter
		t[i] = t[n-i-1]
		t[n-i-1] = temp
	}
}

// Fungsi untuk memeriksa apakah array membentuk palindrom
func palindrom(t tabel, n int) bool {
	/* Mengembalikan true apabila susunan karakter di dalam t membentuk palindrom,
	   dan false apabila sebaliknya. Petunjuk: Manfaatkan prosedur balikanArray */
	var t2 tabel
	// Menyalin array t ke array t2
	copy(t2[:], t[:])
	// Membalikkan array t2
	balikanArray(&t2, n)

	// Memeriksa apakah array t dan t2 sama
	for i := 0; i < n; i++ {
		if t2[i] != t[i] {
			// Jika ada perbedaan, berarti bukan palindrom
			return false
		}
	}
	// Jika semua elemen sama, berarti palindrom
	return true
}

func main() {
	// Deklarasi array untuk menyimpan karakter dan jumlah elemen
	var tab tabel
	var n int
	// Meminta pengguna untuk memasukkan jumlah karakter
	fmt.Println("Masukkan jumlah karakter : ")
	fmt.Scanln(&n)
	// Mengisi array dengan karakter-karakter yang dimasukkan
	isiArray(&tab, n)
	// Menampilkan karakter-karakter yang dimasukkan
	fmt.Println("Karakter yang dimasukkan adalah:")
	cetakArray(tab, n)

	// Memeriksa apakah array membentuk palindrom
	if palindrom(tab, n) {
		// Jika ya, tampilkan pesan palindrom
		fmt.Println("Array ini adalah palindrom.")
	} else {
		// Jika tidak, tampilkan pesan bukan palindrom
		fmt.Println("Array ini bukan palindrom.")
	}
}
