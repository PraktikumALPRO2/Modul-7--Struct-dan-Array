package main

import (
    "fmt"
    "math"
)

// Struct untuk menyimpan data lingkaran
type Rectangle_158 struct {
    xc, yc, radius float64
}

// Struct untuk menyimpan data titik
type Point_158 struct {
    x, y float64
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(xc, yc, x, y float64) float64 {
    return math.Sqrt(math.Pow(x-xc, 2) + math.Pow(y-yc, 2))
}

// Fungsi untuk menentukan posisi titik relatif terhadap dua lingkaran
func posisiTitik(lingkaran1, lingkaran2 Rectangle_158, titik Point_158) string {
    jarak1 := jarak(lingkaran1.xc, lingkaran1.yc, titik.x, titik.y)
    jarak2 := jarak(lingkaran2.xc, lingkaran2.yc, titik.x, titik.y)

    if jarak1 <= lingkaran1.radius && jarak2 <= lingkaran2.radius {
        return "Titik di dalam lingkaran 1 dan 2"
    } else if jarak1 <= lingkaran1.radius {
        return "Titik di dalam lingkaran 1"
    } else if jarak2 <= lingkaran2.radius {
        return "Titik di dalam lingkaran 2"
    } else {
        return "Titik di luar lingkaran 1 dan 2"
    }
}

func main() {
    var xc1, yc1, r1, xc2, yc2, r2, x, y float64
    fmt.Scan(&xc1, &yc1, &r1)
    fmt.Scan(&xc2, &yc2, &r2)
    fmt.Scan(&x, &y)

    lingkaran1 := Rectangle_158{xc: xc1, yc: yc1, radius: r1}
    lingkaran2 := Rectangle_158{xc: xc2, yc: yc2, radius: r2}
    titik := Point_158{x: x, y: y}

    hasil := posisiTitik(lingkaran1, lingkaran2, titik)
    fmt.Println(hasil)
}
