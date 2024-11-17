
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
  <strong>STRUCT & ARRAY</strong>
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
  Mutia Rani Zahra Meilani
  <br>
  2311102182
  <br>
  S1IF1105
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

## <strong> DASAR TEORI </strong>

<span style="font-size:16px"><strong> ── PENGERTIAN ARRAY</strong></span>
<br>
Array dalam bahasa pemrograman Go (Golang) adalah sekumpulan elemen yang memiliki tipe data yang sama dan disimpan dalam urutan yang berurutan. Setiap array memiliki ukuran tetap yang ditentukan saat deklarasi, dan ukuran ini tidak dapat diubah setelahnya. Array sangat berguna untuk menyimpan data dalam jumlah yang banyak dan memudahkan akses serta pengelolaan data tersebut

**Karakteristik Array :**
- Ukuran Tetap: Ukuran array ditentukan saat deklarasi dan tidak dapat diubah.
- Indeks: Indeks array dimulai dari 0, sehingga elemen pertama berada pada indeks 0, elemen kedua pada indeks 1, dan seterusnya.
- Tipe Data Sama: Semua elemen dalam array harus memiliki tipe data yang sama.

Deklarasi :
```go
var numbers [5]int // Mendeklarasikan array dengan 5 elemen bertipe int
```
Inisialisasi :
```go
var fruits = [3]string{"Apple", "Banana", "Cherry"} // Inisialisasi saat deklarasi
```

<span style="font-size:16px"><strong> ── PENGERTIAN STRUCT</strong></span>
<br>
Struct adalah tipe data kompleks yang digunakan untuk mengelompokkan beberapa variabel dengan tipe data yang berbeda menjadi satu kesatuan. Struct sangat berguna untuk merepresentasikan objek dunia nyata dengan atribut yang berbeda-beda. Dalam Golang, struct dideklarasikan dengan kata kunci type diikuti oleh nama struct dan definisi dari field-field yang ada di dalamnya

Deklarasi :
```go
type Person struct {
    Name string
    Age  int
}
```
Inisialisasi :
```go
john := Person{Name: "John", Age: 30} // Inisialisasi struct dengan nilai
```

<span style="font-size:16px"><strong> ── PENGGUNAAN STRUCT DAN ARRAY</strong></span>
<br>
Struct dan array dapat digunakan bersama-sama untuk menyimpan koleksi objek. Misalnya, jika kita ingin menyimpan daftar orang, kita bisa menggunakan array dari struct Person.

Contoh Penggunaan:
```go
var people [2]Person

people[0] = Person{Name: "Alice", Age: 25}
people[1] = Person{Name: "Bob", Age: 30}

for _, person := range people {
    fmt.Println(person.Name, person.Age)
}
```
<span style="font-size:16px"><strong> ── KELEBIHAN DAN KEKURANGAN</strong></span>
<br>

**Kelebihan Array:**
- Akses cepat ke elemen berdasarkan indeks.
- Efisiensi memori karena ukuran tetap.

**Kekurangan Array:**
- Ukuran tetap membuatnya kurang fleksibel dibandingkan struktur data lain seperti slice.

**Kelebihan Struct:**
- Memungkinkan pengelompokan data yang beragam menjadi satu kesatuan logis.

**Kekurangan Struct:**
- Tidak mendukung pewarisan (inheritance) seperti pada OOP tradisional.


## <strong> GUIDED </strong>

### ── Guided 1

#### Source Code

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
	fmt.Printf("Lama parkir: %d jam %d menit %d detik",
		durasi.jam, durasi.menit, durasi.detik)
}

```

#### Output
![image](https://github.com/user-attachments/assets/821a30c2-dc8b-4639-a7ed-635e82456ca0)



#### Deskripsi Program :
Program ini digunakan untuk menghitung lama waktu parkir berdasarkan waktu masuk dan waktu keluar yang diinput user. Program menggunakan struktur waktu yang terdiri dari jam, menit, dan detik. Program akan meminta user memasukkan jam, menit, dan detik untuk waktu masuk dan keluar parkir, kemudian menghitung durasi parkir dengan mengkonversi semua waktu ke dalam detik untuk mempermudah perhitungan selisih waktu. Hasil perhitungan akan dikonversi kembali menjadi format jam, menit, dan detik yang lebih mudah dibaca, lalu ditampilkan ke layar.

### ── Guided 2

#### Source Code

```go
package main

import (
	"fmt"
)

