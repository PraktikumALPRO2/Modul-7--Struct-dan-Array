package main

import (
	"bufio"
	"fmt"
	"os"
)

const NMAX int = 127

type tabel [NMAX]rune

// Fungsi untuk mengisi array
func isiArray(t *tabel, n *int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan teks: ")

	var char rune
	*n = 0
	for *n < NMAX {
		input, _, _ := reader.ReadRune()
		char = input
		if char == '.' {
			break
		}
		t[*n] = char
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

// Fungsi untuk membalik array
func balikanArray(t *tabel, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}

// Fungsi untuk mengecek apakah array adalah palindrome
func palindrome(t tabel, n int) bool {
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

	// Mengisi array dengan memanggil fungsi isiArray
	isiArray(&tab, &m)

	// Mencetak array asli
	fmt.Print("\nTeks: ")
	cetakArray(tab, m)

	// Membalik array
	balikanArray(&tab, m)

	// Mencetak array yang telah dibalik
	fmt.Print("Reverse teks: ")
	cetakArray(tab, m)

	// Mengecek apakah array adalah palindrome dan mencetak hasilnya
	isPalindrome := palindrome(tab, m)
	fmt.Printf("Palindrome? %t\n", isPalindrome)
}
