package main

import "fmt"

func main() {
	var namaKlubA, namaKlubB string
	var skorKlubA, skorKlubB int
	var daftarHasil []string
	
	fmt.Print("Masukkan nama klub A : ")
	fmt.Scanln(&namaKlubA)
	fmt.Print("Masukkan nama klub B : ")
	fmt.Scanln(&namaKlubB)
	for {
		fmt.Printf("Masukkan skor untuk %s dan %s : ", namaKlubA, namaKlubB)
		_, err := fmt.Scanf("%d %d\n", &skorKlubA, &skorKlubB)
		
		if err != nil {
			fmt.Println("Inputan tidak valid")
			tampilkanRekapan(daftarHasil)
			var temp string
			fmt.Scanln(&temp)
			continue
		}
		if skorKlubA < 0 || skorKlubB < 0 {
			break
		}
		if skorKlubA > skorKlubB {
			daftarHasil = append(daftarHasil, namaKlubA) // Klub A menang
			} else if skorKlubA < skorKlubB {
				daftarHasil = append(daftarHasil, namaKlubB) // Klub B menang
				} else {
					daftarHasil = append(daftarHasil, "Seri") // Hasil seri
					}
				}
				fmt.Println("Rekapan hasil pertandingan :")
				tampilkanRekapan(daftarHasil)
				fmt.Println("Pertandingan selesai")
}

func tampilkanRekapan(daftarHasil []string) {
	if len(daftarHasil) == 0 {
		fmt.Println("Belum ada hasil pertandingan.")
		} else {
			for i, hasil := range daftarHasil {
				fmt.Printf("Pertandingan %d : %s\n", i+1, hasil)
			}
		}
}