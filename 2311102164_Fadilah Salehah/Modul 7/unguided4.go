// Fadilah Salehah
// 2311102164

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Konstanta dan tipe data
const MAKSIMUM int = 127

type arrayKarakter [MAKSIMUM]rune

// Fungsi untuk mengisi array berdasarkan input pengguna
func masukkanTeks(arr *arrayKarakter, panjang *int) {
	var teks string
	fmt.Println("Masukkan teks, diakhiri dengan tanda titik ('.'):")
	reader := bufio.NewReader(os.Stdin)
	teks, _ = reader.ReadString('\n')
	teks = strings.TrimSpace(teks)

	for _, karakter := range teks {
		if karakter == '.' || *panjang >= MAKSIMUM {
			break
		}
		arr[*panjang] = karakter
		*panjang++
	}
}

// Fungsi untuk mencetak isi array
func tampilkanTeks(arr arrayKarakter, panjang int) {
	for i := 0; i < panjang; i++ {
		fmt.Printf("%c", arr[i])
	}
	fmt.Println()
}

// Fungsi untuk membalik isi array
func balikArray(arr *arrayKarakter, panjang int) {
	for i := 0; i < panjang/2; i++ {
		arr[i], arr[panjang-1-i] = arr[panjang-1-i], arr[i]
	}
}

// Fungsi untuk memeriksa apakah array membentuk palindrom
func cekPalindrom(arr arrayKarakter, panjang int) bool {
	for i := 0; i < panjang/2; i++ {
		if arr[i] != arr[panjang-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var teks arrayKarakter
	var panjang int

	// Isi array menggunakan fungsi masukkanTeks
	masukkanTeks(&teks, &panjang)

	// Tampilkan teks yang dimasukkan
	fmt.Print("Teks: ")
	tampilkanTeks(teks, panjang)

	// Balikkan isi array
	balikArray(&teks, panjang)

	// Tampilkan teks setelah dibalik
	fmt.Print("Reverse teks: ")
	tampilkanTeks(teks, panjang)

	// Periksa apakah teks tersebut adalah palindrom
	if cekPalindrom(teks, panjang) {
		fmt.Println("Palindrom? true")
	} else {
		fmt.Println("Palindrom? false")
	}
}