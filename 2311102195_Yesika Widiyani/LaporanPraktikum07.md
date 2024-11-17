<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL VI</strong></h2>
<h2 align="center"><strong> STRUCT DAN ARRAY </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Yesika Widiyani / 2311102195 <br>
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

# DASAR TEORI
## STRUCT DAN ARRAY

### STRUCT
#### Pengertian
Struct (atau struktur) adalah tipe data yang digunakan untuk mengelompokkan beberapa tipe data yang berbeda menjadi satu entitas logis. Struct memberikan cara untuk membuat tipe data kompleks dengan menggabungkan berbagai atribut (field).</br>

#### Karakteristik Struct
1. Heterogen: Field dalam struct dapat memiliki tipe data yang berbeda.</br>
2. Custom: Struct memungkinkan pengguna membuat tipe data baru yang sesuai dengan kebutuhan.</br>
3. Fleksibel: Struct dapat digunakan untuk merepresentasikan data yang kompleks, seperti objek dunia nyata (misalnya mahasiswa, mobil, dll).</br>

#### Kegunaan Struct
1. Mengorganisasikan data yang kompleks.</br>
2. Membuat tipe data yang lebih deskriptif dan terstruktur.</br>
3. Mempermudah manipulasi data yang memiliki banyak atribut.</br>

### ARRAY
#### Pengertian
Array adalah sebuah struktur data yang menyimpan kumpulan elemen dengan tipe data yang sama dalam satu variabel. Setiap elemen array diidentifikasi menggunakan indeks, yang dimulai dari 0 hingga panjang array - 1.</br>

#### Karakteristik Array
1. Homogen: Semua elemen dalam array harus memiliki tipe data yang sama.</br>
2. Indeks Tetap: Elemen dalam array diakses menggunakan indeks yang tetap dan tidak dapat diubah.</br>
3. Ukurannya Tetap: Ukuran array ditentukan pada saat deklarasi dan tidak dapat diubah selama runtime.</br>
4. Linear: Elemen dalam array diorganisasikan dalam urutan linier.</br>

#### Kegunaan Array
1. Menyimpan data yang sejenis secara efisien.</br>
2. Memungkinkan akses cepat ke elemen menggunakan indeks.</br>
3. Cocok digunakan untuk pemrosesan data yang memiliki jumlah elemen tetap.</br>

------

# GUIDED
## GUIDED - 1
### Study Case
1. Sebuah program yang menerima input waktu masuk parkir dan waktu keluar parkir dalam format jam, menit, dan detik. Program ini akan menghitung dan menampilkan lama waktu parkir dalam jam, menit, dan detik.</br>

