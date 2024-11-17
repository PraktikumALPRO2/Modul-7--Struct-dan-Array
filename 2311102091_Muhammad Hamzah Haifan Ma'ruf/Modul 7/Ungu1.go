package main

import (
	"fmt"
	"math"
)

type Titik struct {
	X int
	Y int
}
 
type Lingkaran struct {
	Tengah Titik
	Radius int
}

func hitungJarak(t1, t2 Titik) float64 {
	return math.Sqrt(math.Pow(float64(t1.X-t2.X), 2) + math.Pow(float64(t1.Y-t2.Y), 2))
}

func tentukanPosisiTitik(lingkaran1, lingkaran2 Lingkaran, titik Titik) string {
	jarakKeLingkaran1 := hitungJarak(lingkaran1.Tengah, titik)
	jarakKeLingkaran2 := hitungJarak(lingkaran2.Tengah, titik)

	if jarakKeLingkaran1 <= float64(lingkaran1.Radius) && jarakKeLingkaran2 <=
	float64(lingkaran2.Radius) {
		return "Titik berada di dalam lingkaran 1 dan 2"
		} else if jarakKeLingkaran1 <= float64(lingkaran1.Radius) {
			return "Titik berada di dalam lingkaran 1"
			} else if jarakKeLingkaran2 <= float64(lingkaran2.Radius) {
				return "Titik berada di dalam lingkaran 2"
				} else {
					return "Titik berada di luar lingkaran 1 dan 2"
				}
}

func main() {
	var lingkaran1, lingkaran2 Lingkaran
	var titik Titik
	fmt.Print("Masukkan pusat dan radius lingkaran 1 : ")
	fmt.Scanln(&lingkaran1.Tengah.X, &lingkaran1.Tengah.Y,&lingkaran1.Radius)
	fmt.Print("Masukkan pusat dan radius lingkaran 2 : ")
	fmt.Scanln(&lingkaran2.Tengah.X, &lingkaran2.Tengah.Y,&lingkaran2.Radius)
	fmt.Print("Masukkan koordinat titik sembarang : ")
	fmt.Scanln(&titik.X, &titik.Y)
	hasil := tentukanPosisiTitik(lingkaran1, lingkaran2, titik)
	fmt.Println("Posisi titik : ", hasil)
}