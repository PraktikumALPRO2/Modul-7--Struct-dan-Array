package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 127

type tabel [NMAX]rune

// Fungsi untuk mengisi array dengan karakter
func isiArray(t *tabel, n *int) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Teks: ")
	scanner.Scan()
	text := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), " ", ""), ".", ""), ",", "")
	*n = len([]rune(text))
	for i, char := range text {
		if i < NMAX {
			t[i] = char
		}
	}
}

// Fungsi untuk mencetak isi array
func cetakArray(t tabel, n int) {
	fmt.Print("Reverse teks: ")
	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// Fungsi untuk mengecek apakah array membentuk palindrom
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
	var m int

	// Input dan proses array
	isiArray(&tab, &m)

	// Tampilkan array yang dibalik
	cetakArray(tab, m)

	// Cek palindrom
	fmt.Print("Palindrom? ")
	fmt.Println(palindrom(tab, m))
}