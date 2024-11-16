//Haifa Zahra Azzimmi
//2311102163

package main

import (
	"fmt"
	"strings"
)

func reverseArray(arr []rune) []rune {
	n := len(arr)
	reversed := make([]rune, n)
	for i := 0; i < n; i++ {
		reversed[i] = arr[n-i-1]
	}
	return reversed
}

func isPalindrome(arr []rune) bool {
	reversed := reverseArray(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] != reversed[i] {
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Print("Masukkan sekumpulan karakter: ")
	fmt.Scanln(&input)

	charArray := []rune(strings.ToLower(input)) // Ignore case

	reversedArray := reverseArray(charArray)

	fmt.Printf("Array asli: %s\n", string(charArray))
	fmt.Printf("Array yang dibalik: %s\n", string(reversedArray))

	if isPalindrome(charArray) {
		fmt.Println("Array membentuk palindrom.")
	} else {
		fmt.Println("Array tidak membentuk palindrom.")
	}
}
