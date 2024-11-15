package main
import "fmt"
type waktu struct {
	 jam, menit, detik int
}

func main(){
 	var wParkir, wPulang, durasi waktu // awal = 2 45 14 
 	var dParkir, dPulang, lParkir int // pulang = 4 30 12

 	fmt.Scan(&wParkir.jam, &wParkir.menit, &wParkir.detik)
 	fmt.Scan(&wPulang.jam, &wPulang.menit, &wPulang.detik)

 	dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600 // konversimke detik
 	dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600 // detik 
 	lParkir = dPulang - dParkir // detik dari pulang-datang
	
	durasi.jam = lParkir / 3600
 	durasi.menit = lParkir % 3600 / 60
 	durasi.detik = lParkir % 3600 % 60 // 17 
 	fmt.Printf("Lama parkir: %d jam %d menit %d detik",durasi.jam, durasi.menit, durasi.detik)
}