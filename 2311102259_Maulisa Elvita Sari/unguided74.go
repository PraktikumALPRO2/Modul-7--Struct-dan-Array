package main

import (
	"fmt"
)

const NMAX = 127

type Tabel [NMAX]rune

func isArray(t *Tabel, n *int) {
	var input string
	fmt.Print("Masukkan teks (akhiri dengan TITIK '.'): ")
	fmt.Scanln(&input)

	*n = 0
	for i := 0; i < len(input) && input[i] != '.'; i++ {
		(*t)[*n] = rune(input[i])
		*n++
	}
}

func cetakArray(t Tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *Tabel, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}

func palindrom(t Tabel, n int) bool {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if t[i] != t[j] {
			return false
		}
	}
	return true
}

func main() {
	var tab Tabel
	var n int

	isArray(&tab, &n)

	fmt.Print("Teks: ")
	cetakArray(tab, n)

	if palindrom(tab, n) {
		fmt.Println("Palindrom? true")
	} else {
		fmt.Println("Palindrom? false")
	}

	balikanArray(&tab, n)

	fmt.Print("Reverse teks: ")
	cetakArray(tab, n)
}
