// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

// Mendefinisikan tipe data struct bernama 'waktu' yang memiliki atribut jam, menit, dan detik bertipe int
type waktu struct {
	jam, menit, detik int
}

func main() {
	// Mendeklarasikan variabel wParkir, wPulang, dan durasi dengan tipe data waktu
	var wParkir, wPulang, durasi waktu

	// Mendeklarasikan variabel dParkir, dPulang, dan lParkir sebagai variabel integer
	// dParkir dan dPulang akan menyimpan waktu dalam detik, sedangkan lParkir akan menyimpan lama parkir dalam detik
	var dParkir, dPulang, lParkir int

	// Mengambil input dari pengguna untuk waktu parkir dalam bentuk jam, menit, dan detik
	fmt.Scan(&wParkir.jam, &wParkir.menit, &wParkir.detik)

	// Mengambil input dari pengguna untuk waktu pulang dalam bentuk jam, menit, dan detik
	fmt.Scan(&wPulang.jam, &wPulang.menit, &wPulang.detik)

	// Mengonversi waktu parkir (wParkir) menjadi detik, dan menyimpannya ke variabel dParkir
	dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600

	// Mengonversi waktu pulang (wPulang) menjadi detik, dan menyimpannya ke variabel dPulang
	dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600

	// Menghitung lama parkir dalam detik dengan mengurangi dPulang dengan dParkir
	lParkir = dPulang - dParkir

	// Mengonversi total durasi parkir dari detik ke jam, menit, dan detik
	durasi.jam = lParkir / 3600        // Menghitung jam dengan membagi total detik dengan 3600
	durasi.menit = lParkir % 3600 / 60 // Menghitung menit dari sisa detik setelah diambil jam
	durasi.detik = lParkir % 3600 % 60 // Menghitung detik dari sisa setelah diambil jam dan menit

	// Menampilkan lama parkir dalam format "jam menit detik"
	fmt.Printf("Lama Parkir : %d jam %d menit %d detik", durasi.jam, durasi.menit, durasi.detik)
}
