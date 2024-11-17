package main

import "fmt"

type waktu struct {
	jam, menit, detik int
}

func main() {
	var wParkir, wPulang, durasi waktu
	var dParkir, dPulang, lparkir int

	fmt.Scan(&wParkir.jam, &wParkir.menit, &wParkir.detik)
	fmt.Scan(&wPulang.jam, &wPulang.menit, &wPulang.detik)

	dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600
	dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600
	lparkir = dPulang - dParkir

	durasi.jam = lparkir / 3600
	durasi.menit = lparkir % 3600 / 60
	durasi.detik = lparkir % 3600 % 60
	fmt.Printf("LAMA PARKIR : %d jam %d menit %d detik", durasi.jam, durasi.menit, durasi.detik)
}