func sudahAda(daftarTeman []string, nama string) bool {
	for _, teman := range daftarTeman {
		if teman == nama {
			return true
		}
	}
	return false
}

func main() {
	daftarTeman := []string{"Andi", "Budi", "Cici"}

	namaBaru := []string{"Dewi", "Budi", "Eka"}

	for _, nama := range namaBaru {
		if !sudahAda(daftarTeman, nama) {
			daftarTeman = append(daftarTeman, nama)
		} else {
			fmt.Println("Nama", nama, "sudah ada dalam daftar.")
		}
	}

	fmt.Println("Daftar Teman:", daftarTeman)
}

```

#### Output
![image](https://github.com/user-attachments/assets/2fab5fa0-5ad8-46df-8f38-e50cdea285c9)

#### Deskripsi Program :
Program ini digunakan untuk mengelola daftar nama teman dengan mencegah duplikasi nama dalam daftar. Program memiliki daftar nama awal yang terdiri dari "Andi", "Budi", dan "Cici", kemudian akan mencoba menambahkan nama-nama baru ke dalam daftar tersebut. Sebelum menambahkan nama baru, program akan menggunakan fungsi sudahAda untuk memeriksa apakah nama tersebut sudah ada dalam daftar. Jika nama belum ada, maka akan ditambahkan ke daftar, namun jika sudah ada maka program akan menampilkan pesan bahwa nama tersebut sudah terdaftar sebelumnya.

### ── Guided 3

#### Source Code

```go
package main

import (
	"fmt"
)

func main() {
	hargaBuah := map[string]int{
		"Apel":   5000,
		"Pisang": 3000,
		"Mangga": 7000,
	}

	fmt.Println("Harga Buah:")
	for buah, harga := range hargaBuah {
		fmt.Printf("%s: Rp%d\n", buah, harga)
	}

	fmt.Print("Harga buah Mangga = ", hargaBuah["Mangga"])
}

```

#### Output
![image](https://github.com/user-attachments/assets/21e83d0e-90ac-4bfd-b779-a97c017a24e3)

#### Deskripsi Program :
Program ini digunakan untuk menampilkan daftar harga buah yang tersimpan dalam sebuah map dengan key berupa nama buah dan value berupa harga. Program menyimpan data harga beberapa jenis buah seperti Apel, Pisang, dan Mangga dalam struktur map, kemudian menampilkan seluruh daftar harga buah menggunakan perulangan range untuk mengakses setiap pasangan key-value dalam map. Program juga secara khusus menampilkan harga buah Mangga dengan mengakses nilai map menggunakan key "Mangga".

## <strong>  UNGUIDED </strong>

### ── Unguided 1

#### Study Case
Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.

Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.

Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".


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
    pusat Titik
    radius int
}

func jarak(p, q Titik) float64 {
    return math.Sqrt(float64((p.x - q.x)*(p.x - q.x) + (p.y - q.y) * (p.y - q.y)))
}

func inside(c Lingkaran, p Titik) bool {
    return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {
    var l1, l2 Lingkaran
    var p Titik

    fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
    fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
    fmt.Scan(&p.x, &p.y)

    insideL1 := inside(l1, p)
    insideL2 := inside(l2, p)

    if insideL1 && insideL2 {
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    } else if insideL1 {
        fmt.Println("Titik di dalam lingkaran 1")
    } else if insideL2 {
        fmt.Println("Titik di dalam lingkaran 2")
    } else {
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}

```

#### Output
![image](https://github.com/user-attachments/assets/b76c4548-dd96-4e37-82bf-4d292b56af43)

#### Deskripsi Program :
Program ini digunakan untuk menentukan posisi sebuah titik terhadap dua buah lingkaran dalam bidang koordinat. Program menggunakan struktur Titik untuk menyimpan koordinat x dan y, serta struktur Lingkaran yang terdiri dari pusat dan radius. Program meminta input koordinat pusat dan radius untuk dua lingkaran serta koordinat sebuah titik yang akan dicek posisinya. Menggunakan fungsi jarak yang menghitung jarak Euclidean antara dua titik, program menentukan apakah titik tersebut berada di dalam lingkaran 1, lingkaran 2, kedua lingkaran, atau di luar kedua lingkaran.

### ── Unguided 2

#### Study Case

Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut :
- Menampilkan keseluruhan isi dari array.
- Menampilkan elemen-elemen array dengan indeks ganjil saja.
- Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indeks ke-0 adalah genap).
- Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna.
- Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil.
- Menampilkan rata-rata dari bilangan yang ada di dalam array.
- Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.
- Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.

#### Source Code

```go
package main

import (
    "fmt"
    "math"
)

type Array []int

func show_array(arr []int) {
    fmt.Println("ARRAY :", arr)
}

func index_ganjil(arr []int) {
    fmt.Print("Indeks Ganjil : ")
    for i := 1; i < len(arr); i += 2 {
        fmt.Print(arr[i], " ")
    }
    fmt.Println()
}

func index_genap(arr []int) {
    fmt.Print("Indeks Genap : ")
    for i := 0; i < len(arr); i += 2 {
        fmt.Print(arr[i], " ")
    }
    fmt.Println()
}

func kelipatan_index(arr []int, x int) {
    fmt.Printf("Indeks Kelipatan %d : ", x)
    for i := 0; i < len(arr); i += x {
        fmt.Print(arr[i], " ")
    }
    fmt.Println()
}

func hapus_index(arr []int, index int) []int {
    if index < 0 || index >= len(arr) {
        fmt.Println("Indeks Tidak Valid")
        return arr
    }
    fmt.Printf("Elemen di Indeks %d (nilai : %d) dihapus!\n", index, arr[index])
    return append(arr[:index], arr[index+1:]...)
}

func rata_rata(arr []int) float64 {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    return float64(sum) / float64(len(arr))
}

func deviasi(arr []int) float64 {
    mean := rata_rata(arr)
    var sum float64
    for _, v := range arr {
        sum += math.Pow(float64(v)-mean, 2)
    }
    return math.Sqrt(sum / float64(len(arr)))
}

func frekuensi(arr []int, nilai int) int {
    count := 0
    for _, v := range arr {
        if v == nilai {
            count++
        }
    }
    return count
}

func main() {
    var n, x, index, nilai int

    fmt.Print("Masukkan jumlah elemen array : ")
    fmt.Scan(&n)

    arr := make(Array, n)
    fmt.Print("Masukkan elemen array : ")
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    show_array(arr)
    index_ganjil(arr)
    index_genap(arr)

    fmt.Print("\nMasukkan nilai x untuk kelipatan indeks : ")
    fmt.Scan(&x)
    kelipatan_index(arr, x)

    fmt.Print("\nMasukkan indeks untuk menghapus elemen : ")
    fmt.Scan(&index)
    arr = hapus_index(arr, index)
    show_array(arr)

    fmt.Printf("\nRATA RATA : %.2f\n", rata_rata(arr))

    fmt.Printf("DEVIASI : %.2f\n", deviasi(arr))

    fmt.Print("\nMasukkan nilai untuk menghitung frekuensi : ")
    fmt.Scan(&nilai)
    fmt.Printf("FREKUENSI %d : %d\n", nilai, frekuensi(arr, nilai))
}
```

#### Output
![image](https://github.com/user-attachments/assets/ddd666f1-b454-44eb-97ca-fc4b9d79a911)

#### Deskripsi Program :
Program ini digunakan untuk melakukan berbagai operasi analisis dan manipulasi pada array bilangan bulat. Program meminta user memasukkan jumlah elemen array dan nilai-nilainya, kemudian menyediakan berbagai fungsi untuk menganalisis array tersebut. Fungsi-fungsi yang tersedia meliputi menampilkan elemen pada indeks ganjil/genap, menampilkan elemen pada indeks kelipatan tertentu, menghapus elemen pada indeks tertentu, menghitung rata-rata array, menghitung deviasi standar, dan menghitung frekuensi kemunculan nilai tertentu dalam array.

### ── Unguided 3

#### Study Case

Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga.

Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja.

Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.

Perhatikan sesi interaksi pada contoh berikut ini (teks bergaris bawah adalah Input/read).

#### Source Code

```go
package main

import (
	"fmt"
)

type Klub struct {
	klubA, klubB string
}

func input_klub() Klub {
	var klub Klub
	fmt.Print("Klub A : ")
	fmt.Scan(&klub.klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klub.klubB)
	return klub
}

func input_skor(pertandingan int) (int, int) {
	var skorA, skorB int
	fmt.Printf("Pertandingan %d (Skor A, Skor B) : ", pertandingan)
	fmt.Scan(&skorA, &skorB)
	return skorA, skorB
}

func hasil_pemenang(skorA, skorB int, klub Klub) string {
	if skorA > skorB {
		return klub.klubA
	} else if skorB > skorA {
		return klub.klubB
	}
	return "Draw"
}

func hasil_pertandingan(pemenang []string) {
	fmt.Println("\nHasil pertandingan :")
	for i, winner := range pemenang {
		fmt.Printf("Hasil %d : %s\n", i+1, winner)
	}
	fmt.Println("Pertandingan selesai.")
}

func main() {
	// Input nama klub
	klub := input_klub()

	var pemenang []string
	pertandingan := 1

	for {
		skorA, skorB := input_skor(pertandingan)

		if skorA < 0 || skorB < 0 {
			break
		}

		winner := hasil_pemenang(skorA, skorB, klub)
		pemenang = append(pemenang, winner)

		pertandingan++
	}
	hasil_pertandingan(pemenang)
}

```

#### Output
![image](https://github.com/user-attachments/assets/beeb5a0f-0c45-4906-9c3e-d963f7e1b519)

#### Deskripsi Program :
Program ini digunakan untuk mencatat dan menampilkan hasil pertandingan antara dua klub olahraga secara berurutan. Program menggunakan struktur Klub untuk menyimpan nama kedua klub yang bertanding. Program akan meminta input nama kedua klub dan kemudian meminta skor pertandingan secara berulang hingga dimasukkan skor negatif sebagai kondisi berhenti. Setiap pertandingan akan ditentukan pemenangnya berdasarkan skor, dengan hasil "Draw" jika skornya sama. Di akhir program, akan ditampilkan rekap hasil seluruh pertandingan yang telah dilakukan secara berurutan.
### ── Unguided 4

#### Study Case

Sebuah array digunakan untuk menampung sekumpulan karakter. Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.

Modifikasi program tersebut dengan menambahkan fungsi palindrom. Tambahkan instruksi untuk menangani fungsi tersebut dan menampilkan hasilnya pada program utama.

Palindrom adalah teks yang dibaca dari awal atau akhir adalah sama, contoh: KATAK, APA, KASUR_RUSAK.

#### Source Code

```go
package main

import (
	"fmt"
)

const NMAX int = 127
type tabel [NMAX]rune

func array() (t tabel, n int) {
	fmt.Println("Masukkan karakter (akhiri dengan titik) :")
	for n < NMAX {
		var char rune
		fmt.Printf("Karakter ke-%d : ", n+1)
		fmt.Scanf("%c\n", &char)
		if char == '.' {
			break
		}
		t[n] = char
		n++
	}
	return
}

func cetak_array(t tabel, n int, label string) {
	fmt.Printf("%s: ", label)
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balik_array(t tabel, n int) tabel {
	var reversed tabel
	for i := 0; i < n; i++ {
		reversed[i] = t[n-1-i]
	}
	return reversed
}

func palindrom(t tabel, n int) bool {
	reversed := balik_array(t, n)
	for i := 0; i < n; i++ {
		if t[i] != reversed[i] {
			return false
		}
	}
	return true
}

func main() {
	tab, m := array()
	cetak_array(tab, m, "Original")

	reversed := balik_array(tab, m)
	cetak_array(reversed, m, "Reversed")

	if palindrom(tab, m) {
		fmt.Println("Palindrom  True")
	} else {
		fmt.Println("Palindrom : False")
	}
}

```

#### Output
![image](https://github.com/user-attachments/assets/cc1d431b-0157-4b0a-b34b-0bf85dce2a9e)

#### Deskripsi Program :
Program ini digunakan untuk melakukan operasi pada array karakter dan mengecek apakah array tersebut membentuk palindrom (kata yang dibaca sama dari depan maupun belakang). Program akan meminta input karakter satu per satu dari user hingga dimasukkan karakter titik sebagai penanda akhir input. Program kemudian akan menampilkan array karakter asli, menampilkan array yang sudah dibalik urutannya menggunakan fungsi balik_array, dan mengecek apakah susunan karakternya membentuk palindrom dengan membandingkan array asli dengan array yang sudah dibalik menggunakan fungsi palindrom.

## <strong> REFERENSI </strong>

#### [1] Ruang Developer, "Tipe Data Array - Belajar Golang Dari Dasar". diakses dari https://blog.ruangdeveloper.com/golang-array/
#### [2] GEPCODE, "Golang - Tipe Data Array". diakses dari https://gepcode.com/post/golang-tipe-data-array
#### [3] Rumah Coding, "Array: Konsep, Implementasi, dan Penggunaan". diakses dari https://rumahcoding.co.id/array-konsep-implementasi-dan-penggunaan/
