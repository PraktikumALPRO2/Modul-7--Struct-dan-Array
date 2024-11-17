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
	fmt.Print("Teks : ")
	scanner.Scan()
	text := scanner.Text()

	// Menghilangkan spasi dan tanda baca
	text = strings.ReplaceAll(text, " ", "")
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, ",", "")

	*n = 0
	// Mengisi array dengan karakter dari input
	for _, char := range text {
		if *n < NMAX {
			t[*n] = char
			*n++
		}
	}
}

// Fungsi untuk mencetak isi array
func cetakArray(t tabel, n int) {
	fmt.Print("Reverse teks : ")
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// Fungsi untuk membalikkan array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

// Fungsi untuk mengecek apakah array membentuk palindrom
func palindrom(t tabel, n int) bool {
	// Membuat salinan array untuk dibalik
	var temp tabel
	copy(temp[:], t[:n])

	// Balik array salinan
	for i := 0; i < n/2; i++ {
		temp[i], temp[n-1-i] = temp[n-1-i], temp[i]
	}

	// Bandingkan array asli dengan array yang sudah dibalik
	for i := 0; i < n; i++ {
		if t[i] != temp[i] {
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
	var reversed tabel
	copy(reversed[:], tab[:m])
	balikanArray(&reversed, m)
	cetakArray(reversed, m)

	// Cek palindrom
	fmt.Print("Palindrom ? ")
	if palindrom(tab, m) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
