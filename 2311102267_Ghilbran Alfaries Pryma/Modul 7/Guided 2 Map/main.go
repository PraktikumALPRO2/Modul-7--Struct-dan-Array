package main

import (
	"fmt"
)

func main() {
	// Membuat map dengan nama makanan sebagai kunci dan harga sebagai nilai
	hargaMakanan := map[string]int{
		"Nasi Goreng": 20000,
		"Rendang":     35000,
		"Sate":        25000,
	}

	// Menampilkan harga dari setiap makanan
	fmt.Println("Harga Makanan:")
	for makanan, harga := range hargaMakanan {
		fmt.Printf("%s: Rp%d\n", makanan, harga)
	}

	fmt.Print("Harga makanan Sate = ", hargaMakanan["Sate"])
}
