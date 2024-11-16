<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br> 

<h2 align="center"><strong>MODUL VII</strong></h2>
<h2 align="center"><strong> STRUCK & ARRAY </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Ria Wulandari / 2311102173<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh,S.Kom.,M.Kom.
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
#### STRUCK
Struct digunakan untuk mengelompokkan berbagai elemen data dengan tipe yang berbeda ke dalam satu entitas.
##### Deklarasi Struct
```go
type Mahasiswa struct {
    Nama    string
    Umur    int
    Jurusan string
}
```
##### Inisialisasi Struct
Struct dapat diinisialisasi dengan dua cara:

1. Langsung: Menentukan nilai untuk setiap field.
```go
mhs := Mahasiswa{"Budi", 21, "Informatika"}
```

2. Dengan Penamaan Field: Mempermudah pembacaan.
```go
mhs := Mahasiswa{
    Nama:    "Budi",
    Umur:    21,
    Jurusan: "Informatika",
}
```
##### Mengakses dan Memodifikasi Field
```go
fmt.Println(mhs.Nama)  // Mengakses nilai field Nama
mhs.Nama = "Andi"      // Mengubah nilai field Nama
```
##### Pointer ke Struct
```go
mhsPtr := &mhs
mhsPtr.Umur = 22        // Memperbarui nilai melalui pointer
```
##### Struct di dalam Struct
Struct dapat memiliki struct lain sebagai field.
```go
type Alamat struct {
    Kota   string
    Jalan  string
}

type Mahasiswa struct {
    Nama   string
    Umur   int
    Alamat Alamat
}

mhs := Mahasiswa{
    Nama: "Budi",
    Umur: 21,
    Alamat: Alamat{
        Kota:  "Jakarta",
        Jalan: "Jl. Merdeka",
    },
}
fmt.Println(mhs.Alamat.Kota) // Mengakses field dari struct di dalam struct
```
#### ARRAY
Array digunakan untuk menyimpan kumpulan elemen dengan tipe data yang sama. Ukuran array ditentukan saat deklarasi dan tidak dapat diubah.
##### Deklarasi dan Inisialisasi Array
```go
var angka [5]int               // Array dengan 5 elemen bertipe int
angka[0] = 10                  // Menetapkan nilai pada elemen pertama
angka := [5]int{1, 2, 3, 4, 5} // Inisialisasi langsung

// Array dengan jumlah elemen ditentukan otomatis
angka := [...]int{1, 2, 3}
```
##### Mengakses Elemen
Elemen dalam array diakses menggunakan indeks:
```go
fmt.Println(angka[0]) // Mengakses elemen pertama
angka[1] = 20         // Mengubah nilai elemen kedua
```
##### Iterasi pada Array
```go
for i, v := range angka {
    fmt.Printf("Indeks %d: %d\n", i, v)
}
```
## Guided
### 1. Program sederhana untuk menghitung lama waktu parkir berdasarkan waktu kedatangan dan waktu pulang
#### Source code :
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

	dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600 // konversi ke detik
	dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600 // detik
	lParkir = dPulang - dParkir                                   // detik dari pulang-datang

	durasi.jam = lParkir / 3600
	durasi.menit = lParkir % 3600 / 60
	durasi.detik = lParkir % 3600 % 60 // 17
	fmt.Printf("Lama parkir: %d jam %d menit %d detik", durasi.jam, durasi.menit, durasi.detik)
}
```
#### Output
![image](https://github.com/user-attachments/assets/011a33c7-dfe0-4aff-a7ea-b2e07ff0f154)

#### Deskripsi program
Program ini dirancang untuk menghitung durasi waktu parkir dalam format jam, menit, dan detik berdasarkan waktu datang dan waktu pulang kendaraan. Input berupa waktu datang dan pulang dalam format jam, menit, dan detik.
#### Algoritma program
1. Program meminta pengguna untuk memasukkan waktu parkir (datang) dan waktu pulang.

2. Konversi kedua waktu tersebut ke dalam satuan detik untuk mempermudah perhitungan.

3. Hitung selisih waktu dengan mengurangi total detik waktu pulang dari total detik waktu datang.

4. Konversi durasi selisih waktu tersebut kembali ke format jam, menit, dan detik.

5. Tampilkan durasi waktu parkir ke layar.
#### Cara kerja program
1. Waktu datang dan waktu pulang dikonversi ke detik dengan formula:
   Total Detik=(Jam×3600)+(Menit×60)+Detik
2. Selisih detik dihitung dengan formula:
   Selisih Detik=Detik Pulang−Detik Datang
3. Selisih detik diubah ke jam, menit, dan detik dengan rumus:
- Jam = Selisih Detik÷3600
- Menit = (Selisih Detikmod3600)÷60
- Detik = (Selisih Detikmod3600)mod60
### 2. Program sederhana untuk validasi duplikasi nama pada daftar teman
#### Source code :
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
#### Output
![image](https://github.com/user-attachments/assets/a33471a5-e70c-469a-a9d6-31fd590dc12a)

#### Deskripsi program
Program ini mengelola daftar nama teman. Nama baru hanya akan ditambahkan jika belum ada dalam daftar sebelumnya. Jika nama sudah ada, program akan memberikan peringatan bahwa nama tersebut telah terdaftar.
#### Algoritma program
1. Inisialisasi slice dengan daftar nama teman awal.

2. Inisialisasi slice dengan nama-nama baru yang akan ditambahkan.

3. Iterasi setiap nama dalam daftar nama baru:
- Periksa apakah nama tersebut sudah ada di dalam daftar teman menggunakan fungsi `sudahAda`.
- Jika belum ada, tambahkan nama tersebut ke daftar teman.
- Jika sudah ada, tampilkan pesan bahwa nama tersebut sudah terdaftar.

4. Tampilkan daftar teman akhir setelah semua nama baru diproses.
#### Cara kerja program
1. Slice `daftarTeman` diinisialisasi dengan nilai awal: `{"Andi", "Budi", "Cici"}`.

2. Slice `namaBaru` diinisialisasi dengan nama-nama yang akan ditambahkan: `{"Dewi", "Budi", "Eka"}`.

3. Program memproses nama-nama dalam `namaBaru` satu per satu:
- Nama "Dewi" tidak ada di `daftarTeman`, sehingga ditambahkan.
- Nama "Budi" sudah ada di `daftarTeman`, sehingga muncul peringatan.
- Nama "Eka" tidak ada di `daftarTeman`, sehingga ditambahkan.

4. Setelah semua nama baru diproses, daftar teman akhir ditampilkan: `{"Andi", "Budi", "Cici", "Dewi", "Eka"}`.
### 3. Program sederhana untuk menampilkan daftar harga buah menggunakan map.
#### Source code :
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
#### Output
![image](https://github.com/user-attachments/assets/b378870b-c5ef-48aa-aae9-b1c2adf336cd)

#### Deskripsi program
Program ini menggunakan struktur data map untuk menyimpan informasi harga beberapa jenis buah. Program menampilkan daftar buah beserta harganya dan menampilkan harga untuk buah tertentu secara spesifik.
#### Algoritma program
1. Buat sebuah map bernama `hargaBuah` dengan kunci berupa nama buah (string) dan nilai berupa harga (integer).

2. Isi map tersebut dengan pasangan nama buah dan harga:
- "Apel" dengan harga 5000
- "Pisang" dengan harga 3000
- "Mangga" dengan harga 7000

3. Iterasi seluruh pasangan kunci-nilai dalam map menggunakan perulangan `for-range` untuk menampilkan semua buah dan harganya.

4. Ambil harga spesifik untuk buah "Mangga" dan tampilkan di layar.
#### Cara kerja program
1. Program membuat map bernama `hargaBuah` untuk menyimpan data harga buah.
- Map berisi pasangan seperti:
  - `"Apel": 5000`
  - `"Pisang": 3000`
  - `"Mangga": 7000`
2. Program mencetak seluruh daftar buah beserta harganya dengan menggunakan perulangan `for-range`.
- Dalam iterasi ini, setiap kunci (nama buah) dan nilai (harga buah) ditampilkan.

3. Program mengambil harga buah "Mangga" dari map dan mencetaknya secara spesifik menggunakan akses langsung melalui kunci `"Mangga"`.
## Unguided
### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥,𝑐𝑦) dengan radius 𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥,𝑦) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.

Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
#### Source code :
```go
package main

