<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL VII</strong></h2>
<h2 align="center"><strong> STRUCK & ARRAY </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Andreas Besar Wibowo / 2311102198<br>
  S1 IF11-05
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

## I. Dasar Teori
### Definisi Struct
Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu. Property dalam struct, tipe datanya bisa bervariasi. Mirip seperti map, hanya saja key-nya sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda.

Dari sebuah struct, kita bisa buat variabel baru, yang memiliki atribut sesuai skema struct tersebut. Variabel tersebut dipanggil dengan istilah object atau variabel object. Dengan memanfaatkan struct, penyimpanan data yang sifatnya kolektif menjadi lebih mudah, lebih rapi, dan mudah untuk dikelola.

**Deklarasi Struct**
```go
type student struct {
    name string
    grade int
}
```
Struct student dideklarasikan memiliki 2 property, yaitu name dan grade. Property adalah istilah untuk variabel yang menempel ke struct[1].

### Definisi Array
Array adalah urutan item atau data (elemen) bernomor dan panjang yang tetapdari jenis tipe data tunggal yang sama (itu adalah struktur data yang homogen),tipe ini bisa apa saja dari tipe primitive seperti integer, string, dan hingga tipeyang ditentukan sendiri. Panjangnya harus berupa ekspresi konstan, yang harusdievaluasi ke nilai integer non-negatif. Indikasi penggunaan dari array yaitudengan tanda kurung siku (square brackets) “[]”[2]. 

### Definisi Slice
Slice atau irisan adalah segmen dari array. Seperti array, irisan dapat diindeksdan memiliki panjang. Berbeda dengan array panjang dari slice diperbolehkanuntuk berubah-ubah. Slice dapat berupa seluruh elemen atau subset dari item atauelemen yang ditunjukkan oleh indeks awal dan akhir[2].

### Definisi Map
Maps adalah tipe data asosiatif. Maps adalah kumpulan pasangan item yangtidak berurutan, dimana pasangan item terdiri dari kata kunci dan data atau nilai(maps juga kadang-kadang disebut array asosiatif, tabel hash, atau kamus). Mapsdigunakan untuk mencari nilai dengan kata kunci yang terkait dengan cepat[2].

