<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL VII </strong></h2>
<h2 align="center"><strong> STRUCK & ARRAY </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Fadilah Salehah /2311102164<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh S.Kom., M.Kom. 
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------

## Daftar Isi



1. [Dasar Teori](#dasar-teori)
2. [Guided](#guided)
3. [Unguided](#unguided)

## Dasar Teori
### 1. Struktur (Struck)

Struktur atau struct adalah tipe data buatan yang memungkinkan pengguna untuk mengelompokkan berbagai jenis data ke dalam satu unit. Struktur digunakan untuk menyimpan informasi yang saling terkait tetapi mungkin memiliki tipe data yang berbeda. Elemen-elemen di dalam struktur disebut sebagai field atau anggota.

Karakteristik Struktur:

-Memiliki satu atau lebih field dengan tipe data yang berbeda.

-Struktur memungkinkan pengelompokan data yang logis, seperti data mahasiswa (nama, NIM, IPK).
### Contoh Deklarasi Struck:
```go
type Mahasiswa struct {
    Nama string
    NIM  string
    IPK  float64
}
```
### Penggunaan struck
- Deklarasi
```go
var mhs Mahasiswa
```
- Pengisian Data
```go
mhs.Nama = "Budi"
mhs.NIM = "12345678"
mhs.IPK = 3.75
```
- Akses Data
```go
fmt.Println(mhs.Nama, mhs.NIM, mhs.IPK)
```

### 2. Array
Array adalah tipe data terstruktur yang digunakan untuk menyimpan sekumpulan elemen dengan tipe data yang sama. Elemen-elemen dalam array diakses menggunakan indeks yang dimulai dari 0.

### Karakteristik Array:

- Hanya dapat menyimpan elemen dengan tipe data yang sama.
- Ukuran array harus ditentukan saat deklarasi.
- Indeks array dimulai dari nol (0).

### Contoh Deklarasi Array:
```go
var nilai [5]int // Array dengan 5 elemen bertipe integer
Pengisian dan Akses Elemen:
``` 
- Pengisian Elemen:
```go
nilai[0] = 90
nilai[1] = 85
nilai[2] = 88
```
- Akses Elemen:
```go
fmt.Println(nilai[0]) // Output: 90
```
### 3. Kombinasi Struct dan Array
Struktur dan array dapat digunakan bersama untuk mengelola data yang lebih kompleks, seperti daftar mahasiswa, daftar buku, atau catatan penjualan.

Contoh:
```go
type Mahasiswa struct {
    Nama string
    NIM  string
    IPK  float64
}

func main() {
    var dataMahasiswa [3]Mahasiswa

    dataMahasiswa[0] = Mahasiswa{"Budi", "12345678", 3.75}
    dataMahasiswa[1] = Mahasiswa{"Ani", "87654321", 3.90}
    dataMahasiswa[2] = Mahasiswa{"Tono", "56781234", 3.60}

    for _, mhs := range dataMahasiswa {
        fmt.Println(mhs.Nama, mhs.NIM, mhs.IPK)
    }
}
```
### Penjelasan:
- Deklarasi Struktur: Struktur Mahasiswa digunakan untuk menyimpan data seperti nama, NIM, dan IPK.
- Array Struktur: Array dataMahasiswa digunakan untuk menyimpan banyak objek Mahasiswa.
- Pengisian Data: Data diisi ke dalam array, di mana setiap elemen adalah objek Mahasiswa.

## Guided
### 1.  Program sederhana untuk menghitung lama waktu parkir berdasarkan waktu kedatangan dan waktu pulang
### Source Code 
```go
package main

import "fmt"

type waktu struct {
	jam, menit, detik int
}

func main() {
	var wParkir, wPulang, durasi waktu
	var dParkir, dPulang, lParkir int

	fmt.Scan(&wParkir.jam, &wParkir.menit, &wParkir.detik)
	fmt.Scan(&wPulang.jam, &wPulang.menit, &wPulang.detik)

	dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600 //konversi ke detik
	dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600 //detik
	lParkir = dPulang - dParkir                                   // detik dari pulang-datang

	durasi.jam = lParkir / 3600
	durasi.menit = lParkir % 3600 / 60
	durasi.detik = lParkir % 3600 % 60 //17
	fmt.Printf("Lama parkir: %d jam %d menit %d detik", durasi.jam, durasi.menit, durasi.detik)
}

```
#### Output Program
![image](https://github.com/user-attachments/assets/6e0d086a-e5dd-4515-8f68-f5457ed28b70)

### Deskripsi Program
Program di atas digunakan untuk menghitung durasi parkir kendaraan berdasarkan waktu datang dan waktu pulang yang dimasukkan oleh pengguna. Waktu tersebut dihitung dalam format jam, menit, dan detik, kemudian selisihnya diproses untuk mendapatkan lama parkir dalam format yang sama.


### 2. Program sederhana untuk validasi duplikasi nama pada daftar teman
### Source code
```go
//Guided 2 - Slice

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
#### Output Program
![image](https://github.com/user-attachments/assets/8ca7d85b-b82a-4262-ab06-07babc61fb33)

### Deskripsi Program
Program di atas untuk mengelola daftar teman menggunakan slice. Program memastikan tidak ada nama yang duplikat dengan memeriksa terlebih dahulu apakah nama yang ingin ditambahkan sudah ada di daftar. Jika nama sudah ada, program akan memberikan pesan pemberitahuan. Dengan demikian, daftar teman tetap unik.


### 3. Program sederhana untuk menampilkan daftar harga buah menggunakan map.
### Source code
```go
//Guided 3 - Map

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
#### Output Program
![image](https://github.com/user-attachments/assets/12e8e8a3-91b8-4962-83b3-a24a722a5edd)

### Deskripsi Program
Program ini digunakan untuk mengelola informasi harga buah. Nama buah digunakan sebagai kunci, sementara harga buah menjadi nilainya. Program ini menampilkan daftar harga semua buah yang tersimpan dalam map, serta menampilkan harga spesifik untuk buah tertentu.


## Unguided
### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥,𝑐𝑦) dengan radius 𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥,𝑦) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.

Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
### Source code
```go
// Fadilah Salehah
// 2311102164


package main

import (
	"fmt"
	"math"
)

// Struktur data untuk titik
type Titik struct {
	x, y int
}

// Struktur data untuk lingkaran
type Lingkaran struct {
	titikPusat Titik
	radius     int
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(p, q Titik) float64 {
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))
}

// Fungsi untuk mengecek apakah sebuah titik berada di dalam lingkaran
func diDalamLingkaran(lingkaran Lingkaran, titik Titik) bool {
	return jarak(lingkaran.titikPusat, titik) <= float64(lingkaran.radius)
}

func main() {
	var lingkaran1, lingkaran2 Lingkaran
	var titik Titik

	// Input untuk lingkaran pertama
	fmt.Print("Masukkan titik pusat dan radius lingkaran 1 (cx cy r): ")
	fmt.Scan(&lingkaran1.titikPusat.x, &lingkaran1.titikPusat.y, &lingkaran1.radius)

	// Input untuk lingkaran kedua
	fmt.Print("Masukkan titik pusat dan radius lingkaran 2 (cx cy r): ")
	fmt.Scan(&lingkaran2.titikPusat.x, &lingkaran2.titikPusat.y, &lingkaran2.radius)

	// Input untuk titik sembarang
	fmt.Print("Masukkan koordinat titik (x y): ")
	fmt.Scan(&titik.x, &titik.y)

	// Mengecek posisi titik terhadap lingkaran 1 dan lingkaran 2
	dalamLingkaran1 := diDalamLingkaran(lingkaran1, titik)
	dalamLingkaran2 := diDalamLingkaran(lingkaran2, titik)

	// Menentukan output berdasarkan posisi titik
	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```
#### Output Program
![image](https://github.com/user-attachments/assets/eac4bf07-e617-49f1-acc6-95477b6c35ac)

### Deskripsi Program
Program di atas digunakan untuk memeriksa posisi suatu titik terhadap dua lingkaran. Program akan menerima input berupa koordinat titik pusat dan radius dari dua lingkaran, serta koordinat titik sembarang. Berdasarkan perhitungan jarak, program menentukan apakah titik tersebut berada di dalam lingkaran pertama, lingkaran kedua, kedua lingkaran, atau di luar keduanya.


### 2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut:
a. Menampilkan keseluruhan isi dari array.

b. Menampilkan elemen-elemen array dengan indeks ganjil saja.

c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap).

d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna.

e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid.Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak       tampil

f. Menampilkan rata-rata dari bilangan yang ada di dalam array.

g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.

h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.

### Source code
```go
// Fadilah Salehah
// 2311102164

package main

import (
	"fmt"
	"math"
)

func main() {
	var jumlah int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&jumlah)

	dataArray := make([]int, jumlah)

	// Input elemen array
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&dataArray[i])
	}

	for {
		// Menu opsi
		fmt.Println("\n=== MENU ===")
		fmt.Println("a. Tampilkan seluruh elemen array")
		fmt.Println("b. Tampilkan elemen array pada indeks ganjil")
		fmt.Println("c. Tampilkan elemen array pada indeks genap")
		fmt.Println("d. Tampilkan elemen array pada indeks kelipatan tertentu")
		fmt.Println("e. Hapus elemen array berdasarkan indeks")
		fmt.Println("f. Hitung rata-rata elemen array")
		fmt.Println("g. Hitung standar deviasi elemen array")
		fmt.Println("h. Hitung frekuensi elemen tertentu")
		fmt.Println("i. Keluar")
		fmt.Print("Pilih menu: ")
		var pilihan string
		fmt.Scan(&pilihan)

		switch pilihan {
		case "a":
			// Tampilkan seluruh elemen array
			fmt.Println("Isi array:", dataArray)

		case "b":
			// Tampilkan elemen dengan indeks ganjil
			fmt.Println("Elemen pada indeks ganjil:")
			for i := 1; i < jumlah; i += 2 {
				fmt.Printf("Indeks %d: %d\n", i, dataArray[i])
			}

		case "c":
			// Tampilkan elemen dengan indeks genap
			fmt.Println("Elemen pada indeks genap:")
			for i := 0; i < jumlah; i += 2 {
				fmt.Printf("Indeks %d: %d\n", i, dataArray[i])
			}

		case "d":
			// Tampilkan elemen pada indeks kelipatan tertentu
			var kelipatan int
			fmt.Print("Masukkan nilai kelipatan: ")
			fmt.Scan(&kelipatan)
			fmt.Printf("Elemen pada indeks kelipatan %d:\n", kelipatan)
			for i := 0; i < jumlah; i++ {
				if i%kelipatan == 0 {
					fmt.Printf("Indeks %d: %d\n", i, dataArray[i])
				}
			}

		case "e":
			// Hapus elemen berdasarkan indeks
			var indeks int
			fmt.Print("Masukkan indeks yang ingin dihapus: ")
			fmt.Scan(&indeks)
			if indeks >= 0 && indeks < jumlah {
				dataArray = append(dataArray[:indeks], dataArray[indeks+1:]...)
				jumlah-- // Kurangi jumlah elemen
				fmt.Println("Array setelah penghapusan:", dataArray)
			} else {
				fmt.Println("Indeks tidak valid!")
			}

		case "f":
			// Hitung rata-rata elemen array
			total := 0
			for _, nilai := range dataArray {
				total += nilai
			}
			rataRata := float64(total) / float64(len(dataArray))
			fmt.Printf("Rata-rata: %.2f\n", rataRata)

		case "g":
			// Hitung standar deviasi elemen array
			total := 0
			for _, nilai := range dataArray {
				total += nilai
			}
			rataRata := float64(total) / float64(len(dataArray))

			var variansi float64
			for _, nilai := range dataArray {
				variansi += math.Pow(float64(nilai)-rataRata, 2)
			}
			variansi /= float64(len(dataArray))
			stdDeviasi := math.Sqrt(variansi)
			fmt.Printf("Standar deviasi: %.2f\n", stdDeviasi)

		case "h":
			// Hitung frekuensi elemen tertentu
			var elemen int
			fmt.Print("Masukkan elemen yang ingin dihitung frekuensinya: ")
			fmt.Scan(&elemen)
			frekuensi := 0
			for _, nilai := range dataArray {
				if nilai == elemen {
					frekuensi++
				}
			}
			fmt.Printf("Frekuensi elemen %d: %d kali\n", elemen, frekuensi)

		case "i":
			// Keluar dari program
			fmt.Println("Program selesai. Terima kasih!")
			return

		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi!")
		}
	}
}

```
#### Output Program
![image](https://github.com/user-attachments/assets/70439cf0-e037-4b3e-b3c2-42aeb5d70cac)
![image](https://github.com/user-attachments/assets/0bdddf11-c525-4ca2-8978-dd75907ae30e)
![image](https://github.com/user-attachments/assets/d6f40670-8ea7-4579-9e5a-9acc6593e166)

### Deskripsi Program
Program di atas digunakan untuk melakukan berbagai operasi pada data yang disimpan dalam array. Pengguna dapat memasukkan sejumlah elemen ke dalam array, kemudian menggunakan menu interaktif untuk menjalankan berbagai fungsi seperti menampilkan elemen tertentu (berdasarkan indeks), menghitung rata-rata, standar deviasi, frekuensi kemunculan suatu bilangan, atau menghapus elemen tertentu. Program berjalan secara interaktif hingga pengguna memilih untuk keluar.


### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.

### Source code
```go
// Fadilah Salehah
// 2311102164

package main

import "fmt"

func main() {
	// Deklarasi variabel
	var timA, timB string
	var skorTimA, skorTimB int
	var hasilPertandingan []string
	var nomorPertandingan int = 1

	// Input nama tim
	fmt.Print("Nama Tim A: ")
	fmt.Scan(&timA)
	fmt.Print("Nama Tim B: ")
	fmt.Scan(&timB)

	fmt.Println("\nMasukkan skor untuk setiap pertandingan. Masukkan skor negatif untuk menghentikan.")

	// Perulangan untuk menerima skor dan menentukan pemenang
	for {
		fmt.Printf("Pertandingan %d - Skor %s: ", nomorPertandingan, timA)
		fmt.Scan(&skorTimA)
		fmt.Printf("Pertandingan %d - Skor %s: ", nomorPertandingan, timB)
		fmt.Scan(&skorTimB)

		// Kondisi untuk menghentikan jika skor negatif
		if skorTimA < 0 || skorTimB < 0 {
			fmt.Println("Input skor dihentikan.")
			break
		}

		// Menentukan pemenang atau hasil seri
		if skorTimA > skorTimB {
			hasilPertandingan = append(hasilPertandingan, timA)
		} else if skorTimB > skorTimA {
			hasilPertandingan = append(hasilPertandingan, timB)
		} else {
			hasilPertandingan = append(hasilPertandingan, "Draw")
		}

		nomorPertandingan++
	}

	// Menampilkan hasil akhir pertandingan
	fmt.Println("\nRekap Hasil Pertandingan:")
	for i, hasil := range hasilPertandingan {
		fmt.Printf("Pertandingan %d: %s\n", i+1, hasil)
	}
}

```
#### Output Program
![image](https://github.com/user-attachments/assets/91c0e085-eb26-4400-a2d3-b920aa828800)

### Deskripsi Program
Program di atas digunakan untuk untuk mencatat skor pertandingan antara dua klub dan menentukan pemenang dari setiap pertandingan. Program akan terus menerima skor pertandingan secara berurutan hingga pengguna memasukkan skor negatif, yang menjadi tanda untuk mengakhiri pencatatan. Setelah itu, program menampilkan hasil dari semua pertandingan yang telah dimainkan.


### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom
### Source code
```go
// Fadilah Salehah
// 2311102164

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Konstanta dan tipe data
const MAKSIMUM int = 127

type arrayKarakter [MAKSIMUM]rune

// Fungsi untuk mengisi array berdasarkan input pengguna
func masukkanTeks(arr *arrayKarakter, panjang *int) {
	var teks string
	fmt.Println("Masukkan teks, diakhiri dengan tanda titik ('.'):")
	reader := bufio.NewReader(os.Stdin)
	teks, _ = reader.ReadString('\n')
	teks = strings.TrimSpace(teks)

	for _, karakter := range teks {
		if karakter == '.' || *panjang >= MAKSIMUM {
			break
		}
		arr[*panjang] = karakter
		*panjang++
	}
}

// Fungsi untuk mencetak isi array
func tampilkanTeks(arr arrayKarakter, panjang int) {
	for i := 0; i < panjang; i++ {
		fmt.Printf("%c", arr[i])
	}
	fmt.Println()
}

// Fungsi untuk membalik isi array
func balikArray(arr *arrayKarakter, panjang int) {
	for i := 0; i < panjang/2; i++ {
		arr[i], arr[panjang-1-i] = arr[panjang-1-i], arr[i]
	}
}

// Fungsi untuk memeriksa apakah array membentuk palindrom
func cekPalindrom(arr arrayKarakter, panjang int) bool {
	for i := 0; i < panjang/2; i++ {
		if arr[i] != arr[panjang-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var teks arrayKarakter
	var panjang int

	// Isi array menggunakan fungsi masukkanTeks
	masukkanTeks(&teks, &panjang)

	// Tampilkan teks yang dimasukkan
	fmt.Print("Teks: ")
	tampilkanTeks(teks, panjang)

	// Balikkan isi array
	balikArray(&teks, panjang)

	// Tampilkan teks setelah dibalik
	fmt.Print("Reverse teks: ")
	tampilkanTeks(teks, panjang)

	// Periksa apakah teks tersebut adalah palindrom
	if cekPalindrom(teks, panjang) {
		fmt.Println("Palindrom? true")
	} else {
		fmt.Println("Palindrom? false")
	}
}

```
#### Output Program
![image](https://github.com/user-attachments/assets/7f316f1b-d173-4391-8ac2-47e2913d1fc0)


### Deskripsi Program
Program di atas digunakan untuk memproses teks yang dimasukkan pengguna hingga mencapai tanda titik (.). Program menyimpan teks dalam sebuah array dan memiliki tiga fitur utama: menampilkan teks asli, membalik urutan teks, serta memeriksa apakah teks tersebut merupakan palindrom (teks yang sama saat dibaca dari depan maupun belakang).
## Kesimpulan 
Berdasarkan hasil praktikum, dapat disimpulkan bahwa:
1. Praktikum ini memberikan pemahaman tentang penggunaan array untuk menyimpan dan mengelola sejumlah data dalam program. Array memudahkan pengelolaan data yang memiliki tipe sama dalam satu wadah.
2. Dengan menggunakan struktur data seperti array dan tipe data custom (seperti struktur), kita dapat menangani data yang lebih kompleks dan mengorganisirnya dengan lebih baik dalam program.
3. Dalam praktikum, kita belajar bagaimana mengelola input dari pengguna untuk mengisi array, serta cara menangani input yang tidak valid menggunakan pemeriksaan error.
   
## Daftar Pustaka

[1] A. A. A. Donovan and B. W. Kernighan, *The Go Programming Language*. Boston, MA: Addison-Wesley, 2015.