import (
	"fmt"
	"math"
)

// Struct untuk menyimpan informasi lingkaran
type Lingkaran struct {
	x, y, radius float64
}

// Fungsi untuk menghitung jarak titik dari pusat lingkaran
func hitungJarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// Fungsi untuk mengecek apakah titik berada di dalam lingkaran
func diDalamLingkaran(l Lingkaran, x, y float64) bool {
	return hitungJarak(l.x, l.y, x, y) <= l.radius
}

func main() {
	// Input koordinat pusat dan radius dua lingkaran
	var lingkaran1, lingkaran2 Lingkaran
	fmt.Println("Masukkan titik pusat dan radius lingkaran 1 (cx cy r):")
	fmt.Scan(&lingkaran1.x, &lingkaran1.y, &lingkaran1.radius)

	fmt.Println("Masukkan titik pusat dan radius lingkaran 2 (cx cy r):")
	fmt.Scan(&lingkaran2.x, &lingkaran2.y, &lingkaran2.radius)

	// Input titik yang akan dicek
	var x, y float64
	fmt.Println("Masukkan koordinat titik (x y):")
	fmt.Scan(&x, &y)

	// Cek posisi titik terhadap lingkaran
	diLingkaran1 := diDalamLingkaran(lingkaran1, x, y)
	diLingkaran2 := diDalamLingkaran(lingkaran2, x, y)

	// Output hasil
	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```
#### Output
![image](https://github.com/user-attachments/assets/fd4fe164-5c32-4a8d-9546-c23af1597f8e)

#### Deskripsi program
Program ini menghitung posisi suatu titik terhadap dua lingkaran berdasarkan koordinat pusat dan radius masing-masing lingkaran. Program memeriksa apakah titik tersebut berada di dalam salah satu lingkaran, kedua lingkaran, atau di luar keduanya.
#### Algoritma program
1. Input Data Lingkaran dan Titik

- Minta pengguna memasukkan koordinat pusat dan radius untuk dua lingkaran.
- Minta pengguna memasukkan koordinat titik yang akan diperiksa.

2. Perhitungan Posisi Titik

- Program menghitung jarak antara titik dan pusat lingkaran untuk menentukan apakah titik berada di dalam lingkaran.

3. Cek Kondisi

- Program memeriksa apakah jarak tersebut menunjukkan bahwa titik berada di dalam lingkaran 1, lingkaran 2, atau di keduanya.

4. Output Hasil

- Program menampilkan informasi mengenai posisi titik:
  - "Titik di dalam lingkaran 1 dan 2."
  - "Titik di dalam lingkaran 1."
  - "Titik di dalam lingkaran 2."
  - "Titik di luar lingkaran 1 dan 2."

#### Cara kerja program
1. Penggunaan Struct

- Struct `Lingkaran` digunakan untuk menyimpan informasi pusat dan radius dari lingkaran.

2. Fungsi Pendukung

- Fungsi `hitungJarak` digunakan untuk menghitung jarak antara dua titik.
- Fungsi `diDalamLingkaran` mengevaluasi apakah titik berada di dalam lingkaran tertentu.

3. Proses Utama

- Program meminta data lingkaran pertama dan kedua dari pengguna.
- Program menerima koordinat titik dari pengguna.
- Program memeriksa posisi titik terhadap kedua lingkaran menggunakan fungsi yang sudah dibuat.

4. Hasil Akhir

- Program mencetak hasil posisi titik berdasarkan lingkaran-lingkaran tersebut.
### 2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut:
a. Menampilkan keseluruhan isi dari array.

b. Menampilkan elemen-elemen array dengan indeks ganjil saja.

c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap).

d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna.

e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid.Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil

f. Menampilkan rata-rata dari bilangan yang ada di dalam array.

g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.

h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.
#### Source code :
```go
package main

import (
	"fmt"
	"math"
)

func tampilkanSeluruhArray(arr []int) {
	fmt.Println("Seluruh elemen array:", arr)
}

func tampilkanIndeksGanjil(arr []int) {
	fmt.Println("Elemen dengan indeks ganjil:")
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}
}

func tampilkanIndeksGenap(arr []int) {
	fmt.Println("Elemen dengan indeks genap:")
	for i := 0; i < len(arr); i += 2 {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}
}

func tampilkanIndeksKelipatan(arr []int, x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	for i := x; i < len(arr); i += x {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}
}

func hapusElemenPadaIndeks(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		fmt.Println("Indeks tidak valid")
		return arr
	}
	fmt.Printf("Menghapus elemen pada indeks %d\n", index)
	arr = append(arr[:index], arr[index+1:]...)
	return arr
}

func hitungRataRata(arr []int) float64 {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return float64(sum) / float64(len(arr))
}

func hitungStandarDeviasi(arr []int) float64 {
	rataRata := hitungRataRata(arr)
	sum := 0.0
	for _, val := range arr {
		sum += math.Pow(float64(val)-rataRata, 2)
	}
	return math.Sqrt(sum / float64(len(arr)))
}

func hitungFrekuensi(arr []int, bilangan int) int {
	count := 0
	for _, val := range arr {
		if val == bilangan {
			count++
		}
	}
	return count
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	tampilkanSeluruhArray(arr)
	tampilkanIndeksGanjil(arr)
	tampilkanIndeksGenap(arr)

	var x int
	fmt.Print("Masukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	tampilkanIndeksKelipatan(arr, x)

	var index int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&index)
	arr = hapusElemenPadaIndeks(arr, index)
	tampilkanSeluruhArray(arr)

	rataRata := hitungRataRata(arr)
	fmt.Printf("Rata-rata elemen array: %.2f\n", rataRata)

	standarDeviasi := hitungStandarDeviasi(arr)
	fmt.Printf("Standar deviasi elemen array: %.2f\n", standarDeviasi)

	var bilangan int
	fmt.Print("Masukkan bilangan untuk menghitung frekuensinya: ")
	fmt.Scan(&bilangan)
	frekuensi := hitungFrekuensi(arr, bilangan)
	fmt.Printf("Frekuensi %d di dalam array: %d kali\n", bilangan, frekuensi)
}
```
#### Output
![image](https://github.com/user-attachments/assets/f0cb214d-dac6-42af-b252-e8843a61eebc)

#### Deskripsi program
Program ini memungkinkan pengguna untuk mengelola dan menganalisis elemen-elemen dalam sebuah array. Fungsionalitas program meliputi menampilkan elemen array berdasarkan kriteria tertentu, menghitung rata-rata dan standar deviasi elemen, serta menghitung frekuensi kemunculan sebuah bilangan dalam array.
#### Algoritma program
1. Inisialisasi Array

- Minta pengguna memasukkan jumlah elemen array.
- Gunakan loop untuk menerima input elemen-elemen array dari pengguna.

2. Operasi pada Array

- Menampilkan Elemen Array
  - Tampilkan seluruh elemen array.
  - Tampilkan elemen dengan indeks ganjil.
  - Tampilkan elemen dengan indeks genap.
- Tampilkan Elemen dengan Indeks Kelipatan
  - Minta pengguna memasukkan bilangan kelipatan x.
  - Tampilkan elemen pada indeks yang merupakan kelipatan x.

3. Hapus Elemen pada Indeks

- Minta pengguna memasukkan indeks untuk menghapus elemen.
- Hapus elemen tersebut dari array dan tampilkan array terbaru.

4. Perhitungan Statistik

- Rata-rata: Hitung rata-rata elemen dalam array.
- Standar Deviasi: Hitung standar deviasi elemen dalam array.

5. Frekuensi Bilangan

- Minta pengguna memasukkan sebuah bilangan.
- Hitung berapa kali bilangan tersebut muncul dalam array.

6. Output Hasil

- Tampilkan hasil dari operasi yang dilakukan pada array.
#### Cara kerja program
1. Input Data Array

- Program meminta pengguna memasukkan ukuran array dan elemen-elemennya. Data disimpan dalam slice `arr`.

2. Pengolahan Array

- Fungsi `tampilkanSeluruhArray` digunakan untuk menampilkan semua elemen.
- Fungsi `tampilkanIndeksGanjil` dan `tampilkanIndeksGenap` digunakan untuk menampilkan elemen dengan indeks ganjil/genap.
- Fungsi `tampilkanIndeksKelipatan` menampilkan elemen yang berada pada indeks kelipatan bilangan tertentu.

3. Penghapusan Elemen

- Fungsi `hapusElemenPadaIndeks` menghapus elemen pada indeks tertentu setelah memastikan indeks valid.

4. Perhitungan Statistik

- Rata-rata dihitung menggunakan fungsi `hitungRataRata`.
- Standar Deviasi dihitung menggunakan fungsi `hitungStandarDeviasi`.

5. Frekuensi Bilangan

- Fungsi `hitungFrekuensi` menghitung jumlah kemunculan bilangan tertentu dalam array.

6. Output

- Program menampilkan array hasil pengolahan, statistik rata-rata dan standar deviasi, serta jumlah kemunculan bilangan tertentu.
### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.
#### Source code :
```go
package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var menangA, menangB, seri int

	// Meminta input nama klub
	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	// Meminta input skor dari pengguna hingga 9 kali atau skor negatif diberikan
	for i := 1; i <= 9; i++ {
		fmt.Printf("Pertandingan %d (masukkan skorA skorB): ", i)
		_, err := fmt.Scan(&skorA, &skorB)

		// Menangani error input
		if err != nil {
			fmt.Println("Input tidak valid, coba lagi.")
			i-- // Mengulangi pertandingan yang sama jika input tidak valid
			continue
		}

		// Cek jika ada skor negatif untuk berhenti
		if skorA < 0 || skorB < 0 {
			fmt.Println("Input negatif terdeteksi, pertandingan selesai.")
			break
		}

		// Menentukan pemenang dari pertandingan
		if skorA > skorB {
			fmt.Printf("Hasil %d: %s menang\n", i, klubA)
			menangA++
		} else if skorB > skorA {
			fmt.Printf("Hasil %d: %s menang\n", i, klubB)
			menangB++
		} else {
			fmt.Printf("Hasil %d: Draw\n", i)
			seri++
		}
	}

	// Menampilkan hasil akhir
	fmt.Println("\nHasil Akhir:")
	fmt.Printf("%s menang: %d kali\n", klubA, menangA)
	fmt.Printf("%s menang: %d kali\n", klubB, menangB)
	fmt.Printf("Seri: %d kali\n", seri)
}