## II. GUIDED
**1. Menghitung Lama Parkir (Struck)**
#### Source Code
```go
//Guided 1 - Struck

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
#### Screenshoot Source Code
![guided1 carbon](https://github.com/user-attachments/assets/944a7fdb-0f97-4af8-ace6-c7d3d682eef3)

#### Screenshoot Output
![Guided1 go](https://github.com/user-attachments/assets/789f0462-aa66-45f9-8132-9f8ebb19293c)

#### Deskripsi Program
Program ini merupakan program untuk menghitung dan menampilkan durasi waktu parkir berdasarkan jam masuk atau datang dan jam keluar atau pulang. Program ini menggunakan tipe data struct untuk waktu dengan format jam, menit, dan detik.

Algoritma Program :
1. Input waktu datang dan waktu pulang dengan format jam, menit, dan detik.
2. Konversi waktunya dalam satuan detik (second).
3. Hitung durasi parkir dalam detik dengan selisih jam datang dan jam pulang.
4. Konversi durasi parkir ke format jam, menit, dan detik.
5. Tampilkan durasi parkir.

Cara kerja dari program ini yaitu pengguna memasukkan waktu kedatangan (jam, menit, dan detik) dengan waktu kepulangan yang sama. Program mengkonversi waktu ke dalam detik dan menghitung selisih untuk menemukan durasi total parkir. Setelah itu dikonversi ke dalam format jam, menit, dan detik. Tampilkan hasilnya ke output.

**2. Menampilkan Harga Buah (Map)**
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
		"Apel":   5000,
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
#### Screenshoot Source Code
![guided2 carbon](https://github.com/user-attachments/assets/f0fd9bbd-6ad4-41e1-abe4-78ff591d8237)

#### Screenshoot Output
![Guided2 go](https://github.com/user-attachments/assets/8b005cea-0a32-4a65-a253-79991a1976d5)

#### Deskripsi Program
Program ini merupakan program untuk menampilkan daftar harga dari buah dengan menggunakan tipe data Map. Program ini akan menampilkan harga dari setiap buah di dalam Map. 

Algoritma Program :
1. Inisialisasi Map 'hargaBuah' dengan buah dan harga.
2. Gunakan perulangan untuk menampilkan setiap buah dan harganya.
3. Tampilkan harga untuk buah tertentu.

Cara kerja dari program ini yaitu menggunakan Map untuk menyimpan data harga dari buah. Setelah Map diisi, program mencetak daftar harga buah. Program juga menampilkan harga untuk buah secara spesifik.

**3. Daftar Teman (Slice)**
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
#### Screenshoot Source Code
![guided3 carbon](https://github.com/user-attachments/assets/579f1f84-144f-47cc-aa3a-fa2096aacedc)

#### Screenshoot Output
![Guided3 go](https://github.com/user-attachments/assets/df59faad-273e-4d40-9a3b-e62a751e6f24)

#### Deskripsi Program
Program ini merupakan program untuk menampilkan daftar teman dengan tipe data slice dan menambahkan nama baru ke dalam daftar jika belum ada. Jika nama sudah ada, program akan menampilkan nama tersebut jika sudah di dalam daftar.

Algoritma Program :
1. Inisialisasi slice 'daftarTeman'.
2. Buat slice 'namaBaru' yang berisi nama yang akan ditambahkan ke dalam 'daftarTeman'.
3. Menggunakan perulangan untuk mengecek nama.
4. Tampilkan daftar teman akhir setelah adanya proses penambahan.

Cara kerja dari program ini yaitu menyimpan daftar nama teman. Setiap menambahkan nama baru akan diperiksa. Jika ada nama baru akan ditambahkan ke slice. Jika sudah ada, program akan menampilkan nama tersebut sudah ada.

## III. UNGUIDED
**1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.**                         
*Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.*                                 
*Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".*

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / IF-11-05

package main

import (
	"fmt"
	"math"
)

// Mendefinisikan Tipe Alias
type koordinat = int
type jarakReal = float64

// Mendefinisikan Titik
type Titik struct {
	x koordinat
	y koordinat
}

// Mendefinisikan Lingkaran
type Lingkaran struct {
	tengah Titik
	radius koordinat
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(p, q Titik) jarakReal {
	return math.Sqrt(math.Pow(float64(q.x-p.x), 2) + math.Pow(float64(q.y-p.y), 2))
}

// Fungsi untuk memeriksa apakah titik berada di dalam lingkaran
func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.tengah, p) <= float64(c.radius)
}

func main() {
	// Input data lingkaran 1
	var cx1, cy1, r1 koordinat
	fmt.Print("Koordinat pusat dan radius pada lingkaran 1\n")
	fmt.Print("Koordinat pusat x: ")
	fmt.Scan(&cx1)
	fmt.Print("Koordinat pusat y: ")
	fmt.Scan(&cy1)
	fmt.Print("Radius lingkaran 1: ")
	fmt.Scan(&r1)

	// Input data lingkaran 2
	var cx2, cy2, r2 koordinat
	fmt.Print("Koordinat pusat dan radius pada lingkaran 2\n")
	fmt.Print("Koordinat pusat x: ")
	fmt.Scan(&cx2)
	fmt.Print("Koordinat pusat y: ")
	fmt.Scan(&cy2)
	fmt.Print("Radius lingkaran 2: ")
	fmt.Scan(&r2)

	// Input data titik sembarang
	var x, y koordinat
	fmt.Print("Koordinat titik sembarang (x, y)\n")
	fmt.Print("titik x: ")
	fmt.Scan(&x)
	fmt.Print("titik y: ")
	fmt.Scan(&y)

	// Membuat objek lingkaran dan titik
	lingkaran1 := Lingkaran{Titik{cx1, cy1}, r1}
	lingkaran2 := Lingkaran{Titik{cx2, cy2}, r2}
	titik := Titik{x, y}

	// Memeriksa posisi titik pada dua lingkaran
	diDalam1 := didalam(lingkaran1, titik)
	diDalam2 := didalam(lingkaran2, titik)

	// Mencetak Output Akhir
	fmt.Print("\nPosisi titik: ")
	if diDalam1 && diDalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
#### Screenshoot Source Code
![unguided1 carbon](https://github.com/user-attachments/assets/b040162c-268e-4893-b327-c2657ee6a2ff)

#### Screenshoot Output
![Unguided1(1) go](https://github.com/user-attachments/assets/de609317-3ac7-46c4-bb6c-5454e9a71c81)
![Unguided1(2) go](https://github.com/user-attachments/assets/4aae5a50-5831-4542-bb72-8f5f84c9ba1f)
![Unguided1(3) go](https://github.com/user-attachments/assets/e23d0b60-8523-46f9-b4a7-76984ebc1f04)
![Unguided1(4) go](https://github.com/user-attachments/assets/b0948942-935e-48a5-8605-5c172b18e33d)

#### Deskripsi Program
Program ini merupakan program untuk memeriksa apakah titik berada di dalam atau di luar dua lingkaran berdasarkan koordinat pusat dan radius lingkaran. Program meminta input dari user untuk menentukan pusat dan radius lingkaran, koordinatnya akan diperiksa.

Algoritma Program :
1. Inisialisasi tipe data 'lingkaran' untuk menyimpan pusat dan radius lingkaran dan 'Titik' untuk menyimpan koordinat.
2. Gunakan fungsi 'jarak' untuk menghitung jarak antar dua titik.
3. Ambil input dari pengguna untuk koordinat pusat dan radius.
4. Tentukan apakah titik berada di dalam lingkaran pertama, kedua, keduanya, atau di luar kedua lingkaran.
5. Tampilkan hasil posisi titik.

Cara kerja dari program ini yaitu menggunakan fungsi 'jarak' untuk menghitung jarak antara pusat lingkaran dan titik sembarang. Dengan ini, program memeriksa titik berada di dalam lingkaran dengan membandingkan jarak dengan radius lingkaran. Hasilnya akan ditampilkan di akhir.

**2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilal. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa Informasi berikut:**

_a. Menampilkan keseluruhan isi dari array._

_b. Menampilkan elemen-elemen array dengan indeks ganjil saja._

_c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke - 0 adalah genap)._

_d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna._

_e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid.. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil_

_f. Menampilkan rata-rata dari bilangan yang ada di dalam array._

_g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut._

_h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut._

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / IF-11-05

package main

import "fmt"

// Mendefinisikan tipe data alias
type bilangan int
type rataRata float64
type selisih float64

func main() {
	const maxElemen = 100
	var array [maxElemen]bilangan
	var n, x, indeksHapus, cariAngka bilangan

	// Meminta jumlah elemen dalam array
	fmt.Print("Masukkan jumlah elemen dalam array (maksimum 100): ")
	fmt.Scan(&n)

	// Mengecek jumlah elemen tidak lebih dari 100
	if n > maxElemen || n <= 0 {
		fmt.Println("Jumlah elemen harus antara 1 dan 100.")
		return
	}

	// Meminta inputan user
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < int(n); i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	// Menampilkan seluruh isi array
	fmt.Println("\nIsi array:")
	for i := 0; i < int(n); i++ {
		fmt.Printf("%d ", array[i])
	}
	fmt.Println()

	// Menampilkan elemen-elemen indeks ganjil
	fmt.Println("\nElemen dengan indeks ganjil:")
	for i := 1; i < int(n); i += 2 {
		fmt.Printf("Indeks %d: %d\n", i, array[i])
	}

	// Menampilkan elemen-elemen indeks genap
	fmt.Println("\nElemen dengan indeks genap:")
	for i := 0; i < int(n); i += 2 {
		fmt.Printf("Indeks %d: %d\n", i, array[i])
	}

	// Menampilkan elemen-elemen array dengan indeks kelipatan x
	fmt.Print("\nMasukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	for i := int(x); i < int(n); i += int(x) {
		fmt.Printf("Indeks %d: %d\n", i, array[i])
	}

	// Menghapus elemen array pada indeks tertentu
	fmt.Print("\nMasukkan indeks elemen yang akan dihapus: ")
	fmt.Scan(&indeksHapus)
	if indeksHapus >= 0 && indeksHapus < n {
		for i := int(indeksHapus); i < int(n)-1; i++ {
			array[i] = array[i+1]
		}
		n-- // Mengurangi jumlah elemen array
		fmt.Println("Isi array setelah penghapusan:")
		for i := 0; i < int(n); i++ {
			fmt.Printf("%d ", array[i])
		}
		fmt.Println()
	} else {
		fmt.Println("Indeks yang dimasukkan tidak valid.")
	}

	// Menghitung rata-rata pada array
	var total bilangan
	for i := 0; i < int(n); i++ {
		total += array[i]
	}
	rata := rataRata(float64(total) / float64(n))
	fmt.Printf("\nRata-rata dari array: %.2f\n", rata)

	// Menghitung selisih dari array
	var jumlahKuadrat float64
	for i := 0; i < int(n); i++ {
		selisih := float64(array[i]) - float64(rata)
		jumlahKuadrat += selisih * selisih
	}
	standarselisih := selisih(sqrt(jumlahKuadrat / float64(n)))
	fmt.Printf("Standar selisih dari array: %.2f\n", standarselisih)

	// Menampilkan frekuensi dari suatu bilangan tertentu di dalam array
	fmt.Print("\nMasukkan angka untuk mengetahui frekuensinya dalam array: ")
	fmt.Scan(&cariAngka)
	frekuensi := 0
	for i := 0; i < int(n); i++ {
		if array[i] == cariAngka {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi angka %d dalam array: %d kali\n", cariAngka, frekuensi)
}

// Fungsi untuk menghitung akar kuadrat
func sqrt(x float64) float64 {
	z := x
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
```
#### Screenshoot Source Code
![unguided2 carbon](https://github.com/user-attachments/assets/504c07c2-3349-418f-9077-b9a021ff8ab1)

#### Screenshoot Output
![Unguided2 go](https://github.com/user-attachments/assets/21e5468d-56ae-4d52-8404-07b981a9d6f0)

#### Deskripsi Program
Program ini merupakan program untuk mengelola data bilangan dalam array, pengguna menginputkan elemen-elemen array, menampilkan elemen berdasarkan ganjil,genap, dan kelipatan nilai x, menghapus elemen, menghitung rata-rata dan selisih, dan menghitung frekuensi suatu angka dalam array.

Algoritma Program :
1. Inisialisasi array dengan tipe 'bilangan' dan jumlah maksimum elemen (100).
2. Ambil input dari user.
3. Minta user untuk memasukkan elemen-elemen array satu per satu.
4. Tampilkan seluruh elemen array.
5. Tampilkan elemen yang berada pada indeks ganjil dan genap.
6. Ambil input 'x' dan menampilkan elemen pada indeks kelipatan 'x'.
7. Ambil input untuk menghapus elemen pada posisi tersebut dan memperbarui array.
8. Hitung rata-rata elemen.
9. Hitung selisih dari array berdasarkan rata-rata elemen.
10. Ambil input angka dari user dan hitung frekuensi angka dalam array.

Cara kerja program ini yaitu meminta user memasukkan elemen array dan menampilkannya. Program dapat mengidentifikasi elemen yang berada pada ganjil, genap, dan kelipatan dari suatu nilai 'x'. User dapat menghapus elemen. Program menghitung rata-rata dan selisih elemen-elemen dalam array. Pada akhir, program menghitung dan menampilkan frekuensi angka dalam array.

**3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga.**                       

*Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja.*                                  
*Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.*

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / IF-11-05

package main

import "fmt"

// Mendefinisikan tipe data Alias
type namaKlub string
type skor int

func main() {
	var klubA, klubB namaKlub
	var skorA, skorB skor
	var pemenang []namaKlub

	// Input nama-nama klub
	fmt.Print("Masukkan nama klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Masukkan nama klub B: ")
	fmt.Scan(&klubB)

	// Menampilkan nama klub
	fmt.Printf("\nKlub A: %s\n", klubA)
	fmt.Printf("Klub B: %s\n\n", klubB)

	// Proses dari inputan skor
	var i int
	for {
		i++
		fmt.Printf("Pertandingan %d - Masukkan skor %s: ", i, klubA)
		fmt.Scan(&skorA)
		fmt.Printf("Pertandingan %d - Masukkan skor %s: ", i, klubB)
		fmt.Scan(&skorB)

		// Hentikan jika skor salah satu atau kedua klub tidak valid
		if skorA < 0 || skorB < 0 {
			fmt.Println("\nPertandingan selesai")
			break
		}

		// Menentukan pemenang pertandingan
		if skorA > skorB {
			pemenang = append(pemenang, klubA)
			fmt.Printf("Hasil %d: %s\n", i, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
			fmt.Printf("Hasil %d: %s\n", i, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
			fmt.Printf("Hasil %d: Draw\n", i)
		}
	}

	// Menampilkan klub yang memenangkan pertandingan
	fmt.Println("\n Klub pemenang:")
	for j, klub := range pemenang {
		fmt.Printf("Hasil %d: %s\n", j+1, klub)
	}
}
```
#### Screenshoot Source Code
![unguided3 carbon](https://github.com/user-attachments/assets/cdeb9882-f33f-4939-89e6-620539a2de8a)

#### Screenshoot Output
![Unguided3 go](https://github.com/user-attachments/assets/6acb2777-0740-483c-9dbb-f87e3b98163a)

#### Deskripsi Program
Program ini merupakan program untuk mencatat dan menampilkan hasil pertandingan antara dua klub sepak bola. Skor nya diinputkan oleh user, dan program akan menentukan pemenang dari pertandingan itu dan menampilkan hasil imbang. Program berakhir jika skornya tidak valid (negatif). Diakhir, akan ditampilkan daftar hasil seluruh pertandingan.

Algoritma Program :
1. Inisialisasi variabel 'klubA' dan 'klubB' untuk menyimpan nama klub, dan 'skorA' dan 'skorB' untuk menyimpan skor setiap klub.
2. Ambil input nama dari kedua klub.
3. Mulai perulangan untuk input skor dari masing-masing klub di setiap pertandingan.
4. Tentukan pemenang berdasarkan skor yang lebih tinggi atau hasilnya imbang, lalu simpan hasilnya.
5. Hentikan perulangan jika skor tidak valid.
6. Tampilkan hasil setiap pertandingan dari daftar pemenang.

Cara kerja dari program yaitu meminta user untuk memasukkan nama dua klub yang akan bertanding. User memasukkan skor untuk kedua klub di setiap pertandingan. Program untuk menentukan pemenang dari skor yang lebih unggul atau mencatat hasil imbang jika skornya sama. Ketika skor tidak valid, program akan berhenti dan menampilan hasil pertandingan.

**4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan Isi array dan memeriksa apakah membentuk palindrom.**                         
*Modifikasi program tersebut dengan menambahkan fungsi palindrom. Tambahkan instruksi untuk memanggil fungsi tersebut dan menampilkan hasilnya pada program utama.*

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / IF-11-05

package main

import "fmt"

const NMAX int = 999

// Mendefinisikan tipe Data Alias
type tabel [NMAX]byte

// Fungsi untuk mengisi array dari inputan user
func isiArray(t *tabel, n *int) {
	var input string
	fmt.Print("Teks: ")
	fmt.Scanln(&input) 
	*n = len(input) 
	for i := 0; i < *n; i++ {
		t[i] = input[i]
	}
}

// Fungsi untuk memeriksa apakah array merupakan palindrom
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var a int

	// Mengisi array tab
	isiArray(&tab, &a)

	// Menampilkan Output apakah palindrom
	isPalindrom := palindrom(tab, a)
	fmt.Printf("Palindrom: %t\n", isPalindrom)
}
```
#### Screenshoot Source Code
![unguided4 carbon](https://github.com/user-attachments/assets/ec97f80c-e7cf-49c5-8f94-90953f796759)

#### Screenshoot Output
![Unguided4(1) go](https://github.com/user-attachments/assets/c865104a-3ffb-4516-b701-0827687c42d3)
![Unguided4(2) go](https://github.com/user-attachments/assets/6b2ce938-a633-4d78-a325-16e806054b76)

#### Deskripsi Program
Program ini merupakan program untuk mengisi array karakter dengan input teks dari user dan mengecek apakah teks tersebut merupakan palindrom atau tidak. Program akan mengecek apakah teks palindrom atau tidak.

Algoritma Program :
1. Definisikan array 'tabel' untuk menampung karakter yang diinputkan user.
2. Gunakan fungsi 'isiArray' untuk mengambil input dari user.
3. Fungsi 'palindrom' untuk memeriksa apakah teks di dalam array merupakan palindrom.
4. Tampilkan hasil pemeriksaan apakah teks palindrom atau tidak.

Cara kerja dari program ini yaitu meminta inputan dari user untuk memasukkan teks karakter per karakter. Karakter yang diinputkan akan di salin ke array. Program memeriksa apakah termasuk palindrom. Hasil pemeriksaan ini ditampilkan dengan mencetak 'true' jika teks adalah palindrom atau 'false' jika tidak palindrom.

## Referensi
[1] Novalagung. Variabel dalam Go. Diakses dari https://dasarpemrogramangolang.novalagung.com/A-variabel.html
                                                                        
[2] Ahmad Faisal. (2021). Dasar-Dasar Bahasa Pemrograman Golang. Diakses dari https://pta.pilkommedia.org/progress/upload/AhmadFaisal_A1C615001_Dasar-DasarBahasaPemrogramanGolang.pdf 
