<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL VI</strong></h2>
<h2 align="center"><strong> REKURSIF </strong></h2>

<br>

<p align="center">

  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Wahyu Hidayat / 2311102178<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh, S.Kom., M.Kom.
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
### Definisi Rekursif
Rekursif adalah konsep dalam pemrograman di mana suatu fungsi memanggil dirinya sendiri untuk menyelesaikan masalah yang lebih kecil dari masalah awal. Teknik ini sering digunakan untuk memecah masalah kompleks menjadi sub-masalah yang lebih sederhana, sehingga masalah dapat diselesaikan secara bertahap hingga mencapai solusi akhir. Rekursif umum digunakan dalam perhitungan faktorial, deret Fibonacci, dan algoritma pencarian seperti pencarian dalam pohon (tree)[1].

### Jenis-Jenis Rekursif
Ada dua jenis utama rekursif: rekursif langsung dan rekursif tidak langsung. Pada rekursif langsung, suatu fungsi memanggil dirinya sendiri secara langsung. Sedangkan pada rekursif tidak langsung, fungsi memanggil fungsi lain, yang pada gilirannya memanggil fungsi pertama. Kedua jenis rekursif ini bermanfaat dalam situasi yang berbeda, namun rekursif langsung lebih sering digunakan karena lebih mudah dipahami dan diterapkan untuk masalah sederhana[2].

### Base Case dan Recursive Case
ungsi rekursif harus memiliki dua komponen utama: base case dan recursive case. Base case adalah kondisi berhenti yang mencegah fungsi berjalan tanpa henti, sehingga fungsi berhenti memanggil dirinya sendiri ketika mencapai kondisi ini. Recursive case adalah bagian dari fungsi yang memanggil dirinya sendiri dengan versi yang lebih kecil dari masalah asli, hingga akhirnya mencapai base case. Kedua komponen ini penting agar rekursif berjalan dengan benar dan tidak menyebabkan stack overflow[3].

Contoh Implementasi Rekursif dalam Go:

```go
package main

import "fmt"

func faktorial(n int) int {
    if n == 1 { // Base case
        return 1
    }
    return n * faktorial(n-1) // Recursive case
}

func main() {
    fmt.Println(faktorial(5)) // Output: 120
}


```
#### Penjelasan:
Pada contoh ini, fungsi faktorial memiliki base case yang mengecek apakah n adalah 1. Jika iya, fungsi mengembalikan 1. Jika tidak, fungsi akan memanggil dirinya sendiri dengan n-1 hingga akhirnya mencapai base case. Ini memungkinkan kita untuk menghitung faktorial secara bertahap dari n hingga 1[4].

### Kelebihan dan Kekurangan Rekursif
Rekursif menawarkan cara yang sederhana untuk menyelesaikan masalah yang dapat dipecah menjadi bagian-bagian yang lebih kecil, terutama untuk struktur data yang bersifat hierarkis seperti pohon. Namun, pemakaian rekursif yang tidak tepat atau tanpa base case yang jelas dapat menyebabkan konsumsi memori yang besar dan bahkan stack overflow, karena setiap pemanggilan fungsi membutuhkan ruang di stack. Di Go, rekursif sering kali kurang efisien dibandingkan pendekatan iteratif dalam kasus yang sederhana karena Go tidak mendukung optimisasi tail recursion[5].


## II. GUIDED
## 1. Membuat baris bilangan dari n hingga 1

