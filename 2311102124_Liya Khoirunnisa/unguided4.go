/* Liya Khoirunnisa - 2311102124 */
package main

import "fmt"

// Konstanta maksimum array
const NMAX int = 127

// Tipe data array menyimpan karakter
type tabel [NMAX]rune

// Fungsi untuk mengisi array
func isiArray(t *tabel) int {
	// Deklarasi variabel
	var input string
	
	// Input teks dari pengguna
	fmt.Print("Teks: ")
	fmt.Scanln(&input)

	// Jika input adalah titik
	if input == "." {
		return 0 
	}

	// Panjang input
	n := len(input)

	// Mengisi array dengan karakter inputan pengguna
	for i := 0; i < n && i < NMAX; i++ {
		(*t)[i] = rune(input[i])
	}
	return n
}

// Fungsi untuk mencetak array sebagai string
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]))
	}
	fmt.Println()
}

// Fungsi untuk membalikkan isi array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

// Fungsi untuk memeriksa apakah array merupakan palindrom
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	// Deklarasi variabel
	var tab tabel
	var m int

	// Mengisi array
	m = isiArray(&tab) 

	// Jika input adalah titik
	if m == 0 {
		return 
	}

	// Membalikkan isi array
	balikanArray(&tab, m) 

	// Menampilkan teks yang sudah dibalik
	fmt.Print("Reverse teks: ")
	cetakArray(tab, m)

	// Mengecek teks apakah palindrom
	if palindrom(tab, m) {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}
}