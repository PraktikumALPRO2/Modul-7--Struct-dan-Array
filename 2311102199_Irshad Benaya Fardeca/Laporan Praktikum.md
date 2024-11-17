<h2 align="center">LAPORAN PAKTIKUM ALGORITMA DAN PEMROGRAMAN 2</h2>
<h2 align="center">MODUL 7</h2>
<h2 align="center">Struct dan Array</h2>

<p align="center"><img src=https://github.com/user-attachments/assets/37e9c953-078b-4ef4-97e7-a5d25344e50b alt="Logo" width="300"/><p>

<p align="center">Disusun Oleh : </p>
<p align="center">Irshad Benaya Fardeca / 2311102199</p>
<p align="center">IF-11-05</p>
<br></br>
<p align="center">Dosen Pengampu : </p>
<p align="center">Arif Amrulloh, S.Kom., M.Kom.</p>
<br></br>
<h3 align="center">PROGRAM STUDI S1 TEKNIK INFORMATIKA</h3>
<h3 align="center">FAKULTAS INFORMATIKA</h3>
<h3 align="center">TELKOM UNIVERSITY PURWOKERTO</h3>
<h3 align="center">2024</p>

<br></br>

#### I. DASAR TEORI
#### Struct
Go tidak mengadopsi konsep class seperti pada beberapa bahasa pemrograman OOP lainnya. Namun Go memiliki tipe data struktur Struct.
Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu. Property dalam struct, tipe datanya bisa bervariasi. Mirip seperti map , hanya saja key-nya sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda.
Dari sebuah struct, kita bisa buat variabel baru, yang memiliki atribut sesuai skema struct tersebut. Kita sepakati dalam buku ini, variabel tersebut dipanggil dengan istilah object atau variabel object.
Dengan memanfaatkan struct, penyimpanan data yang sifatnya kolektif menjadi
lebih mudah, lebih rapi, dan mudah untuk dikelola.
Kombinasi keyword type dan struct digunakan untuk deklarasi struct. Di bawah ini merupakan contoh cara penerapannya.
```
type student struct {
 name string
 grade int
}
```
Struct student dideklarasikan memiliki 2 property, yaitu name dan grade. Property adalah istilah untuk variabel yang menempel ke struct.


#### Array
Array adalah kumpulan data bertipe sama, yang disimpan dalam sebuah variabel. Array memiliki kapasitas yang nilainya ditentukan pada saat pembuatan, menjadikan elemen/data yang disimpan di array tersebut jumlahnya tidak boleh melebihi yang sudah dialokasikan.
Default nilai tiap elemen array pada awalnya tergantung dari tipe datanya. Jika int maka tiap element zero value-nya adalah 0 , jika bool maka false , dan seterusnya. Setiap elemen array memiliki indeks berupa angka yang merepresentasikan posisi urutan elemen tersebut. Indeks array dimulai dari 0.
Contoh penerapan array:
```
var names [4]string
names[0] = "trafalgar"
names[1] = "d"
names[2] = "water"
names[3] = "law"
fmt.Println(names[0], names[1], names[2], names[3])
```
Variabel names dideklarasikan sebagai array string dengan alokasi kapasitas elemen adalah 4 slot. Cara mengisi slot elemen array bisa dilihat di kode di atas, yaitu dengan langsung mengakses elemen menggunakan indeks, lalu mengisinya.


#### II. GUIDED

##### 1. Sruct
##### Source code
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

	dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600
	dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600
	lParkir = dPulang - dParkir

	durasi.jam = lParkir / 3600
	durasi.menit = lParkir % 3600 / 60
	durasi.detik = lParkir % 3600 % 60
	fmt.Printf("Lama Parkir : %d jam %d menit %d detik", durasi.jam, durasi.menit, durasi.detik)
}

