<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br> 

<h2 align="center"><strong>MODUL 7</strong></h2>
<h2 align="center"><strong> STRUCT DAN ARRAY </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Fatik Nurimamah / 2311102190<br>
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
4. [Daftar Pustaka](#daftar-pustaka)


## Dasar Teori
**ARRAY**

Array adalah kumpulan data bertipe sama, yang disimpan dalam sebuah variabel. Array memiliki kapasitas yang nilainya ditentukan pada saat pembuatan, menjadikan elemen/data yang disimpan di array tersebut jumlahnya tidak boleh melebihi yang sudah dialokasikan.

Default nilai tiap elemen array pada awalnya tergantung dari tipe datanya. Jika int maka tiap element zero value-nya adalah `0`, jika `bool` maka `false`, dan seterusnya. Setiap elemen array memiliki indeks berupa angka yang merepresentasikan posisi urutan elemen tersebut. Indeks array dimulai dari 0.

Contoh penerapan array:
```go
var names [4]string
names[0] = "trafalgar"
names[1] = "d"
names[2] = "water"
names[3] = "law"

fmt.Println(names[0], names[1], names[2], names[3])

```

Variabel `names` dideklarasikan sebagai `array string` dengan alokasi kapasitas elemen adalah 4 slot. Cara mengisi slot elemen array bisa dilihat di kode di atas, yaitu dengan langsung mengakses elemen menggunakan indeks, lalu mengisinya.

**1. Slice**

Slice adalah reference elemen array. Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya. Karena slice merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama.

**Inisialisasi Slice**

Cara pembuatan slice mirip seperti pembuatan array, bedanya tidak perlu mendefinisikan jumlah elemen ketika awal deklarasi. Pengaksesan nilai elemen-nya juga sama. Contoh pembuatan slice:
```go
var fruits = []string{"apple", "grape", "banana", "melon"}
fmt.Println(fruits[0]) // "apple"
```

- Fungsi `len()`
  
Fungsi `len()` digunakan untuk menghitung jumlah elemen slice yang ada. Sebagai contoh jika sebuah variabel adalah slice dengan data 4 buah, maka fungsi ini pada variabel tersebut akan mengembalikan angka 4.
```go
var fruits = []string{"apple", "grape", "banana", "melon"}
fmt.Println(len(fruits)) // 4
```
- Fungsi `cap()`
  
Fungsi `cap()` digunakan untuk menghitung lebar atau kapasitas maksimum slice. Nilai kembalian fungsi ini untuk slice yang baru dibuat pasti sama dengan len, tapi bisa berubah seiring operasi slice yang dilakukan.
```go
var fruits = []string{"apple", "grape", "banana", "melon"}
fmt.Println(len(fruits))  // len: 4
fmt.Println(cap(fruits))  // cap: 4

var aFruits = fruits[0:3]
fmt.Println(len(aFruits)) // len: 3
fmt.Println(cap(aFruits)) // cap: 4

var bFruits = fruits[1:4]
fmt.Println(len(bFruits)) // len: 3
fmt.Println(cap(bFruits)) // cap: 3
```
- Fungsi `append()`
  
Fungsi `append()` digunakan untuk menambahkan elemen pada slice. Elemen baru tersebut diposisikan setelah indeks paling akhir. Nilai balik fungsi ini adalah slice yang sudah ditambahkan nilai barunya.
```go
var fruits = []string{"apple", "grape", "banana"}
var cFruits = append(fruits, "papaya")

fmt.Println(fruits)  // ["apple", "grape", "banana"]
fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]
```
- Fungsi `copy()`
  
Fungsi `copy()` digunakan untuk men-copy elements slice pada src (parameter ke-2), ke dst (parameter pertama).
```go
copy(dst, src)
```

**2. Map**

Map adalah tipe data asosiatif dalam Go yang menyimpan pasangan key-value. Setiap data (value) di dalam map selalu memiliki key yang menyertainya. Key ini harus unik karena berfungsi sebagai identifier untuk mengakses value yang tersimpan. Berbeda dengan slice yang menggunakan indeks numerik sebagai identifier, map memungkinkan key memiliki tipe data yang dapat disesuaikan sesuai kebutuhan.

**Penggunaan Map**

Penggunaan map sangat sederhana, cukup dengan menuliskan keyword `map` diikuti oleh tipe data untuk key dan value-nya.
```go
var chicken map[string]int
chicken = map[string]int{}

chicken["januari"] = 50
chicken["februari"] = 40

fmt.Println("januari", chicken["januari"]) // januari 50
fmt.Println("mei",     chicken["mei"])     // mei 0
```

**Inisialisasi Map**

Inisialisasi map dilakukan dengan menambahkan tanda kurung kurawal `{}` setelah penulisan tipe datanya, misalnya `map[string]int{}`. Zero value dari map adalah `nil`, sehingga disarankan untuk menginisialisasinya secara eksplisit agar tidak bernilai `nil`. Nilai variabel bertipe map dapat didefinisikan sejak awal dengan menambahkan kurung kurawal setelah tipe data, lalu menuliskan pasangan key-value di dalamnya. Cara ini mirip dengan inisialisasi array atau slice, namun dalam format key-value. Jika inisialisasi menggunakan keyword `new`, yang dihasilkan adalah pointer ke map. Untuk mengakses nilai aslinya, gunakan tanda asterisk (`*`).

```go
var data map[string]int
data["one"] = 1
// akan muncul error!

data = map[string]int{}
data["one"] = 1
// tidak ada error
```

```go
var chicken3 = map[string]int{}
var chicken4 = make(map[string]int)
var chicken5 = *new(map[string]int)
```

**STRUCT**

Struct adalah kumpulan variabel (atau properti) dan/atau fungsi (atau metode) yang dikelompokkan menjadi tipe data baru dengan nama tertentu. Properti dalam struct dapat memiliki tipe data yang beragam. Struct mirip dengan map, tetapi key-nya telah ditentukan sejak awal, dan setiap itemnya dapat memiliki tipe data yang berbeda. Dari sebuah struct, kita dapat membuat variabel baru yang memiliki atribut sesuai dengan struktur tersebut. Dalam pembahasan ini, variabel tersebut akan disebut sebagai objek atau variabel objek.

**Deklarasi Struct**

Kombinasi keyword `type` dan `struct` digunakan untuk deklarasi struct.
```go
type student struct {
    name string
    grade int
}
``` 

## Guided 

### 1. Program sederhana untuk menghitung lama waktu parkir berdasarkan waktu kedatangan dan waktu pulang
### Source Code :
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

### Output:
![image](https://github.com/user-attachments/assets/cd256395-fcbe-4840-a8c8-34112e688979)

### Full code Screenshot:
![image](https://github.com/user-attachments/assets/68227973-95dd-4f4d-bedd-f37dcf91a8ef)

### Deskripsi Program : 
Program tersebut adalah program yang digunakan untuk menghitung durasi parkir kendaraan. Program ini menggunakan struktur data `waktu` yang berisi atribut `jam`, `menit`, dan `detik`. Pertama, program meminta input untuk waktu kedatangan dan waktu kepulangan kendaraan dalam satuan jam, menit, dan detik. Kedua, waktu kedatangan dan kepulangan dikonversi ke dalam detik untuk memudahkan perhitungan selisihnya sebagai durasi parkir. Setelah itu, durasi parkir dalam detik diubah kembali menjadi jam, menit, dan detik agar hasilnya lebih mudah dipahami, kemudian hasil tersebut ditampilkan ke layar.

### 2. Program sederhana untuk validasi duplikasi nama pada daftar teman

### Source Code :
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

### Output:
![image](https://github.com/user-attachments/assets/189a5a74-8955-41dd-bfff-7a85003a4bad)

### Full code Screenshot:
![image](https://github.com/user-attachments/assets/55d9c939-f418-4663-9c4e-ec3b3b3fa866)

### Deskripsi Program : 
Program tersebut adalah program yang digunakan untuk mengelola daftar teman tanpa duplikasi nama. Program ini diawali dengan slice `daftarTeman` berisi beberapa nama awal, seperti "Andi," "Budi," dan "Cici," serta `namaBaru` yang memuat nama-nama tambahan yang ingin ditambahkan. Fungsi `sudahAda` dibuat untuk memeriksa apakah suatu nama sudah ada di dalam `daftarTeman`, dan mengembalikan `true` jika ada atau `false` jika tidak. Di dalam fungsi `main`, program memeriksa setiap nama di `namaBaru`: jika nama belum ada di `daftarTeman`, nama tersebut ditambahkan menggunakan `append`; namun, jika sudah ada, program menampilkan pesan bahwa nama tersebut sudah terdaftar. Di akhir program, daftar teman yang telah diperbarui dicetak ke layar, dan memastikan tidak ada nama yang terduplikasi.

### 3. Program sederhana untuk menampilkan daftar harga buah menggunakan map.
### Source Code :
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
### Output:
![image](https://github.com/user-attachments/assets/9b5574b8-6bf9-40aa-8e4e-9c0db34ece17)

### Full code Screenshot:
![image](https://github.com/user-attachments/assets/976f9ee5-6f16-4dae-9bb3-a248639e299a)

### Deskripsi Program : 
Program tersebut menggunakan `map` dalam bahasa Go untuk menyimpan data harga buah-buahan. Struktur `map` dengan nama `hargaBuah` memiliki kunci berupa nama buah (`string`) dan nilai berupa harga buah (`int`). `hargaBuah` diinisialisasi dengan beberapa data, seperti "Apel" dengan harga Rp5000, "Pisang" Rp3000, dan "Mangga" Rp7000. Program kemudian menggunakan loop `for` dengan `range` untuk menampilkan setiap buah beserta harganya. Setelah itu, harga khusus untuk "Mangga" ditampilkan secara langsung dengan mengakses kuncinya di `map`. Program ini menunjukkan bagaimana `map` memudahkan penyimpanan dan akses cepat terhadap data harga buah-buahan.




## Unguided 

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥,𝑐𝑦) dengan radius 𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥,𝑦) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.

Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
### Source Code :
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
	titikPusat Titik
	radius int
}

// Fungsi untuk menghitung jarak antara dua titik
func jarakAntarTitik(a, b Titik) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2))
}

// Fungsi untuk mengecek apakah titik berada di dalam lingkaran
func cekTitikDiDalam(lingkaran Lingkaran, titik Titik) bool {
	return jarakAntarTitik(lingkaran.titikPusat, titik) <= float64(lingkaran.radius)
}

func main() {
	var lingkaran1, lingkaran2 Lingkaran
	var titikSembarang Titik

	// Input titik pusat dan radius lingkaran 1
	fmt.Print("Masukkan titik pusat dan radius lingkaran 1: ")
	fmt.Scan(&lingkaran1.titikPusat.x, &lingkaran1.titikPusat.y, &lingkaran1.radius)


	// Input titik pusat dan radius lingkaran 2
	fmt.Print("Masukkan titik pusat dan radius lingkaran 2: ")
	fmt.Scan(&lingkaran2.titikPusat.x, &lingkaran2.titikPusat.y, &lingkaran2.radius)

	// Input titik sembarang
	fmt.Print("Masukkan titik sembarang: ")
	fmt.Scan(&titikSembarang.x, &titikSembarang.y)

	// Mengecek posisi titik terhadap lingkaran 1 dan 2
	diDalamLingkaran1 := cekTitikDiDalam(lingkaran1, titikSembarang)
	diDalamLingkaran2 := cekTitikDiDalam(lingkaran2, titikSembarang)

	// Menentukan output berdasarkan posisi titik
	if diDalamLingkaran1 && diDalamLingkaran2 {
		fmt.Println("\nTitik di dalam lingkaran 1 dan 2")
	} else if diDalamLingkaran1 {
		fmt.Println("\nTitik di dalam lingkaran 1")
	} else if diDalamLingkaran2 {
		fmt.Println("\nTitik di dalam lingkaran 2")
	} else {
		fmt.Println("\nTitik di luar lingkaran 1 dan 2")
	}
}

```
### Output:
![image](https://github.com/user-attachments/assets/9907fef2-147c-4c1d-b16f-d28e21c6fe23)
![image](https://github.com/user-attachments/assets/2c236f32-0c41-42c5-9581-b4b3ecd0ff35)
![image](https://github.com/user-attachments/assets/bede38f8-4185-447b-950b-eb45e64b874b)

### Full code Screenshot:
![code](https://github.com/user-attachments/assets/cf4a5aee-23fd-4ddf-81d9-7f16d7223dbe)

### Deskripsi Program : 
Program tersebut adalah program yang digunakan untuk menentukan apakah sebuah titik berada di dalam salah satu atau kedua lingkaran. Lingkaran didefinisikan sebagai struct dengan properti berupa titik pusat dan radius, sementara titik direpresentasikan sebagai struct dengan koordinat x dan y. Program meminta input berupa koordinat titik pusat dan radius untuk dua lingkaran, serta koordinat sebuah titik sembarang. Fungsi `jarakAntarTitik` menghitung jarak antara dua titik, yang kemudian digunakan oleh fungsi `cekTitikDiDalam` untuk mengecek apakah titik tersebut berada di dalam lingkaran. Berdasarkan hasil pengecekan, program menampilkan output apakah titik berada di dalam lingkaran pertama, kedua, keduanya, atau di luar kedua lingkaran.


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
package main

import (
	"fmt"
	"math"
)

func main() {
	var jumlahElemen, kelipatan, indeksHapus int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&jumlahElemen)

	daftarBilangan := make([]int, jumlahElemen)
	for i := 0; i < jumlahElemen; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&daftarBilangan[i])
	}

	fmt.Println("\nPilihan operasi:")
	fmt.Println("a. Tampilkan keseluruhan isi array")
	fmt.Println("b. Tampilkan elemen dengan indeks ganjil")
	fmt.Println("c. Tampilkan elemen dengan indeks genap")
	fmt.Println("d. Tampilkan elemen dengan indeks kelipatan tertentu")
	fmt.Println("e. Hapus elemen pada indeks tertentu")
	fmt.Println("f. Tampilkan rata-rata elemen")
	fmt.Println("g. Tampilkan standar deviasi elemen")
	fmt.Println("h. Tampilkan frekuensi elemen tertentu")

	var pilihan string
	fmt.Print("\nPilih operasi: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case "a":
		fmt.Println("Isi array:", daftarBilangan)
	case "b":
		fmt.Println("Elemen dengan nilai ganjil:")
		for i := 0; i < jumlahElemen; i++ {
			if daftarBilangan[i]%2 != 0 { // Cek jika elemen adalah bilangan ganjil
				fmt.Print(daftarBilangan[i], " ")
			}
		}
		fmt.Println()
	case "c":
		fmt.Println("Elemen dengan nilai genap:")
		for i := 1; i < jumlahElemen; i += 2 { 
			fmt.Print(daftarBilangan[i], " ")
		}
		fmt.Println()
	case "d":
		fmt.Print("Masukkan nilai kelipatan: ")
		fmt.Scan(&kelipatan)
		fmt.Printf("Elemen dengan kelipatan indeks %d:\n", kelipatan)
		for i := 0; i < jumlahElemen; i++ {
			if daftarBilangan[i] % kelipatan == 0 { // Memeriksa apakah nilai elemen adalah kelipatan dari angka yang dimasukkan
				fmt.Print(daftarBilangan[i], " ")
			}
		}
		fmt.Println()	
	case "e":
		fmt.Print("Masukkan indeks yang akan dihapus: ")
		fmt.Scan(&indeksHapus)
		if indeksHapus >= 0 && indeksHapus < jumlahElemen {
			daftarBilangan = append(daftarBilangan[:indeksHapus], daftarBilangan[indeksHapus+1:]...)
			fmt.Println("Array setelah penghapusan:", daftarBilangan)
		} else {
			fmt.Println("Indeks tidak valid.")
		}
	case "f":
		total := 0
		for _, bilangan := range daftarBilangan {
			total += bilangan
		}
		rataRata := float64(total) / float64(jumlahElemen)
		fmt.Println("Rata-rata:", rataRata)
	case "g":
		total := 0
		for _, bilangan := range daftarBilangan {
			total += bilangan
		}
		rataRata := float64(total) / float64(jumlahElemen)
		var jumlahVarian float64
		for _, bilangan := range daftarBilangan {
			jumlahVarian += math.Pow(float64(bilangan)-rataRata, 2)
		}
		standarDeviasi := math.Sqrt(jumlahVarian / float64(jumlahElemen))
		fmt.Println("Standar deviasi:", standarDeviasi)
	case "h":
		fmt.Print("Masukkan elemen yang ingin dicari frekuensinya: ")
		fmt.Scan(&kelipatan)
		frekuensi := 0
		for _, bilangan := range daftarBilangan {
			if bilangan == kelipatan {
				frekuensi++
			}
		}
		fmt.Printf("Frekuensi %d: %d\n", kelipatan, frekuensi)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

```
### Output:
![image](https://github.com/user-attachments/assets/540505ad-fb8b-4f13-ba04-e0e8b032c4a3)
![image](https://github.com/user-attachments/assets/637776b3-c893-43dc-bb9c-6419f047e37b)
![image](https://github.com/user-attachments/assets/007bd5d1-69c4-4f13-a440-badf4fde80ec)
![image](https://github.com/user-attachments/assets/0eafce78-7957-4e9c-bcca-85991a576c02)
![image](https://github.com/user-attachments/assets/c6346911-3bd7-463a-a956-5b9768f46073)
![image](https://github.com/user-attachments/assets/be020cb3-3b09-4993-b698-36a22e612e05)
![image](https://github.com/user-attachments/assets/cbacd229-6613-4067-90e5-e185bb6a1a99)
![image](https://github.com/user-attachments/assets/716f9b2e-1012-49ab-9d46-4426c079c4b8)

### Full code Screenshot:
![code](https://github.com/user-attachments/assets/52b620b1-e1fb-4e90-b8dc-87271317e3ba)

### Deskripsi Program : 
Program tersebut dibuat untuk mengelola array berisi bilangan bulat dan memiliki beberapa operasi yang dapat dilakukan. Pertama, pengguna akan diminta untuk mengisi jumlah elemen dalam array serta memasukkan setiap angkanya secara satu per satu. Setelah data terisi, pengguna bisa memilih beberapa opsi seperti menampilkan semua isi array, menampilkan elemen dengan indeks ganjil atau genap, dan menampilkan elemen pada indeks tertentu yang merupakan kelipatan dari angka yang ditentukan. Selain itu, program juga memiliki operasi untuk menghapus elemen pada indeks tertentu, menghitung rata-rata dan standar deviasi dari seluruh elemen, serta mencari seberapa sering suatu angka muncul dalam array. Program ini dirancang dengan antarmuka yang sederhana, sehingga pengguna dapat dengan mudah memilih opsi dan langsung melihat hasilnya.

  
### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.
### Source Code :
```go
package main

import (
	"fmt"
)

func main() {
	var KlubA, KlubB string
	fmt.Print("Masukkan nama Tim A: ")
	fmt.Scanln(&KlubA)
	fmt.Print("Masukkan nama Tim B: ")
	fmt.Scanln(&KlubB)

	rekapHasilPertandingan := []string{}
	jumlahPertandingan := 1

	for {
		var skorKlubA, skorKlubB int
		fmt.Printf("\nPertandingan %d - Masukkan skor %s: ", jumlahPertandingan, KlubA)
		fmt.Scan(&skorKlubA)
		fmt.Printf("Pertandingan %d - Masukkan skor %s: ", jumlahPertandingan, KlubB)
		fmt.Scan(&skorKlubB)

		// Cek apakah skor valid (tidak negatif)
		if skorKlubA < 0 || skorKlubB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		// Menentukan hasil pertandingan
		if skorKlubA > skorKlubB {
			fmt.Printf("\nHasil %d: %s menang\n", jumlahPertandingan, KlubA)
			rekapHasilPertandingan = append(rekapHasilPertandingan, KlubA)
		} else if skorKlubB > skorKlubA {
			fmt.Printf("\nHasil %d: %s menang\n", jumlahPertandingan, KlubB)
			rekapHasilPertandingan = append(rekapHasilPertandingan, KlubB)
		} else {
			fmt.Printf("\nHasil %d: Draw\n", jumlahPertandingan)
			rekapHasilPertandingan = append(rekapHasilPertandingan, "Draw")
		}

		jumlahPertandingan++
	}

	fmt.Println("\nDaftar tim yang memenangkan pertandingan:")
	for i, hasil := range rekapHasilPertandingan {
    	fmt.Printf("Hasil %d: %s\n", i+1, hasil)
	}

}

```
### Output:
![image](https://github.com/user-attachments/assets/411fa03a-ed70-4939-9f1a-e8e3de481c35)
![image](https://github.com/user-attachments/assets/8a3db6f4-2bc3-4ae9-bd6a-4c2c589549dc)

### Full code Screenshot:
![code](https://github.com/user-attachments/assets/6a4c881c-119a-4c4b-98f6-8bcd8433b521)

### Deskripsi Program : 
Program tersebut digunakan untuk mencatat hasil pertandingan antara dua tim yang dimasukkan oleh pengguna. Pengguna diminta untuk memasukkan nama tim A dan tim B, lalu memasukkan skor untuk setiap pertandingan. Jika skor negatif dimasukkan, program akan berhenti dan menampilkan pesan "Pertandingan selesai". Berdasarkan skor yang dimasukkan, program menentukan pemenang atau jika pertandingan berakhir seri. Setelah semua pertandingan selesai, program menampilkan daftar tim yang menang di setiap pertandingan, termasuk jika ada hasil seri. Program ini menggunakan array untuk menyimpan rekap hasil pertandingan.


### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom
### Source Code :
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

const NMAX int = 127

type tabel [NMAX]rune

// Fungsi untuk mengisi array
func isiArray(t *tabel, n *int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan teks: ")

	var char rune
	*n = 0
	for *n < NMAX {
		input, _, _ := reader.ReadRune()
		char = input
		if char == '.' {
			break
		}
		t[*n] = char
		*n++
	}
}

// Fungsi untuk mencetak array
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// Fungsi untuk membalik array
func balikanArray(t *tabel, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}

// Fungsi untuk mengecek apakah array adalah palindrome
func palindrome(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	// Mengisi array dengan memanggil fungsi isiArray
	isiArray(&tab, &m)

	// Mencetak array asli
	fmt.Print("\nTeks: ")
	cetakArray(tab, m)

	// Membalik array
	balikanArray(&tab, m)

	// Mencetak array yang telah dibalik
	fmt.Print("Reverse teks: ")
	cetakArray(tab, m)

	// Mengecek apakah array adalah palindrome dan mencetak hasilnya
	isPalindrome := palindrome(tab, m)
	fmt.Printf("Palindrome? %t\n", isPalindrome)
}
		
```
### Output:
![image](https://github.com/user-attachments/assets/76cdcaef-dd9d-4e5a-98b8-25abd60d955d)

### Full code Screenshot:
![code](https://github.com/user-attachments/assets/970bf7b7-fccd-4d40-9ec0-4733264945b0)

### Deskripsi Program : 
Program tersebut digunakan untuk menerima input teks dari pengguna, kemudian mengecek apakah teks tersebut merupakan palindrome. Program ini juga menampilkan teks asli dan versi terbaliknya. Pada fungsi pertama, teks yang dimasukkan pengguna disimpan dalam array hingga tanda titik (`.`) ditemukan sebagai penanda akhir input. Fungsi kedua digunakan untuk menampilkan teks dalam array, baik dalam urutan aslinya maupun setelah dibalik. Fungsi selanjutnya membalikkan urutan karakter dalam array dan menghasilkan teks terbalik. Terakhir, program memeriksa apakah teks yang dimasukkan sama ketika dibaca dari depan dan belakang, lalu menampilkan hasilnya dengan teks asli dan yang terbalik.

## Daftar Pustaka
[1] Agung, N. (n.d.). Array. Retrieved from Dasar Pemrograman Golang: https://dasarpemrogramangolang.novalagung.com/A-array.html