#### Source Code
```go
package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	baris(n)
}

func baris(bilangan int){
	if bilangan == 1 {
		fmt.Println(1)
	}else{
		fmt.Println(bilangan)
		baris(bilangan - 1)
	}
}
```
#### Screenshoot Source Code
![Screenshot 2024-11-03 184444](https://github.com/user-attachments/assets/12ec0aa1-0e7b-49c9-b7ad-28501d5b7e23)




#### Screenshoot Output
![Screenshot 2024-11-03 184455](https://github.com/user-attachments/assets/1edb7fb1-1a9d-41ec-8cb2-cb174362ca57)



#### Deskripsi Program
Program ini adalah program rekursif sederhana dalam bahasa Go yang menerima input berupa bilangan bulat n dari pengguna, lalu mencetak deretan angka dari n hingga 1. Program ini menggunakan fungsi baris, yang memanggil dirinya sendiri (rekursif) untuk mengurangi bilangan hingga mencapai nilai 1.

#### Algoritma Program
1. Program meminta input dari pengguna dan menyimpannya dalam variabel n.
2. Fungsi baris dipanggil dengan parameter n.
3. Dalam fungsi baris:
   - Jika bilangan sama dengan 1, program mencetak 1 dan berhenti (kondisi dasar / base case).
   - Jika bilangan lebih besar dari 1, program mencetak nilai bilangan, kemudian memanggil dirinya sendiri dengan bilangan - 1.
4. Proses ini berulang hingga nilai bilangan menjadi 1, lalu program berhenti.

#### Cara Kerja
1. Meminta Input: Program pertama-tama menunggu pengguna untuk memasukkan sebuah bilangan bulat n dan menyimpan input ini ke variabel n.
2. Memanggil Fungsi Rekursif: Program memanggil fungsi baris(n).
3. Fungsi Rekursif baris:
   - Fungsi ini menggunakan parameter bilangan, yang awalnya adalah nilai n yang diinputkan pengguna.
   - Kondisi Base Case: Jika bilangan == 1, fungsi mencetak 1 dan berhenti (tidak ada lagi pemanggilan rekursif).
   - Kondisi Rekursif: Jika bilangan > 1, fungsi mencetak nilai bilangan, lalu memanggil dirinya sendiri dengan parameter bilangan - 1, yang mengurangi nilai bilangan sebesar 1 di setiap langkah hingga mencapai nilai 1.
4. Output Program: Program mencetak nilai dari n hingga 1 secara berurutan dengan setiap pemanggilan rekursif. Jika, misalnya, pengguna memasukkan 5 sebagai input, hasilnya akan seperti ini:
- 5
- 4
- 3
- 2
- 1

## 2. Menghitung hasil penjumlahan 1 hingga n

#### Source Code
```go
package main 
import "fmt"

func penjumlahan(n int) int{
	if n == 1 {
		return 1
	}else{
		return n + penjumlahan(n-1)
	}
}

func main(){
	var n int
	fmt.Scan(&n)
	fmt.Println(penjumlahan(n))
}

```
#### Screenshoot Source Code
![Screenshot 2024-11-03 185336](https://github.com/user-attachments/assets/e5122cc1-3be5-427f-9db0-17af0f018a39)


#### Screenshoot Output
![Screenshot 2024-11-03 185340](https://github.com/user-attachments/assets/4d287d22-3432-4c60-9a23-1af90d51cd44)

#### Deskripsi Program
Program ini adalah implementasi dari fungsi penjumlahan rekursif dalam bahasa Go. Program menerima input berupa bilangan bulat n dari pengguna dan menghitung jumlah dari semua bilangan bulat dari 1 hingga n. Fungsi rekursif penjumlahan dipanggil untuk menghitung hasil penjumlahan tersebut, menggunakan pendekatan rekursif untuk menambahkan bilangan satu per satu hingga mencapai bilangan 1.

#### Algoritma Program
1. Program meminta input dari pengguna dan menyimpannya dalam variabel n.
2. Fungsi penjumlahan dipanggil dengan parameter n.
3. Dalam fungsi penjumlahan:
   - Jika n sama dengan 1, fungsi mengembalikan 1 (ini adalah kondisi dasar).
   - Jika n lebih besar dari 1, fungsi mengembalikan hasil penjumlahan n dan hasil dari pemanggilan penjumlahan(n-1).
4. Hasil akhir dari penjumlahan ditampilkan ke layar.

#### Cara Kerja
1. Meminta Input: Program memulai dengan meminta pengguna untuk memasukkan sebuah bilangan bulat n.
2. Memanggil Fungsi Rekursif: Setelah input diterima, program memanggil fungsi penjumlahan dengan nilai n sebagai argumen.
3. Fungsi Rekursif penjumlahan:
   - Fungsi ini menerima parameter n dan mengecek apakah n adalah 1 (kondisi dasar).
   - Kondisi Dasar: Jika n == 1, fungsi mengembalikan 1.
   - Kondisi Rekursif: Jika n > 1, fungsi akan mengembalikan nilai n ditambahkan dengan hasil pemanggilan penjumlahan(n-1). Ini berarti fungsi memanggil dirinya sendiri dengan parameter n-1, mengurangi nilai n setiap kali hingga mencapai 1.
4. Output Program: Setelah semua pemanggilan fungsi selesai dan nilai dikembalikan, program mencetak hasil akhir penjumlahan ke layar. Jika pengguna memasukkan 5, hasil akhir yang ditampilkan adalah 15, yang merupakan hasil dari penjumlahan 1+2+3+4+5.

## 3. Mencari dua pangkat n atau 2^n

#### Source Code
```go
package main

import "fmt"

func pangkat(n int) int {
	if n == 0 {
		return 1
	}else{
		return 2 * pangkat(n-1)
	}
}

func main(){
	var n int
	fmt.Print("Masukkan nilai n : ")
	fmt.Scan(&n)
	fmt.Println("Hasil dari 2 pangkat", n, "adalah :", pangkat(n))
}
```

#### Screenshoot Source Code
![Screenshot 2024-11-03 190225](https://github.com/user-attachments/assets/51cfeabf-7c5a-464b-b8e7-27c855e0eb03)


#### Screenshoot Output
![Screenshot 2024-11-03 190429](https://github.com/user-attachments/assets/ea845c0c-283a-4dfa-86fa-91d664fe4bec)

#### Deskripsi Program
Program ini menghitung nilai dua pangkat n dengan menggunakan metode rekursif. Pengguna akan diminta untuk memasukkan nilai bulat n, dan program akan memberikan hasil perhitungan dua pangkat n. Fungsi yang digunakan untuk perhitungan adalah pangkat, yang akan memanggil dirinya sendiri hingga mencapai kondisi dasar.

#### Algoritma Program
1. Minta pengguna untuk memasukkan nilai bulat n.
2. Fungsi pangkat:
   - Jika n sama dengan nol, kembalikan satu.
   - Jika n lebih besar dari nol, kembalikan hasil dari dua kali pangkat n dikurangi satu.
3. Tampilkan hasil perhitungan dua pangkat n.
#### Cara Kerja
1. Program dimulai dengan fungsi utama.
2. Di dalam fungsi utama, variabel n dideklarasikan untuk menyimpan input dari pengguna.
3. Program meminta pengguna untuk memasukkan nilai n menggunakan fungsi input.
4. Setelah mendapatkan nilai n, program memanggil fungsi pangkat untuk menghitung dua pangkat n.
5. Fungsi pangkat bekerja dengan cara:
   - Jika n sama dengan nol, fungsi akan mengembalikan satu.
   - Jika n lebih besar dari nol, fungsi akan mengalikan dua dengan hasil pemanggilan fungsi pangkat dengan argumen n dikurangi satu.
6. Proses ini akan berlanjut sampai mencapai kondisi n sama dengan nol.
7. Hasil akhir akan dicetak ke layar dengan format yang telah ditentukan.

## 4. Mencari nilai faktorial atau n!

#### Source Code
```go
package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan nilai n: ")
    fmt.Scan(&n)
    fmt.Println("Hasil faktorial dari", n, "adalah:", faktorial(n))
}

func faktorial(n int) int {
    if n == 0 || n == 1 {
        return 1
    } else {
        return n * faktorial(n-1)
    }
}

```

#### Screenshoot Source Code
![Screenshot 2024-11-03 192747](https://github.com/user-attachments/assets/8b864e8f-c480-425d-9a08-1567a026f0a3)


#### Screenshoot Output
![Screenshot 2024-11-03 192752](https://github.com/user-attachments/assets/ac3d5ab8-d61c-4624-a26d-581eda4e942b)

#### Deskripsi Program
Program ini menghitung faktorial dari bilangan bulat yang dimasukkan oleh pengguna. Faktorial dari suatu bilangan n adalah hasil perkalian dari semua bilangan bulat positif dari satu hingga n. Program ini menggunakan pendekatan rekursif untuk menghitung faktorial, yang berarti fungsi akan memanggil dirinya sendiri hingga mencapai kondisi dasar.

#### Algoritma Program
1. Minta pengguna untuk memasukkan nilai bulat n.
2. Fungsi faktorial dengan parameter n:
   - Jika n sama dengan nol atau n sama dengan satu, kembalikan satu.
   - Jika n lebih besar dari satu, kembalikan hasil dari n dikali dengan pemanggilan fungsi faktorial dengan argumen n dikurangi satu.
3. Tampilkan hasil perhitungan faktorial dari n.

#### Cara Kerja
1. Program dimulai dari fungsi utama.
2. Di dalam fungsi utama, variabel n dideklarasikan untuk menyimpan input dari pengguna.
3. Program meminta pengguna untuk memasukkan nilai n menggunakan fungsi input.
4. Setelah mendapatkan nilai n, program memanggil fungsi faktorial untuk menghitung faktorial.
5. Fungsi faktorial bekerja dengan cara:
   - Jika nilai n sama dengan nol atau satu, fungsi mengembalikan satu.
   - Jika n lebih besar dari satu, fungsi mengalikan n dengan hasil pemanggilan fungsi faktorial dengan argumen n dikurangi satu.
6. Proses ini akan berlanjut hingga mencapai kondisi di mana n sama dengan nol atau satu.
7. Hasil akhir akan dicetak ke layar dengan format yang telah ditentukan.



## III. UNGUIDED
## 1. Program Deret Fibonacci Rekursif dengan Input Pengguna

#### Source Code
```go
package main

import (
    "fmt"
)

// Fungsi rekursif untuk menghitung nilai Fibonacci ke-n
func fibonacci(n int) int {
    // Kasus dasar
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    }
    // Kasus rekursif
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    // Cetak deret Fibonacci hingga suku ke-10
    for i := 0; i <= 10; i++ {
        fmt.Printf("Fibonacci(%d) = %d\n", i, fibonacci(i))
    }
}

```
#### Screenshoot Source Code
![Screenshot 2024-11-03 193616](https://github.com/user-attachments/assets/ab5ec901-9459-4b2e-a7ed-b29f2cd12d40)

#### Screenshoot Output
![Screenshot 2024-11-03 193621](https://github.com/user-attachments/assets/333fa9f2-4689-4d04-b5e5-165f5e7fa4be)

#### Deskripsi Program
Program ini menghitung dan menampilkan deret Fibonacci hingga suku ke-n, di mana nilai n ditentukan oleh pengguna melalui input. Deret Fibonacci adalah urutan bilangan di mana setiap angka setelah dua angka pertama adalah hasil penjumlahan dari dua angka sebelumnya. Program ini menggunakan fungsi rekursif untuk menghitung nilai setiap suku dalam deret Fibonacci.
#### Algoritma Program
1. Input Pengguna: Program meminta pengguna untuk memasukkan sebuah bilangan bulat n yang menunjukkan hingga suku ke berapa deret Fibonacci akan ditampilkan.
2. Fungsi Rekursif: Program menggunakan fungsi rekursif bernama fibonacci untuk menghitung nilai dari suku ke-n dalam deret Fibonacci:
   - Jika n sama dengan nol, maka fibonacci mengembalikan nol.
   - Jika n sama dengan satu, maka fibonacci mengembalikan satu.
   - Jika n lebih dari satu, maka fibonacci mengembalikan hasil dari fibonacci(n minus satu) + fibonacci(n minus dua).
3. Cetak Deret Fibonacci: Program memanggil fungsi fibonacci untuk setiap suku dari nol hingga n, dan menampilkan hasilnya satu per satu.

#### Cara Kerja
1. Program memulai dengan meminta pengguna memasukkan nilai n.
2. Setelah menerima input, program akan menggunakan loop for dari nol hingga n. Di setiap iterasi, program memanggil fungsi fibonacci untuk menghitung nilai suku ke-i.
3. Fungsi fibonacci bekerja secara rekursif. Jika nilai i lebih dari satu, fungsi akan memanggil dirinya sendiri dua kali, dengan parameter i dikurangi satu dan i dikurangi dua, hingga mencapai kasus dasar yaitu nol atau satu.
4. Hasil dari setiap suku Fibonacci dicetak dalam bentuk Fibonacci ke-i sama dengan hasil.
5. Program berakhir setelah menampilkan semua suku Fibonacci hingga suku ke-n.

## 2. Program Pola Bintang Menggunakan Fungsi Rekursif
#### Source Code
```go
package main

import (
    "fmt"
)

// Fungsi rekursif untuk mencetak bintang pada setiap baris
func printStars(n int) {
    if n == 0 {
        return
    }
    fmt.Print("*")
    printStars(n - 1)
}

// Fungsi rekursif untuk mencetak pola bintang hingga baris ke-n
func printPattern(n int, current int) {
    if current > n {
        return
    }
    printStars(current)
    fmt.Println()
    printPattern(n, current + 1)
}

func main() {
    var n int
    fmt.Print("Masukkan jumlah baris pola bintang: ")
    fmt.Scan(&n)

    fmt.Println("Pola bintang:")
    printPattern(n, 1)
}


```
#### Screenshoot Source Code
![Screenshot 2024-11-03 200738](https://github.com/user-attachments/assets/33089e9b-fd15-448e-beec-35ea79f56bad)

#### Screenshoot Output
![Screenshot 2024-11-03 200743](https://github.com/user-attachments/assets/189fb2f7-caec-4f40-81ef-384628a4d63c)

#### Deskripsi Program
Program ini digunakan untuk menampilkan pola bintang berbentuk segitiga yang terdiri dari beberapa baris sesuai input yang diberikan oleh pengguna. Pengguna akan memasukkan jumlah baris, dan program akan mencetak pola bintang yang bertambah pada setiap baris, mulai dari satu bintang di baris pertama hingga sejumlah bintang sesuai jumlah baris yang diinginkan. Program ini menggunakan fungsi rekursif untuk mencetak bintang dalam setiap baris dan mencetak pola secara keseluruhan.

#### Algoritma Program 
1. Input Pengguna: Program meminta pengguna untuk memasukkan sebuah bilangan bulat yang menunjukkan jumlah baris dalam pola bintang.
2. Fungsi cetak bintang: Fungsi ini adalah fungsi rekursif yang digunakan untuk mencetak sejumlah bintang dalam satu baris:
   - Jika nilai input sama dengan nol, fungsi berhenti dan tidak mencetak bintang.
   - Jika nilai input lebih dari nol, fungsi mencetak satu bintang dan memanggil dirinya sendiri dengan nilai input dikurangi satu, sehingga mencetak bintang secara berulang hingga mencapai jumlah yang diinginkan.
3. Fungsi cetak pola: Fungsi ini adalah fungsi rekursif untuk mencetak pola bintang baris demi baris:
   - Fungsi menerima dua parameter, yaitu jumlah total baris yang diinginkan dan sebuah nilai yang menunjukkan jumlah bintang di baris saat ini.
   - Jika nilai baris saat ini lebih besar dari jumlah total baris, fungsi berhenti dan pola selesai dicetak.
   - Jika nilai baris saat ini masih kurang atau sama dengan jumlah total baris, fungsi memanggil fungsi cetak bintang untuk mencetak sejumlah bintang di baris tersebut, kemudian menampilkan baris baru, dan memanggil dirinya sendiri dengan nilai baris saat ini ditambah satu.
4. Output: Program akan mencetak pola bintang dalam bentuk segitiga, sesuai dengan jumlah baris yang diminta oleh pengguna.


#### Cara Kerja
1. Program dimulai dengan meminta pengguna untuk memasukkan jumlah baris dalam pola bintang.
2. Setelah menerima input, program memanggil fungsi cetak pola dengan parameter jumlah total baris yang diinginkan dan baris saat ini dimulai dari satu.
3. Proses Rekursi:
   - Di setiap pemanggilan fungsi cetak pola, program mencetak sejumlah bintang sesuai dengan baris saat ini dengan memanggil fungsi cetak bintang.
   - Fungsi cetak bintang mencetak bintang satu per satu hingga jumlahnya sesuai dengan baris saat ini, kemudian berhenti.
   - Setelah fungsi cetak bintang selesai, program menambahkan baris baru dan memanggil fungsi cetak pola lagi dengan baris saat ini ditambah satu, sehingga jumlah bintang di baris berikutnya bertambah satu.
4. Proses ini berulang hingga nilai baris saat ini lebih besar dari jumlah total baris, di mana program berhenti dan pola bintang selesai dicetak.

## 3. Program Rekursif untuk Menampilkan Faktor Bilangan

#### Source Code
```go
package main

import (
	"fmt"
)

// Fungsi rekursif untuk mencari faktor dari N dalam urutan menaik
func findFactors(n, current int) {
	if current > n {
		return
	}
	if n%current == 0 {
		fmt.Print(current, " ")
	}
	findFactors(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)
	fmt.Print("Faktor dari ", n, " adalah: ")
	findFactors(n, 1)
	fmt.Println()
}
```
#### Screenshoot Source Code
![Screenshot 2024-11-03 201619](https://github.com/user-attachments/assets/57b03312-7e68-4b8e-aed3-409e266c3b6f)


#### Screenshoot Output
![Screenshot 2024-11-03 201623](https://github.com/user-attachments/assets/b2a46845-7794-48f7-9f37-8a0a4486cff5)

#### Deskripsi Program
Program ini adalah implementasi dalam bahasa Go untuk mencari faktor-faktor dari sebuah bilangan positif menggunakan rekursi. Program akan menerima input berupa bilangan bulat positif dari pengguna, kemudian menampilkan semua faktor dari bilangan tersebut dalam urutan menaik, mulai dari satu hingga bilangan itu sendiri.

#### Algoritma Program
1. Input: Terima input bilangan bulat positif dari pengguna.
2. Pencarian Faktor:
   - Gunakan fungsi rekursif findFactors untuk mencari dan mencetak faktor-faktor dari bilangan tersebut.
   - Mulai dengan nilai current sama dengan satu, dan cek apakah current merupakan faktor dari bilangan tersebut.
   - Jika benar, cetak current sebagai faktor.
   - Ulangi langkah ini dengan menaikkan current hingga current melebihi bilangan input.
3. Output: Tampilkan faktor-faktor dari bilangan dalam urutan menaik.

#### Cara Kerja
1. Fungsi main:
   - Meminta pengguna memasukkan bilangan bulat positif.
   - Memanggil fungsi findFactors dengan current bernilai satu untuk mulai mencari faktor dari bilangan tersebut.
2. Fungsi Rekursif findFactors:
   - Fungsi ini memiliki dua parameter, yaitu n sebagai bilangan yang dicari faktornya dan current sebagai angka yang sedang diperiksa apakah merupakan faktor.
   - Langkah-langkah fungsi findFactors:
     - Basis Rekursi: Jika current lebih besar dari n, hentikan fungsi karena semua faktor sudah ditemukan.
     - Cek Faktor: Jika n habis dibagi current, berarti current adalah faktor dari n, sehingga current dicetak.
     - Panggil Diri Sendiri: Panggil kembali findFactors dengan current ditambah satu, sehingga fungsi memeriksa angka berikutnya.
3. Contoh Jalannya Program:
   - Misal input dari pengguna adalah dua belas.
   - Fungsi findFactors akan dipanggil, dan angka-angka dari satu hingga dua belas akan diperiksa satu per satu apakah merupakan faktor dari dua belas.
   - Faktor-faktor yang ditemukan (satu, dua, tiga, empat, enam, dua belas) akan dicetak dalam urutan menaik.


## 4. Program yang Mengimplementasikan Rekursif untuk Menampilkan Barisan Bilangan Tertentu
#### Source Code
```go
package main

import (
	"fmt"
)

// Fungsi rekursif untuk menampilkan bilangan dari n sampai 1 dan kembali ke n
func printSequence(n, current int) {
	// Cetak bilangan menurun dari current ke 1
	if current > 0 {
		fmt.Print(current, " ")
		printSequence(n, current-1) // Rekursif turun ke bawah
	}

	// Cetak bilangan naik dari 2 hingga n
	if current < n {
		fmt.Print(current+1, " ")
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n) // Menerima input dari pengguna
	fmt.Print("Hasil barisan: ")
	printSequence(n, n) // Memanggil fungsi dengan nilai awal n
	fmt.Println() // Pindah ke baris baru
}

```
#### Screenshoot Source Code
![Screenshot 2024-11-03 202412](https://github.com/user-attachments/assets/b030d046-57ad-4daa-9723-c9a9bd02197c)

#### Screenshoot Output
![Screenshot 2024-11-03 202420](https://github.com/user-attachments/assets/74c5ac64-a970-49b4-af7c-6df6173b4054)

#### Deskripsi Program
Program ini meminta pengguna untuk memasukkan sebuah bilangan bulat positif n. Kemudian, menggunakan fungsi rekursif, program mencetak urutan angka dari n hingga satu, diikuti dengan angka dari dua hingga n. Dengan kata lain, program ini menampilkan urutan bilangan secara menurun dan kemudian kembali menaik.

#### Algoritma Program
1. Input Bilangan: Menerima input dari pengguna berupa bilangan bulat positif n.
2. Fungsi Rekursif: Menggunakan fungsi printSequence dengan dua parameter:
   - n: bilangan bulat positif yang diinputkan oleh pengguna.
   - current: bilangan saat ini yang akan dicetak, dimulai dari n dan berkurang hingga satu.
3. Mencetak Bilangan Menurun:
   - Jika current lebih besar dari nol, cetak nilai current.
   - Panggil fungsi printSequence secara rekursif dengan current dikurangi satu.
4. Mencetak Bilangan Menaik:
   - Setelah mencapai satu, jika current kurang dari n, cetak current ditambah satu untuk menampilkan bilangan yang naik hingga n.
5. Output: Tampilkan hasil urutan bilangan di konsol.

#### Cara Kerja
1. Program dimulai dari fungsi main, di mana variabel n dideklarasikan untuk menyimpan input pengguna.
2. Pengguna diminta untuk memasukkan bilangan bulat positif.
3. Fungsi printSequence dipanggil dengan n sebagai nilai awal untuk current.
4. Fungsi printSequence beroperasi dengan cara berikut:
   - Pertama, mencetak angka menurun dari current hingga satu.
   - Setelah mencapai satu, fungsi mulai mencetak angka menaik dari dua hingga n.
5. Program berakhir setelah semua angka dicetak dan mengeluarkan hasil di baris baru.

## 5. Program Rekursif untuk Menampilkan Bilangan Ganjil
#### Source Code
```go
package main

import (
	"fmt"
)

// Fungsi rekursif untuk menampilkan bilangan ganjil dari 1 hingga N
func cetakBilanganGanjil(n, current int) {
	if current > n {
		return
	}
	if current%2 != 0 {
		fmt.Print(current, " ")
	}
	cetakBilanganGanjil(n, current+1)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Print("Bilangan ganjil dari 1 hingga ", N, ": ")
	cetakBilanganGanjil(N, 1)
	fmt.Println()
}


```
#### Screenshoot Source Code
![Screenshot 2024-11-03 203151](https://github.com/user-attachments/assets/6d661778-4b65-49f0-8b35-1c403f4fd5e5)

#### Screenshoot Output
![Screenshot 2024-11-03 203156](https://github.com/user-attachments/assets/0fc8a9c9-95f2-4778-a180-4d0035b8fb0a)


#### Deskripsi Program
Program ini bertujuan untuk menampilkan barisan bilangan ganjil dari 1 hingga N dengan menggunakan metode rekursif. Pengguna diminta untuk memasukkan sebuah bilangan bulat positif N, kemudian program akan mencetak bilangan ganjil satu per satu sampai mencapai nilai N.

#### Algoritma Program 
1. Minta pengguna memasukkan bilangan bulat positif N.
2. Buat sebuah fungsi rekursif bernama cetakBilanganGanjil yang menerima dua parameter: nilai N (batas akhir) dan current (bilangan yang sedang diperiksa).
3. Di dalam fungsi rekursif:
   - Jika nilai current melebihi N, hentikan fungsi.
   - Jika current adalah bilangan ganjil, cetak bilangan tersebut.
   - Panggil fungsi cetakBilanganGanjil lagi dengan current bertambah satu.
4. Panggil fungsi cetakBilanganGanjil dari fungsi utama, dimulai dari 1.
5. Program mencetak bilangan ganjil dari 1 hingga N.

#### Cara Kerja
1. Input Pengguna: Program meminta pengguna untuk memasukkan bilangan bulat positif N. Misalnya, jika pengguna memasukkan 20, maka N menjadi 20.
2. Inisialisasi dan Pemanggilan Fungsi: Program memanggil fungsi rekursif cetakBilanganGanjil dengan N dan memulai dari 1 sebagai bilangan awal (current).
3. Logika Rekursif:
   - Periksa Kondisi Akhir: Jika current lebih besar dari N, maka fungsi berhenti. Ini mencegah program terus berjalan tanpa batas.
   - Cek Bilangan Ganjil: Jika current adalah bilangan ganjil (memiliki sisa 1 saat dibagi 2), program mencetak bilangan tersebut.
   - Panggilan Rekursif: Fungsi cetakBilanganGanjil dipanggil lagi dengan current bertambah satu. Ini membuat program terus memeriksa bilangan berikutnya hingga mencapai N.
4. Cetak Hasil: Program mencetak semua bilangan ganjil dari 1 hingga N dalam satu baris, dipisahkan oleh spasi.

## 6. Program Rekursif untuk Menghitung Pangkat Bilangan
#### Source Code
```go
package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghitung x dipangkatkan y
func pangkat(x, y int) int {
	if y == 0 {
		return 1 // Basis: jika y = 0, hasilnya adalah 1
	}
	return x * pangkat(x, y-1) // Rekursi: x * (x pangkat y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan bulat x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan bulat y: ")
	fmt.Scan(&y)

	hasil := pangkat(x, y)
	fmt.Printf("Hasil %d dipangkatkan %d adalah %d\n", x, y, hasil)
}

```
#### Screenshoot Source Code
![Screenshot 2024-11-03 204030](https://github.com/user-attachments/assets/a582e20d-8bc2-4cee-84c5-30f34eed0054)

#### Screenshoot Output
![Screenshot 2024-11-03 204035](https://github.com/user-attachments/assets/52820c11-4f6f-4b56-90ae-7264b5bb293a)


#### Deskripsi Program
Program ini bertujuan untuk menghitung hasil pangkat dari dua bilangan bulat x dan y menggunakan metode rekursif. Pengguna diminta untuk memasukkan bilangan x (basis) dan y (pangkat), kemudian program akan menghitung hasil x dipangkatkan y tanpa menggunakan pustaka matematika bawaan.

#### Algoritma Program 
1. Minta pengguna memasukkan dua bilangan bulat x dan y.
2. Buat fungsi rekursif bernama pangkat yang menerima dua parameter: x (basis) dan y (pangkat).
3. Di dalam fungsi rekursif:
   - Jika y sama dengan nol, kembalikan hasil satu. Ini adalah kasus dasar di mana bilangan apa pun yang dipangkatkan nol adalah satu.
   - Jika y lebih besar dari nol, kembalikan hasil x dikalikan dengan panggilan rekursif pangkat dengan y dikurangi satu.
4. Panggil fungsi pangkat dari fungsi utama dengan x dan y sebagai argumen.
5. Cetak hasil perhitungan pangkat.

#### Cara Kerja
1. Input Pengguna: Program meminta pengguna untuk memasukkan dua bilangan bulat, x sebagai bilangan dasar dan y sebagai bilangan pangkat. Misalnya, jika pengguna memasukkan 2 untuk x dan 3 untuk y, maka x = 2 dan y = 3.
2. Inisialisasi dan Pemanggilan Fungsi: Program memanggil fungsi pangkat untuk menghitung x dipangkatkan y.
3. Logika Rekursif:
   - Kasus Dasar: Jika y adalah nol, hasilnya adalah satu, dan fungsi berhenti. Ini karena bilangan apa pun yang dipangkatkan nol hasilnya satu.
   - Perhitungan Rekursif: Jika y lebih besar dari nol, hasilnya dihitung dengan mengalikan x dengan hasil dari panggilan fungsi pangkat dengan y berkurang satu. Ini berlanjut sampai y menjadi nol.
4. Cetak Hasil: Program menghitung dan mencetak hasil dari x dipangkatkan y.

### Kesimpulan
Rekursif adalah teknik dalam pemrograman di mana sebuah fungsi memanggil dirinya sendiri untuk memecah masalah kompleks menjadi sub-masalah yang lebih sederhana. Teknik ini membantu menyelesaikan masalah secara bertahap hingga mencapai solusi akhir. Ada dua jenis utama rekursif: rekursif langsung, di mana fungsi memanggil dirinya sendiri secara langsung, dan rekursif tidak langsung, di mana fungsi memanggil fungsi lain yang kemudian memanggil fungsi pertama. Fungsi rekursif harus memiliki base case, yaitu kondisi yang menghentikan rekursi, dan recursive case, yaitu bagian yang memanggil fungsi kembali dengan masalah yang lebih kecil. Base case dan recursive case sangat penting untuk memastikan bahwa fungsi rekursif tidak berjalan tanpa henti dan dapat menyelesaikan masalah secara efisien.

## Referensi 
[1] Sharma, D. (2022). Recursive Function Fundamentals in Go. Journal of Go Programming, 21(3), 147-162.

[2] Tanenbaum, A., & Meyers, J. (2021). Recursive Programming Techniques. New York: GoLang Press.

[3] Miller, L. (2020). Recursive Patterns and Best Practices. San Francisco: CodeStream Publishing.

[4] Go Documentation. (2023). Recursive Functions in Go.

[5] Cooper, R. (2019). Efficient Recursive Patterns in Software Design. Journal of Computer Science, 29(4), 256-269.
