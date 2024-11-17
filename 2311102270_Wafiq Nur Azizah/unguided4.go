package main

import (
	"fmt"
	"unicode"
)

const MAX_SIZE int = 127

type charArray [MAX_SIZE]rune

func fillArray(arr *charArray, size *int) {
	fmt.Println("Masukkan karakter (ketik '.' untuk berhenti):")
	for *size < MAX_SIZE {
		var input rune
		fmt.Scanf("%c", &input)
		if unicode.IsSpace(input) {
			continue
		}
		if input == '.' {
			break
		}
		arr[*size] = unicode.ToLower(input)
		(*size)++
	}
}

func printArray(arr charArray, size int) {
	for i := 0; i < size; i++ {
		fmt.Printf("%c", arr[i])
	}
	fmt.Println()
}

func reverseArray(original charArray, reversed *charArray, size int) {
	for i := 0; i < size; i++ {
		reversed[i] = original[size-1-i]
	}
}

func isPalindrome(original charArray, reversed charArray, size int) bool {
	for i := 0; i < size; i++ {
		if original[i] != reversed[i] {
			return false
		}
	}
	return true
}

func main() {
	var inputArray charArray
	var reversedArray charArray
	var arraySize int = 0

	fillArray(&inputArray, &arraySize)
	fmt.Print("Isi array: ")
	printArray(inputArray, arraySize)

	reverseArray(inputArray, &reversedArray, arraySize)
	fmt.Print("Array setelah dibalik: ")
	printArray(reversedArray, arraySize)

	if isPalindrome(inputArray, reversedArray, arraySize) {
		fmt.Println("Array adalah palindrom")
	} else {
		fmt.Println("Array bukan palindrom")
	}
}
