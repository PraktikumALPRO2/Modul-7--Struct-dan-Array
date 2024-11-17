package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Konstanta dan tipe data
const NMAX235 int = 127

type tabel235 [NMAX235]rune

// Fungsi untuk mengisi array dari input pengguna
func isiArray235(t235 *tabel235, n235 *int) {
	var input235 string
	fmt.Println("Masukkan teks, akhiri dengan karakter titik ('.'):")
	reader235 := bufio.NewReader(os.Stdin)
	input235, _ = reader235.ReadString('\n')
	input235 = strings.TrimSpace(input235)

	for _, char235 := range input235 {
		if char235 == '.' || *n235 >= NMAX235 {
			break
		}
		t235[*n235] = char235
		*n235++
	}
}

// Fungsi untuk mencetak isi array
func cetakArray235(t235 tabel235, n235 int) {
	for i235 := 0; i235 < n235; i235++ {
		fmt.Printf("%c", t235[i235])
	}
	fmt.Println()
}

// Fungsi untuk membalikkan isi array
func balikanArray235(t235 *tabel235, n235 int) {
	for i235 := 0; i235 < n235/2; i235++ {
		t235[i235], t235[n235-1-i235] = t235[n235-1-i235], t235[i235]
	}
}

// Fungsi untuk memeriksa apakah array membentuk palindrom
func palindrom235(t235 tabel235, n235 int) bool {
	for i235 := 0; i235 < n235/2; i235++ {
		if t235[i235] != t235[n235-1-i235] {
			return false
		}
	}
	return true
}

func main() {
	var tab235 tabel235
	var n235 int

	// Isi array dengan memanggil prosedur isiArray235
	isiArray235(&tab235, &n235)

	// Cetak isi array
	fmt.Print("Teks: ")
	cetakArray235(tab235, n235)

	// Balikkan isi array
	balikanArray235(&tab235, n235)

	// Cetak isi array setelah dibalik
	fmt.Print("Reverse teks: ")
	cetakArray235(tab235, n235)

	// Periksa apakah array membentuk palindrom
	if palindrom235(tab235, n235) {
		fmt.Println("Palindrom? true")
	} else {
		fmt.Println("Palindrom? false")
	}
}
