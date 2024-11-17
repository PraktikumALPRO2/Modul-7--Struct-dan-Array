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
  <strong>STRUCT DAN ARRAY</strong>
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
  Abdee Alamsyah Noer Siyam
  <br>
  2311102247
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


## <strong> Dasar Teori </strong>

<strong><h3> STRUCT</h3></strong>

<strong><h4>PENGERTIAN STRUCT</h4></strong>

Struct adalah tipe data yang memungkinkan pengembang untuk mengelompokkan beberapa variabel dengan tipe yang berbeda dalam satu unit logis. Struct sangat berguna untuk merepresentasikan objek dunia nyata dengan atribut-atributnya. Struct sering digunakan untuk merepresentasikan objek dalam bentuk yang lebih terstruktur, seperti data entitas di aplikasi (contohnya: user, product, atau order).

<strong><h4>DEKLARASI STRUCT</h4></strong>

Struct dideklarasikan menggunakan kata kunci type, diikuti dengan nama struct dan definisi field-nya. Contoh deklarasi struct untuk merepresentasikan mobil adalah sebagai berikut:

```go
type Car struct {
    Brand  string
    Model  string
    Year   int

		// Tambahkan lebih banyak field sesuai kebutuhan
}
```

<strong><h4>INISIALISASI dan AKSES STRUCT</h4></strong>

Struct dapat diinisialisasi dengan cara berikut:

```go
myCar := Car{Brand: "Toyota", Model: "Corolla", Year: 2020}
```

Atau bisa juga menggunakan nilai default tanpa menyebutkan nama field:

```go
anotherCar := Car{"Honda", "Civic", 2021}
```

Field dalam struct diakses menggunakan titik (.). Contohnya:
```go
fmt.Println(myCar.Brand) // Output: Toyota
```

<strong><h4>CONTOH PENGGUNAAN</h4></strong>

Contoh penggunaan struct pada bahasa golang secara langsung adalah sebagai berikut:
```go
package main

import "fmt"

// Definisi struct
type User struct {
    ID       int
    Name     string
    IsActive bool
}

func main() {
    // Inisialisasi struct
    user1 := User{
        ID:       1,
        Name:     "Alice",
        IsActive: true,
    }

    // Menampilkan nilai field
    fmt.Println("ID:", user1.ID)
    fmt.Println("Name:", user1.Name)
    fmt.Println("IsActive:", user1.IsActive)

    // Mengubah nilai field
    user1.IsActive = false
    fmt.Println("Updated IsActive:", user1.IsActive)
}

```

<strong><h2>ARRAY</h2></strong>

<strong><h4>PENGERTIAN ARRAY</h4></strong>

Array adalah kumpulan elemen yang memiliki tipe data yang sama, disimpan dalam urutan yang tetap. Ukuran array ditentukan pada saat deklarasi dan tidak dapat diubah setelahnya. Ini menjadikan array ideal untuk situasi di mana jumlah elemen sudah diketahui dan tetap. Semua elemen di dalam array diakses menggunakan indeks yang dimulai dari 0.

<strong><h4>DEKLARASI DAN INISIALISASI</h4></strong>

Untuk mendeklarasikan array, kita harus menentukan ukuran dan tipe elemen. Contoh deklarasi array integer dengan 5 elemen adalah sebagai berikut:

```go
var numbers [5]int
```

Pengguna juga dapat menginisialisasi array saat deklarasi:

```go
var primes = [5]int{2, 3, 5, 7, 11}
```

Jika ingin agar Go menghitung jumlah elemen secara otomatis, bisa menggunakan sintaksis ...:

```go
var languages = [...]string{"Go", "Python", "Java", "C++"}
```

<strong><h4>AKSES DAN MODIFIKASI ELEMEN</h4></strong>

Elemen dalam array diakses menggunakan indeks yang dimulai dari 0. Misalnya, untuk mengakses elemen pertama dari array primes:

```go
firstPrime := primes[0]
```

<strong><h4>FITUR ARRAY</h4></strong>

1. Array dengan Panjang Tetap

	Panjang array harus ditentukan saat deklarasi dan tidak dapat diubah.

