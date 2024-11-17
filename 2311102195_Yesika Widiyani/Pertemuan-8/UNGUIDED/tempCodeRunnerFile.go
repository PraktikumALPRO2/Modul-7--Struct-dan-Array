package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX = 127

type tabChar [MAX]rune

func isArrayPalindrome(tab tabChar, n int) bool {
	for i := 0; i < n/2; i++ {
		if tab[i] != tab[n-1-i] {
			return false
		}
	}
	return true
}

func balikArray(tab tabChar, n int) tabChar {
	var reversed tabChar
	for i := 0; i < n; i++ {
		reversed[i] = tab[n-1-i]
	}
	return reversed
}

func cetakArray(tab tabChar, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(tab[i]))
	}
	fmt.Println()
}

func main() {
	var tab tabChar
	var n int

	// Membaca input dari pengguna
	fmt.Print("Masukkan karakter (tanpa spasi): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Membersihkan karakter newline dari input
	input = input[:len(input)-1]
	n = len(input)

	// Validasi input agar tidak melebihi batas MAX
	if n > MAX {
		fmt.Printf("Jumlah karakter terlalu banyak! Maksimum %d karakter.\n", MAX)
		return
	}

	// Memasukkan karakter ke array
	for i, char := range input {
		tab[i] = char
	}

	// Menampilkan array asli
	fmt.Println("\nArray asli:")
	cetakArray(tab, n)

	// Membalik array dan menampilkan hasil
	fmt.Println("Array setelah dibalik:")
	reversed := balikArray(tab, n)
	cetakArray(reversed, n)

	// Memeriksa apakah array adalah palindrome
	if isArrayPalindrome(tab, n) {
		fmt.Println("Array adalah palindrome.")
	} else {
		fmt.Println("Array bukan palindrome.")
	}
}