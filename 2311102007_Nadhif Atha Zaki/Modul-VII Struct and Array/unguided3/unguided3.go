package main

import "fmt"

func main() {
	var namaKlubA, namaKlubB string
	var skorKlubA, skorKlubB int
	var daftarHasil []string

	// Meminta input nama klub
	fmt.Print("Masukkan nama klub A: ")
	fmt.Scanln(&namaKlubA)
	fmt.Print("Masukkan nama klub B: ")
	fmt.Scanln(&namaKlubB)

	// Loop untuk menerima input skor
	for {
		fmt.Printf("Masukkan skor untuk %s dan %s: ", namaKlubA, namaKlubB)
		_, err := fmt.Scanf("%d %d\n", &skorKlubA, &skorKlubB)

		// Periksa jika input tidak valid
		if err != nil {
			fmt.Println("Input tidak valid. Harap masukkan dua angka.")
			// Menampilkan rekapan hasil pertandingan sejauh ini
			tampilkanRekapan(daftarHasil)
			// Mengosongkan buffer input agar input berikutnya tidak terpengaruh
			var temp string
			fmt.Scanln(&temp)
			continue
		}

		// Keluar dari loop jika ada skor negatif
		if skorKlubA < 0 || skorKlubB < 0 {
			break
		}

		// Tentukan hasil pertandingan
		if skorKlubA > skorKlubB {
			daftarHasil = append(daftarHasil, namaKlubA) // Klub A menang
		} else if skorKlubA < skorKlubB {
			daftarHasil = append(daftarHasil, namaKlubB) // Klub B menang
		} else {
			daftarHasil = append(daftarHasil, "Seri") // Hasil seri
		}
	}

	// Menampilkan hasil akhir
	fmt.Println("Rekapan hasil pertandingan:")
	tampilkanRekapan(daftarHasil)
	fmt.Println("Pertandingan selesai")
}

// Fungsi untuk menampilkan rekapan hasil pertandingan
func tampilkanRekapan(daftarHasil []string) {
	if len(daftarHasil) == 0 {
		fmt.Println("Belum ada hasil pertandingan.")
	} else {
		for i, hasil := range daftarHasil {
			fmt.Printf("Pertandingan %d: %s\n", i+1, hasil)
		}
	}
}
