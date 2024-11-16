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
  	
      o	Melakukan iterasi melalui elemen-elemen dalam `daftarTeman`.
  	
      o	Jika nama ditemukan, mengembalikan `true`; jika tidak, mengembalikan `false`.
  	
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
  	
      o	Kunci: nama buah (contoh:"`Apel"`, `"Pisang"`).
  	
      o	Nilai: harga buah dalam bentuk angka.
  	
2.	Menampilkan Harga Semua Buah
   
    •	Gunakan perulangan `for` dengan range untuk menelusuri setiap elemen dalam map hargaBuah:
  	
      o	Cetak nama buah dan harganya dengan format `buah: Rp harga`.
  	
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
#### Output Program
### Deskripsi Program
### Algoritma Program:
### Cara Kerja Program:


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
#### Output Program
### Deskripsi Program
### Algoritma Program:
### Cara Kerja Program:

### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.

### Source code
#### Output Program
### Deskripsi Program
### Algoritma Program:
### Cara Kerja Program:

### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom
### Source code
#### Output Program
### Deskripsi Program
### Algoritma Program:
### Cara Kerja Program:



