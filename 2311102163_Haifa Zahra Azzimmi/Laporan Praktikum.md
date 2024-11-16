<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL 7</strong></h2>
<h2 align="center"><strong> STRUCK & ARRAY </strong></h2> 

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Haifa Zahra Azzimmi / 2311102163<br>
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

### Array
Array adalah struktur data yang menyimpan elemen-elemen dengan tipe yang sama dalam satu variabel. Elemen diakses menggunakan indeks mulai dari 0. Array bisa statis atau dinamis tergantung bahasa pemrograman.

**Karakteristik Array**
- Tipe Data Seragam: Semua elemen dalam array memiliki tipe data yang sama.
- Akses melalui Indeks: Elemen dalam array diakses menggunakan indeks, mulai dari 0.
- Ukuran Tetap: Ukuran array biasanya tetap setelah dideklarasikan (di beberapa bahasa pemrograman).
- Lokasi Memori Kontigu: Elemen-elemen disimpan dalam blok memori yang berurutan.
- Efisien dalam Akses: Akses elemen cepat karena menggunakan indeks langsung.
- Iterasi Mudah: Dapat dengan mudah diiterasi menggunakan loop.

**Deklarasi Array:**

- Deklarasi dengan ukuran tetap
  ```go
  var arr [5]int
  ```
  Ini mendeklarasikan array arr dengan 5 elemen bertipe int.
  
- Deklarasi dan Inisialisasi
  ```go
  arr := [5]int{1, 2, 3, 4, 5}
  ```
  Ini mendeklarasikan dan menginisialisasi array arr dengan nilai-nilai 1, 2, 3, 4, dan 5.

- Deklarasi tanpa Menentukan Ukuran
    ```go
    arr := [...]int{1, 2, 3, 4, 5}
    ```
    Ini memungkinkan compiler untuk menentukan ukuran array berdasarkan jumlah elemen yang diberikan.

- Contoh Lengkap
    ```go
    package main

  import (
    "fmt"
  )

  func main() {
    // Deklarasi array dengan ukuran tetap
    var arr1 [5]int
    fmt.Println("Array 1:", arr1)

    // Deklarasi dan inisialisasi array
    arr2 := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Array 2:", arr2)

    // Deklarasi array tanpa menentukan ukuran
    arr3 := [...]int{1, 2, 3, 4, 5}
    fmt.Println("Array 3:", arr3)
  }
  ```

    Program ini menunjukkan berbagai cara untuk mendeklarasikan dan menginisialisasi array di Go.

### Struck
Struct adalah tipe data komposit dalam pemrograman yang memungkinkan pengelompokan beberapa variabel dengan tipe data yang berbeda di bawah satu nama. Setiap variabel dalam struct disebut "field" atau "member".


**Karakteristik Struct:**
- Pengelompokan Data: Struct memungkinkan pengelompokan berbagai tipe data di bawah satu nama.

- Field Berbeda Tipe: Setiap field dalam struct dapat memiliki tipe data yang berbeda.

- Akses Menggunakan Dot Notation: Field dalam struct diakses menggunakan operator titik (.).

- Deklarasi Custom: Struct memungkinkan definisi tipe data yang lebih kompleks dan spesifik sesuai kebutuhan program.

- Inisialisasi Fleksibel: Struct dapat diinisialisasi dengan berbagai cara, termasuk inisialisasi kosong, inisialisasi langsung dengan nilai, atau menggunakan fungsi.

**Contoh Struct dalam Go**
```go
package main

import (
    "fmt"
)

// Definisi struct
type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    // Inisialisasi struct
    p := Person{Name: "John", Age: 30, City: "New York"}

    // Akses field menggunakan dot notation
    fmt.Println("Name:", p.Name)
    fmt.Println("Age:", p.Age)
    fmt.Println("City:", p.City)
}
```

## Guided 

### 1. Program sederhana untuk menghitung lama waktu parkir berdasarkan waktu kedatangan dan waktu pulang.

### Source Code :

