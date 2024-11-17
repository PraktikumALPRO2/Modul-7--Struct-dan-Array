package main

import "fmt"

func main() {
    var klubA, klubB string
    var skorA, skorB int
    klubMenang := []string{}

    fmt.Print("Masukkan nama Klub A: ")
    fmt.Scanln(&klubA)

    fmt.Print("Masukkan nama Klub B: ")
    fmt.Scanln(&klubB)

    for pertandingan := 1; pertandingan <= 9; pertandingan++ {
        fmt.Printf("Masukkan skor pertandingan %d: ", pertandingan)
        fmt.Scanln(&skorA, &skorB)

        // Cek validitas skor
        if skorA < 0 || skorB < 0 {
            fmt.Println("Skor tidak valid. Program berhenti.")
            break
        }

        // Tentukan pemenang
        if skorA > skorB {
            klubMenang = append(klubMenang, fmt.Sprintf("Hasil %d: %s", pertandingan, klubA))
        } else if skorB > skorA {
            klubMenang = append(klubMenang, fmt.Sprintf("Hasil %d: %s", pertandingan, klubB))
        } else {
            klubMenang = append(klubMenang, fmt.Sprintf("Hasil %d: Draw", pertandingan))
        }
    }

    fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
    for _, klub := range klubMenang {
        fmt.Println(klub)
    }
}