2. Iterasi Elemen Array

	Gunakan perulangan for atau for range untuk iterasi elemen array.

	```go
	for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
	}

	for index, value := range fruits {
			fmt.Printf("Index %d: %s\n", index, value)
	}
	```

3. Array Multidimensi

	Array dapat memiliki lebih dari satu dimensi.

	```go
	matrix := [2][2]int{
    {1, 2},
    {3, 4},
	}
	fmt.Println(matrix)
	fmt.Println("Element [1][0]:", matrix[1][0])
	```

4. Pass by Value 

	Array di Go bersifat pass by value, artinya ketika sebuah array dikirimkan ke fungsi, salinan dari array tersebut yang digunakan.

	```go
	func updateArray(arr [3]int) {
    arr[0] = 100
	}

	func main() {
			nums := [3]int{1, 2, 3}
			updateArray(nums)
			fmt.Println(nums) // Tetap [1, 2, 3]
	}
	```


## <strong> Guided </strong>

### <h2>GUIDED 1</h2>

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
	
	fmt.Printf("Lama Parkir : %d jam %d menit %d detik", durasi.jam,
		durasi.menit, durasi.detik)
}

```

#### Output

![image](https://github.com/user-attachments/assets/3183bd21-b5a4-47b0-99ff-94db64ba412d)


Deskripsi Program : 
Program ini menghitung durasi waktu parkir berdasarkan input waktu masuk dan waktu keluar dalam format jam, menit, dan detik. Struktur data `waktu` digunakan untuk merepresentasikan waktu dengan atribut jam, menit, dan detik. Pengguna diminta memasukkan waktu masuk (`wParkir`) dan waktu keluar (`wPulang`), yang kemudian dikonversi menjadi total detik (`dParkir` dan `dPulang`) untuk mempermudah perhitungan. Selisih total detik ini dihitung sebagai durasi parkir (`lParkir`). Durasi tersebut dikonversi kembali ke format jam, menit, dan detik menggunakan operasi pembagian dan modulus. Hasil akhirnya ditampilkan dalam format yang mudah dibaca, misalnya "Lama Parkir: 2 jam 15 menit 30 detik".

### <h2>GUIDED 2</h2>

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

![image](https://github.com/user-attachments/assets/65c8412e-ccb4-4257-8c9e-c99818d17036)

Deskripsi Program : 
Program ini bertujuan untuk memperbarui daftar teman dengan menambahkan nama-nama baru jika belum ada di dalam daftar, sekaligus memberikan notifikasi jika nama yang dimasukkan sudah ada. Fungsi `sudahAda` digunakan untuk memeriksa apakah sebuah nama sudah terdapat dalam daftar teman dengan melakukan iterasi melalui elemen-elemen di dalam slice `daftarTeman`. Dalam fungsi `main`, terdapat dua slice: `daftarTeman`, yang berisi nama-nama awal, dan `namaBaru`, yang berisi nama-nama yang akan ditambahkan. Program melakukan iterasi pada `namaBaru` dan memeriksa setiap nama menggunakan fungsi `sudahAda`. Jika nama belum ada, maka ditambahkan ke dalam `daftarTeman` menggunakan fungsi `append`. Jika nama sudah ada, program mencetak pesan bahwa nama tersebut sudah ada dalam daftar. Pada akhir eksekusi, program menampilkan daftar teman yang telah diperbarui.

### <h2>GUIDED 3</h2>

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

![image](https://github.com/user-attachments/assets/da6cc609-0ca2-448a-80a1-fb452b2d3d51)

Deskripsi Program : 
Program ini bertujuan untuk menampilkan daftar harga buah dan memungkinkan akses langsung terhadap harga buah tertentu. Menggunakan tipe data map, program menyimpan pasangan nama buah dan harganya dalam variabel `hargaBuah`. Program kemudian mencetak daftar harga buah dengan iterasi menggunakan perulangan `for range`, yang menampilkan nama buah beserta harganya dalam format "Nama: RpHarga". Selain itu, program secara spesifik mencetak harga buah "Mangga" dengan mengakses nilai dari map menggunakan kunci "Mangga". Hasil akhirnya adalah daftar lengkap harga buah dan harga spesifik buah "Mangga" yang ditampilkan secara terpisah.


## <strong> Unguided </strong>

### <h2>UNGUIDED 1</h2>

#### Question

Suatu lingkaran didefinisikan dengan titik pusat `(cx, cy)` dengan radius `r`. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang `(x, y)` berdasarkan dua lingkaran tersebut. **Gunakan tipe bentukan `titik` untuk menyimpan koordinat, dan tipe bentukan `lingkaran` untuk menyimpan titik pusat lingkaran dan radiusnya.**

Masukan
Masukan terdiri dari beberapa tiga baris:
1. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2.
2. Baris ketiga adalah koordinat titik sembarang.

Asumsi sumbu `x` dan `y` dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.

Keluaran
Keluaran berupa string yang menyatakan posisi titik:
- **"Titik di dalam lingkaran 1 dan 2"**
- **"Titik di dalam lingkaran 1"**
- **"Titik di dalam lingkaran 2"**
- **"Titik di luar lingkaran 1 dan 2"**

---

Contoh Masukan dan Keluaran

| No | Masukan               | Keluaran                       |
|----|-----------------------|--------------------------------|
| 1  | 8 8 4                | Titik di dalam lingkaran 1     |
|    | 5 6 2                |                                |
|    | 6 6                  |                                |
|----|-----------------------|--------------------------------|
| 2  | 2 1 3                | Titik di dalam lingkaran 2     |
|    | 5 4 2                |                                |
|    | 6 4                  |                                |
|----|-----------------------|--------------------------------|
| 3  | 5 5 6                | Titik di dalam lingkaran 1 dan 2 |
|    | 5 5 6                |                                |
|    | 5 5                  |                                |
|----|-----------------------|--------------------------------|
| 4  | 15 -15 20            | Titik di luar lingkaran 1 dan 2 |
|    | -15 15 20            |                                |
|    | 20 20                |                                |

---

Fungsi Perhitungan

Untuk menghitung jarak titik `(a, b)` dan `(c, d)` digunakan rumus jarak:


$$\text{jarak} = \sqrt{(a - c)^2 + (b - d)^2}$$


Fungsi ini juga digunakan untuk menentukan posisi sebuah titik sembarang berada di dalam atau di luar lingkaran tertentu.

---

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
    return math.Sqrt(float64((p.x - q.x)*(p.x - q.x) + (p.y - q.y)*(p.y - q.y)))
}

func didalam(c Lingkaran, p Titik) bool {
    return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {
    var l1, l2 Lingkaran
    var p Titik

    fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
    fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
    fmt.Scan(&p.x, &p.y)

    insideL1 := didalam(l1, p)
    insideL2 := didalam(l2, p)

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

![image](https://github.com/user-attachments/assets/51d94b9e-e8fd-4631-8d24-bfbcaec447f5)

Deskripsi Program : 
Program ini digunakan untuk menentukan apakah sebuah titik berada di dalam satu atau dua lingkaran berdasarkan posisi pusat lingkaran, radius, dan koordinat titik yang diberikan. Program memanfaatkan dua struktur data: `Titik` untuk merepresentasikan koordinat (x, y), dan `Lingkaran` yang terdiri dari atribut pusat lingkaran (`Titik`) dan radius.  Fungsi `jarak` menghitung jarak antara dua titik menggunakan rumus jarak Euclidean. Fungsi `didalam` mengevaluasi apakah sebuah titik berada di dalam lingkaran dengan memeriksa apakah jarak titik tersebut ke pusat lingkaran lebih kecil atau sama dengan radius lingkaran. Dalam fungsi `main`, program membaca data untuk dua lingkaran (`l1` dan `l2`) dan satu titik (`p`). Kemudian, program mengevaluasi apakah titik tersebut berada di dalam lingkaran pertama (`l1`), lingkaran kedua (`l2`), atau keduanya, menggunakan fungsi `didalam`. Berdasarkan hasil evaluasi, program mencetak pesan yang sesuai: apakah titik berada di dalam salah satu lingkaran, keduanya, atau di luar keduanya. Program ini sangat berguna untuk analisis geometri sederhana.

### <h2>UNGUIDED 2</h2>

#### Question

Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak **N elemen nilai**. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut:

1. Menampilkan keseluruhan isi dari array.
2. Menampilkan elemen-elemen array dengan indeks ganjil saja.  
3. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indeks ke-0 adalah genap).  
4. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x.`x` bisa diperoleh dari masukan pengguna.  
5. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid.Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil.  
6. Menampilkan rata-rata dari bilangan yang ada di dalam array.  
7. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.  
8. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi sebelumnya.  

#### Source Code

```go
package main