```
#### Output
![image](https://github.com/user-attachments/assets/77fce4e3-429a-4d9d-b6c9-7aac5ed226ee)

#### Deskripsi program
Program ini berfungsi untuk merekam hasil pertandingan antara dua klub sepak bola. Pengguna diminta untuk memasukkan nama kedua klub dan skor mereka untuk setiap pertandingan. Program kemudian akan menghitung berapa kali masing-masing klub menang, berapa kali pertandingan berakhir seri, dan menampilkan hasil akhir setelah 9 pertandingan atau jika ada skor negatif yang dimasukkan, yang menandakan akhir pertandingan.
#### Algoritma program
1. Input Nama Klub:

- Program meminta input nama klub A dan klub B.

2. Input Skor Pertandingan:

- Program meminta input skor untuk 9 pertandingan atau sampai skor negatif diberikan.
- Jika input tidak valid (misalnya, skor bukan angka), program akan meminta pengguna untuk mencoba lagi dan tidak menghitung pertandingan tersebut.
- Jika skor negatif dimasukkan, program akan menghentikan input lebih lanjut dan menampilkan bahwa pertandingan selesai.

3. Penentuan Pemenang:

- Program membandingkan skor dari kedua klub:
  - Jika skor klub A lebih besar dari klub B, maka klub A menang.
  - Jika skor klub B lebih besar dari klub A, maka klub B menang.
  - Jika skor keduanya sama, pertandingan berakhir seri.

4. Hasil Akhir:

- Setelah semua pertandingan atau jika skor negatif dimasukkan, program menampilkan hasil akhir:
  - Jumlah kemenangan untuk klub A dan klub B.
  - Jumlah pertandingan yang berakhir seri.

#### Cara kerja program
1. Input Nama Klub:

- Program meminta pengguna untuk memasukkan nama klub A dan klub B melalui `fmt.Scanln`.

2. Input Skor Pertandingan:

- Program menggunakan loop `for i := 1; i <= 9; i++` untuk meminta skor untuk 9 pertandingan.
- Di dalam loop, program membaca skor pertandingan menggunakan `fmt.Scan`. Jika input tidak valid (misalnya bukan angka), program akan menampilkan pesan error dan meminta pengguna untuk mengulang input skor yang sama.
- Program memeriksa apakah skor yang dimasukkan adalah angka negatif, jika ya, program berhenti meminta input lebih lanjut dan keluar dari loop.

3. Penentuan Pemenang:

- Setelah mendapatkan skor untuk setiap pertandingan, program menentukan pemenang berdasarkan perbandingan antara `skorA` dan `skorB`.
- Hasil pertandingan (pemenang atau seri) dicatat dengan menambah variabel penghitung `menangA`, `menangB`, atau `seri` sesuai kondisi pertandingan.

4. Menampilkan Hasil Akhir:

- Setelah 9 pertandingan atau setelah skor negatif dimasukkan, program menampilkan hasil akhir dengan menggunakan `fmt.Printf` untuk menampilkan jumlah kemenangan dan jumlah seri.
### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom
#### Source code :
```go
package main

