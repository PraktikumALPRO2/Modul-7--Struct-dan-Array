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
  Lutfiana Isnaeni L /2311102165<br>
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
![image](https://github.com/user-attachments/assets/11d29e57-7089-4411-91c6-328ab1c40e17)

### Deskripsi Program
Program di atas digunakan untuk menghitung durasi parkir kendaraan berdasarkan waktu datang dan waktu pulang yang dimasukkan oleh pengguna. Waktu tersebut dihitung dalam format jam, menit, dan detik, kemudian selisihnya diproses untuk mendapatkan lama parkir dalam format yang sama.
### Algoritma Program:
1.	Definisi Struktur Data
   
    •	Buat struktur `waktu` dengan atribut `jam`, `menit`, dan `detik`.
  	
2.	Deklarasi Variabel
   
    •	Tentukan variabel untuk menyimpan waktu kedatangan `(wParkir)` dan waktu kepulangan `(wPulang)`.
  	
    •	Sediakan variabel `durasi` untuk menyimpan hasil selisih waktu dalam bentuk struktur `waktu`.
  	
    •	Tambahkan variabel `dParkir` dan `dPulang` untuk menyimpan waktu dalam bentuk total detik, serta `lParkir` untuk menyimpan selisih total detik.
  	
3.	Masukan Data
   
    •	Terima input berupa jam, menit, dan detik untuk waktu kedatangan `(wParkir)` dan `waktu` kepulangan `(wPulang)`.
  	
4.	Konversi Waktu ke Detik
   
    •	Hitung waktu kedatangan dalam detik:
  	
     `dParkir = wParkir.detik + (wParkir.menit * 60) + (wParkir.jam * 3600).`
   
    •	Hitung waktu kepulangan dalam detik:
  	
      `dPulang = wPulang.detik + (wPulang.menit * 60) + (wPulang.jam * 3600).`
  	
5.	Hitung Selisih Waktu
    
    •	Dapatkan selisih waktu dalam detik:
  	
      `lParkir = dPulang - dParkir.`
  	
    •	Konversi kembali total detik ke format jam, menit, dan detik:
  	
      -	Jam: `durasi.jam = lParkir / 3600.`
        
      -	Menit: `durasi.menit = (lParkir % 3600) / 60.`
        
      -	Detik: `durasi.detik = lParkir % 60.`
        
6.	Tampilkan Hasil
    
    •	Cetak lama parkir dalam format:
   	
        Lama parkir: `X jam Y menit Z detik.`


### Cara Kerja Program:
1.	Input Data
   
    •	Pengguna memasukkan waktu datang dan waktu pulang masing-masing dalam tiga komponen: jam, menit, dan detik.
  	
2.	Konversi ke Total Detik
   
    •	Kedua waktu diubah menjadi jumlah total detik sejak pukul 00:00 untuk memudahkan perhitungan.
  	
3.	Menghitung Lama Parkir
   
    •	Selisih waktu dihitung dalam bentuk detik.

4.	Konversi Kembali ke Format Jam, Menit, dan Detik
   
    •	Selisih waktu yang sebelumnya dalam bentuk detik dikembalikan ke format jam, menit, dan detik.

5.	Output Hasil
    
    •	Program menampilkan lama parkir dalam format jam, menit, dan detik.




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
![image](https://github.com/user-attachments/assets/9249da54-523b-4cfe-ac91-1a823aec22fc)

### Deskripsi Program
Program di atas untuk mengelola daftar teman menggunakan slice. Program memastikan tidak ada nama yang duplikat dengan memeriksa terlebih dahulu apakah nama yang ingin ditambahkan sudah ada di daftar. Jika nama sudah ada, program akan memberikan pesan pemberitahuan. Dengan demikian, daftar teman tetap unik.
### Algoritma Program:
1.	Inisialisasi Daftar Awal
   
    •	Buat slice `daftarTeman` dengan beberapa nama sebagai data awal.
  	
2.	Nama Baru untuk Ditambahkan
   
    •	Definisikan slice namaBaru yang berisi nama-nama yang akan ditambahkan ke dalam `daftarTeman`.
  	
3.	Fungsi Pemeriksaan Nama
   
    •	Buat fungsi sudahAda untuk memeriksa apakah nama sudah terdapat di `daftarTeman`. Fungsi ini akan:
  	
      -	Melakukan iterasi melalui elemen-elemen dalam `daftarTeman`.
  	
      -	Jika nama ditemukan, mengembalikan `true`; jika tidak, mengembalikan `false`.
  	
3.	Proses Menambahkan Nama
   
    •	Iterasi setiap nama dalam `namaBaru`:
  	
      o	Jika nama belum ada di daftarTeman (dengan memanfaatkan fungsi sudahAda), tambahkan nama tersebut menggunakan `append`.
  	
      o	Jika nama sudah ada, tampilkan pesan bahwa nama tersebut sudah ada dalam daftar.
  	
5.	Menampilkan Hasil Akhir
    
    •	Cetak daftar teman setelah nama-nama baru diproses.


### Cara Kerja Program:
1.	Mempersiapkan Daftar Awal
   
    •	Program memulai dengan daftar teman awal, misalnya `Andi`, `Budi`, dan `Cici`.
  	
2.	Memproses Nama Baru
   
    •	Nama-nama dalam slice `namaBaru` (misalnya `Dewi`, `Budi`, `Eka`) akan diproses satu per satu.
  	
4.	Pengecekan Keberadaan Nama
   
    •	Setiap nama dalam namaBaru diperiksa dengan fungsi sudahAda untuk memastikan apakah nama tersebut sudah ada di `daftarTeman`.
  	
5.	Penambahan Nama
   
    •	Jika nama belum ada, program menambahkannya ke daftar menggunakan `append`.
  	
    •	Jika nama sudah ada, program memberikan pesan, misalnya: `Nama Budi sudah ada dalam daftar`.
  	
6.	Hasil Akhir
    
    •	Setelah semua nama diproses, daftar teman yang diperbarui ditampilkan.



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
![image](https://github.com/user-attachments/assets/a05eeeef-cc24-469b-8b87-e954099d254a)

### Deskripsi Program
Program ini digunakan untuk mengelola informasi harga buah. Nama buah digunakan sebagai kunci, sementara harga buah menjadi nilainya. Program ini menampilkan daftar harga semua buah yang tersimpan dalam map, serta menampilkan harga spesifik untuk buah tertentu.

### Algoritma Program
1.	Inisialisasi Map
   
    •	Buat map bernama `hargaBuah` yang berisi pasangan kunci dan nilai:
  	
      -	Kunci: nama buah (contoh:"`Apel"`, `"Pisang"`).
  	
      -	Nilai: harga buah dalam bentuk angka.
  	
2.	Menampilkan Harga Semua Buah
   
    •	Gunakan perulangan `for` dengan range untuk menelusuri setiap elemen dalam map hargaBuah:
  	
      -	Cetak nama buah dan harganya dengan format `buah: Rp harga`.
  	
3.	Menampilkan Harga Buah Tertentu
   
    •	Ambil harga buah tertentu secara langsung menggunakan kunci, misalnya `"Mangga"`, dan tampilkan hasilnya.


### Cara Kerja Program
1.	Deklarasi Map
   
    •	Program dimulai dengan mendeklarasikan map bernama hargaBuah, yang menghubungkan nama buah dengan harga.
  	
2.	Mengisi Data ke Map
   
    •	Map hargaBuah diisi dengan data:
  	
      o	`"Apel"` dengan harga `5000`.
  	
      o	`"Pisang"` dengan harga `3000`.
  	
      o	`"Mangga"` dengan harga `7000`.
  	
3.	Menampilkan Semua Harga Buah
   
    •	Program menggunakan perulangan for untuk mengiterasi isi map. Setiap iterasi mencetak nama buah dan harganya.
  	
4.	Mengakses Harga Spesifik
   
    •	Harga buah "Mangga" diambil menggunakan kunci `"Mangga"` dan ditampilkan dengan format Harga `buah Mangga = harga`.


## Unguided
### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥,𝑐𝑦) dengan radius 𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥,𝑦) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.

Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
### Source code
```go
// Lutfiana Isnaeni Lathifah
// 2311102165

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
![image](https://github.com/user-attachments/assets/bfa0c9b1-e127-4707-80c1-1c2af02314e6)

### Deskripsi Program
Program di atas digunakan untuk memeriksa posisi suatu titik terhadap dua lingkaran. Program akan menerima input berupa koordinat titik pusat dan radius dari dua lingkaran, serta koordinat titik sembarang. Berdasarkan perhitungan jarak, program menentukan apakah titik tersebut berada di dalam lingkaran pertama, lingkaran kedua, kedua lingkaran, atau di luar keduanya.

### Algoritma Program:
1.	Definisi Struktur Data
   
	•	Titik: Struktur untuk menyimpan koordinat `(x, y)` sebuah titik.

	•	Lingkaran: Struktur untuk menyimpan koordinat pusat `(cx, cy)` dan radius `r` dari lingkaran.

2.	Fungsi `jarak`
   
	•	Menghitung jarak antara dua titik
 
3.	Fungsi `diDalamLingkaran`
   
	•	Mengecek apakah sebuah titik berada di dalam lingkaran:

	o	Periksa apakah jarak dari titik ke pusat lingkaran lebih kecil atau sama dengan radius lingkaran.

4.	Input Data
   
	•	Masukkan koordinat titik pusat dan radius lingkaran pertama.

	•	Masukkan koordinat titik pusat dan radius lingkaran kedua.

	•	Masukkan koordinat titik sembarang yang akan diuji.

5.	Pengecekan Posisi
    
	•	Gunakan fungsi diDalamLingkaran untuk memeriksa posisi titik sembarang:

	-	Apakah berada di dalam lingkaran pertama.
	-	Apakah berada di dalam lingkaran kedua.
	-	Apakah berada di dalam kedua lingkaran.
	-	Apakah berada di luar keduanya.

6.	Output Hasil
    
	•	Program mencetak posisi titik berdasarkan hasil pengecekan:

	-	"Titik di dalam lingkaran 1 dan 2".
	-	"Titik di dalam lingkaran 1".
	-	"Titik di dalam lingkaran 2".
	-	"Titik di luar lingkaran 1 dan 2".


### Cara Kerja Program:
1.	Inisialisasi:
   
	•	Deklarasikan variabel untuk menyimpan data lingkaran pertama, lingkaran kedua, dan titik sembarang.

2.	Input Data:
   
	•	Pengguna diminta memasukkan data lingkaran pertama (koordinat pusat dan radius).

	•	Pengguna memasukkan data lingkaran kedua (koordinat pusat dan radius).

	•	Pengguna memasukkan koordinat titik sembarang yang ingin diperiksa.

3.	Perhitungan Jarak:
   
	•	Hitung jarak antara titik sembarang dan titik pusat lingkaran menggunakan fungsi `jarak`.

4.	Pengecekan Posisi:
   
	•	Gunakan fungsi `diDalamLingkaran` untuk memeriksa apakah jarak tersebut lebih kecil atau sama dengan radius lingkaran pertama atau kedua.

5.	Output Hasil:
    
	•	Berdasarkan hasil pengecekan:

	-	Jika titik berada di dalam kedua lingkaran: cetak "Titik di dalam lingkaran 1 dan 2".
	-	Jika hanya berada di lingkaran pertama: cetak "Titik di dalam lingkaran 1".
	-	Jika hanya berada di lingkaran kedua: cetak "Titik di dalam lingkaran 2".
	-	Jika tidak berada di keduanya: cetak "Titik di luar lingkaran 1 dan 2".



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
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	array := make([]int, n)

	// Input array
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	for {
		// Menu
		fmt.Println("\n=== MENU ===")
		fmt.Println("a. Tampilkan keseluruhan isi array")
		fmt.Println("b. Tampilkan elemen array dengan indeks ganjil")
		fmt.Println("c. Tampilkan elemen array dengan indeks genap")
		fmt.Println("d. Tampilkan elemen array dengan indeks kelipatan x")
		fmt.Println("e. Hapus elemen array pada indeks tertentu")
		fmt.Println("f. Tampilkan rata-rata dari bilangan dalam array")
		fmt.Println("g. Tampilkan standar deviasi dari bilangan dalam array")
		fmt.Println("h. Tampilkan frekuensi dari bilangan tertentu")
		fmt.Println("i. Keluar")
		fmt.Print("Pilih opsi: ")
		var opsi string
		fmt.Scan(&opsi)

		switch opsi {
		case "a":
			// Tampilkan seluruh isi array
			fmt.Println("Isi array:", array)

		case "b":
			// Tampilkan elemen array dengan indeks ganjil
			fmt.Println("Elemen dengan indeks ganjil:")
			for i := 1; i < n; i += 2 {
				fmt.Printf("Index %d: %d\n", i, array[i])
			}

		case "c":
			// Tampilkan elemen array dengan indeks genap
			fmt.Println("Elemen dengan indeks genap:")
			for i := 0; i < n; i += 2 {
				fmt.Printf("Index %d: %d\n", i, array[i])
			}

		case "d":
			// Tampilkan elemen array dengan indeks kelipatan x
			var x int
			fmt.Print("Masukkan nilai x: ")
			fmt.Scan(&x)
			fmt.Println("Elemen dengan indeks kelipatan", x, ":")
			for i := 0; i < n; i++ {
				if i%x == 0 {
					fmt.Printf("Index %d: %d\n", i, array[i])
				}
			}

		case "e":
			// Hapus elemen array pada indeks tertentu
			var index int
			fmt.Print("Masukkan indeks yang ingin dihapus: ")
			fmt.Scan(&index)
			if index >= 0 && index < n {
				array = append(array[:index], array[index+1:]...)
				n-- // Kurangi jumlah elemen
				fmt.Println("Array setelah penghapusan:", array)
			} else {
				fmt.Println("Indeks tidak valid!")
			}

		case "f":
			// Tampilkan rata-rata
			sum := 0
			for _, val := range array {
				sum += val
			}
			avg := float64(sum) / float64(len(array))
			fmt.Printf("Rata-rata: %.2f\n", avg)

		case "g":
			// Tampilkan standar deviasi
			sum := 0
			for _, val := range array {
				sum += val
			}
			avg := float64(sum) / float64(len(array))

			var variance float64
			for _, val := range array {
				variance += math.Pow(float64(val)-avg, 2)
			}
			variance /= float64(len(array))
			stdDev := math.Sqrt(variance)
			fmt.Printf("Standar deviasi: %.2f\n", stdDev)

		case "h":
			// Tampilkan frekuensi dari bilangan tertentu
			var num int
			fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
			fmt.Scan(&num)
			count := 0
			for _, val := range array {
				if val == num {
					count++
				}
			}
			fmt.Printf("Frekuensi bilangan %d: %d kali\n", num, count)

		case "i":
			// Keluar
			fmt.Println("Program selesai.")
			return

		default:
			fmt.Println("Opsi tidak valid!")
		}
	}
}
```
#### Output Program
![image](https://github.com/user-attachments/assets/9d12702f-6554-4ee5-8654-bade420d22c8)
![image](https://github.com/user-attachments/assets/67b3c687-b23f-4572-a705-77a29f4b2eff)
![image](https://github.com/user-attachments/assets/526e0552-378a-41b4-83df-89f460d62c3b)

### Deskripsi Program
Program di atas digunakan untuk melakukan berbagai operasi pada data yang disimpan dalam array. Pengguna dapat memasukkan sejumlah elemen ke dalam array, kemudian menggunakan menu interaktif untuk menjalankan berbagai fungsi seperti menampilkan elemen tertentu (berdasarkan indeks), menghitung rata-rata, standar deviasi, frekuensi kemunculan suatu bilangan, atau menghapus elemen tertentu. Program berjalan secara interaktif hingga pengguna memilih untuk keluar.
### Algoritma Program:
1.	Inisialisasi Array:
   
	•	Pengguna diminta memasukkan jumlah elemen yang akan dimasukkan ke dalam array.

	•	Array diinisialisasi, dan pengguna memasukkan elemen-elemen array satu per satu.

2.	Tampilan Menu:
   
	•	Program menampilkan menu pilihan yang memungkinkan pengguna memilih operasi yang ingin dilakukan, seperti:

	-	Menampilkan seluruh elemen array.
	-	Menampilkan elemen array berdasarkan indeks ganjil atau genap.
	-	Menampilkan elemen array berdasarkan indeks kelipatan tertentu.
	-	Menghapus elemen array berdasarkan indeks tertentu.
	-	Menghitung rata-rata atau standar deviasi dari elemen array.
	-	Menghitung frekuensi kemunculan suatu bilangan dalam array.
	-	Keluar dari program.

3.	Pemrosesan Menu:
   
	•	Berdasarkan pilihan pengguna, program melakukan operasi berikut:

	-	a: Menampilkan seluruh elemen array.
	-	b: Menampilkan elemen array dengan indeks ganjil.
	-	c: Menampilkan elemen array dengan indeks genap.
	-	d: Menampilkan elemen array dengan indeks yang merupakan kelipatan bilangan tertentu.
	-	e: Menghapus elemen array pada indeks tertentu.
	-	f: Menghitung rata-rata elemen array.
	-	g: Menghitung standar deviasi elemen array.
	-	h: Menghitung jumlah kemunculan bilangan tertentu dalam array.
	-	i: Keluar dari program.

4.	Perulangan:

	•	Setelah setiap operasi selesai, program kembali ke menu utama hingga pengguna memilih untuk keluar dengan opsi "i".


### Cara Kerja Program:
1.	Inisialisasi:
   
	•	Program dimulai dengan meminta pengguna untuk menentukan jumlah elemen array.

	•	Pengguna memasukkan nilai elemen array satu per satu.

2.	Operasi pada Array:
   
	•	Menampilkan elemen array:

	-	a: Menampilkan seluruh elemen array.
	-	b: Menampilkan elemen pada indeks ganjil.
	-	c: Menampilkan elemen pada indeks genap.

	•	Operasi berdasarkan indeks:

	-	d: Menampilkan elemen array pada indeks yang merupakan kelipatan bilangan tertentu.
 	
	•	Manipulasi array:

	-	e: Menghapus elemen array pada indeks tertentu, kemudian mencetak array yang telah diperbarui.
 	
	•	Statistik array:

	-	f: Menghitung rata-rata elemen array.
	-	g: Menghitung standar deviasi elemen array.

	•	Frekuensi bilangan:

	-	h: Menghitung jumlah kemunculan bilangan tertentu dalam array.
 	
3.	Keluar Program:
  	
	•	Program terus berjalan hingga pengguna memilih opsi "i" untuk keluar.


### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.

### Source code
```go
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import "fmt"

func main() {
	// Deklarasi variabel
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string
	var pertandingan int = 1

	// Input nama klub
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	fmt.Println("\nMasukkan skor untuk setiap pertandingan. Masukkan skor negatif untuk menghentikan.")

	// Loop untuk menerima skor dan menentukan pemenang
	for {
		fmt.Printf("Pertandingan %d - Skor %s: ", pertandingan, klubA)
		fmt.Scan(&skorA)
		fmt.Printf("Pertandingan %d - Skor %s: ", pertandingan, klubB)
		fmt.Scan(&skorB)

		// Kondisi untuk menghentikan program jika skor negatif
		if skorA < 0 || skorB < 0 {
			fmt.Println("Pertandingan selesai.")
			break
		}

		// Menentukan hasil pertandingan
		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}

		pertandingan++
	}

	// Menampilkan hasil pertandingan
	fmt.Println("\nHasil pertandingan:")
	for i, hasil := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, hasil)
	}
}
```
#### Output Program
![image](https://github.com/user-attachments/assets/2fc62f6f-dd68-4953-b22f-39f581f8f89e)

### Deskripsi Program
Program di atas digunakan untuk untuk mencatat skor pertandingan antara dua klub dan menentukan pemenang dari setiap pertandingan. Program akan terus menerima skor pertandingan secara berurutan hingga pengguna memasukkan skor negatif, yang menjadi tanda untuk mengakhiri pencatatan. Setelah itu, program menampilkan hasil dari semua pertandingan yang telah dimainkan.

### Algoritma Program:
1.	Deklarasi Variabel:
   
	•	Variabel string `klubA` dan `klubB` digunakan untuk menyimpan nama masing-masing klub.

 	•	Variabel integer `skorA` dan `skorB` untuk mencatat skor kedua klub dalam setiap pertandingan.

	•	Slice string `pemenang` digunakan untuk menyimpan pemenang dari setiap pertandingan.

	•	Variabel `pertandingan` untuk mencatat nomor pertandingan.

3.	Input Nama Klub:
   
	•	Pengguna diminta untuk memasukkan nama klub pertama `(klubA)` dan klub kedua `(klubB)`.

4.	Proses Input Skor:
   
	•	Program menjalankan loop untuk menerima skor pertandingan.

	•	Setiap pertandingan meminta skor untuk `klubA` dan `klubB`.

	•	Jika salah satu skor negatif, loop dihentikan, menandakan akhir pencatatan.

5.	Penentuan Pemenang:
   
	•	Jika skor `klubA` lebih tinggi dari `klubB`, nama `klubA` ditambahkan ke daftar pemenang.

	•	Jika sebaliknya, nama `klubB` ditambahkan.

	•	Jika skornya sama, program menambahkan "Draw" ke daftar.

6.	Output Hasil:
	•	Setelah loop selesai, program menampilkan hasil pertandingan berdasarkan daftar pemenang.


### Cara Kerja Program:
1.	Memulai Program:
   
	•	Pengguna diminta memasukkan nama dua klub, misalnya `Klub A` dan `Klub B`.

2.	Pencatatan Skor:
   
	•	Untuk setiap pertandingan, pengguna memasukkan skor untuk kedua klub.

3.	Penentuan dan Pencatatan Hasil:
   
	•	Program membandingkan skor yang dimasukkan:

	-	Jika `skorA` lebih tinggi, `klubA` tercatat sebagai pemenang.

	-	Jika `skorB` lebih tinggi, `klubB` tercatat sebagai pemenang.

	-	Jika skor seimbang, hasil "Draw" dicatat.

4.	Mengakhiri Input:
  	
	•	Pengguna dapat menghentikan proses dengan memasukkan skor negatif.

5.	Menampilkan Hasil:
   
	•	Program mencetak hasil setiap pertandingan 


### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom
### Source code
```go
// Lutfiana Isnaeni Lthifah
// 2311102165

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Konstanta dan tipe data
const NMAX int = 127

type tabel [NMAX]rune

// Fungsi untuk mengisi array dari input pengguna
func isiArray(t *tabel, n *int) {
	var input string
	fmt.Println("Masukkan teks, akhiri dengan karakter titik ('.'):")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	for _, char := range input {
		if char == '.' || *n >= NMAX {
			break
		}
		t[*n] = char
		*n++
	}
}

// Fungsi untuk mencetak isi array
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// Fungsi untuk membalikkan isi array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

// Fungsi untuk memeriksa apakah array membentuk palindrom
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	// Isi array dengan memanggil prosedur isiArray
	isiArray(&tab, &n)

	// Cetak isi array
	fmt.Print("Teks: ")
	cetakArray(tab, n)

	// Balikkan isi array
	balikanArray(&tab, n)

	// Cetak isi array setelah dibalik
	fmt.Print("Reverse teks: ")
	cetakArray(tab, n)

	// Periksa apakah array membentuk palindrom
	if palindrom(tab, n) {
		fmt.Println("Palindrom? true")
	} else {
		fmt.Println("Palindrom? false")
	}
}
```
#### Output Program
![image](https://github.com/user-attachments/assets/5247310d-db93-4080-ac46-aed967099410)

### Deskripsi Program
Program di atas digunakan untuk memproses teks yang dimasukkan pengguna hingga mencapai tanda titik (.). Program menyimpan teks dalam sebuah array dan memiliki tiga fitur utama: menampilkan teks asli, membalik urutan teks, serta memeriksa apakah teks tersebut merupakan palindrom (teks yang sama saat dibaca dari depan maupun belakang).

### Algoritma Program:
1. Inisialisasi Variabel
   
	Program mendeklarasikan variabel untuk menyimpan teks, jumlah karakter teks, dan array untuk menampung karakter yang diinput.

2. Pengisian Teks
   
	Pengguna diminta memasukkan teks hingga tanda titik (.). Setiap karakter yang dimasukkan akan disimpan ke dalam array hingga karakter titik ditemukan 		atau jumlah karakter mencapai batas maksimal.

3. Menampilkan Teks Asli
   
	Program mencetak teks yang dimasukkan pengguna tanpa perubahan.

4. Membalik Urutan Teks
   
	Program membalik urutan teks dengan menukar elemen awal dengan elemen akhir secara berpasangan.

5. Pemeriksaan Palindrom
   
	Program membandingkan elemen teks dari depan ke belakang. Jika semua pasangan elemen cocok, teks tersebut adalah palindrom. Jika tidak, teks tersebut 		bukan palindrom.

6. Output Akhir
   
	Program menampilkan teks asli, teks yang sudah dibalik, serta informasi apakah teks tersebut merupakan palindrom.

### Cara Kerja Program:
1. Memulai Program
- Pengguna diminta untuk memasukkan teks melalui terminal.
- Input dihentikan ketika pengguna mengetikkan tanda titik (`.`) atau jumlah karakter mencapai batas maksimal yang telah ditentukan.
2. Menyimpan Teks
- Teks yang dimasukkan pengguna disimpan ke dalam sebuah array, karakter demi karakter, hingga mencapai tanda titik.
3. Menampilkan Teks Asli
- Program mencetak teks yang telah disimpan ke dalam array tanpa perubahan, sesuai dengan urutan aslinya.
4. Membalik Urutan Teks
- Program membalik urutan teks yang telah disimpan di array, di mana karakter pertama ditukar dengan karakter terakhir, dan seterusnya hingga semua karakter dalam array berbalik posisi.
5. Menampilkan Teks yang Dibalik
- Program mencetak teks setelah proses pembalikan urutan, menunjukkan hasil dari teks yang dibalik.
6. Memeriksa Palindrom
- Program membandingkan karakter dari awal hingga tengah teks dengan karakter dari akhir ke tengah. Jika setiap pasangan karakter sama, maka teks tersebut adalah palindrom.
- Jika ditemukan pasangan karakter yang tidak sama, maka teks tersebut bukan palindrom.
7. Menampilkan Hasil Palindrom
- Program memberi tahu pengguna apakah teks yang dimasukkan adalah palindrom atau bukan.
8. Mengakhiri Program
- Setelah menampilkan hasil pemeriksaan, program selesai dan kembali ke terminal.


