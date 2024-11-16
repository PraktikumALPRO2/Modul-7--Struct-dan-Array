# <h2 align="center">LAPORAN PRAKTIKUM</h2>
# <h2 align="center">ALGORITMA DAN PEMROGRAMAN 2</h2>
# <h2 align="center">MODUL 7</h2>
# <h2 align="center"> STRUCT & ARRAY </h2>
<p align="center">
    <img src="https://github.com/user-attachments/assets/3ccfed0b-72d1-4349-ac08-c4dce379c827" alt="Gambar">
</p>
 <h3  align="center" >Disusun Oleh : </h3>
<p align="center">Amanda Windhu Gustyas - 2311102121</p>
<p align="center">IF-11-05</p>
 <h3 <p align="center" >Dosen Pengampu : </h3> </p>
 <p align="center">Arif Amrulloh, S.Kom., M.Kom.</p>

# <h3 align="center"> PROGRAM STUDI S1 TEKNIK INFORMATIKA </h3>
# <h3 align="center"> FAKULTAS INFORMATIKA </h3>
# <h3 align="center"> TELKOM UNIVERSITY PURWOKERTO </h3>
# <h3 align="center"> 2024 </h3>

## I. DASAR TEORI

### 1. Tipe Bentukan
Tipe bentukan memungkinkan pemrograman untuk mendefinisikan suatu tipe data baru pada suatu bahasa pemrograman. Tipe bentukan ini dapat dibedakan atas dua jenis, yaitu Alias dan Struct.</br>
    - Alias (Type)<br/>
    Bahasa pemrograman pada umumnya mengizinkan pemrograman untuk mengubah nama suatu tipe data dengan nama baru yang lebih ringkas dan familiar. Sebagai contoh
    "integer" dapat diubah dengan nama alias "bilangan". Caranya dengan menggunakan kata kunci "type".<br/>
    ![image](https://github.com/user-attachments/assets/13275380-ce9e-4c1e-bce0-0affa3239e62)
    Sebagai contoh perhatikan program Go berikut beserta  hasil eksekusinya!<br/>
    ![image](https://github.com/user-attachments/assets/887c4efc-3d81-4201-9d80-75db2e59f166)
    - Struct atau Record<br/>
    Structure memungkinkan pemrograman untuk mengelompokkan beberapa data atau nilai yang memiliki relasi atau keterkaitan tertentu menjadi suatu kesatuan. Masing
    masing nilai tersimpan dalam field dari structure tersebut.<br/>
    ![image](https://github.com/user-attachments/assets/ee368f95-8266-4a14-87da-1f85bcf2cc3e)
    Berbeda dengan bahasa pemrograman lain, kesamaan tipe dari dua variabel berjenis structure bukan karena namanya tetapi karena strukturnya. Dua variabel dengan 
    nama-nama field dan tipe field yang sama (dan dalam urutan yang sama) dianggap mempunyai tipe yang sama. Tentunya akan lebih memudahkan jika structure tersebut
    didefinisikan sebagai sebuah tipe baru, sehingga deklarasi structure tidak perlu lagi seluruh field-nya ditulis ulang berkali-kali.<br/>
    ![image](https://github.com/user-attachments/assets/afbe3b7c-6617-4765-9a43-459f408d9f63)

### 2. Array<br/>
Array mempunyai ukuran (jumlah elemen) yang tetap (statis) selama eksekusi program, sehingga jumlah elemen array menjadi bagian dari deklarasi variabel dengan tipe array.<br/>
![image](https://github.com/user-attachments/assets/bf257fb1-0618-4938-adfc-f4421d286980)
Jumlah elemen array dapat diminta dengan fungsi len yang tersedia. Sebagai contoh len(arr) akan menghasilkan 73 untuk contoh di atas.<br/>
Indeks array dimulai dari 0, sehingga indeks array pada contoh adalah 0, 1.. len(arr)-1<br/>
Contoh:<br/>
![image](https://github.com/user-attachments/assets/9d6a3239-2ae2-4fcb-b98f-adf08c23df36)<br/>
    - Slice (Array Dinamik)<br/>
    Array dalam Go juga dapat mempunyai ukuran yang dinamik. (Tidak digunakan di kelas Algoritma Pemrograman). Deklarasinya mirip dengan deklarasi array, tetapi
    jumlah elemennya dikosongkan.<br/>
    ![image](https://github.com/user-attachments/assets/fd8d3f2a-0494-429f-82aa-f1c42870e8d0)
    Sebuah slice dapat direalokasi menggunakan fungsi built-in make:<br/>
    ![image](https://github.com/user-attachments/assets/eccdd951-b904-4162-a9de-86b03439c704)
    Fungsi built-in len dapat digunakan untuk mengetahui ukuran slice. Fungsi lain, cap, dapat digunakan untuk mengetahui total tempat yang disediakan untuk slice
    tersebut.<br/>
    ![image](https://github.com/user-attachments/assets/f910cb7c-0062-45a2-b702-0113e6e81fde)
    Fungsi built-in append dapat digunakan untuk menambahkan elemen ke suatu slice, dan bila perlu memperbesar tempat untuk slice tersebut.<br/>
    ![image](https://github.com/user-attachments/assets/f36ecae1-0f8b-4a6a-8a48-18ab7b97a883)
    Sebuah slice baru juga dapat terbentuk dengan mengambil slice dari suatu array atau slice yang lain.<br/>
    ![image](https://github.com/user-attachments/assets/ab36eea8-d2ef-4bf5-9ad8-f3d9059baaad)
    - Map<br/>
    Tipe array lain, sebuah array dinamik. Indeksnya (di sini disebut kunci) tidak harus berbentuk integer. Indeks dapat berasal dari tipe apa saja. Struktur ini
    disebut map.<br/>
    ![image](https://github.com/user-attachments/assets/7916e324-fdc1-48f4-a304-0f805bd3c920)<br/>
    ![image](https://github.com/user-attachments/assets/7a9c1cdb-dc40-4083-ae3d-22210e26ebce)

## II. GUIDED

### 1. Program Menghitung Lama Parkir

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
	dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600 // Konversi ke detik
	dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600 // detik
	lParkir = dPulang - dParkir                                   //detik dari pulang-datang

	durasi.jam = lParkir / 3600
	durasi.menit = lParkir % 3600 / 60
	durasi.detik = lParkir % 3600 % 60 //17
	fmt.Printf("Lama Parkir : %d jam %d menit %d detik", durasi.jam, durasi.menit, durasi.detik)
}
```
## Output: ![image](https://github.com/user-attachments/assets/a2c3298e-bffc-45dd-a9d1-2a809e2cd487)

Program ini dirancang untuk menghitung durasi waktu parkir kendaraan berdasarkan waktu datang dan waktu pulang. Data waktu datang dan pulang dimasukkan dalam format jam, menit, dan detik. Program mengonversi kedua waktu tersebut menjadi detik untuk mempermudah perhitungan selisihnya. Setelah mendapatkan selisih waktu dalam bentuk detik, program mengubahnya kembali menjadi format jam, menit, dan detik yang mudah dipahami. Hasil akhirnya ditampilkan sebagai durasi parkir dalam format yang jelas dan terstruktur.

### 2. Program Slice

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
## Output: ![image](https://github.com/user-attachments/assets/ce78a6fa-3263-449a-82f0-bd6acb183a0a)

Program ini menambahkan nama baru ke daftar teman, tapi hanya jika belum ada di dalamnya. Nama baru dicek menggunakan fungsi sudahAda. Jika belum ada, nama ditambahkan; jika sudah ada, program memberi pemberitahuan. Hasil akhirnya adalah daftar teman yang diperbarui.

### 2. Program Map

```go
// Guided 2 - Map
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

	// Menampilkan harga buah Mangga secara spesifik
	if harga, ada := hargaBuah["Mangga"]; ada {
		fmt.Printf("Harga buah Mangga = Rp%d\n", harga)
	} else {
		fmt.Println("Buah Mangga tidak ditemukan dalam daftar.")
	}
}
```
## Output: ![image](https://github.com/user-attachments/assets/f6c3009d-704e-4561-a00e-3575eb2b72e3)

Program ini menyimpan daftar harga buah menggunakan map dengan nama buah sebagai kunci dan harga sebagai nilai. Program pertama-tama menampilkan harga semua buah yang ada dalam map. Kemudian, program mencari harga buah Mangga dan menampilkannya. Jika buah Mangga tidak ada, program akan memberi tahu bahwa buah tersebut tidak ditemukan dalam daftar.

## III. UNGUIDED

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx,cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x,y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya. <br> Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.<br/> Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2". <br/> Fungsi untuk menghitung jarak titik (a,b) dan (c,d) di mana rumus jarak adalah:<br/> ![image](https://github.com/user-attachments/assets/1ec76514-808a-436f-9d61-a908461f0954) <br/> Dan juga fungsi untuk menentukan posisi sebuah titik sembarang berada di dalam suatu lingkaran atau tidak.<br/> ![image](https://github.com/user-attachments/assets/862f3db5-0399-4b10-a06b-08f053461f68)

```go
package main

import (
	"fmt"
	"math"
)

// Struct untuk merepresentasikan titik
type Titik struct {
	x, y int
}

// Struct untuk merepresentasikan lingkaran
type Lingkaran struct {
	titikPusat Titik
	radius     int
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(p, q Titik) float64 {
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))
}

// Fungsi untuk mengecek apakah sebuah titik berada di dalam lingkaran
func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.titikPusat, p) <= float64(c.radius)
}

func main() {
	// Input lingkaran 1 dan 2
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	fmt.Scan(&cx1, &cy1, &r1) // Input lingkaran 1
	fmt.Scan(&cx2, &cy2, &r2) // Input lingkaran 2

	lingkaran1 := Lingkaran{Titik{cx1, cy1}, r1}
	lingkaran2 := Lingkaran{Titik{cx2, cy2}, r2}

	// Input titik sembarang
	var x, y int
	fmt.Scan(&x, &y)
	titik := Titik{x, y}

	// Menentukan posisi titik
	inLingkaran1 := didalam(lingkaran1, titik)
	inLingkaran2 := didalam(lingkaran2, titik)

	if inLingkaran1 && inLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```
## Output: ![image](https://github.com/user-attachments/assets/d976d4dc-281b-4bcf-a44d-39d2515ddda5)

Program ini berfungsi untuk menentukan letak sebuah titik sembarang terhadap dua lingkaran. Input yang diperlukan adalah koordinat pusat dan jari-jari kedua lingkaran, serta koordinat titik yang akan diperiksa. Program akan menghitung jarak titik tersebut ke pusat masing-masing lingkaran dengan rumus Euclidean, lalu membandingkannya dengan jari-jari lingkaran. Hasil perbandingan tersebut akan digunakan untuk menentukan apakah titik berada di dalam lingkaran pertama, lingkaran kedua, kedua lingkaran, atau justru di luar keduanya. Hasil akhirnya kemudian ditampilkan.

### 2. 

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// Input jumlah elemen array
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	// Input elemen array
	array := make([]int, n)
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	// a. Menampilkan keseluruhan isi array
	fmt.Println("a. Keseluruhan isi array:", array)

	// b. Menampilkan elemen-elemen array dengan indeks ganjil saja
	fmt.Print("b. Elemen dengan indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	// c. Menampilkan elemen-elemen array dengan indeks genap saja
	fmt.Print("c. Elemen dengan indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	// d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x
	var x int
	fmt.Print("Masukkan nilai x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	fmt.Print("d. Elemen dengan indeks kelipatan", x, ": ")
	for i := x; i < n; i += x {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	// e. Menghapus elemen array pada indeks tertentu
	var delIndex int
	fmt.Print("Masukkan indeks elemen yang akan dihapus: ")
	fmt.Scan(&delIndex)
	if delIndex >= 0 && delIndex < n {
		array = append(array[:delIndex], array[delIndex+1:]...)
		fmt.Println("e. Array setelah penghapusan elemen:", array)
	} else {
		fmt.Println("Indeks tidak valid!")
	}

	// f. Menampilkan rata-rata dari elemen array
	sum := 0
	for _, val := range array {
		sum += val
	}
	avg := float64(sum) / float64(len(array))
	fmt.Printf("f. Rata-rata elemen array: %.2f\n", avg)

	// g. Menampilkan standar deviasi dari elemen array
	var variance float64
	for _, val := range array {
		variance += math.Pow(float64(val)-avg, 2)
	}
	stdDev := math.Sqrt(variance / float64(len(array)))
	fmt.Printf("g. Standar deviasi elemen array: %.2f\n", stdDev)

	// h. Menampilkan frekuensi dari setiap bilangan dalam array
	freq := make(map[int]int)
	for _, val := range array {
		freq[val]++
	}
	fmt.Println("h. Frekuensi elemen dalam array:")
	for key, val := range freq {
		fmt.Printf("   %d: %d kali\n", key, val)
	}
}
```
## Output: ![image](https://github.com/user-attachments/assets/3c92f1e8-00e6-411f-a051-6fccdcd50f41)

Program ini mengolah array dengan beberapa operasi, seperti menampilkan seluruh elemen, elemen dengan indeks ganjil/genap, elemen dengan indeks kelipatan tertentu, menghapus elemen, menghitung rata-rata dan standar deviasi, serta menghitung frekuensi tiap elemen dalam array. Pengguna bisa memasukkan nilai array dan indeks yang ingin dimanipulasi.