import (
    "fmt"
    "math"
)

type Array []int

func tampilkanSeluruhArray(arr []int) {
    fmt.Println("Seluruh isi array:", arr)
}

func tampilkanIndeksGanjil(arr []int) {
    fmt.Print("Elemen dengan indeks ganjil: ")
    for i := 1; i < len(arr); i += 2 {
        fmt.Print(arr[i], " ")
    }
    fmt.Println()
}

func tampilkanIndeksGenap(arr []int) {
    fmt.Print("Elemen dengan indeks genap: ")
    for i := 0; i < len(arr); i += 2 {
        fmt.Print(arr[i], " ")
    }
    fmt.Println()
}

func tampilkanKelipatanIndeks(arr []int, x int) {
    fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
    for i := 0; i < len(arr); i += x {
        fmt.Print(arr[i], " ")
    }
    fmt.Println()
}

func hapusPadaIndeks(arr []int, index int) []int {
    if index < 0 || index >= len(arr) {
        fmt.Println("Indeks tidak valid")
        return arr
    }
    fmt.Printf("Menghapus elemen di indeks %d (nilai: %d)\n", index, arr[index])
    return append(arr[:index], arr[index+1:]...)
}

func rataRata(arr []int) float64 {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    return float64(sum) / float64(len(arr))
}