### Source Code
![carbon](https://github.com/user-attachments/assets/e84c8a06-9ab7-4d2f-b5d5-afbd87c82fe0)

### Screenshot Output
![image](https://github.com/user-attachments/assets/404af871-390d-48f2-8f71-6443c16b0c12)

### Deskripsi Program
Program ini digunakan untuk menghitung durasi parkir kendaraan berdasarkan waktu kedatangan dan waktu pulang. Program memanfaatkan struct untuk merepresentasikan waktu dalam bentuk jam, menit, dan detik. Input berupa waktu kedatangan (wParkir) dan waktu pulang (wPulang), kemudian program akan menghitung selisih waktu tersebut sebagai lama parkir dalam format jam, menit, dan detik.</br>

**Algoritma Singkat**
1. **Deklarasi struct `waktu`** untuk menyimpan jam, menit, dan detik.  
2. Buat variabel `wParkir`, `wPulang` untuk input waktu, serta `durasi` untuk hasil selisih waktu.  
3. Input jam, menit, detik untuk waktu kedatangan (`wParkir`) dan pulang (`wPulang`).  
4. Konversi `wParkir` dan `wPulang` ke total detik:
   - Total Detik = (jam × 3600) + (menit × 60) + detik.  
5. Hitung selisih waktu dalam detik:  
   - `lParkir = dPulang - dParkir`.  
6. Konversi `lParkir` kembali ke jam, menit, detik:  
   - Jam = `lParkir / 3600`.  
   - Menit = `(lParkir % 3600) / 60`.  
   - Detik = `(lParkir % 3600) % 60`.  
7. Cetak hasil lama parkir dalam format jam, menit, detik.

**Cara Kerja Singkat**
1. Program menerima input waktu kedatangan dan pulang. </br>
2. Waktu dikonversi ke detik untuk mempermudah perhitungan. </br>
3. Selisih total detik dihitung untuk mendapatkan durasi parkir.</br> 
4. Selisih detik dikembalikan ke format jam, menit, dan detik. </br>
5. Hasil lama parkir ditampilkan ke layar. </br>

## GUIDED - 2
### Studi Case  
Sebuah toko buah memiliki daftar harga untuk beberapa jenis buah. Pemilik toko ingin menyimpan data harga buah dan menampilkannya ke pelanggan. Selain itu, pelanggan juga bisa meminta harga buah tertentu, misalnya "Mangga". </br>

### Source Code
![carbon (1)](https://github.com/user-attachments/assets/90285ae9-d165-4cc8-b631-3cde00030d5f)

### Screenshot Output
![image](https://github.com/user-attachments/assets/855b6c34-83ac-42f6-a260-25380d44c90b)

### Deskripsi Singkat 
Program menggunakan **map** untuk menyimpan nama buah sebagai kunci dan harga sebagai nilai. Map mempermudah pengelolaan data pasangan kunci-nilai, seperti pencarian harga berdasarkan nama buah.</br>

**Algoritma**  
1. Buat map `hargaBuah` dengan nama buah sebagai kunci dan harga sebagai nilai.</br>
2. Isi map dengan data harga beberapa buah (contoh: Apel, Pisang, Mangga).</br>  
3. Gunakan loop `for` untuk menampilkan seluruh harga buah yang tersedia.</br> 
4. Tampilkan harga spesifik buah tertentu (contoh: Mangga) menggunakan kunci dari map.</br>

**Cara Kerja**  
1. Map `hargaBuah` dibuat dan diisi dengan data buah beserta harganya.</br>  
2. Program mencetak semua data di dalam map menggunakan loop `for range`.</br>  
3. Harga spesifik buah tertentu dicetak langsung dengan mengakses nilai berdasarkan kunci (contoh: `hargaBuah["Mangga"]`).</br>

## GUIDED 3
### Studi case
Sebuah aplikasi ingin mengelola daftar nama teman. Saat menambahkan nama baru, aplikasi harus memastikan bahwa nama tersebut belum ada di daftar untuk menghindari duplikasi.</br>

### Source Code
![carbon (2)](https://github.com/user-attachments/assets/564cb99c-5ff8-48d2-9689-c1f62b2b44bb)

### Screenshot Output
![image](https://github.com/user-attachments/assets/7b2c1837-d672-4905-9627-8c1f3372c19e)

### Deskripsi Singkat
Program ini menggunakan **slice** untuk menyimpan daftar nama teman. Nama baru hanya ditambahkan jika tidak ada di dalam daftar, dengan memanfaatkan fungsi pengecekan sederhana.</br> 

**Algoritma**  
1. Buat slice awal `daftarTeman` berisi beberapa nama teman.</br> 
2. Siapkan slice `namaBaru` berisi nama-nama yang ingin ditambahkan.</br>   
3. Iterasi melalui `namaBaru`:</br> 
   - Cek apakah nama sudah ada di `daftarTeman` menggunakan fungsi `sudahAda`.</br> 
   - Jika nama belum ada, tambahkan ke `daftarTeman` dengan `append`.</br> 
   - Jika sudah ada, beri informasi bahwa nama tersebut sudah ada.</br>  
4. Tampilkan daftar teman akhir.</br> 

**Cara Kerja**  
1. Program dimulai dengan slice `daftarTeman` berisi data awal.</br>  
2. Fungsi `sudahAda` digunakan untuk memeriksa apakah nama sudah terdapat di slice.</br>  
3. Nama dari `namaBaru` yang belum ada ditambahkan menggunakan `append`.</br> 
4. Setelah iterasi selesai, daftar teman yang telah diperbarui ditampilkan.</br>

------

# UNGUIDED
## UNGUIDED - 1
### Study Case
1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.</br>
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".</br>


Fungsi untuk menghitung jarak titik (a, b) dan (c, d) dimana rumus jarak adalah:</br>
![image](https://github.com/user-attachments/assets/065e1989-390f-47c2-bbd2-ea7088126e49)

dan juga fungsi untuk menentukan posisi sebuah titik sembarang berada di dalam suatu lingkaran atau tidak.</br>
![image](https://github.com/user-attachments/assets/64942433-dde2-43a9-adcd-e2aecf5fd429)

### Source Code
![carbon (4)](https://github.com/user-attachments/assets/dabc536a-b574-4131-b8cc-77ac3460929b)

### Screenshot Output
![image](https://github.com/user-attachments/assets/4972a9f8-0364-46fe-9202-6d0626801297)

#### Deskripsi Program
Program ini mengevaluasi apakah sebuah titik berada di dalam lingkaran, di luar lingkaran, atau di kedua lingkaran sekaligus. Dua lingkaran didefinisikan oleh pusat koordinat `(cx, cy)` dan radius `r`. Program menggunakan fungsi jarak Euclidean untuk memeriksa posisi titik.

**Algoritma:**
1. **Input:**
   - Koordinat pusat dan radius untuk dua lingkaran.
   - Daftar koordinat titik-titik yang akan diuji.

2. **Proses:**
   - Untuk setiap titik dalam daftar:
     - Hitung jarak titik ke pusat lingkaran menggunakan rumus \( \text{jarak} = \sqrt{(x_2 - x_1)^2 + (y_2 - y_1)^2} \).
     - Periksa apakah jarak lebih kecil atau sama dengan radius lingkaran:
       - Jika ya, titik berada di dalam lingkaran.
       - Jika tidak, titik berada di luar lingkaran.
     - Evaluasi apakah titik berada di dalam salah satu, kedua lingkaran, atau tidak sama sekali.

3. **Output:**
   - Status posisi setiap titik, seperti "Titik di dalam lingkaran 1", "Titik di luar lingkaran 1 dan 2", dll.

**Cara Kerja:**
1. Program meminta input pusat dan radius untuk dua lingkaran.
2. Menggunakan fungsi `jarak`, program menghitung jarak Euclidean antara pusat lingkaran dan titik.
3. Fungsi `diDalamLingkaran` mengevaluasi posisi titik dengan membandingkan jarak dengan radius lingkaran.
4. Hasil posisi tiap titik ditampilkan berdasarkan kondisi evaluasi.

## GUIDED - 2
### Studi Case
2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut:
a. Menampilkan keseluruhan isi dari array.
b. Menampilkan elemen-elemen array dengan indeks ganjil saja.
c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-O adalah genap).
d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa dari masukan pengguna.
e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil
f. Menampilkan rata-rata dari bilangan yang ada di dalam array.
g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.
h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.

### Source Code
![carbon (5)](https://github.com/user-attachments/assets/b7c6d87a-4afc-43bf-b46a-d833dd69a1fe)

### Screenshot Output
![image](https://github.com/user-attachments/assets/60583ddd-73d4-4c0a-8929-118c22bd999d)

#### Deskripsi Program
Program ini melakukan berbagai operasi pada array, seperti menampilkan elemen berdasarkan kriteria tertentu (indeks genap, ganjil, kelipatan indeks), menghapus elemen tertentu, menghitung rata-rata, standar deviasi, dan frekuensi bilangan tertentu.

**Algoritma:**

1. **Input:**
   - Jumlah elemen array.
   - Elemen-elemen array.
   - Parameter tambahan seperti indeks, kelipatan, atau bilangan tertentu untuk operasi tertentu.

2. **Proses:**
   - **a.** Tampilkan semua elemen array.
   - **b.** Filter elemen berdasarkan indeks ganjil.
   - **c.** Filter elemen berdasarkan indeks genap.
   - **d.** Filter elemen berdasarkan kelipatan indeks yang diberikan.
   - **e.** Hapus elemen pada indeks tertentu:
     - Gabungkan elemen sebelum dan sesudah indeks menggunakan `append`.
   - **f.** Hitung rata-rata:
     - Jumlahkan semua elemen lalu bagi dengan panjang array.
   - **g.** Hitung standar deviasi:
     - Hitung variansi \( \sigma^2 = \frac{\sum{(x - \mu)^2}}{n} \).
     - Ambil akar kuadrat untuk mendapatkan standar deviasi \( \sigma \).
   - **h.** Hitung frekuensi:
     - Iterasi array untuk menghitung kemunculan bilangan tertentu.

3. **Output:**
   - Hasil operasi sesuai pilihan (misalnya, elemen ganjil, rata-rata, frekuensi, dll.).

**Cara Kerja:**
1. Program meminta input jumlah elemen dan nilai elemen array.
2. Program menyediakan opsi untuk berbagai operasi:
   - Menampilkan elemen berdasarkan kriteria.
   - Menghapus elemen di indeks tertentu.
   - Menghitung rata-rata atau standar deviasi.
   - Menghitung frekuensi bilangan tertentu.
3. Menggunakan perulangan, operasi matematis (seperti penjumlahan, pembagian, akar kuadrat), dan manipulasi array untuk menyelesaikan setiap permintaan.
4. Hasil setiap operasi ditampilkan ke layar.

## GUIDED - 3
### Studi Case
3) Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga.</br>
Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja.</br>

Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.
Perhatikan sesi interaksi pada contoh berikut ini </br>
![image](https://github.com/user-attachments/assets/0d04f520-a2fa-490e-848d-675ecbd89455)

### Source Code
![carbon (6)](https://github.com/user-attachments/assets/d47a35c8-ad9f-4b3c-9c48-9e29331cdbc6)

### Screenshot Output
![image](https://github.com/user-attachments/assets/862df858-2f9c-4067-a020-a827fad3a351)

### Deskripsi Program
Program ini merekam hasil pertandingan antara dua klub sepak bola berdasarkan skor. Jika salah satu skor negatif, program berhenti menerima input dan menampilkan daftar hasil pertandingan.

**Algoritma**
1. Input nama Klub A dan Klub B.  
2. Lakukan input skor kedua klub dalam loop:  
   - Jika skor salah satu klub negatif, akhiri loop.  
   - Tentukan pemenang berdasarkan skor:  
     - Klub dengan skor lebih tinggi dimasukkan ke daftar pemenang.  
     - Jika skor sama, tambahkan "Draw".  
3. Setelah loop, tampilkan daftar pemenang seluruh pertandingan.

**Cara Kerja**
1. Program menerima nama kedua klub.  
2. Skor pertandingan diinput dalam loop hingga ditemukan skor negatif.  
3. Setiap skor dibandingkan, pemenang atau hasil "Draw" disimpan dalam slice.  
4. Setelah semua skor diinput, hasil pertandingan ditampilkan secara berurutan.

## GUIDED - 4
### Studi Case

4) Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.</br>
Lengkapi potongan algoritma berikut ini!</br>
package main irport "fit"
![image](https://github.com/user-attachments/assets/7a1914c9-9fb1-476d-96e0-fef514a8dbfa)
![image](https://github.com/user-attachments/assets/d7c6cbc6-1029-4428-8d3f-cdbd8eb4fdb5)

![image](https://github.com/user-attachments/assets/f3fa2f02-775f-48a6-94cf-88257664a459)

Modifikasi program tersebut dengan menambahkan fungsi palindrom. Tambahkan instruksi untuk memanggil fungsi tersebut dan menampilkan hasilnya pada program utama.
*Palindrom adalah teks yang dibaca dari awal atau akhir adalah sama, contoh: KATAK, APA, KASUR RUSAK.

![image](https://github.com/user-attachments/assets/5f0f47a8-fa7d-426f-b6aa-3bd0fe62dcbd)

### Source Code
![carbon (7)](https://github.com/user-attachments/assets/66d38b4b-f0e9-4ca9-bff2-1dd554f10ecb)

### Screenshot Output
![image](https://github.com/user-attachments/assets/bbc0b9d9-fb92-4fea-9521-7ddc41db3344)

![image](https://github.com/user-attachments/assets/7068e52c-c10a-49e2-a633-f06da1917f57)

### Deskripsi Singkat
Program ini bertujuan untuk menerima sekumpulan karakter dari pengguna, menampilkan array asli, membalik urutannya, dan memeriksa apakah array tersebut merupakan palindrome. Palindrome adalah array yang jika dibaca dari depan atau belakang tetap sama.

**Algoritma**
1. *Input Data*: Program menerima jumlah karakter dan karakter itu sendiri.
2. *Menampilkan Array Asli*: Menampilkan karakter yang dimasukkan oleh pengguna.
3. *Membalik Array*: Membalik urutan elemen array menggunakan fungsi balikArray.
4. *Memeriksa Palindrome*: Mengecek apakah elemen array sama saat dibaca dari depan dan belakang menggunakan fungsi isArrayPalindrome.
5. *Output Hasil*:
   - Menampilkan array asli dan array yang sudah dibalik.
   - Memberi tahu apakah array tersebut adalah palindrome atau bukan.

**Cara Kerja**
1. Masukkan jumlah karakter dan karakter array.
2. Program memproses array dengan fungsi:
   - cetakArray: Untuk mencetak array ke layar.
   - balikArray: Untuk membalik elemen array.
   - isArrayPalindrome: Untuk memeriksa apakah array adalah palindrome.
3. Tampilkan hasil ke layar sesuai dengan operasi yang dilakukan.





