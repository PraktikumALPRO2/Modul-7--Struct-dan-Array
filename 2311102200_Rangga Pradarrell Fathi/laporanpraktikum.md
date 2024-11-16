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
 Rangga Pradarrell Fathi
  <br>
  2311102200
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
Array adalah kumpulan data yang memiliki tipe yang sama dan disimpan dalam memori secara berurutan. Setiap elemen dalam array dapat diakses melalui indeks.

### <strong> Beberapa karakteristik utama dari Array:
- Semua elemen memiliki tipe data yang sama.
- Array memiliki ukuran tetap yang ditentukan saat deklarasi.
- Setiap elemen dapat diakses dengan menggunakan indeks numerik, di mana indeks pertama adalah 0.

<strong><h2>Definisi Struct</h2></strong>
Array adalah kumpulan data yang memiliki tipe yang sama dan disimpan dalam memori secara berurutan. Setiap elemen dalam array dapat diakses melalui indeks.

### <strong> Beberapa karakteristik utama dari Struct:
- Dapat menyimpan variabel dengan tipe data yang berbeda.
- Digunakan untuk membuat tipe data yang lebih kompleks, seperti data mahasiswa yang terdiri dari nama, usia, dan nilai.


# <strong> Guided </strong>
## Guided - 1
### Study Case
**Sebuah program yang menerima input waktu masuk parkir dan waktu keluar parkir dalam format jam, menit, dan detik. Program ini akan menghitung dan menampilkan lama waktu parkir dalam jam, menit, dan detik.** 

### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
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

