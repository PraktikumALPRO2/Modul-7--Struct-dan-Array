 package main

import (
	"fmt"
	"strings"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	/* I.S. Data tersedia dalam piranti masukan
	   F.S. Array t berisi sejumlah n karakter yang dimasukkan user,
	   Proses input selama karakter bukanlah TITIK dan n <= NMAX */
	var input string
	fmt.Println("Masukkan kata atau teks (akhiri dengan titik):")
	fmt.Scanln(&input)

	// Hapus titik jika ada
	input = strings.TrimSuffix(input, ".")
	*n = len(input)

	for i := 0; i < *n; i++ {
		t[i] = rune(input[i]) // Mengisi array dengan karakter-karakter input
	}
}

func cetakArray(t tabel, n int) {
	/* I.S. Terdefinisi array t yang berisi sejumlah n karakter
	   F.S. n karakter dalam array muncul di layar */
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	/* I.S. Terdefinisi array t yang berisi sejumlah n karakter
	   F.S. Urutan isi array t terbalik */
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}

func isPalindrome(t tabel, n int) bool {
	/* Mengembalikan true apabila susunan karakter di dalam t membentuk palindrom,
	   dan false apabila sebaliknya. */
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	// Isi array tab dengan memanggil prosedur isiArray
	isiArray(&tab, &m)

	// Cetak isi array tab
	fmt.Print("Teks: ")
	cetakArray(tab, m)

	// Balikkan isi array tab dengan memanggil balikanArray
	fmt.Print("Reverse teks: ")
	balikanArray(&tab, m)
	cetakArray(tab, m)

	// Periksa apakah array adalah palindrom
	fmt.Print("Palindrome: ")
	if isPalindrome(tab, m) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