func standarDeviasi(arr []int) float64 {
    mean := rataRata(arr)
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

    fmt.Print("Masukkan jumlah elemen array: ")
    fmt.Scan(&n)

    arr := make(Array, n)
    fmt.Println("Masukkan elemen array:")
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    tampilkanSeluruhArray(arr)
    tampilkanIndeksGanjil(arr)
    tampilkanIndeksGenap(arr)

    fmt.Print("Masukkan nilai x untuk kelipatan indeks: ")
    fmt.Scan(&x)
    tampilkanKelipatanIndeks(arr, x)

    fmt.Print("Masukkan indeks untuk menghapus elemen: ")
    fmt.Scan(&index)
    arr = hapusPadaIndeks(arr, index)
    tampilkanSeluruhArray(arr)

    fmt.Printf("Rata-rata elemen array: %.2f\n", rataRata(arr))

    fmt.Printf("Standar deviasi elemen array: %.2f\n", standarDeviasi(arr))

    fmt.Print("Masukkan nilai untuk menghitung frekuensi: ")
    fmt.Scan(&nilai)
    fmt.Printf("Frekuensi %d dalam array: %d\n", nilai, frekuensi(arr, nilai))
}

```

#### Output

![image](https://github.com/user-attachments/assets/d94758de-8b4a-45d2-a678-12bbdd3c8b49)

Deskripsi Program :
Program ini merupakan aplikasi serbaguna untuk melakukan berbagai operasi pada array integer. Program menerima input array dari pengguna, kemudian memberikan sejumlah fungsi untuk menganalisis dan memodifikasi array tersebut. Berikut adalah fitur utama program ini, Tampilkan Isi Array: Fungsi `tampilkanSeluruhArray` mencetak semua elemen dalam array. Elemen Berdasarkan Indeks: Fungsi `tampilkanIndeksGanjil` dan `tampilkanIndeksGenap` mencetak elemen array yang berada pada indeks ganjil dan genap. Elemen pada Indeks Kelipatan X: Fungsi `tampilkanKelipatanIndeks` menampilkan elemen pada indeks yang merupakan kelipatan dari nilai `x` yang dimasukkan oleh pengguna. Hapus Elemen: Fungsi `hapusPadaIndeks` menghapus elemen pada indeks tertentu, jika indeks valid, dengan menggunakan operasi slicing. Rata-rata: Fungsi `rataRata` menghitung rata-rata nilai elemen dalam array. Standar Deviasi: Fungsi `standarDeviasi` menghitung standar deviasi dari elemen array menggunakan formula statistik. Frekuensi Elemen: Fungsi `frekuensi` menghitung berapa kali sebuah nilai tertentu muncul dalam array. Dalam fungsi `main`, program meminta pengguna untuk memasukkan jumlah elemen array (`n`) dan nilai-nilainya. Program selanjutnya menampilkan berbagai informasi berdasarkan fungsi-fungsi yang disediakan, seperti elemen berdasarkan indeks tertentu, rata-rata, standar deviasi, serta frekuensi kemunculan nilai tertentu. Dengan pendekatan interaktif, program ini mempermudah pengguna untuk melakukan eksplorasi data pada array integer.

### <h2>UNGUIDED 3</h2>

#### Question

ebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga.

Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja.

Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.

Perhatikan sesi interaksi pada contoh berikut ini (**teks bergaris bawah adalah Input/read**)

Klub A : **MU**  
Klub B : **Inter**  
Pertandingan 1 : **2 0**         // MU = 2 sedangkan Inter = 0  
Pertandingan 2 : **1 2**  
Pertandingan 3 : **2 2**  
Pertandingan 4 : **0 1**  
Pertandingan 5 : **3 2**  
Pertandingan 6 : **1 0**  
Pertandingan 7 : **5 2**  
Pertandingan 8 : **2 3**  
Pertandingan 9 : **-1 2**  

Hasil 1 : MU  
Hasil 2 : Inter  
Hasil 3 : Draw  
Hasil 4 : Inter  
Hasil 5 : MU  
Hasil 6 : MU  
Hasil 7 : MU  
Hasil 8 : Inter  
Pertandingan selesai

#### Source Code

```go
package main

