package main

import (
	"fmt"
)

func main() {
	elemenArray := make([]int, 10)
	var jumlahElemen int
	fmt.Print("Masukkan jumlah elemen (maksimal 10) : ")
	fmt.Scanln(&jumlahElemen)
	if jumlahElemen > 10 {
		jumlahElemen = 10
		fmt.Println("Jumlah elemen dibatasi maksimal 10. Menggunakan 10 elemen.")
	}
	fmt.Println("Masukkan elemen-elemen array : ")
	for i := 0; i < jumlahElemen; i++ {
		var elemen int
		fmt.Scanln(&elemen)
		elemenArray[i] = elemen
	}
	for {
		fmt.Println("\nMenu pilihan : ")
		fmt.Println("1. Menampilkan seluruh isi array")
		fmt.Println("2. Menampilkan elemen array dengan indeks ganjil")
		fmt.Println("3. Menampilkan elemen array dengan indeks genap")
		fmt.Println("4. Menampilkan elemen array dengan indeks kelipatan bilangan x")
		fmt.Println("5. Menghapus elemen array pada indeks tertentu")
		fmt.Println("6. Menampilkan rata-rata nilai dalam array")
		fmt.Println("7. Menampilkan standar deviasi dari nilai dalam array")
		fmt.Println("8. Menampilkan frekuensi kemunculan suatu bilangan dalam array")
		fmt.Println("9. Keluar")
		
		var pilihan int
		fmt.Print("\nMasukkan pilihan : ")
		fmt.Scanln(&pilihan)
		
		switch pilihan {
			case 1 : tampilkanSeluruhIsiArray(elemenArray, jumlahElemen)
		case 2 :
			tampilkanElemenIndeksGanjil(elemenArray, jumlahElemen)
		case 3:
			tampilkanElemenIndeksGenap(elemenArray, jumlahElemen)
		case 4:
			tampilkanElemenKelipatan(elemenArray, jumlahElemen)
		case 5:
			jumlahElemen = hapusElemenArray(elemenArray, jumlahElemen) 
		case 6:
			tampilkanRataRata(elemenArray, jumlahElemen)
		case 7:
			tampilkanStandarDeviasi(elemenArray, jumlahElemen)
		case 8:
			tampilkanFrekuensi(elemenArray, jumlahElemen)
		case 9:
			fmt.Println("Terima kasih! Program selesai.")
			return // Keluar dari program
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func tampilkanSeluruhIsiArray(array []int, jumlahElemen int) {
	fmt.Println("Isi array:")
	for i := 0; i < jumlahElemen; i++ {
		fmt.Println(array[i])
	}
}

func tampilkanElemenIndeksGanjil(array []int, jumlahElemen int) {
	fmt.Println("Elemen dengan indeks ganjil:")
	for i := 1; i < jumlahElemen; i += 2 {
		fmt.Println(array[i])
	}
}

func tampilkanElemenIndeksGenap(array []int, jumlahElemen int) {
	fmt.Println("Elemen dengan indeks genap:")
	for i := 0; i < jumlahElemen; i += 2 {
		fmt.Println(array[i])
	}
}

func tampilkanElemenKelipatan(array []int, jumlahElemen int) {
	var x int
	fmt.Print("Masukkan bilangan kelipatan: ")
	fmt.Scanln(&x)
	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	for i := 0; i < jumlahElemen; i++ {
		if i%x == 0 {
			fmt.Println(array[i])
		}
	}
}

func hapusElemenArray(array []int, jumlahElemen int) int {
	var indeks int
	fmt.Print("Masukkan indeks elemen yang ingin dihapus: ")
	fmt.Scanln(&indeks)
	if indeks < 0 || indeks >= jumlahElemen {
		fmt.Println("Indeks tidak valid")
		return jumlahElemen // Kembalikan jumlahElemen tanpa perubahan
		}
		for i := indeks; i < jumlahElemen-1; i++ {
			array[i] = array[i+1]
		}
		jumlahElemen-- 
		fmt.Println("Elemen pada indeks", indeks, "telah dihapus")
		return jumlahElemen // Kembalikan jumlah elemen yang baru
}

func tampilkanRataRata(array []int, jumlahElemen int) {
	var total int
	for i := 0; i < jumlahElemen; i++ {
		total += array[i]
	}
	rataRata := float64(total) / float64(jumlahElemen)
	fmt.Printf("Rata-rata: %.2f\n", rataRata)
}

func tampilkanStandarDeviasi(array []int, jumlahElemen int) {
	var total, rataRata float64
	for i := 0; i < jumlahElemen; i++ {
		total += float64(array[i])
	}
	rataRata = total / float64(jumlahElemen)
	var totalDeviasi float64
	for i := 0; i < jumlahElemen; i++ {
		deviasi := float64(array[i]) - rataRata
		totalDeviasi += deviasi * deviasi
	}
	standarDeviasi := totalDeviasi / float64(jumlahElemen)
	standarDeviasi = hitungAkarKuadrat(standarDeviasi)
	fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi)
}

func hitungAkarKuadrat(x float64) float64 {
	if x < 0 {
		return 0
	}
	z := x / 2.0
	for i := 0; i < 10; i++ {
		z = (z + x/z) / 2.0
	}
	return z
}

func tampilkanFrekuensi(array []int, jumlahElemen int) {
	var bilangan int
	fmt.Print("Masukkan bilangan untuk mencari frekuensi: ")
	fmt.Scanln(&bilangan)
	frekuensi := 0
	for i := 0; i < jumlahElemen; i++ {
		if array[i] == bilangan {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi dari bilangan %d: %d\n", bilangan, frekuensi)
}