import (
	"fmt"
	"strings"
)

// Fungsi untuk membalikkan urutan elemen dalam array karakter
func reverseCharacters(input []rune) []rune {
	length := len(input)
	reversedArray := make([]rune, length)
	for i := 0; i < length; i++ {
		reversedArray[i] = input[length-i-1]
	}
	return reversedArray
}

// Fungsi untuk memeriksa apakah array karakter membentuk palindrom
func isPalindrome(input []rune) bool {
	reversedArray := reverseCharacters(input)
	for i := 0; i < len(input); i++ {
		if input[i] != reversedArray[i] {
			return false
		}
	}
	return true
}

func main() {
	var userInput string
	fmt.Print("Masukkan sekumpulan karakter: ")
	fmt.Scanln(&userInput)

	// Mengonversi input ke dalam array karakter (rune) dan mengabaikan perbedaan huruf kapital
	convertedArray := []rune(strings.ToLower(userInput))

	// Membalikkan array karakter
	reversedConvertedArray := reverseCharacters(convertedArray)

	// Menampilkan array asli dan array yang dibalik
	fmt.Printf("Array asli: %s\n", string(convertedArray))
	fmt.Printf("Array yang dibalik: %s\n", string(reversedConvertedArray))

	// Memeriksa apakah array tersebut membentuk palindrom
	if isPalindrome(convertedArray) {
		fmt.Println("Array membentuk palindrom.")
	} else {
		fmt.Println("Array tidak membentuk palindrom.")
	}
}

