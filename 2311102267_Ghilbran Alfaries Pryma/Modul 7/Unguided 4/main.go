package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

// Fungsi untuk mengisi array
func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0
	fmt.Println("Masukkan karakter (akhiri dengan '.' atau maksimal 127 karakter):")
	for *n < NMAX {
		fmt.Scanf("%c", &ch)
		if ch == '.' {
			break
		}
		(*t)[*n] = ch
		*n++
	}
}

// Fungsi untuk mencetak array
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// Fungsi untuk membalikkan array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

// Fungsi untuk memeriksa apakah array adalah palindrom
func palindrom(t tabel, n int) bool {
	var reversed tabel
	copy(reversed[:n], t[:n])
	balikanArray(&reversed, n)
	for i := 0; i < n; i++ {
		if t[i] != reversed[i] {
			return false
		}
	}
	return true
}

// Fungsi utama
func main() {
	var tab tabel
	var n int

	// Mengisi array
	isiArray(&tab, &n)

	// Menampilkan teks asli
	fmt.Print("Teks : ")
	cetakArray(tab, n)

	// Membalikkan isi array
	balikanArray(&tab, n)
	fmt.Print("Reverse teks : ")
	cetakArray(tab, n)

	// Memeriksa apakah palindrom
	if palindrom(tab, n) {
		fmt.Println("Teks tersebut adalah palindrom.")
	} else {
		fmt.Println("Teks tersebut bukan palindrom.")
	}
}
