package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	cx, cy, r int
}

func jarak(t1, t2 titik) float64 {
	return math.Sqrt(math.Pow(float64(t1.x-t2.x), 2) + math.Pow(float64(t1.y-t2.y), 2))
}

func lokasititik(t titik, l lingkaran) bool {
	return jarak(titik{t.x, t.y}, titik{l.cx, l.cy}) <= float64(l.r)
}

func main() {
	var l1, l2 lingkaran
	var t titik

	fmt.Print("Masukkan Titik Pusat Lingkaran 1 (cx cy): ")
	fmt.Scan(&l1.cx, &l1.cy)
	fmt.Print("Masukkan Radius Lingkaran 1: ")
	fmt.Scan(&l1.r)
	fmt.Print("Masukkan Titik Pusat Lingkaran 2 (cx cy): ")
	fmt.Scan(&l2.cx, &l2.cy)
	fmt.Print("Masukkan Radius Lingkaran 2: ")
	fmt.Scan(&l2.r)
	fmt.Print("Masukkan Koordinat Titik (x y): ")
	fmt.Scan(&t.x, &t.y)

	if lokasititik(t, l1) && lokasititik(t, l2) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if lokasititik(t, l1) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if lokasititik(t, l2) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
