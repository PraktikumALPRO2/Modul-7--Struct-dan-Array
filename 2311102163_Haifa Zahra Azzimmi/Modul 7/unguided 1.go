//Haifa Zahra Azzimmi
//2311102163

package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat  Titik
	radius int
}

func jarak(t1, t2 Titik) float64 {
	return math.Sqrt(float64((t1.x-t2.x)*(t1.x-t2.x) + (t1.y-t2.y)*(t1.y-t2.y)))
}

func titikDiDalamLingkaran(t Titik, lingkaran Lingkaran) bool {
	return jarak(t, lingkaran.pusat) <= float64(lingkaran.radius)
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 1 (cx1 cy1 r1): ")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 2 (cx2 cy2 r2): ")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Print("Masukkan koordinat titik sembarang (x y): ")
	fmt.Scan(&x, &y)

	lingkaran1 := Lingkaran{pusat: Titik{x: cx1, y: cy1}, radius: r1}
	lingkaran2 := Lingkaran{pusat: Titik{x: cx2, y: cy2}, radius: r2}
	titik := Titik{x: x, y: y}

	diDalamL1 := titikDiDalamLingkaran(titik, lingkaran1)
	diDalamL2 := titikDiDalamLingkaran(titik, lingkaran2)

	if diDalamL1 && diDalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