### Screenshot Output
![image](https://github.com/user-attachments/assets/2e8f2f66-e80b-4bd6-863d-2a15c542bbbe)


### Deskripsi Program
Program di atas menghitung durasi waktu parkir dalam jam, menit, dan detik. Pengguna memasukkan waktu masuk (saat memulai parkir) dan waktu keluar (saat meninggalkan parkir) dalam format jam, menit, dan detik, kemudian program akan menghitung selisih waktu antara kedua waktu tersebut untuk mendapatkan durasi parkir dalam format yang sama.


## Guided - 2
#### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
//IF-11-05
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
### Screenshot Output
![image](https://github.com/user-attachments/assets/41a822b4-e57b-4d0d-9d88-be77b537dc9e)

### Deskripsi Program
Program berikut adalah program sederhana dalam bahasa Go yang melakukan beberapa tugas utama, yaitu memeriksa dan menambahkan nama baru ke dalam daftar teman yang sudah ada, dengan syarat nama tersebut belum ada dalam daftar. Jika nama sudah ada, program akan memberi tahu pengguna bahwa nama tersebut telah terdaftar.

## Guided - 3
#### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
//IF-11-05
//Guided 2 - Map

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
### Screenshot Output
![image](https://github.com/user-attachments/assets/9ad9ce96-7aa3-441f-aee7-5dd903c01450)


### Deskripsi Program

Program ini menggunakan struktur data map dalam bahasa Go untuk menyimpan dan menampilkan harga berbagai buah. Dalam map, setiap nama buah digunakan sebagai kunci, dan harga buah tersebut digunakan sebagai nilai.


# <strong> Unguided </strong>
## Unguided - 1
### Study Case
**1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.**                         
*Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.*                                 
*Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".*

#### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
//IF-11-05

package main

import (
	"fmt"
	"math"
)

// Definisi tipe data untuk Titik dan Lingkaran
type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat Titik
	radius int
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(p, q Titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

// Fungsi untuk mengecek apakah titik berada di dalam lingkaran
func diDalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {
	// Input: Data lingkaran 1, lingkaran 2, dan titik sembarang
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var px, py int

	// Masukkan koordinat lingkaran 1
	fmt.Scan(&cx1, &cy1, &r1)
	lingkaran1 := Lingkaran{Titik{cx1, cy1}, r1}

	// Masukkan koordinat lingkaran 2
	fmt.Scan(&cx2, &cy2, &r2)
	lingkaran2 := Lingkaran{Titik{cx2, cy2}, r2}

	// Masukkan koordinat titik sembarang
	fmt.Scan(&px, &py)
	titik := Titik{px, py}

	// Cek posisi titik terhadap lingkaran
	diLingkaran1 := diDalam(lingkaran1, titik)
	diLingkaran2 := diDalam(lingkaran2, titik)

	// Output hasil berdasarkan kondisi
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
#### Screenshoot Output


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
//Rangga Pradarrell Fathi
//2311102200
//IF-11-05
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Println("Masukkan jumlah elemen array:")
	fmt.Scan(&n)

	// Input array
	arr := make([]int, n)
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// Menu operasi
	for {
		fmt.Println("\nPilih operasi:")
		fmt.Println("a. Tampilkan seluruh isi array")
		fmt.Println("b. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("c. Tampilkan elemen dengan indeks genap")
		fmt.Println("d. Tampilkan elemen dengan kelipatan bilangan tertentu")
		fmt.Println("e. Hapus elemen dengan indeks tertentu")
		fmt.Println("f. Hitung rata-rata")
		fmt.Println("g. Hitung standar deviasi")
		fmt.Println("h. Hitung frekuensi elemen tertentu")
		fmt.Println("i. Keluar")
		
		var choice string
		fmt.Scan(&choice)

		switch choice {
		case "a":
			// Tampilkan seluruh isi array
			fmt.Println("Isi array:", arr)
		case "b":
			// Tampilkan elemen dengan indeks ganjil
			fmt.Println("Elemen dengan indeks ganjil:")
			for i := 1; i < len(arr); i += 2 {
				fmt.Print(arr[i], " ")
			}
			fmt.Println()
		case "c":
			// Tampilkan elemen dengan indeks genap
			fmt.Println("Elemen dengan indeks genap:")
			for i := 0; i < len(arr); i += 2 {
				fmt.Print(arr[i], " ")
			}
			fmt.Println()
		case "d":
			// Tampilkan elemen dengan kelipatan bilangan tertentu
			var x int
			fmt.Println("Masukkan bilangan kelipatan:")
			fmt.Scan(&x)
			fmt.Println("Elemen dengan kelipatan", x, ":")
			for _, val := range arr {
				if val%x == 0 {
					fmt.Print(val, " ")
				}
			}
			fmt.Println()
		case "e":
			// Hapus elemen dengan indeks tertentu
			var idx int
			fmt.Println("Masukkan indeks elemen yang akan dihapus:")
			fmt.Scan(&idx)
			if idx >= 0 && idx < len(arr) {
				arr = append(arr[:idx], arr[idx+1:]...)
				fmt.Println("Array setelah elemen dihapus:", arr)
			} else {
				fmt.Println("Indeks tidak valid!")
			}
		case "f":
			// Hitung rata-rata
			sum := 0
			for _, val := range arr {
				sum += val
			}
			rataRata := float64(sum) / float64(len(arr))
			fmt.Printf("Rata-rata: %.2f\n", rataRata)
		case "g":
			// Hitung standar deviasi
			sum := 0
			for _, val := range arr {
				sum += val
			}
			rataRata := float64(sum) / float64(len(arr))

			// Variansi
			var variansi float64
			for _, val := range arr {
				variansi += math.Pow(float64(val)-rataRata, 2)
			}
			variansi /= float64(len(arr))

			// Standar deviasi
			stdDev := math.Sqrt(variansi)
			fmt.Printf("Standar deviasi: %.2f\n", stdDev)
		case "h":
			// Hitung frekuensi elemen tertentu
			var target int
			fmt.Println("Masukkan elemen yang akan dihitung frekuensinya:")
			fmt.Scan(&target)
			frekuensi := 0
			for _, val := range arr {
				if val == target {
					frekuensi++
				}
			}
			fmt.Printf("Frekuensi %d: %d kali\n", target, frekuensi)
		case "i":
			// Keluar
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

```
#### Screenshoot Source Code


#### Screenshoot Output


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
//Rangga Pradarrell Fathi
//2311102200
//IF-11-05
package main

import "fmt"

func main() {
	// Memasukkan nama klub
	var klub1, klub2 string
	fmt.Print("Klub 1: ")
	fmt.Scanln(&klub1)
	fmt.Print("Klub 2: ")
	fmt.Scanln(&klub2)

	// Inisialisasi array untuk menyimpan pemenang
	var pemenang []string
	pertandinganKe := 1

	for {
		// Meminta input skor
		var skor1, skor2 int
		fmt.Printf("Pertandingan %d (masukkan skor %s dan %s): ", pertandinganKe, klub1, klub2)
		fmt.Scan(&skor1, &skor2)

		// Periksa jika ada skor negatif
		if skor1 < 0 || skor2 < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		// Menentukan hasil pertandingan
		if skor1 > skor2 {
			pemenang = append(pemenang, klub1)
			fmt.Printf("Hasil %d: %s\n", pertandinganKe, klub1)
		} else if skor1 < skor2 {
			pemenang = append(pemenang, klub2)
			fmt.Printf("Hasil %d: %s\n", pertandinganKe, klub2)
		} else {
			fmt.Printf("Hasil %d: Draw\n", pertandinganKe)
		}

		pertandinganKe++
	}

	// Menampilkan hasil akhir
	fmt.Println("\nDaftar klub yang menang:")
	for _, klub := range pemenang {
		fmt.Println(klub)
	}
}
```
#### Screenshoot Source Code


#### Screenshoot Output


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
//Rangga Pradarrell Fathi
//2311102200
//IF-11-05
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isArray(t *tabel, n *int) {
	var char rune
	fmt.Println("Masukkan teks karakter per karakter (akhiri dengan tanda titik '.'): ")
	*n = 0
	for {
		fmt.Scanf("%c", &char)
		if char == '.' || *n >= NMAX {
			break
		}
		t[*n] = char
		*n++
	}
}

func balikkanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func palindrome(t *tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func cetakArray(t *tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]))
	}
	fmt.Println()
}

func main() {
	var tab tabel
	var n int

	// Mengisi array
	isArray(&tab, &n)

	// Menampilkan array asli
	fmt.Print("Teks: ")
	cetakArray(&tab, n)

	// Membalikkan array
	balikkanArray(&tab, n)

	// Menampilkan array terbalik
	fmt.Print("Reverse teks: ")
	cetakArray(&tab, n)

	// Mengecek apakah palindrom
	if palindrome(&tab, n) {
		fmt.Println("Palindrome: true")
	} else {
		fmt.Println("Palindrome: false")
	}
}

```
#### Screenshoot Output


#### Deskripsi Program
Program ini merupakan program untuk mengisi array karakter dengan input teks dari user dan mengecek apakah teks tersebut merupakan palindrom atau tidak. Program akan mengecek apakah teks palindrom atau tidak.

Algoritma Program :
1. Definisikan array 'tabel' untuk menampung karakter yang diinputkan user.
2. Gunakan fungsi 'isiArray' untuk mengambil input dari user.
3. Fungsi 'palindrom' untuk memeriksa apakah teks di dalam array merupakan palindrom.
4. Tampilkan hasil pemeriksaan apakah teks palindrom atau tidak.

Cara kerja dari program ini yaitu meminta inputan dari user untuk memasukkan teks karakter per karakter. Karakter yang diinputkan akan di salin ke array. Program memeriksa apakah termasuk palindrom. Hasil pemeriksaan ini ditampilkan dengan mencetak 'true' jika teks adalah palindrom atau 'false' jika tidak palindrom.

