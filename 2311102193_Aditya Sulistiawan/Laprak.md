<p align="center">
  <strong>LAPORAN PRAKTIKUM</strong>
  <br>
  <strong>ALGORITMA DAN PEMROGRAMAN 2</strong>
  <br>
</p>

<br>

<p align="center">
  <strong>MODUL VII</strong>
  <br>
  <strong>Struck & Array</strong>
  <br>
</p>

<br>

<p align="center">
  <img src="https://github.com/user-attachments/assets/296eb3c2-bd6b-401f-8256-3661150ec39e" alt="Logo" width="200"/>
</p>

<br>

<p align="center">
  <strong>Disusun Oleh :</strong>
  <br>
 Aditya Sulistiawan
  <br>
  2311102193
  <br>
  IF-11-05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu :</strong><br>
  Arif Amrulloh, S.Kom., M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong>
  <br>
  <strong>FAKULTAS INFORMATIKA</strong>
  <br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong>
  <br>
  <strong>2024</strong>
</p>


# <strong> Dasar Teori </strong>

<strong><h2>Definisi Array</h2></strong>
Array adalah kumpulan data yang memiliki tipe yang sama dan disimpan dalam memori secara berurutan. Setiap elemen dalam array dapat diakses melalui indeks.

### <strong> Beberapa karakteristik utama dari Array:
- Semua elemen memiliki tipe data yang sama.
- Array memiliki ukuran tetap yang ditentukan saat deklarasi.
- Setiap elemen dapat diakses dengan menggunakan indeks numerik, di mana indeks pertama adalah 0.

<strong><h2>Definisi Struct</h2></strong>
Array adalah kumpulan data yang memiliki tipe yang sama dan disimpan dalam memori secara berurutan. Setiap elemen dalam array dapat diakses melalui indeks.

### <strong> Beberapa karakteristik utama dari Struct:
- Dapat menyimpan variabel dengan tipe data yang berbeda.
- Digunakan untuk membuat tipe data yang lebih kompleks, seperti data mahasiswa yang terdiri dari nama, usia, dan nilai.


# <strong> Guided </strong>
## Guided - 1
### Study Case
**Sebuah program yang menerima input waktu masuk parkir dan waktu keluar parkir dalam format jam, menit, dan detik. Program ini akan menghitung dan menampilkan lama waktu parkir dalam jam, menit, dan detik.** 

### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
//IF-11-05
package main 

import "fmt"

type waktu struct {
jam, menit, detik int
}