```
#### Output
![image](https://github.com/user-attachments/assets/e9cf1d11-9a73-417b-83e7-d7eb5db6e53e)

#### Deskripsi program
Program ini berfungsi untuk memeriksa apakah input sekumpulan karakter yang dimasukkan oleh pengguna membentuk palindrom. Palindrom adalah urutan karakter yang terbaca sama dari depan maupun dari belakang, mengabaikan perbedaan huruf kapital.
#### Algoritma program
1. Fungsi `reverseCharacters(input []rune) []rune`:

- Mengambil input berupa array karakter (`rune`).
- Membuat array baru (`reversedArray`) dengan panjang yang sama seperti input.
- Membalikkan urutan elemen dalam array, sehingga elemen pertama dari input menjadi elemen terakhir dalam `reversedArray` dan seterusnya.
- Mengembalikan `reversedArray`.

2. Fungsi `isPalindrome(input []rune) bool`:

- Memanggil fungsi `reverseCharacters` untuk membalikkan input.
- Membandingkan setiap karakter di dalam input dengan karakter yang sesuai dalam `reversedArray`.
- Jika semua karakter dalam `input` sama dengan karakter dalam `reversedArray` pada posisi yang sama, maka `isPalindrome` akan mengembalikan `true` (array membentuk palindrom), jika tidak maka `false`.

3. Fungsi `main()`:

- Meminta input sekumpulan karakter dari pengguna.
- Mengonversi input menjadi array karakter `rune` untuk mengabaikan perbedaan huruf kapital.
- Memanggil `reverseCharacters` untuk membalikkan array karakter.
- Menampilkan array asli dan array yang sudah dibalik.
- Memanggil `isPalindrome` untuk memeriksa apakah array membentuk palindrom.
- Menampilkan hasil dari pemeriksaan palindrom.
#### Cara kerja program
1. Input: Program ini memulai dengan meminta pengguna untuk memasukkan sekumpulan karakter. Input ini kemudian diubah menjadi array karakter `rune` menggunakan `strings.ToLower()` untuk mengabaikan perbedaan huruf kapital.

2. Pemrosesan:
- Fungsi `reverseCharacters` digunakan untuk membalikkan urutan elemen dalam array karakter.
- Fungsi `isPalindrome` memeriksa apakah karakter di array input sama dengan karakter yang sesuai dalam array yang sudah dibalik.

3. Output:
- Menampilkan array asli dan array yang sudah dibalik.
- Menunjukkan hasil pemeriksaan apakah array tersebut membentuk palindrom atau tidak.
