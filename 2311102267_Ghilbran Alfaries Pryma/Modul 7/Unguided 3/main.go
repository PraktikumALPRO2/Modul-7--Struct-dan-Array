package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struct untuk merepresentasikan pertandingan
type Pertandingan struct {
	KlubA   string
	KlubB   string
	SkorA   int
	SkorB   int
	Pemenang string
}

func main() {
	// Membuat scanner untuk input
	scanner := bufio.NewScanner(os.Stdin)

	// Input nama klub
	fmt.Print("Klub A: ")
	scanner.Scan()
	klubA := scanner.Text()

	fmt.Print("Klub B: ")
	scanner.Scan()
	klubB := scanner.Text()

	// Daftar pemenang dan semua pertandingan
	var pemenang []string
	var pertandinganList []Pertandingan

	// Mulai loop untuk input skor pertandingan
	pertandinganKe := 1
	for {
		fmt.Printf("Pertandingan %d: ", pertandinganKe)
		scanner.Scan()
		input := scanner.Text()
		skor := strings.Split(input, " ")

		// Validasi input
		if len(skor) != 2 {
			fmt.Println("Input tidak valid. Masukkan dua angka dipisahkan dengan spasi.")
			continue
		}

		// Parsing skor
		skorA, errA := strconv.Atoi(skor[0])
		skorB, errB := strconv.Atoi(skor[1])
		if errA != nil || errB != nil {
			fmt.Println("Input tidak valid. Masukkan angka.")
			continue
		}

		// Periksa validitas skor
		if skorA < 0 || skorB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		// Tentukan pemenang
		pemenangPertandingan := ""
		if skorA > skorB {
			pemenangPertandingan = klubA
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenangPertandingan = klubB
			pemenang = append(pemenang, klubB)
		} else {
			pemenangPertandingan = "Draw"
		}

		// Simpan data pertandingan ke struct
		pertandingan := Pertandingan{
			KlubA:   klubA,
			KlubB:   klubB,
			SkorA:   skorA,
			SkorB:   skorB,
			Pemenang: pemenangPertandingan,
		}
		pertandinganList = append(pertandinganList, pertandingan)

		// Tampilkan hasil pertandingan
		fmt.Printf("Hasil %d: %s\n", pertandinganKe, pemenangPertandingan)

		pertandinganKe++
	}

	// Tampilkan daftar semua pertandingan
	fmt.Println("\nDetail pertandingan:")
	for i, pertandingan := range pertandinganList {
		fmt.Printf("Pertandingan %d: %s %d-%d %s, Pemenang: %s\n",
			i+1, pertandingan.KlubA, pertandingan.SkorA, pertandingan.SkorB, pertandingan.KlubB, pertandingan.Pemenang)
	}

	// Tampilkan daftar klub yang menang
	if len(pemenang) > 0 {
		fmt.Println("\nDaftar klub pemenang:")
		for _, klub := range pemenang {
			fmt.Println(klub)
		}
	} else {
		fmt.Println("\nTidak ada klub yang menang.")
	}
}
