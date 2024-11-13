// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import (
	"fmt"
)

// Fungsi untuk mengecek apakah nama sudah ada di dalam slice
func sudahAda(daftarTeman []string, nama string) bool {
	// Loop untuk mengecek setiap nama dalam daftarTeman
	for _, teman := range daftarTeman {
		// Jika nama yang dimasukkan sudah ada dalam daftar, return true
		if teman == nama {
			return true
		}
	}
	// Jika tidak ditemukan, return false
	return false
}

func main() {
	// Slice awal untuk daftar teman dengan beberapa data
	daftarTeman := []string{"Andi", "Budi", "Cici"}

	// Nama-nama baru yang ingin ditambahkan ke daftar
	namaBaru := []string{"Dewi", "Budi", "Eka"}

	// Menambahkan nama baru hanya jika belum ada di daftar
	for _, nama := range namaBaru {
		// Mengecek apakah nama sudah ada dalam daftar
		if !sudahAda(daftarTeman, nama) {
			// Jika belum ada, menambahkannya ke dalam daftar
			daftarTeman = append(daftarTeman, nama)
		} else {
			// Jika nama sudah ada, menampilkan pesan
			fmt.Println("Nama", nama, "sudah ada dalam daftar.")
		}
	}

	// Menampilkan daftar teman akhir setelah penambahan
	fmt.Println("Daftar Teman:", daftarTeman)
}
