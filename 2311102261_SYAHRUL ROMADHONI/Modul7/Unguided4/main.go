package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 127

type Tabel struct {
	tab [NMAX]rune
	m   int
}

func isiArray(t *Tabel) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan teks (akhiri dengan '.'): ")
	input, _ := reader.ReadString('.')
	input = strings.TrimSuffix(input, ".")

	for i, r := range input {
		if i >= NMAX {
			break
		}
		t.tab[i] = r
		t.m++
	}
}

func cetakArray(t Tabel) {
	fmt.Print("Array: ")
	for i := 0; i < t.m; i++ {
		fmt.Print(string(t.tab[i]))
	}
	fmt.Println()
}

func balikkanArray(t *Tabel) {
	for i, j := 0, t.m-1; i < j; i, j = i+1, j-1 {
		t.tab[i], t.tab[j] = t.tab[j], t.tab[i]
	}
}

func palindrome(t Tabel) bool {
	for i := 0; i < t.m/2; i++ {
		if t.tab[i] != t.tab[t.m-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab Tabel

	isiArray(&tab)

	fmt.Println("\nArray asli:")
	cetakArray(tab)

	isPalindrome := palindrome(tab)
	if isPalindrome {
		fmt.Println("Palindrome: true")
	} else {
		fmt.Println("Palindrome: false")
	}

	balikkanArray(&tab)
	fmt.Println("\nArray setelah dibalik:")
	cetakArray(tab)
}