```
##### Screenshoot Output
![Screenshot 2024-11-17 204913](https://github.com/user-attachments/assets/f27e95b2-e784-43ec-9bba-f8c4a5437690)

##### Deskripsi Program
Program ini dirancang untuk menghitung durasi waktu parkir kendaraan. user diminta untuk memasukkan waktu kedatangan dan keberangkatan dalam format jam, menit, dan detik. Program kemudian akan mengkonversi waktu tersebut menjadi detik untuk memudahkan perhitungan selisih waktu. Setelah itu, selisih waktu dalam detik dikonversi kembali ke format jam, menit, dan detik untuk ditampilkan sebagai hasil akhir, yaitu lama waktu kendaraan tersebut parkir.

##### 2. Slice
##### Source code
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
##### Screenshoot Output
![Screenshot 2024-11-17 205119](https://github.com/user-attachments/assets/2a960f8f-c7fc-4821-a9f7-d4b96060fb42)

##### Deskripsi Program
Program ini dirancang untuk mengelola daftar teman secara dinamis menggunakan slice. Program memulai dengan sebuah slice yang sudah berisi beberapa nama teman. Kemudian, program mencoba menambahkan nama-nama baru dari slice lain. Sebelum menambahkan, program akan memeriksa terlebih dahulu apakah nama tersebut sudah ada dalam daftar menggunakan fungsi sudahAda. Jika nama belum ada, maka nama tersebut akan ditambahkan ke dalam slice daftarTeman menggunakan fungsi append. Namun, jika nama sudah ada, program akan memberikan pesan bahwa nama tersebut sudah terdaftar. Akhirnya, program akan mencetak daftar teman yang sudah diperbarui. Dengan demikian, program ini memastikan bahwa setiap nama hanya muncul satu kali dalam daftar.

##### 3. Map
##### Source code
```go
/// Guided 3 - Map
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
##### Screenshoot Output
![Screenshot 2024-11-17 205246](https://github.com/user-attachments/assets/2556d433-fc64-483c-bcbc-48f13fedbec2)

##### Deskripsi Program
Program ini dibuat dengan menggunakan fungsi map, map digunakan untuk menyimpan informasi harga berbagai jenis buah. Program memulai dengan membuat map bernama hargaBuah yang memiliki kunci berupa nama buah (string) dan nilai berupa harga buah (integer). Selanjutnya, program menampilkan semua pasangan kunci-nilai dalam map tersebut. Terakhir, program mengakses dan menampilkan harga buah mangga secara spesifik, menunjukkan bagaimana cara mengakses nilai dalam map berdasarkan kuncinya.

<br></br>


#### III. UNGUIDED

##### 1. Unguided 1
Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radlusnya. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

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
	pusat  Titik
	radius int
}

func jarak(p1, p2 Titik) float64 {
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)))
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {

	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Println("Masukkan lingkaran 1 (cx1, cy1, r1):")
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Println("Masukkan lingkaran 2 (cx2, cy2, r2):")
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Println("Masukkan titik (x, y):")
	fmt.Scan(&x, &y)

	lingkaran1 := Lingkaran{Titik{cx1, cy1}, r1}
	lingkaran2 := Lingkaran{Titik{cx2, cy2}, r2}
	titik := Titik{x, y}

	dalamLingkaran1 := didalam(lingkaran1, titik)
	dalamLingkaran2 := didalam(lingkaran2, titik)

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
##### Screenshoot Output
![Screenshot 2024-11-17 212021](https://github.com/user-attachments/assets/5703fa1d-104a-4909-9cc1-e00a105c3143)

##### Deskripsi Program
Program ini dirancang untuk menentukan posisi suatu titik terhadap dua lingkaran yang diberikan. Program dimulai dengan mendefinisikan struktur data Titik dan Lingkaran untuk merepresentasikan titik di bidang kartesian dan lingkaran dengan pusat dan jari-jari tertentu. Fungsi jarak digunakan untuk menghitung jarak antara dua titik, sedangkan fungsi didalam mengecek apakah suatu titik berada di dalam lingkaran.

##### 2. Unguided 2
Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa Informasi berikut:
a. Menampilkan keseluruhan isi dari array.
b. Menampilkan elemen-elemen array dengan indeks ganjil saja.
c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-O adalah genap).
d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan user.
e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil
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

func main() {
	var n int
	fmt.Print("Tentukan jumlah elemen array: ")
	fmt.Scan(&n)

	array := make([]int, n)

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	for {
		fmt.Println("\nPilih operasi:")
		fmt.Println("1. Menampilkan seluruh isi array")
		fmt.Println("2. Menampilkan elemen array dengan indeks ganjil")
		fmt.Println("3. Menampilkan elemen array dengan indeks genap")
		fmt.Println("4. Menampilkan elemen array dengan indeks kelipatan x")
		fmt.Println("5. Hapus elemen array pada indeks tertentu")
		fmt.Println("6. Menampilkan rata-rata elemen array")
		fmt.Println("7. Menampilkan standar deviasi array")
		fmt.Println("8. Menampilkan frekuensi suatu bilangan")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih Nomer : ")

		var pilih int
		fmt.Scan(&pilih)

		switch pilih {
		case 1:

			fmt.Println("Isi array:", array)

		case 2:

			fmt.Println("Elemen dengan indeks ganjil:")
			for i := 1; i < len(array); i += 2 {
				fmt.Printf("array[%d] = %d\n", i, array[i])
			}

		case 3:

			fmt.Println("Elemen dengan indeks genap:")
			for i := 0; i < len(array); i += 2 {
				fmt.Printf("array[%d] = %d\n", i, array[i])
			}

		case 4:

			var x int
			fmt.Println("Masukkan bilangan x:")
			fmt.Scan(&x)
			fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
			for i := 0; i < len(array); i++ {
				if i%x == 0 {
					fmt.Printf("array[%d] = %d\n", i, array[i])
				}
			}

		case 5:

			var idx int
			fmt.Println("Masukkan indeks yang ingin dihapus:")
			fmt.Scan(&idx)
			if idx >= 0 && idx < len(array) {
				array = append(array[:idx], array[idx+1:]...)
				fmt.Println("Array setelah penghapusan:", array)
			} else {
				fmt.Println("Indeks tidak valid!")
			}

		case 6:

			sum := 0
			for _, val := range array {
				sum += val
			}
			rataRata := float64(sum) / float64(len(array))
			fmt.Printf("Rata-rata: %.2f\n", rataRata)

		case 7:

			sum := 0
			for _, val := range array {
				sum += val
			}
			rataRata := float64(sum) / float64(len(array))

			var variance float64
			for _, val := range array {
				variance += math.Pow(float64(val)-rataRata, 2)
			}
			variance /= float64(len(array))
			deviasi := math.Sqrt(variance)
			fmt.Printf("Standar deviasi: %.2f\n", deviasi)

		case 8:

			var target int
			fmt.Println("Masukkan bilangan yang ingin dicari frekuensinya:")
			fmt.Scan(&target)

			frekuensi := 0
			for _, val := range array {
				if val == target {
					frekuensi++
				}
			}
			fmt.Printf("Frekuensi bilangan %d: %d kali\n", target, frekuensi)

		case 9:

			fmt.Println("Anda telah keluar dari Program")
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

```
##### Screenshoot Output
![Screenshot 2024-11-17 221526](https://github.com/user-attachments/assets/dc0c9fcd-2f6d-402f-bcea-48124fb5d821)
![Screenshot 2024-11-17 221538](https://github.com/user-attachments/assets/df426983-d840-499a-8b9d-ac3e8d89c0d5)
![Screenshot 2024-11-17 221550](https://github.com/user-attachments/assets/e5156bad-7a6e-4fd6-900e-d05b076b86d0)
![Screenshot 2024-11-17 221600](https://github.com/user-attachments/assets/0e21e245-40bc-4886-a8d8-5921d0e11f74)
![Screenshot 2024-11-17 221608](https://github.com/user-attachments/assets/2f270454-81da-4a9a-95bd-3267efb72838)
![Screenshot 2024-11-17 221616](https://github.com/user-attachments/assets/2fe7a540-2b8f-41da-9100-86e08d30fbdd)
![Screenshot 2024-11-17 221623](https://github.com/user-attachments/assets/9eb48777-b005-42c8-b72f-c2acac2cc52d)
![Screenshot 2024-11-17 221630](https://github.com/user-attachments/assets/226f55f6-50c4-41bd-b240-182c951e9e4b)
![Screenshot 2024-11-17 221639](https://github.com/user-attachments/assets/47c17dd5-4a14-4844-b72c-5177e9fc7ff8)

##### Deskripsi Program
Program ini dibuat untuk memanipulasi array. Pertama user akan menginputkan sejumlah elemen yang ingin dimasukkan ke dalam array. Setelah itu, program akan menampilkan menu interaktif yang memungkinkan user untuk memilih operasi yang ingin dilakukan, mulai dari menampilkan seluruh elemen array, menampilkan elemen dengan indeks tertentu (ganjil, genap, kelipatan x), menghapus elemen, menghitung rata-rata, menghitung standar deviasi, hingga mencari frekuensi kemunculan suatu bilangan dalam array.

##### 3. Unguided 3
Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja.
Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.
#### Source Code
```go
package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string
	var pertandingan int = 1

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	fmt.Println("\nMasukkan skor pertandingan:")

	for {
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			fmt.Println("\nPertandingan selesai.")
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}
		pertandingan++
	}

	fmt.Println("\nHasil pertandingan:")

	for i, klub := range pemenang {
		fmt.Printf("Hasil %d. %s\n", i+1, klub)
	}
}

```
##### Screenshoot Output
![Screenshot 2024-11-17 220314](https://github.com/user-attachments/assets/ca3bcd26-6ab1-4d29-a089-f147ed12c9fb)

##### Deskripsi Program
Program ini digunakan untuk mencatat hasil pertandingan antara dua klub sepak bola dan menampilkan daftar pemenangnya. user diminta untuk memasukkan nama kedua klub yang bertanding. Kemudian, program akan meminta input skor untuk setiap pertandingan secara berulang hingga user memasukkan skor negatif. Setelah semua data pertandingan dimasukkan, program akan menampilkan daftar pemenang dari setiap pertandingan. Jika terjadi seri, maka kata "Draw" akan ditampilkan sebagai hasil.

##### 4. Unguided 4
Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.
#### Source Code
```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	*n = 0
	fmt.Println("Masukkan karakter (akhiri dengan titik '.'): ")
	for {
		var c rune
		fmt.Scanf("%c", &c)
		if c == '.' || *n >= NMAX {
			break
		}
		t[*n] = c
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

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
	var m int

	isiArray(&tab, &m)

	fmt.Print("Teks : ")
	cetakArray(tab, m)

	balikanArray(&tab, m)
	fmt.Print("Reverse Teks: ")
	cetakArray(tab, m)

	isPalindrom := palindrom(tab, m)
	fmt.Printf("Teks Palindrom? %v\n", isPalindrom)
}

```
##### Screenshoot Output
![Screenshot 2024-11-17 214530](https://github.com/user-attachments/assets/0b7067c5-fe04-478a-af74-86e51e0e4d7e)

##### Deskripsi Program
Program ini digunakan untuk menerima input karakter dari user, menyimpannya dalam sebuah array, menampilkannya, membalik urutan elemennya, dan menentukan apakah array tersebut merupakan palindrom. Fungsi isiArray digunakan untuk mengisi array dengan karakter input, cetakArray untuk menampilkan isi array, balikanArray untuk membalik urutan elemen array, dan palindrom untuk menentukan apakah array tersebut merupakan palindrom.

### Referensi
[1] Dasar Pergraman Golang, Noval Agung

