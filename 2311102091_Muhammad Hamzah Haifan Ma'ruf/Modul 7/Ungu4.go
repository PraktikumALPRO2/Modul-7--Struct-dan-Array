package main

import "fmt"

const MaksKarakter int = 127 

type arrayKarakter [MaksKarakter]rune

func inputArrayKarakter(array *arrayKarakter, jumlah int) {
	var indeks int
	var karakter rune
	fmt.Print("Masukkan karakter : ")
	for indeks = 0; indeks < jumlah; indeks++ {
		fmt.Scanf("%c", &karakter)
		if karakter == '.' {
			break
		}
		array[indeks] = karakter
	}
}

func tampilkanKarakter(array arrayKarakter, jumlah int) {
	for i := 0; i < jumlah; i++ {
		fmt.Printf("%c", array[i])
	}
	fmt.Println()
}

func balikkanUrutan(array *arrayKarakter, jumlah int) {
	for i := 0; i < jumlah/2; i++ {
		sementara := array[i]
		array[i] = array[jumlah-i-1]
		array[jumlah-i-1] = sementara
	}
}
func cekApakahPalindrom(array arrayKarakter, jumlah int) bool {
	var arraySalinan arrayKarakter
	copy(arraySalinan[:], array[:])
	balikkanUrutan(&arraySalinan, jumlah)
	for i := 0; i < jumlah; i++ {
		if arraySalinan[i] != array[i] {
			return false
		}
	}
	return true
}
func main() {
	var array arrayKarakter
	var jumlahKarakter int
	fmt.Print("Masukkan jumlah karakter : ")
	fmt.Scanln(&jumlahKarakter)
	inputArrayKarakter(&array, jumlahKarakter)
	fmt.Print("Karakter yang dimasukkan adalah : ")
	tampilkanKarakter(array, jumlahKarakter)
	balikkanUrutan(&array, jumlahKarakter)
	fmt.Print("Teks terbalik (reverse) adalah : ")
	tampilkanKarakter(array, jumlahKarakter)
	if cekApakahPalindrom(array, jumlahKarakter) {
		fmt.Print("Array ini adalah palindrom.")
		} else {
			fmt.Print("Array ini bukan palindrom.")
		}
	}