func main(){
var wParkir, wPulang, durasi waktu 
var dParkir, dPulang, lParkir int

fmt.Scan(&wParkir.jam, &wParkir.menit, &wParkir.detik) 
fmt.Scan(&wPulang.jam, &wPulang.menit, &wPulang.detik) 
dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600 
dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600 
lParkir = dPulang - dParkir

durasi.jam = lParkir / 3600 
durasi.menit = lParkir % 3600 / 60 
durasi.detik = lParkir % 3600 % 60
fmt.Printf("Lama parkir: %d jam %d menit %d detik", durasi.jam, durasi.menit, durasi.detik)
}
```

### Screenshot Output
![image](https://github.com/user-attachments/assets/2e8f2f66-e80b-4bd6-863d-2a15c542bbbe)


### Deskripsi Program
Program di atas menghitung durasi waktu parkir dalam jam, menit, dan detik. Pengguna memasukkan waktu masuk (saat memulai parkir) dan waktu keluar (saat meninggalkan parkir) dalam format jam, menit, dan detik, kemudian program akan menghitung selisih waktu antara kedua waktu tersebut untuk mendapatkan durasi parkir dalam format yang sama.


## Guided - 2
#### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
//IF-11-05
package main

import (
	"fmt"
)

// Fungsi untuk mengecek apakah nama sudah ada di dalam slice
func sudahAda(daftarTeman []string, nama string) bool {
	for _, teman := range daftarTeman {
		if teman == nama {
			return true
		}
	}
	return false
}

func main() {
	// Slice awal untuk daftar teman dengan beberapa data
	daftarTeman := []string{"Andi", "Budi", "Cici"}

	// Nama-nama baru yang ingin ditambahkan
	namaBaru := []string{"Dewi", "Budi", "Eka"}

	// Menambahkan nama baru hanya jika belum ada di daftar
	for _, nama := range namaBaru {
		if !sudahAda(daftarTeman, nama) {
			daftarTeman = append(daftarTeman, nama)
		} else {
			fmt.Println("Nama", nama, "sudah ada dalam daftar.")
		}
	}

	// Menampilkan daftar teman akhir
	fmt.Println("Daftar Teman:", daftarTeman)
}
```
### Screenshot Output
![image](https://github.com/user-attachments/assets/41a822b4-e57b-4d0d-9d88-be77b537dc9e)

### Deskripsi Program
Program berikut adalah program sederhana dalam bahasa Go yang melakukan beberapa tugas utama, yaitu memeriksa dan menambahkan nama baru ke dalam daftar teman yang sudah ada, dengan syarat nama tersebut belum ada dalam daftar. Jika nama sudah ada, program akan memberi tahu pengguna bahwa nama tersebut telah terdaftar.

## Guided - 3
#### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
//IF-11-05
//Guided 2 - Map

import (
	"fmt"
)

func main() {
	// Membuat map dengan nama buah sebagai kunci dan harga sebagai nilai
	hargaBuah := map[string]int{
		"Apel":  5000,
		"Pisang": 3000,
		"Mangga": 7000,
	}

	// Menampilkan harga dari setiap buah
	fmt.Println("Harga Buah:")
	for buah, harga := range hargaBuah {
		fmt.Printf("%s: Rp%d\n", buah, harga)
	}

	fmt.Print("Harga buah Mangga = ", hargaBuah["Mangga"])
}
```
### Screenshot Output
![image](https://github.com/user-attachments/assets/9ad9ce96-7aa3-441f-aee7-5dd903c01450)


### Deskripsi Program

Program ini menggunakan struktur data map dalam bahasa Go untuk menyimpan dan menampilkan harga berbagai buah. Dalam map, setiap nama buah digunakan sebagai kunci, dan harga buah tersebut digunakan sebagai nilai.


# <strong> Unguided </strong>
## Unguided - 1
### Study Case
**Deret fibonacci adalah sebuah deret dengan nilai suku ke-O dan ke-1 adalah O dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan S=Sn-1+Sn-2 Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonассі tersebut.**

#### Source Code
```go
// Rangga Pradarrell Fathi
// 2311102200
// IF-11-05
package main

import "fmt"

func FibonacciAngka(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return FibonacciAngka(n-1) + FibonacciAngka(n-2)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah suku Fibonacci yang ingin dicetak: ")
	fmt.Scanln(&n)

	fmt.Println("Deret Fibonacci hingga suku ke-", n)
	for i := 0; i < n; i++ {
		fmt.Print(FibonacciAngka(i), " ")
	}
	fmt.Println()
}
```
### Screenshot output
![Screenshot 2024-11-03 151606](https://github.com/user-attachments/assets/b97174e8-8961-4852-b73b-d89495c866ff)


### Deskripsi Program
Program ini menghasilkan deret Fibonacci hingga jumlah suku ke-n yang dimasukkan pengguna.

1. Fungsi FibonacciAngka menghitung nilai suku ke-n dalam deret Fibonacci secara rekursif.
2. Fungsi main meminta input n, kemudian menggunakan loop untuk mencetak deret Fibonacci dari suku pertama hingga ke-n.
Contoh: Jika pengguna memasukkan n = 5, program mencetak 0 1 1 2 3.

## Unguided - 2
### Study Case
**Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.**

#### Source Code
```go
// Rangga Pradarrell Fathi
// 2311102200
// IF-11-05
package main

import "fmt"

func printBintang(n int) {
	if n == 0 {
		return
	}
	printBintang(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scanln(&n)

	printBintang(n)
}
```
### Screenshot Code
![Screenshot 2024-11-03 151746](https://github.com/user-attachments/assets/7196b886-e329-4593-8bba-2c06ee01cf8e)


### Deskripsi Program
Program ini mencetak pola segitiga bintang secara rekursif berdasarkan nilai n yang dimasukkan pengguna.

1. Fungsi printBintang: Fungsi ini mencetak n baris bintang, dengan setiap baris memiliki jumlah bintang yang meningkat dari 1 hingga n. Fungsi ini dipanggil secara rekursif, mulai dari n = 1 hingga mencapai jumlah bintang sesuai nilai n.
2. Fungsi main: Meminta input n dari pengguna dan memanggil fungsi printBintang untuk mencetak pola bintang.

## Unguided - 3
### Study Case
**Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.**                                        
*Masukan terdiri dari sebuah bilangan bulat positif N.*                                                        
*Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).*

#### Source Code
```go
// Rangga Pradarrell Fathi
// 2311102200
// IF-11-05
package main

import "fmt"

func findFactors(n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Printf("%d ", i)
	}
	findFactors(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	fmt.Printf("Faktor dari %d adalah: ", n)
	findFactors(n, 1)
	fmt.Println()
}
```
### Screenshot Code
![Screenshot 2024-11-03 151842](https://github.com/user-attachments/assets/9c358852-5498-4db0-81df-c70012166e43)


### Deskripsi Program
Program ini mencari dan menampilkan semua faktor dari bilangan bulat positif n yang dimasukkan oleh pengguna.

1. Fungsi findFactors: Fungsi ini menggunakan rekursi untuk mencari faktor dari n, mulai dari 1 hingga n. Jika n habis dibagi i, maka i adalah faktor dari n, dan program mencetak i.
2. Fungsi main: Meminta input n dari pengguna dan memanggil findFactors untuk menampilkan faktor-faktor dari n.

## Unguided - 4
### Study Case
**Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu.**                             
*Masukan terdiri dari sebuah bilangan bulat positif N.*                                                                 
*Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N.*

#### Source Code
```go
// Rangga Pradarrell Fathi
// 2311102200
// IF-11-05
package main

import "fmt"

func printSekuensial(n int) {
	if n == 1 {
		fmt.Print(n, " ")
	} else {
		fmt.Print(n, " ")
		printSekuensial(n - 1)
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&n)

	fmt.Print("Keluaran: ")
	printSekuensial(n)
	fmt.Println()
}
```
### Screenshot Code
![Screenshot 2024-11-03 151919](https://github.com/user-attachments/assets/b64f5293-30c3-4bf2-ad4f-2164e531efbc)


### Deskripsi Program
Program ini mencetak urutan bilangan secara berurutan dari n ke 1, lalu kembali dari 1 ke n menggunakan rekursi.

1. Fungsi printSekuensial: Jika n adalah 1, program mencetak 1. Jika tidak, program mencetak n, lalu memanggil dirinya sendiri dengan n - 1 untuk mencetak bilangan berikutnya secara berurutan, kemudian kembali mencetak n setelah rekursi selesai.
2. Fungsi main: Meminta input n dari pengguna dan memanggil fungsi printSekuensial untuk mencetak pola.

## Unguided - 5
### Study Case
**Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.**                                 
*Masukan terdiri dari sebuah bilangan bulat positif N.*                                        
*Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.*

#### Source Code
```go
// Rangga Pradarrell Fathi
// 2311102200
// IF-11-05
package main

import "fmt"

func cetakBilanganGanjil(n int) {
	if n >= 1 {
		cetakBilanganGanjil(n - 2)
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&n)

	fmt.Println("Barisan bilangan ganjil dari 1 hingga", n, "adalah:")
	cetakBilanganGanjil(n)
}
```
### Screenshot Code
![Screenshot 2024-11-03 152016](https://github.com/user-attachments/assets/57d34524-189b-4327-bfae-035aad3bfc99)


### Deskripsi Program
Program ini mencetak bilangan ganjil dari 1 hingga n secara berurutan menggunakan rekursi.

1. Fungsi cetakBilanganGanjil: Jika n lebih besar atau sama dengan 1, fungsi ini memanggil dirinya sendiri dengan nilai n - 2 (mengurangi 2 setiap kali untuk mendapatkan bilangan ganjil sebelumnya), lalu mencetak nilai n setelah rekursi. Fungsi ini memastikan bilangan ganjil dicetak dalam urutan menaik.
2. Fungsi main: Meminta input n dari pengguna dan memanggil cetakBilanganGanjil untuk mencetak bilangan ganjil dari 1 hingga n.

## Unguided - 6
### Study Case
**Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan.**                                     
*Masukan terdiri dari bilangan bulat x dan y.*                                        
*Keluaran terdiri dari hasil x dipangkatkan y.*    

#### Source Code
```go
// Rangga Pradarrell Fathi
// 2311102200
// IF-11-05
package main

import "fmt"

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	}
	hasil := 0
	for i := 0; i < y; i++ {
		if i == 0 {
			hasil = x
		} else {
			hasil += x 
		}
	}
	return hasil
}

func main() {
	var x, y int
	fmt.Scan(&x)
	fmt.Scanln(&y)

	hasil := pangkat(x, y)
	fmt.Printf("%d pangkat %d = %d\n", x, y, hasil)
}
```
### Screenshot Code
![Screenshot 2024-11-03 212955](https://github.com/user-attachments/assets/99dd1e1d-e8a8-41b6-81fa-a2d7ac049834)


### Deskripsi Program
Program ini seharusnya menghitung pangkat dari bilangan x dengan eksponen y, tetapi terdapat kesalahan dalam implementasi fungsi pangkat.

1. Fungsi pangkat(x, y int) int: Fungsi ini dirancang untuk menghitung hasil x pangkat y, tetapi menggunakan penjumlahan berulang alih-alih perkalian berulang. Logika dalam fungsi ini hanya menjumlahkan x sebanyak y kali, yang sebenarnya menghasilkan x * y dan bukan x^y.
2. Fungsi main: Meminta pengguna memasukkan nilai x dan y, kemudian memanggil fungsi pangkat untuk menghitung hasil, dan menampilkannya.
