package main

import (
	"fmt"
	"strings"
)

func reverseArray(arr []rune) []rune {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return arr
}

func palindrom(arr []rune) bool {
	reversed := reverseArray(append([]rune{}, arr...))
	return string(arr) == string(reversed)
}

func main() {
	var input string
	fmt.Print("Teks: ")
	fmt.Scanln(&input)

	arr := []rune(strings.ReplaceAll(input, " ", ""))
	reversed := reverseArray(append([]rune{}, arr...))

	fmt.Printf("Reverse teks: %s\n", string(reversed))
	fmt.Printf("Palindrom? %t\n", palindrom(arr))
}