```go
// Haifa Zahra Azzimmi
// 2311102163
// S1 IF 11 5
package main

import "fmt"

type waktu struct {
    jam, menit, detik int
}

func main() {
    var wParkir, wPulang, durasi waktu
    var dParkir, dPulang, lParkir int

    fmt.Print("Masukkan waktu datang parkir (jam menit detik): ")
    fmt.Scan(&wParkir.jam, &wParkir.menit, &wParkir.detik)
    fmt.Print("Masukkan waktu pulang parkir (jam menit detik): ")
    fmt.Scan(&wPulang.jam, &wPulang.menit, &wPulang.detik)

    dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600 // konversi ke detik
    dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600 // detik
    lParkir = dPulang - dParkir                                   // detik dari pulang-datang

    durasi.jam = lParkir / 3600
    durasi.menit = lParkir % 3600 / 60
    durasi.detik = lParkir % 3600 % 60
    fmt.Printf("Lama parkir: %d jam %d menit %d detik\n", durasi.jam, durasi.menit, durasi.detik)
}


```

### Output:
![image](https://github.com/user-attachments/assets/549f47eb-b0f2-46e9-94c7-aeb1eb20a575)

### Full code Screenshot :
![image](https://github.com/user-attachments/assets/142c364e-f960-4c88-92df-99950b0d8b88)

### Deskripsi Program : 
Program ini menghitung durasi parkir berdasarkan waktu datang dan pulang. Menggunakan struct waktu untuk menyimpan jam, menit, dan detik, program ini mengonversi waktu datang dan pulang ke detik, menghitung selisihnya, dan mengonversi kembali ke jam, menit, dan detik untuk menampilkan durasi parkir.

### Algoritma Program :
1. Mulai
   
2. Deklarasi Struct untuk menyimpan waktu dalam format jam, menit, dan detik.
   
3. Deklarasi Variabel untuk menyimpan waktu datang (wParkir), waktu pulang (wPulang), durasi (durasi), total detik waktu datang (dParkir), total detik waktu pulang (dPulang), dan lama parkir dalam detik (lParkir).
   
4. Input Waktu: Minta pengguna memasukkan waktu datang parkir (jam, menit, detik) dan waktu pulang parkir (jam, menit, detik).
   
5. Konversi Waktu ke Detik:
   - Hitung total detik dari waktu datang (dParkir) dan waktu pulang (dPulang).
     
6. Hitung Lama Parkir:
   - Hitung lama parkir dalam detik (lParkir) dengan mengurangi dParkir dari dPulang.
     
7. Konversi Detik ke Format Jam, Menit, Detik:
   - Hitung durasi dalam jam, menit, dan detik.
     
8. Output Hasil: Cetak hasil lama parkir dalam format jam, menit, dan detik.
   
9. Selesai
    
### Cara Kerja Program :
1. Deklarasi Struct dan Variabel: Program dimulai dengan mendeklarasikan struct waktu untuk menyimpan data jam, menit, dan detik. Variabel wParkir, wPulang, durasi, dParkir, dPulang, dan lParkir dideklarasikan.

2. Input Waktu: Program meminta pengguna untuk memasukkan waktu datang parkir (wParkir) dan waktu pulang parkir (wPulang).

3. Konversi Waktu ke Detik: Program mengonversi waktu datang (wParkir) dan waktu pulang (wPulang) ke total detik masing-masing (dParkir dan dPulang) menggunakan rumus:
   - dParkir = wParkir.detik + wParkir.menit * 60 + wParkir.jam * 3600
   - dPulang = wPulang.detik + wPulang.menit * 60 + wPulang.jam * 3600

4. Hitung Lama Parkir: Program menghitung lama parkir dalam detik (lParkir) dengan mengurangi dParkir dari dPulang.
   
5. Konversi Detik ke Format Jam, Menit, Detik: Program mengonversi durasi lama parkir dalam detik (lParkir) menjadi jam (durasi.jam), menit (durasi.menit), dan detik (durasi.detik):
   - durasi.jam = lParkir / 3600
   - durasi.menit = (lParkir % 3600) / 60
   - durasi.detik = (lParkir % 3600) % 60

6. Output Hasil: Program mencetak hasil lama parkir dalam format jam, menit, dan detik dengan menggunakan fmt.Printf.

7. Program Selesai: Setelah mencetak hasil, program selesai menjalankan tugasnya.
   
### 2. Program sederhana untuk validasi duplikasi nama pada daftar teman.

### Source Code :

```go
// Haifa Zahra Azzimmi
// 2311102163
// S1 IF 11 5

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

### Output:
![image](https://github.com/user-attachments/assets/dd970ce5-7dbb-4748-be84-8abcc94b655d)

### Full code Screenshot :
![image](https://github.com/user-attachments/assets/fac8daf4-511c-4fa1-880c-8e1e27eb2686)

### Deskripsi Program : 
Program ini menambahkan nama baru ke dalam daftar teman jika nama tersebut belum ada dalam daftar. Nama-nama baru diperiksa menggunakan fungsi sudahAda, dan jika belum ada, ditambahkan ke daftar teman. Jika sudah ada, program mencetak pesan bahwa nama tersebut sudah ada. Hasil akhir adalah daftar teman yang diperbarui tanpa duplikasi.

### Algoritma Program :
1. Mulai

2. Deklarasi Fungsi sudahAda untuk mengecek apakah nama sudah ada di dalam slice daftarTeman.
   - Loop melalui setiap elemen di daftarTeman.
   - Jika elemen sama dengan nama yang dicari, kembalikan true.
   - Jika tidak ada yang sama, kembalikan false.

3. Deklarasi Slice Awal daftarTeman dengan nama-nama awal.

4. Deklarasi Slice namaBaru dengan nama-nama baru yang ingin ditambahkan.
   
5. Loop melalui namaBaru:
   - Untuk setiap nama dalam namaBaru, panggil fungsi sudahAda:
     - Jika sudahAda mengembalikan false, tambahkan nama ke daftarTeman.
     - Jika sudahAda mengembalikan true, cetak pesan bahwa nama sudah ada dalam daftar.

6. Cetak Daftar Teman Akhir

7. Selesai

### Cara Kerja Program :
1. Program dimulai dengan mendeklarasikan fungsi sudahAda yang mengecek apakah suatu nama sudah ada dalam slice daftarTeman.

2. Program mendeklarasikan slice awal daftarTeman yang berisi nama-nama awal seperti "Andi", "Budi", dan "Cici".

3. Program juga mendeklarasikan slice namaBaru yang berisi nama-nama baru yang ingin ditambahkan.

4. Untuk setiap nama dalam namaBaru, program memeriksa apakah nama tersebut sudah ada di daftarTeman menggunakan fungsi sudahAda.

5. Jika nama belum ada di daftarTeman, program menambahkan nama tersebut ke daftarTeman.

6. Jika nama sudah ada di daftarTeman, program mencetak pesan bahwa nama tersebut sudah ada.

7. Akhirnya, program mencetak daftar teman akhir setelah menambahkan nama-nama baru yang belum ada dalam daftar.

### 3. Program sederhana untuk menampilkan daftar harga buah menggunakan map.

### Source Code :

```go
// Haifa Zahra Azzimmi
// 2311102163
// S1 IF 11 5

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

### Output:
![image](https://github.com/user-attachments/assets/aedcf4b5-a3b2-4a3d-9c4c-5984b42b0f21)

### Full code Screenshot :
![image](https://github.com/user-attachments/assets/d02a9501-3f6e-476f-a177-fc4df50fde52)

### Deskripsi Program : 
Program ini membuat peta (map) yang menyimpan harga buah dengan nama buah sebagai kunci. Program kemudian menampilkan harga setiap buah dan mencetak harga buah Mangga secara spesifik.

### Algoritma Program :
1. Mulai
   
2. Deklarasi Peta hargaBuah dengan nama buah sebagai kunci dan harga sebagai nilai.

3. Inisialisasi peta hargaBuah dengan beberapa pasangan kunci-nilai (nama buah dan harga).

4. Cetak judul "Harga Buah".

5. Iterasi melalui peta hargaBuah untuk setiap pasangan kunci-nilai:
   - Cetak nama buah dan harga.

6. Cetak harga buah Mangga secara spesifik.

7. Selesai
   
### Cara Kerja Program :
1. Program dimulai dengan mendeklarasikan sebuah peta (map) bernama hargaBuah, di mana nama buah digunakan sebagai kunci dan harga buah sebagai nilai.

2. Peta hargaBuah diinisialisasi dengan beberapa nama buah dan harganya, seperti "Apel", "Pisang", dan "Mangga".

3. Program mencetak judul "Harga Buah".

4. Menggunakan loop, program iterasi melalui setiap pasangan kunci-nilai dalam peta hargaBuah, mencetak nama buah dan harganya.

5. Program mencetak harga buah Mangga secara spesifik.

6. Program selesai menjalankan tugasnya dengan menampilkan daftar harga buah dan harga spesifik untuk buah Mangga.


## Unguided 

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥,𝑐𝑦) dengan radius 𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥,𝑦) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.

Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

### Source Code :
```go
//Haifa Zahra Azzimmi
//2311102163

package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat  Titik
	radius int
}

func jarak(t1, t2 Titik) float64 {
	return math.Sqrt(float64((t1.x-t2.x)*(t1.x-t2.x) + (t1.y-t2.y)*(t1.y-t2.y)))
}

func titikDiDalamLingkaran(t Titik, lingkaran Lingkaran) bool {
	return jarak(t, lingkaran.pusat) <= float64(lingkaran.radius)
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 1 (cx1 cy1 r1): ")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 2 (cx2 cy2 r2): ")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Print("Masukkan koordinat titik sembarang (x y): ")
	fmt.Scan(&x, &y)

	lingkaran1 := Lingkaran{pusat: Titik{x: cx1, y: cy1}, radius: r1}
	lingkaran2 := Lingkaran{pusat: Titik{x: cx2, y: cy2}, radius: r2}
	titik := Titik{x: x, y: y}

	diDalamL1 := titikDiDalamLingkaran(titik, lingkaran1)
	diDalamL2 := titikDiDalamLingkaran(titik, lingkaran2)

	if diDalamL1 && diDalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}


```
### Output:
![image](https://github.com/user-attachments/assets/c5e13943-40a3-4219-bb75-e9ab093245e2)

### Full code Screenshot:
![code 1](https://github.com/user-attachments/assets/44168cf2-ae08-41a8-8a30-131daa46ff08)

### Deskripsi Program : 
Program ini menentukan apakah sebuah titik berada di dalam satu atau dua lingkaran yang diberikan. Program meminta input pengguna untuk koordinat pusat dan radius dari dua lingkaran, serta koordinat titik sembarang. Program menggunakan fungsi jarak untuk menghitung jarak antara dua titik dan fungsi titikDiDalamLingkaran untuk memeriksa apakah sebuah titik berada di dalam lingkaran. Program ini kemudian mengevaluasi apakah titik berada di dalam salah satu atau kedua lingkaran, dan mencetak hasilnya. Hasil yang dicetak akan menunjukkan apakah titik tersebut berada di dalam lingkaran 1, lingkaran 2, kedua lingkaran, atau di luar kedua lingkaran

### 2.  Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut:
a. Menampilkan keseluruhan isi dari array.

b. Menampilkan elemen-elemen array dengan indeks ganjil saja.

c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap).

d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna.

e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid.Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil

f. Menampilkan rata-rata dari bilangan yang ada di dalam array.

g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.

h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.

### Source Code :
```go
//Haifa Zahra Azzimmi
//2311102163

package main

import (
	"fmt"
	"math"
)

func tampilkanArray(arr []int) {
	fmt.Println("Isi Array:")
	for _, elemen := range arr {
		fmt.Print(elemen, " ")
	}
	fmt.Println()
}

func tampilkanIndeksGanjil(arr []int) {
	fmt.Println("Elemen dengan indeks ganjil:")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func tampilkanIndeksGenap(arr []int) {
	fmt.Println("Elemen dengan indeks genap:")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func tampilkanIndeksKelipatanX(arr []int, x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

func hapusElemen(arr []int, indeks int) []int {
	arr = append(arr[:indeks], arr[indeks+1:]...)
	return arr
}

func hitungRataRata(arr []int) float64 {
	var sum int
	for _, elemen := range arr {
		sum += elemen
	}
	return float64(sum) / float64(len(arr))
}

func hitungStandarDeviasi(arr []int) float64 {
	mean := hitungRataRata(arr)
	var sumSquares float64
	for _, elemen := range arr {
		sumSquares += math.Pow(float64(elemen)-mean, 2)
	}
	variance := sumSquares / float64(len(arr))
	return math.Sqrt(variance)
}

func hitungFrekuensi(arr []int, bilangan int) int {
	var count int
	for _, elemen := range arr {
		if elemen == bilangan {
			count++
		}
	}
	return count
}

func main() {
	var N, x, indeksHapus, bilanganCari int

	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&N)

	arr := make([]int, N)

	fmt.Println("Masukkan nilai untuk setiap elemen array:")
	for i := 0; i < N; i++ {
		fmt.Printf("Elemen ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	tampilkanArray(arr)

	tampilkanIndeksGanjil(arr)

	tampilkanIndeksGenap(arr)

	fmt.Print("Masukkan nilai x untuk indeks kelipatan x: ")
	fmt.Scan(&x)
	tampilkanIndeksKelipatanX(arr, x)

	fmt.Print("Masukkan indeks elemen yang akan dihapus: ")
	fmt.Scan(&indeksHapus)
	arr = hapusElemen(arr, indeksHapus)
	fmt.Println("Array setelah elemen dihapus:")
	tampilkanArray(arr)

	rataRata := hitungRataRata(arr)
	fmt.Printf("Rata-rata array: %.2f\n", rataRata)

	standarDeviasi := hitungStandarDeviasi(arr)
	fmt.Printf("Standar deviasi array: %.2f\n", standarDeviasi)

	fmt.Print("Masukkan bilangan untuk mencari frekuensinya: ")
	fmt.Scan(&bilanganCari)
	frekuensi := hitungFrekuensi(arr, bilanganCari)
	fmt.Printf("Frekuensi bilangan %d: %d kali\n", bilanganCari, frekuensi)
}


```
### Output:
![image](https://github.com/user-attachments/assets/58601cc7-17f8-417e-8782-8bc923193719)

### Full code Screenshot:
![code 2](https://github.com/user-attachments/assets/df360bf2-300a-41bf-a5b8-e37a6ee564f5)

### Deskripsi Program : 
Program ini meminta pengguna untuk memasukkan jumlah elemen dalam array dan nilainya, lalu menampilkan array tersebut. Selanjutnya, program menampilkan elemen dengan indeks ganjil, genap, dan kelipatan x sesuai masukan pengguna. Program juga dapat menghapus elemen pada indeks tertentu, menghitung rata-rata dan standar deviasi dari array, serta menghitung frekuensi kemunculan angka tertentu dalam array. Tujuannya adalah memanipulasi dan menganalisis data dalam array secara interaktif dengan berbagai fungsi.
  
### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.

### Source Code :
```go
//Haifa Zahra Azzimmi
//2311102163

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var winningClubs []string // Slice untuk menyimpan nama klub yang menang
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Rekap Skor Pertandingan Bola")
	for {
		fmt.Print("Masukkan nama klub pertama: ")
		club1, _ := reader.ReadString('\n')
		club1 = strings.TrimSpace(club1)

		fmt.Print("Masukkan nama klub kedua: ")
		club2, _ := reader.ReadString('\n')
		club2 = strings.TrimSpace(club2)

		fmt.Print("Masukkan skor klub pertama: ")
		score1Str, _ := reader.ReadString('\n')
		score1Str = strings.TrimSpace(score1Str)
		score1, err1 := strconv.Atoi(score1Str)
		if err1 != nil || score1 < 0 {
			fmt.Println("Skor tidak valid! Program berhenti.")
			break
		}

		fmt.Print("Masukkan skor klub kedua: ")
		score2Str, _ := reader.ReadString('\n')
		score2Str = strings.TrimSpace(score2Str)
		score2, err2 := strconv.Atoi(score2Str)
		if err2 != nil || score2 < 0 {
			fmt.Println("Skor tidak valid! Program berhenti.")
			break
		}

		if score1 > score2 {
			winningClubs = append(winningClubs, club1)
			fmt.Printf("Pemenang: %s\n", club1)
		} else if score2 > score1 {
			winningClubs = append(winningClubs, club2)
			fmt.Printf("Pemenang: %s\n", club2)
		} else {
			fmt.Println("Pertandingan seri, tidak ada pemenang.")
		}
		fmt.Println()
	}

	fmt.Println("\nDaftar Klub yang Menang:")
	for i, club := range winningClubs {
		fmt.Printf("%d. %s\n", i+1, club)
	}
}


```
### Output:
![image](https://github.com/user-attachments/assets/fc2a25aa-6c57-4ea7-ade6-3186e4d48960)
![image](https://github.com/user-attachments/assets/6bb816ce-539a-4d4a-b585-b023825ed7b4)

### Full code Screenshot:
![code 3](https://github.com/user-attachments/assets/044200d3-f986-4d7b-9c5d-634db682adc2)

### Deskripsi Program : 
Program ini digunakan untuk merekap skor pertandingan sepak bola antara dua klub. Pengguna diminta untuk memasukkan nama dan skor dari kedua klub. Program akan menentukan pemenang berdasarkan skor yang lebih tinggi, dan mencatat nama klub yang menang ke dalam sebuah slice. Jika skor tidak valid atau negatif, program akan berhenti. Pada akhirnya, program menampilkan daftar klub yang menang sepanjang input diberikan.

### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom

### Source Code :
```go
//Haifa Zahra Azzimmi
//2311102163

package main

import (
	"fmt"
	"strings"
)

func reverseArray(arr []rune) []rune {
	n := len(arr)
	reversed := make([]rune, n)
	for i := 0; i < n; i++ {
		reversed[i] = arr[n-i-1]
	}
	return reversed
}

func isPalindrome(arr []rune) bool {
	reversed := reverseArray(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] != reversed[i] {
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Print("Masukkan sekumpulan karakter: ")
	fmt.Scanln(&input)

	charArray := []rune(strings.ToLower(input)) // Ignore case

	reversedArray := reverseArray(charArray)

	fmt.Printf("Array asli: %s\n", string(charArray))
	fmt.Printf("Array yang dibalik: %s\n", string(reversedArray))

	if isPalindrome(charArray) {
		fmt.Println("Array membentuk palindrom.")
	} else {
		fmt.Println("Array tidak membentuk palindrom.")
	}
}

```
### Output:
![image](https://github.com/user-attachments/assets/16882de6-d34e-4f8f-b40f-da1928357f2c)

### Full code Screenshot:
![code 4](https://github.com/user-attachments/assets/e67bc28b-06d6-40be-9028-a80bc578ce67)

### Deskripsi Program : 
Program ini meminta pengguna untuk memasukkan sekumpulan karakter dan mengubahnya menjadi array karakter dengan mengabaikan perbedaan huruf besar dan kecil. Program kemudian membalik urutan karakter dalam array tersebut. Setelah itu, program memeriksa apakah array asli dan array yang dibalik adalah sama, yang menunjukkan bahwa karakter-karakter tersebut membentuk palindrom. Jika ya, program akan mencetak bahwa array membentuk palindrom; jika tidak, program akan mencetak bahwa array tidak membentuk palindrom. Program ini membantu dalam menentukan simetri dalam sebuah sekumpulan karakter.