import (
  "fmt"
)

type Klub struct {
  klubA, klubB string
}

func main() {
  var klub Klub
  fmt.Print("Klub A: ")
  fmt.Scan(&klub.klubA)
  fmt.Print("Klub B: ")
  fmt.Scan(&klub.klubB)

  pemenang := make([]string, 0)
  pertandingan := 1

  for {
    var skorA, skorB int
    fmt.Printf("pertandingan %d :", pertandingan)
    fmt.Scan(&skorA, &skorB)

    if skorA < 0 || skorB < 0 {
      break
    }

    var winner string
    if skorA > skorB {
      winner = klub.klubA
    } else if skorB > skorA {
      winner = klub.klubB
    } else {
      winner = "Draw"
    }

    pemenang = append(pemenang, winner)
    pertandingan++
  }

  fmt.Println("Hasil pertandingan:")
  for i, winner := range pemenang {
    fmt.Printf("Hasil %d: %s\n", i+1, winner)
  }
  fmt.Println("pertandingan selesai")
}

```

#### Output

![image](https://github.com/user-attachments/assets/569633d2-c80c-4c42-bef3-8423fc8d7487)

Deskripsi Program :
Program ini mensimulasikan pertandingan antara dua klub sepak bola. Pengguna diminta untuk memasukkan nama dua klub (klub A dan klub B), lalu program akan mencatat hasil pertandingan yang berulang kali dimasukkan hingga skor negatif diberikan, yang menandakan akhir input. Program ini kemudian mencetak hasil pertandingan dan menampilkan pemenang dari setiap pertandingan yang telah dicatat. Di awal, tipe data `Klub` didefinisikan sebagai struct yang memiliki dua atribut string, yaitu `klubA` dan `klubB`. Fungsi utama `main` dimulai dengan meminta pengguna memasukkan nama dua klub yang akan bertanding, yang disimpan dalam atribut `klubA` dan `klubB`. Daftar pemenang pertandingan disimpan dalam slice `pemenang`, dan variabel `pertandingan` digunakan untuk melacak jumlah pertandingan yang berlangsung. Dalam sebuah loop tak terbatas, pengguna diminta memasukkan skor kedua tim untuk setiap pertandingan. Jika salah satu skor negatif, loop akan berhenti. Program kemudian menentukan pemenang berdasarkan skor. Jika skor tim A lebih tinggi, pemenangnya adalah `klubA`; jika skor tim B lebih tinggi, pemenangnya adalah `klubB`; dan jika skor seri, hasilnya adalah "Draw". Setelah loop selesai, program mencetak hasil dari semua pertandingan yang tercatat, termasuk pemenang masing-masing pertandingan. Terakhir, pesan "pertandingan selesai" ditampilkan untuk menandai akhir program. Program ini menunjukkan pengolahan input secara dinamis dan penggunaan slice untuk menyimpan data hasil pertandingan.

### <h2>UNGUIDED 4</h2>

#### Question

Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.

Lengkapi potongan algoritma berikut ini!

```go
package main
import "fmt"
const NMAX int = 127
type tabel [NMAX]rune
    tab : tabel
    m : integer

