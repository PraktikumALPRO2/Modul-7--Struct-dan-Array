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
