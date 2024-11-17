package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	cx, cy, r int
}

func jarak(p1, p2 Titik) float64 {
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)))
}

func diDalam(c Lingkaran, p Titik) bool {
	jarakTitik := jarak(Titik{c.cx, c.cy}, p)
	return jarakTitik <= float64(c.r)
}

func main() {
	lingkaran1 := Lingkaran{cx: 2, cy: 1, r: 5}
	lingkaran2 := Lingkaran{cx: 8, cy: 4, r: 2}
	titik := Titik{x: 3, y: 2}

	dalamL1 := diDalam(lingkaran1, titik)
	dalamL2 := diDalam(lingkaran2, titik)

	// Output hasil
	if dalamL1 && dalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