func isiArray(t *tabel, n *int)
/* I.S. Data tersedia dalam piranti masukan
   F.S. Array t berisi sejumlah n karakter yang dimasukkan user,
   Proses input selama karakter bukanlah TITIK dan n <= NMAX */

func cetakArray(t tabel, n int)
/* I.S. Terdefinisi array t yang berisi sejumlah n karakter
   F.S. n karakter dalam array muncul di layar */

func balikanArray(t *tabel, n int)
/* I.S. Terdefinisi array t yang berisi sejumlah n karakter
   F.S. Urutan isi array t terbalik */

func main(){
   var tab tabel
   var n int

   // Isi array tab dengan memanggil prosedur isiArray

   // Balikkan isi array tab dengan memanggil balikanArray

   // Cetak isi array tab
}
```

---

**Perhatikan sesi interaksi pada contoh berikut ini** *(teks bergaris bawah adalah Input/read)*:

Teks  
`:S E N A N G.`  
Reverse teks  
`:G N A N E S`  

Teks  
`:K A T A K.`  
Reverse teks  
`:K A T A K`  

---

Modifikasi program tersebut dengan menambahkan fungsi *palindrom*. Tambahkan instruksi untuk memanggil fungsi tersebut dan menampilkan hasilnya pada program utama.  

**"Palindrom adalah teks yang dibaca dari awal atau akhir adalah sama, contoh: KATAK, APA, KASUR_RUSAK."**

```go
func palindrom(t tabel, n int) bool
/* Mengembalikan true apabila susunan karakter di dalam t membentuk palindrom,
   dan false apabila sebaliknya. Petunjuk: Manfaatkan prosedur balikanArray */
```

---

**Perhatikan sesi interaksi pada contoh berikut ini** *(teks bergaris bawah adalah Input/read)*:

Teks  
`:K A T A K.`  
Palindrom  
`:true`  

Teks  
`:S E N A N G.`  
Palindrom  
`:false`  

#### Source Code

```go
package main

