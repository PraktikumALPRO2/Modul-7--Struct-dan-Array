package main

import (
	"fmt"
	"unicode"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	fmt.Println("Masukkan karakter (tuliskan '.' untuk berhenti):")
	for *n < NMAX {
		var karakter rune
		fmt.Scanf("%c", &karakter)
		if unicode.IsSpace(karakter) {
			continue
		}
		if karakter == '.' {
			break
		}
		t[*n] = unicode.ToLower(karakter)
		(*n)++
	}
}

func cetakArray(t tabel, n int) {
	for indeks := 0; indeks < n; indeks++ {
		fmt.Printf("%c", t[indeks])
	}
	fmt.Println()
}

func balikanArray(awal tabel, terbalik *tabel, n int) {
	for i := 0; i < n; i++ {
		terbalik[i] = awal[n-1-i]
	}
}

func palindrom(awal tabel, terbalik tabel, n int) bool {
	for indeks := 0; indeks < n; indeks++ {
		if awal[indeks] != terbalik[indeks] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var tabkebalik tabel
	var jumlahKarakter int = 0

	isiArray(&tab, &jumlahKarakter)

	fmt.Print("Array: ")
	cetakArray(tab, jumlahKarakter)

	balikanArray(tab, &tabkebalik, jumlahKarakter)

	fmt.Print("Array yang sudah dibalik: ")
	cetakArray(tabkebalik, jumlahKarakter)

	if palindrom(tab, tabkebalik, jumlahKarakter) {
		fmt.Println("Array adalah palindrome.")
	} else {
		fmt.Println("Array bukan palindrome.")
	}
}
