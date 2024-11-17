package main

import "fmt"

const NMAX = 127

type Tabel [NMAX]rune

func isiArray(t *Tabel, n *int) {
	fmt.Print("Teks: ")
	var input string
	fmt.Scanln(&input)

	*n = len(input) 
	for i := 0; i < *n; i++ {
		t[i] = rune(input[i])
	}
}

func cetakArray(t Tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikkanArray(t *Tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func palindrom(t Tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var t Tabel
	var n int

	isiArray(&t, &n)

	fmt.Print("TEKS: ")
	cetakArray(t, n)

	balikkanArray(&t, n)
	fmt.Print("ARRAY TERBALIK: ")
	cetakArray(t, n)

	fmt.Print("PALINDROM ")
	if palindrom(t, n) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}