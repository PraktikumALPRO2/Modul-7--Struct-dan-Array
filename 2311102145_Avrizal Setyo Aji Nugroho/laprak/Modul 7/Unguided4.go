package main

import (
	"fmt"
)

const NMAX = 127

type tabel_145 struct {
	tab [NMAX]rune
}

func isiArray(t *tabel_145, n *int) {
	var input string
	fmt.Println("Masukkan teks (akhiri dengan titik '.'): ")
	fmt.Scanln(&input)

	for i, char := range input {
		if char == '.' || i >= NMAX {
			break
		}
		t.tab[i] = char
		*n++
	}
}

func cetakArray(t tabel_145, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t.tab[i])
	}
	fmt.Println()
}

func balikanArray(t tabel_145, n int) tabel_145 {
	var reversed tabel_145
	for i := 0; i < n; i++ {
		reversed.tab[i] = t.tab[n-i-1]
	}
	return reversed
}

func palindrom(t tabel_145, n int) bool {
	for i := 0; i < n/2; i++ {
		if t.tab[i] != t.tab[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel_145
	var n int

	isiArray(&tab, &n)

	fmt.Print("Teks: ")
	cetakArray(tab, n)

	reversedTab := balikanArray(tab, n)
	fmt.Print("Reverse teks: ")
	cetakArray(reversedTab, n)

	if palindrom(tab, n) {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}
}
