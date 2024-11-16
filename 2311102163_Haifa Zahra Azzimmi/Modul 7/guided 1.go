//Haifa Zahra Azzimmi
//2311102163

package main

import "fmt"

type waktu struct {
    jam, menit, detik int
}

func main() {
    var wParkir, wPulang, durasi waktu
    var dParkir, dPulang, lParkir int

    fmt.Print("Masukkan waktu datang parkir (jam menit detik): ")
    fmt.Scan(&wParkir.jam, &wParkir.menit, &wParkir.detik)
    fmt.Print("Masukkan waktu pulang parkir (jam menit detik): ")
    fmt.Scan(&wPulang.jam, &wPulang.menit, &wPulang.detik)

    dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600 // konversi ke detik
    dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600 // detik
    lParkir = dPulang - dParkir                                   // detik dari pulang-datang

    durasi.jam = lParkir / 3600
    durasi.menit = lParkir % 3600 / 60
    durasi.detik = lParkir % 3600 % 60
    fmt.Printf("Lama parkir: %d jam %d menit %d detik\n", durasi.jam, durasi.menit, durasi.detik)
}
