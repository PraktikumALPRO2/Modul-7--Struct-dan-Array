package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Struct untuk menyimpan koordinat titik
type Point struct {
	x, y float64
}

// Struct untuk menyimpan informasi lingkaran
type Circle struct {
	center Point
	radius float64
}

// Fungsi untuk menghitung jarak antara dua titik
func distance(p1, p2 Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

// Fungsi untuk mengecek apakah titik berada di dalam lingkaran
func isPointInCircle(p Point, c Circle) bool {
	return distance(p, c.center) <= c.radius
}

// Fungsi untuk menentukan posisi titik terhadap dua lingkaran
func determinePosition(p Point, c1, c2 Circle) string {
	inCircle1 := isPointInCircle(p, c1)
	inCircle2 := isPointInCircle(p, c2)

	if inCircle1 && inCircle2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if inCircle1 {
		return "Titik di dalam lingkaran 1"
	} else if inCircle2 {
		return "Titik di dalam lingkaran 2"
	}
	return "Titik di luar lingkaran 1 dan 2"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Membaca data lingkaran 1
	scanner.Scan()
	line1 := strings.Fields(scanner.Text())
	x1, _ := strconv.ParseFloat(line1[0], 64)
	y1, _ := strconv.ParseFloat(line1[1], 64)
	r1, _ := strconv.ParseFloat(line1[2], 64)
	circle1 := Circle{Point{x1, y1}, r1}

	// Membaca data lingkaran 2
	scanner.Scan()
	line2 := strings.Fields(scanner.Text())
	x2, _ := strconv.ParseFloat(line2[0], 64)
	y2, _ := strconv.ParseFloat(line2[1], 64)
	r2, _ := strconv.ParseFloat(line2[2], 64)
	circle2 := Circle{Point{x2, y2}, r2}

	// Membaca koordinat titik yang akan dicek
	scanner.Scan()
	line3 := strings.Fields(scanner.Text())
	x, _ := strconv.ParseFloat(line3[0], 64)
	y, _ := strconv.ParseFloat(line3[1], 64)
	point := Point{x, y}

	// Menentukan dan mencetak hasil
	result := determinePosition(point, circle1, circle2)
	fmt.Println(result)
}
