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
array digunakan untuk menyimpan beberapa nilai dengan tipe data yang sama ke dalam satu variabel

### <strong> Beberapa karakteristik utama dari Array:
- Semua elemen memiliki tipe data yang sama.
- Array memiliki ukuran tetap yang ditentukan saat deklarasi.
- Setiap elemen dapat diakses dengan menggunakan indeks numerik, di mana indeks pertama adalah 0.

<strong><h2>Definisi Struct</h2></strong>
Digunakan untuk menggabungkan nilai-nilai dengan tipe data yang berbeda ke dalam satu variabel. Struct dapat berisi berbagai tipe data, seperti integer, string, dan boolean. Struct mirip dengan kelas dalam pemrograman berorientasi objek, di mana Anda dapat menentukan tipe kompleks Anda sendiri dengan beberapa bidang.

### <strong> Mengakses nilai struct:
Untuk mengakses nilai milik suatu instance dari struct, Anda cukup menggunakan notasi titik diikuti dengan nama field. Perhatikan contoh berikut:

```go
fmt.Println(person.name)
fmt.Println(person.age)
fmt.Println(person.email)
```
Sebagai hasil dari operasi di atas, kita melihat output berikut:

```go
Andy Brinker
40
andy@example.com
```



# <strong> Guided </strong>
## Guided - 1
### Study Case
**Sebuah program yang menerima input waktu masuk parkir dan waktu keluar parkir dalam format jam, menit, dan detik. Program ini akan menghitung dan menampilkan lama waktu parkir dalam jam, menit, dan detik.** 

### Source Code
```go
//Aditya Sulistiawan
//2311102193
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
### Screenshot Source Code
![Source code](https://github.com/user-attachments/assets/909cf2e2-3f8a-4ad9-9191-0b01436a05ab)


### Screenshot Output
![Gui1](https://github.com/user-attachments/assets/98694a45-5f6a-4093-bdc8-bd04d1868d99)


#### Deskripsi Program
Program di atas menghitung berapa lama kendaraan parkir dalam jam, menit, dan detik. Pengguna harus menginput waktu masuk (saat mulai parkir) dan waktu keluar (saat selesai parkir) dalam format jam, menit, dan detik. Kemudian program akan menghitung berapa lama mobil parkir dengan menghitung selisih waktu antara kedua waktu tersebut dalam format yang sama.

#### Algoritma Program
1. Mulai
2. Deklarasi Variabel
3. Input Data
   - Minta pengguna menginput jam, menit, dan detik untuk waktu masuk 
   - Minta pengguna menginput jam, menit, dan detik untuk waktu keluar 
4. Hitung Total Detik
5. Hitung Lama Parkir
6. Output Hasil
7. Selesai

cara kerja program : 
Program ini dibuat untuk menghitung berapa lama mobil parkir berdasarkan waktu masuk dan keluar yang dimasukkan oleh pengguna. Langkah pertama adalah memasukkan waktu parkir dan waktu pulang dalam format jam, menit, dan detik. Setelah itu, program akan mengubah waktu tersebut menjadi jumlah detik agar perhitungan lebih mudah. Dengan begini, program dapat menghitung beda waktu antara pulang dan parkir, lalu diubah ke format jam, menit, dan detik. Akhirnya, pengguna dapat melihat hasil parkir mereka dengan jelas.

## Guided - 2
#### Source Code
```go
//Guided 2 - Map
package main 

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
### Screenshot Source Code
![carbon](https://github.com/user-attachments/assets/0c973ea0-6d93-447f-80b0-6fd279a7c362)


### Screenshot Output
![1](https://github.com/user-attachments/assets/4591490d-2baf-4ada-acb8-88b6b6caec23)


### Deskripsi Program
Program mencetak daftar harga semua buah menggunakan perulangan for yang mengulangi setiap elemen dalam peta. Setiap pasangan kunci dan nilai dicetak menggunakan fungsi fmt.Printf dengan format yang rapi. Terakhir, program secara khusus mengakses harga buah "Mangga" menggunakan kunci langsung dari peta dan mencetaknya dengan pernyataan fmt.Print. Program berjalan dengan baik dan cocok untuk mengelola data berpasangan seperti ini.

### Algoritma Program
1. Buat daftar yang menyimpan nama buah dan harganya.  
2. Masukkan data harga untuk Apel, Pisang, dan Mangga.  
3. Cetak tulisan "Harga Buah:" untuk memberi tahu pengguna.  
4. Tampilkan semua nama buah beserta harganya satu per satu.  
5. Tampilkan harga buah Mangga secara khusus.  

Cara kerja program 
1. Program membuat daftar berisi nama buah dan harganya.  
2. Program menampilkan semua buah dan harga satu per satu.  
3. Program menunjukkan harga buah Mangga secara langsung.  

## Guided - 3
#### Source Code
```go
// Guided 2 - Slice
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
### Screenshot Source Code
![carbon](https://github.com/user-attachments/assets/67ad822f-ceba-4142-9d78-34eab89199ee)


### Screenshot Output
![Screenshot 2024-11-17 173007](https://github.com/user-attachments/assets/7291f717-2869-44d7-b7df-cd5f996cb439)



### Deskripsi Program

Program ini menggunakan struktur data map dalam bahasa Go untuk menyimpan dan menampilkan harga berbagai buah. Dalam map, setiap nama buah digunakan sebagai kunci, dan harga buah tersebut digunakan sebagai nilai.

### Algoritma Program
1. Buat fungsi untuk mengecek apakah nama ada di daftar.  
2. Buat daftar teman awal.  
3. Siapkan nama-nama baru yang ingin ditambahkan.  
4. Periksa setiap nama baru:  
   - Jika belum ada, tambahkan ke daftar.  
   - Jika sudah ada, tampilkan pesan.  
5. Cetak daftar teman akhir.  


Cara kerja program
Program dimulai dengan daftar teman awal. Lalu, setiap nama baru diperiksa apakah sudah ada dalam daftar menggunakan fungsi cek. Nama yang tidak ada ditambahkan, sedangkan yang sudah ada menunjukkan pesan. Akhirnya, daftar teman yang diperbarui ditampilkan.


# <strong> Unguided </strong>
## Unguided - 1
### Study Case
**Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.**

_Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2"._

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
### Screenshot Source Code


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
### Screenshot Source Code


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
### Screenshot Source Code


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
### Screenshot Source Code


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
### Screenshot Source Code


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
### Screenshot Source Code


### Screenshot Code
![Screenshot 2024-11-03 212955](https://github.com/user-attachments/assets/99dd1e1d-e8a8-41b6-81fa-a2d7ac049834)


### Deskripsi Program
Program ini seharusnya menghitung pangkat dari bilangan x dengan eksponen y, tetapi terdapat kesalahan dalam implementasi fungsi pangkat.

1. Fungsi pangkat(x, y int) int: Fungsi ini dirancang untuk menghitung hasil x pangkat y, tetapi menggunakan penjumlahan berulang alih-alih perkalian berulang. Logika dalam fungsi ini hanya menjumlahkan x sebanyak y kali, yang sebenarnya menghasilkan x * y dan bukan x^y.
2. Fungsi main: Meminta pengguna memasukkan nilai x dan y, kemudian memanggil fungsi pangkat untuk menghitung hasil, dan menampilkannya.
