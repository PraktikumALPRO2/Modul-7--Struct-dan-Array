package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	// Membuat array dengan panjang n
	arr := make([]int, n)

	// Input elemen array
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	// Menu
	for {
		fmt.Println("\nMenu:")
		fmt.Println("a. Tampilkan keseluruhan isi array")
		fmt.Println("b. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("c. Tampilkan elemen dengan indeks genap")
		fmt.Println("d. Tampilkan elemen dengan indeks kelipatan bilangan x")
		fmt.Println("e. Hapus elemen array pada indeks tertentu")
		fmt.Println("f. Tampilkan rata-rata dari elemen array")
		fmt.Println("g. Hitung standar deviasi elemen array")
		fmt.Println("h. Hitung frekuensi suatu bilangan dalam array")
		fmt.Println("i. Keluar")

		var pilihan string
		fmt.Print("\nMasukkan pilihan (a-h): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case "a": // Menampilkan keseluruhan isi dari array
			fmt.Println("\nKeseluruhan isi array:")
			fmt.Println(arr)

		case "b": // Menampilkan elemen-elemen array dengan indeks ganjil saja
			fmt.Println("\nElemen array dengan indeks ganjil:")
			for i := 1; i < n; i += 2 {
				fmt.Printf("arr[%d] = %d\n", i, arr[i])
			}

		case "c": // Menampilkan elemen-elemen array dengan indeks genap saja
			fmt.Println("\nElemen array dengan indeks genap:")
			for i := 0; i < n; i += 2 {
				fmt.Printf("arr[%d] = %d\n", i, arr[i])
			}

		case "d": // Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x
			var x int
			fmt.Print("\nMasukkan bilangan untuk kelipatan indeks: ")
			fmt.Scan(&x)

			if x <= 0 || x >= n { // Memastikan x valid untuk digunakan sebagai kelipatan indeks
				continue
			}

			// Menampilkan elemen-elemen array dengan indeks kelipatan x
			fmt.Printf("Elemen array dengan indeks kelipatan %d:\n", x)
			kelipatan := false // Untuk mengecek apakah ada indeks kelipatan x yang valid
			for i := x; i < n; i += x {
				fmt.Printf("arr[%d] = %d\n", i, arr[i])
				kelipatan = true
			}

			// Jika tidak ada elemen dengan indeks kelipatan x
			if !kelipatan {
				fmt.Printf("Tidak ada elemen pada indeks kelipatan %d\n", x)
			}
			
		case "e": // Menghapus elemen array pada indeks tertentu
			var hapusIndeks int // Mendeklarasikan variabel hapusIndeks dengan tipe data integer untuk menyimpan indeks yang akan dihapus
			fmt.Print("\nMasukkan indeks elemen yang akan dihapus: ")
			fmt.Scan(&hapusIndeks)
			if hapusIndeks >= 0 && hapusIndeks < n { // Memastikan indeks yang dimasukkan valid
				arr = append(arr[:hapusIndeks], arr[hapusIndeks+1:]...) // Menghapus elemen pada indeks hapusIndeks dengan menggabungkan bagian array sebelumnya dan sesudah dihapus
				fmt.Println("Array setelah dihapus:")
				fmt.Println(arr)
			} else {
				fmt.Println("Indeks tidak valid.") // Mencetak pesan jika indeks yang dimasukkan tidak valid
			}

		case "f": // Menampilkan rata-rata dari bilangan yang ada di dalam array
			sum := 0
			for _, nilai := range arr {
				sum += nilai
			}
			rataRata := float64(sum) / float64(len(arr))
			fmt.Printf("\nRata-rata elemen array: %.2f\n", rataRata)

		case "g": // Menghitung standar deviasi
			sum := 0
			for _, nilai := range arr {
				sum += nilai
			}
			rataRata := float64(sum) / float64(len(arr))
			variansiSum := 0.0
			for _, nilai := range arr {
				variansiSum += (float64(nilai) - rataRata) * (float64(nilai) - rataRata)
			}
			standarDeviasi := variansiSum / float64(len(arr)) // Variansi
			standarDeviasi = math.Sqrt(standarDeviasi)
			fmt.Printf("Standar deviasi elemen array: %.2f\n", standarDeviasi)

		case "h": // Menampilkan frekuensi dari suatu bilangan tertentu di dalam array
			var cari int
			fmt.Print("\nMasukkan bilangan untuk dihitung frekuensinya: ")
			fmt.Scan(&cari)
			frekuensi := 0
			for _, nilai := range arr {
				if nilai == cari {
					frekuensi++
				}
			}
			fmt.Printf("Frekuensi %d di dalam array: %d kali\n", cari, frekuensi)

		case "i": // Keluar dari program
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