import (
    "fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

// Fungsi untuk mengisi array
func isiArray(t *tabel, n *int) {
    /*I.S. Data tersedia dalam piranti masukan
      F.S. Array t berisi sejumlah n karakter yang dimasukkan user,
           Proses input selama karakter bukanlah TITIK dan n <= NMAX */
    var char rune
    *n = 0
    fmt.Print("Masukkan karakter (akhiri dengan titik) \n")
    fmt.Printf("Karakter ke-%d: ", *n+1)
    fmt.Scanf("%c\n", &char)
    
    for char != '.' && *n < NMAX {
        t[*n] = char
        *n++
        if *n < NMAX {
            fmt.Printf("Karakter ke-%d: ", *n+1)
            fmt.Scanf("%c\n", &char)
        }
    }
}

func cetakArray(t tabel, n int) {
    /*I.S. Terdefinisi array t yang berisi sejumlah n karakter 
        F.S. n karakter dalam array muncul di layar */
    fmt.Print("Teks: ")
    for i := 0; i < n; i++ {
        fmt.Printf("%c", t[i])
    }
    fmt.Println()
}

func balikArray(t *tabel, n int) {
    /*I.S. Terdefinisi array t yang berisi sejumlah n karakter
      F.S. Urutan isi array t terbalik */
    var temp tabel
    for i := 0; i < n; i++ {
        temp[i] = t[n-1-i]
    }
    for i := 0; i < n; i++ {
        t[i] = temp[i]
    }
}

func palindrom(t tabel, n int) bool {
    /* Mengembalikan true apabila susunan karakter di dalam t membentuk palindrom,
       dan false apabila sebaliknya. Memanfaatkan prosedur balikArray */
    var temp tabel = t
    balikArray(&temp, n)

    for i := 0; i < n; i++ {
        if t[i] != temp[i] {
            return false
        }
    }
    return true
}

func main() {
    var tab tabel
    var m int
    
    isiArray(&tab, &m)
    
    fmt.Print("Original ")
    cetakArray(tab, m)
    
    var reversed tabel = tab
    balikArray(&reversed, m)
    
    fmt.Print("Reverse ")
    cetakArray(reversed, m)
    
    if palindrom(tab, m) {
        fmt.Println("Palindrom: True")
    } else {
        fmt.Println("Palindrom: False")
    }
}

```

#### Output

![image](https://github.com/user-attachments/assets/1b80d932-bdc0-42a1-80f8-8737cfa4f9dd)

Deskripsi Program :
Program ini adalah aplikasi konsol yang menggunakan bahasa Go untuk memproses dan menganalisis array karakter. Program dimulai dengan mendeklarasikan konstanta `NMAX`, yang menentukan ukuran maksimum array, dan tipe data `tabel`, berupa array berukuran tetap untuk menyimpan karakter. Fungsi `isiArray` bertugas mengisi array dengan karakter yang dimasukkan oleh pengguna, hingga karakter titik (`.`) ditemukan atau jumlah karakter mencapai batas maksimum. Program kemudian mencetak array asli menggunakan fungsi `cetakArray`. Selanjutnya, fungsi `balikArray` membalikkan urutan elemen dalam array, dan hasilnya dicetak kembali. Untuk memeriksa apakah array membentuk sebuah palindrom (susunan yang sama jika dibaca dari depan maupun belakang), fungsi `palindrom` membandingkan array asli dengan versi terbaliknya. Program utama mengintegrasikan fungsi-fungsi ini untuk menampilkan array asli, array terbalik, dan hasil analisis apakah array adalah palindrom. Program ini memanfaatkan konsep manipulasi array, perulangan, dan pemrosesan karakter secara efisien.

## <strong> Kesimpulan </strong>

## Kesimpulan

1. **Array**:
   - Array adalah kumpulan elemen dengan tipe data yang sama, memiliki ukuran tetap yang ditentukan saat deklarasi, dan cocok untuk situasi di mana jumlah elemen sudah diketahui.
   - Elemen dalam array diakses menggunakan indeks, dan iterasi dapat dilakukan dengan loop `for`, menjadikannya mudah untuk manipulasi data sederhana.

2. **Struct**:
   - Struct memungkinkan pengelompokan beberapa variabel dengan tipe yang berbeda dalam satu unit logis, sangat berguna untuk merepresentasikan objek dunia nyata dengan atribut-atributnya.
   - Field dalam struct diakses menggunakan notasi titik, dan struct meningkatkan keterbacaan serta pengorganisasian kode, terutama dalam aplikasi yang kompleks.

3. **Penggunaan yang Tepat**:
   - Pemilihan antara array dan struct tergantung pada kebutuhan spesifik: gunakan array untuk koleksi data bertipe sama dan struct untuk mengelompokkan data yang berbeda. Memahami karakteristik masing-masing akan membantu pengembang dalam merancang struktur data yang efisien dan efektif.

## <strong> Referensi </strong>

#### [1] Kurniawan, Rizky. 2022. #13: Tipe Data Array - Belajar Golang Dari Dasar. Diakses melalui website https://blog.ruangdeveloper.com/golang-array/#google_vignette

#### [2] Andi. Golang: Arti, Cara Kerja, Struktur, Kelebihan dan Kekurangan. Diakses melalui website https://www.wangs.id/digital-development/golang-adalah/

#### [3] JEDITOR. 2024. Struktur Data di Go Pemahaman dan Implementasi. Diakses melalui website https://jocodev.id/struktur-data-di-go-pemahaman-dan-implementasi/
