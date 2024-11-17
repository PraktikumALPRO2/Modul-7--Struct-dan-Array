package main

import "fmt"

const NMAX int = 127
type tabel [NMAX] rune

var tab tabel
var m int

// Prosedur untuk mengisi array dengan karakter yang dimasukkan user
func isiArray(t *tabel, n *int) {
    var input string
    fmt.Print("Teks : ")
    fmt.Scanln(&input)  // Membaca input dalam bentuk string

    *n = len(input)     // Menghitung panjang string
    for i := 0; i < *n; i++ {
        (*t)[i] = rune(input[i])  // Menyimpan karakter ke dalam array
    }
}

// Prosedur untuk mencetak array
func cetakArray(t tabel, n int) {
    for i := 0; i < n; i++ {
        fmt.Print(string(t[i]))  // Mencetak karakter satu per satu
    }
    fmt.Println()
}

// Prosedur untuk membalikkan urutan array
func balikanArray(t *tabel, n int) {
    for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
        (*t)[i], (*t)[j] = (*t)[j], (*t)[i]
	}
}

// Fungsi untuk memeriksa apakah array membentuk palindrom
func palindrom(t tabel, n int) bool {
    // Membalikkan array sementara untuk pemeriksaan palindrom
    balikanArray(&t, n)

    // Memeriksa apakah array asli sama dengan array yang dibalik
    for i := 0; i < n; i++ {
        if t[i] != tab[i] {
            return false
        }
    }
    return true
}

func main() {
    // Mengisi array dengan memanggil prosedur isiArray
    isiArray(&tab, &m)

    // Membalikkan isi array dan menampilkan hasilnya
    balikanArray(&tab, m)
    fmt.Print("Reverse teks : ")
    cetakArray(tab, m)

    // Memeriksa apakah teks tersebut palindrom
    if palindrom(tab, m) {
        fmt.Println("Teks Palindrom : true")
    } else {
        fmt.Println("Teks Palindrom : false")
    }
}
