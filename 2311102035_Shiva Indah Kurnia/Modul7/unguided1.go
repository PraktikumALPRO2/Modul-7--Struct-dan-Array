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

	switch {
	case inCircle1 && inCircle2:
		return "Titik di dalam lingkaran 1 dan 2"
	case inCircle1:
		return "Titik di dalam lingkaran 1"
	case inCircle2:
		return "Titik di dalam lingkaran 2"
	default:
		return "Titik di luar lingkaran 1 dan 2"
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Membaca data lingkaran 1
	circle1 := readCircle(scanner)

	// Membaca data lingkaran 2
	circle2 := readCircle(scanner)

	// Membaca koordinat titik yang akan dicek
	point := readPoint(scanner)

	// Menentukan dan mencetak hasil
	result := determinePosition(point, circle1, circle2)
	fmt.Println(result)
}

// Fungsi untuk membaca data lingkaran dari input
func readCircle(scanner *bufio.Scanner) Circle {
	scanner.Scan()
	line := strings.Fields(scanner.Text())
	x, _ := strconv.ParseFloat(line[0], 64)
	y, _ := strconv.ParseFloat(line[1], 64)
	radius, _ := strconv.ParseFloat(line[2], 64)
	return Circle{Point{x, y}, radius}
}

// Fungsi untuk membaca koordinat titik dari input
func readPoint(scanner *bufio.Scanner) Point {
	scanner.Scan()
	line := strings.Fields(scanner.Text())
	x, _ := strconv.ParseFloat(line[0], 64)
	y, _ := strconv.ParseFloat(line[1], 64)
	return Point{x, y}
}