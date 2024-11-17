package main

import (
	"fmt"
	"strings"
)

// Fungsi untuk membalikkan urutan elemen dalam array karakter
func reverseCharacters(input []rune) []rune {
	length := len(input)
	reversedArray := make([]rune, length)
	for i := 0; i < length; i++ {
		reversedArray[i] = input[length-i-1]
	}
	return reversedArray
}

// Fungsi untuk memeriksa apakah array karakter membentuk palindrom
func isPalindrome(input []rune) bool {
	reversedArray := reverseCharacters(input)
	for i := 0; i < len(input); i++ {
		if input[i] != reversedArray[i] {
			return false
		}
	}
	return true
}

func main() {
	var userInput string
	fmt.Print("Masukkan sekumpulan karakter: ")
	fmt.Scanln(&userInput)

	// Mengonversi input ke dalam array karakter (rune) dan mengabaikan perbedaan huruf kapital
	convertedArray := []rune(strings.ToLower(userInput))

	// Membalikkan array karakter
	reversedConvertedArray := reverseCharacters(convertedArray)

	// Menampilkan array asli dan array yang dibalik
	fmt.Printf("Array asli: %s\n", string(convertedArray))
	fmt.Printf("Array yang dibalik: %s\n", string(reversedConvertedArray))

	// Memeriksa apakah array tersebut membentuk palindrom
	if isPalindrome(convertedArray) {
		fmt.Println("Array membentuk palindrom.")
	} else {
		fmt.Println("Array tidak membentuk palindrom.")
	}
}
