package main

import (
	"fmt"
)

// Fungsi untuk mengecek apakah nama sudah ada di dalam slice
func sudahTersedia(listNama []string, namaBaru string) bool {
	for _, nama := range listNama {
		if nama == namaBaru {
			return true
		}
	}
	return false
}

func main() {
	// Slice awal untuk daftar kontak dengan beberapa data
	kontakAwal := []string{"Ari", "Bayu", "Citra"}

	// Nama-nama baru yang ingin dimasukkan
	kontakTambahan := []string{"Dian", "Bayu", "Elok"} 

	// Menambahkan nama baru hanya jika belum ada di daftar
	for _, kontak := range kontakTambahan {
		if !sudahTersedia(kontakAwal, kontak) {
			kontakAwal = append(kontakAwal, kontak)
		} else {
			fmt.Println("Nama", kontak, "sudah ada di daftar kontak.")
		}
	}

	// Menampilkan daftar kontak akhir
	fmt.Println("Daftar Kontak:", kontakAwal)
}
