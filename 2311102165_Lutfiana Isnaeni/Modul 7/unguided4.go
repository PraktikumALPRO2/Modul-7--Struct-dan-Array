// Lutfiana Isnaeni Lthifah
// 2311102165

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Konstanta dan tipe data
const NMAX int = 127

type tabel [NMAX]rune

// Fungsi untuk mengisi array dari input pengguna
func isiArray(t *tabel, n *int) {
	var input string
	fmt.Println("Masukkan teks, akhiri dengan karakter titik ('.'):")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	for _, char := range input {
		if char == '.' || *n >= NMAX {
			break
		}
		t[*n] = char
		*n++
	}
}

// Fungsi untuk mencetak isi array
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// Fungsi untuk membalikkan isi array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

// Fungsi untuk memeriksa apakah array membentuk palindrom
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	// Isi array dengan memanggil prosedur isiArray
	isiArray(&tab, &n)

	// Cetak isi array
	fmt.Print("Teks: ")
	cetakArray(tab, n)

	// Balikkan isi array
	balikanArray(&tab, n)

	// Cetak isi array setelah dibalik
	fmt.Print("Reverse teks: ")
	cetakArray(tab, n)

	// Periksa apakah array membentuk palindrom
	if palindrom(tab, n) {
		fmt.Println("Palindrom? true")
	} else {
		fmt.Println("Palindrom? false")
	}
}
