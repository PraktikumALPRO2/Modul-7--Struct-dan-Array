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
