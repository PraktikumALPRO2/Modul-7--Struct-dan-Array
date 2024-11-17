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
package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	radius     float64
	titikPusat Titik
}

func getJarak(lingkaran Lingkaran, titikSembarang Titik) float64 {
	return math.Sqrt(math.Pow(float64(lingkaran.titikPusat.x)-float64(titikSembarang.x), 2) + math.Pow(float64(lingkaran.titikPusat.y)-float64(titikSembarang.y), 2))
}

func cekDidalam(lingkaran Lingkaran, jarak float64) bool {
	return lingkaran.radius >= jarak
}

func main() {
	var lingkaran1, lingkaran2 Lingkaran
	var titikSembarang Titik
	var didalam1, didalam2 bool
	fmt.Scanln(&lingkaran1.titikPusat.x, &lingkaran1.titikPusat.y, &lingkaran1.radius)
	fmt.Scanln(&lingkaran2.titikPusat.x, &lingkaran2.titikPusat.y, &lingkaran2.radius)
	fmt.Scanln(&titikSembarang.x, &titikSembarang.y)
	didalam1 = cekDidalam(lingkaran1, getJarak(lingkaran1, titikSembarang))
	didalam2 = cekDidalam(lingkaran2, getJarak(lingkaran2, titikSembarang))
	switch {
	case didalam2 && didalam1:
		fmt.Print("Titik di dalam lingkaran 1 dan 2")
	case didalam1:
		fmt.Print("Titik di dalam lingkaran 1")
	case didalam2:
		fmt.Print("Titik di dalam lingkaran 2")
	default:
		fmt.Print("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Screenshot Source Code
![carbon](https://github.com/user-attachments/assets/60e7bb01-1036-4108-a2d7-25ebc29edb8e)


### Screenshot output
![Screenshot 2024-11-17 193752](https://github.com/user-attachments/assets/a96620b3-6aac-40ac-8454-3f16ce8d52bc)



### Deskripsi Program
Program ini menghasilkan deret Fibonacci hingga jumlah suku ke-n yang dimasukkan pengguna.

1. Fungsi FibonacciAngka menghitung nilai suku ke-n dalam deret Fibonacci secara rekursif.
2. Fungsi main meminta input n, kemudian menggunakan loop untuk mencetak deret Fibonacci dari suku pertama hingga ke-n.
Contoh: Jika pengguna memasukkan n = 5, program mencetak 0 1 1 2 3.

## Unguided - 2
### Study Case
Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilal. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa Informasi berikut:

a. Menampilkan keseluruhan isi dari array.

b. Menampilkan elemen-elemen array dengan indeks ganjil saja.

c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke - 0 adalah genap).

d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna.

e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid.. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil

f. Menampilkan rata-rata dari bilangan yang ada di dalam array.

g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.

h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.

#### Source Code
```go
package main

import (
	"fmt"
	"math"
)

func showMenu() {
	fmt.Println("1. menampilkan seluruh isi array")
	fmt.Println("2. Menampilkan elemen-elemen array dengan indeks ganjil saja.")
	fmt.Println("3. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap).")
	fmt.Println("4. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x")
	fmt.Println("5. Menghapus elemen array pada indeks tertentu.")
	fmt.Println("6. Menampilkan rata-rata dari bilangan yang ada di dalam array.")
	fmt.Println("7. Menampilkan standar deviasi atau simpangan baku")
	fmt.Println("8. Menampilkan frekuensi dari suatu bilangan")
	fmt.Println("9. Exit")
	fmt.Print("Pilih menu : ")
}

func showData(arr []int, skip int, expect int, skipZero bool) {
	for index, data := range arr {
		if skipZero && index == 0 {
			continue
		}
		if index%skip == expect {
			fmt.Println("Index ke - ", index, " : ", data)
		}
	}
}

func reRata(arr []int) float64 {
	var sum int
	for _, data := range arr {
		sum += data
	}
	return float64(sum) / float64(len(arr))
}

func simpBaku(arr []int) float64 {
	var sum float64
	for _, data := range arr {
		sum += math.Pow(float64(data)-reRata(arr), 2)
	}
	return math.Sqrt(float64(sum) / float64(len(arr)))
}

func frek(arr []int, bil int) int {
	var count int
	for _, data := range arr {
		if data == bil {
			count++
		}
	}
	return count
}

func removeIndex(index int, arr *[]int) {
	if len(*arr)-1 < index {
		fmt.Println("index out of range")
	} else {
		fmt.Println("item", (*arr)[index], "pada index", index, "telah dihapus\nisi Sekarang : ")
		*arr = append((*arr)[:index], (*arr)[index+1:]...)
		showData(*arr, 1, 0, false)
	}
}

func main() {
	var arraySize, input int
	fmt.Print("Masukkan Kapasitas array : ")
	fmt.Scan(&arraySize)
	var newArr []int = make([]int, arraySize)
	for index, _ := range newArr {
		fmt.Print("Masukkan index ke -", index, " : ")
		fmt.Scan(&newArr[index])
	}
	for input != 9 {
		showMenu()
		fmt.Scan(&input)
		switch input {
		case 1:
			showData(newArr, 1, 0, false)
		case 2:
			showData(newArr, 2, 1, false)
		case 3:
			showData(newArr, 2, 0, false)
		case 4:
			fmt.Print("Masukkan Kelipatan : ")
			fmt.Scan(&input)
			showData(newArr, input, 0, true)
		case 5:
			fmt.Print("Masukkan index yang dihapus : ")
			fmt.Scan(&input)
			removeIndex(input, &newArr)
		case 6:
			fmt.Println("Rerata dari isi array adalah :", reRata(newArr))
		case 7:
			fmt.Println("Deviasi dari isi array adalah :", simpBaku(newArr))
		case 8:
			fmt.Print("Masukkan bilangan yang akan dicek : ")
			fmt.Scan(&input)
			fmt.Println("Frekuansi dari", input, "pada array adalah :", frek(newArr, input))
		default:
			fmt.Println("Opsi salah")
		}
	}
}
```
### Screenshot Source Code
![carbon](https://github.com/user-attachments/assets/b6947558-523a-4c3f-8f76-a38a29d6a9cd)


### Screenshot Code
![1](https://github.com/user-attachments/assets/5ca9ee4f-bc3b-4993-ab86-23b2682e9c09)
![3](https://github.com/user-attachments/assets/29de9e02-3eba-45c7-a06c-31b64de2a331)
![2](https://github.com/user-attachments/assets/4b3787ed-9643-4cdb-9cc9-22002ed69a3e)



### Deskripsi Program
Program ini mencetak pola segitiga bintang secara rekursif berdasarkan nilai n yang dimasukkan pengguna.

1. Fungsi printBintang: Fungsi ini mencetak n baris bintang, dengan setiap baris memiliki jumlah bintang yang meningkat dari 1 hingga n. Fungsi ini dipanggil secara rekursif, mulai dari n = 1 hingga mencapai jumlah bintang sesuai nilai n.
2. Fungsi main: Meminta input n dari pengguna dan memanggil fungsi printBintang untuk mencetak pola bintang.

## Unguided - 3
### Study Case
**Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga.**

*Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja.
Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.*                                                        

#### Source Code
```go
package main

import "fmt"

type Pertandingan struct {
	klubA             string
	klubB             string
	hasilPertandingan []string
}

func hitung(dataClub Pertandingan, skor1, skor2 int) string {
	if skor1 > skor2 {
		return dataClub.klubA
	} else if skor1 == skor2 {
		return "Draw"
	} else {
		return dataClub.klubB
	}
}

func showHasil(dataClub Pertandingan) {
	for i, data := range dataClub.hasilPertandingan {
		fmt.Printf("Hasil %d : %s\n", i+1, data)
	}
	fmt.Println("Pertandingan selesai")
}

func main() {
	var skor1, skor2, count int
	var dataTanding Pertandingan
	fmt.Print("Klub A : ")
	fmt.Scan(&dataTanding.klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&dataTanding.klubB)
	for {
		fmt.Printf("Pertandingan %d : ", count+1)
		fmt.Scan(&skor1, &skor2)
		if skor1 >= 0 && skor2 >= 0 {
			dataTanding.hasilPertandingan = append(dataTanding.hasilPertandingan, hitung(dataTanding, skor1, skor2))
			count++
		} else {
			break
		}

	}
	showHasil(dataTanding)
}
```
### Screenshot Source Code
![carbon](https://github.com/user-attachments/assets/bb961cf8-5436-454f-896b-bca9acab8fe9)


### Screenshot Code
![Screenshot 2024-11-17 195422](https://github.com/user-attachments/assets/ca4c703f-eec2-4300-9f8e-7173ac09cac0)



### Deskripsi Program
Program ini mencari dan menampilkan semua faktor dari bilangan bulat positif n yang dimasukkan oleh pengguna.

1. Fungsi findFactors: Fungsi ini menggunakan rekursi untuk mencari faktor dari n, mulai dari 1 hingga n. Jika n habis dibagi i, maka i adalah faktor dari n, dan program mencetak i.
2. Fungsi main: Meminta input n dari pengguna dan memanggil findFactors untuk menampilkan faktor-faktor dari n.

## Unguided - 4
### Study Case
**Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan Isi array dan memeriksa apakah membentuk palindrom.**                             
*Modifikasi program tersebut dengan menambahkan fungsi palindrom. Tambahkan instruksi untuk memanggil fungsi tersebut dan menampilkan hasilnya pada program utama.*

#### Source Code
```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var temp_inp rune
	for *n < NMAX {
		fmt.Scanf("%c", &temp_inp)
		if temp_inp == '.' {
			break
		} else if temp_inp == ' ' {
			continue
		}
		t[*n] = temp_inp
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
		fmt.Print(" ")
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	var dat tabel

	for i := n - 1; i >= 0; i-- {
		dat[(n-1)-i] = t[i]
	}
	copy((*t)[:], dat[:])
}

func main() {
	var tab tabel
	var m int
	fmt.Print("Teks\t\t: ")
	isiArray(&tab, &m)
	balikanArray(&tab, m)
	fmt.Print("Reverse Teks\t: ")
	cetakArray(tab, m)
}
```
### Screenshot Source Code
![carbon](https://github.com/user-attachments/assets/ee1a9e5b-9c3d-4dbf-9f74-352249f25019)


### Screenshot Code
![Screenshot 2024-11-17 195907](https://github.com/user-attachments/assets/d5b4e434-b9e4-4786-883d-d52772487b99)



### Deskripsi Program
Program ini mencetak urutan bilangan secara berurutan dari n ke 1, lalu kembali dari 1 ke n menggunakan rekursi.

1. Fungsi printSekuensial: Jika n adalah 1, program mencetak 1. Jika tidak, program mencetak n, lalu memanggil dirinya sendiri dengan n - 1 untuk mencetak bilangan berikutnya secara berurutan, kemudian kembali mencetak n setelah rekursi selesai.
2. Fungsi main: Meminta input n dari pengguna dan memanggil fungsi printSekuensial untuk mencetak pola.

