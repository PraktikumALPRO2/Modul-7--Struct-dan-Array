package main

import (
    "fmt"
)

// Fungsi untuk membalikkan array
func reverseTeks_158(tab []rune, n int) {
    for i := 0; i < n/2; i++ {
        temp := tab[i]
        tab[i] = tab[n-i-1]
        tab[n-i-1] = temp
    }
}

// Fungsi untuk memeriksa apakah array membentuk palindrome
func Palindrome_158(tab []rune, n int) bool {
    for i := 0; i < n/2; i++ {
        if tab[i] != tab[n-i-1] {
            return false
        }
    }
    return true
}

func main() {
    var input string
    fmt.Print("Teks: ")
    fmt.Scanln(&input)

    // Menghilangkan spasi dari input dan mengubah menjadi array rune
    var tab []rune
    for _, char := range input {
        if char != ' ' {
            tab = append(tab, char)
        }
    }
    n := len(tab)

    reverseTeks_158(tab, n)
    fmt.Print("Reverse: ")
    for _, char := range tab {
        fmt.Printf("%c ", char)
    }
    fmt.Println()

    palindrome := Palindrome_158(tab, n)
    fmt.Printf("Palindrome = %v\n", palindrome)
}
