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
    didefinisikan sebagai sebuah tipe baru, sehingga deklarasi structure tidak perlu lagi seluruh field-nya ditulis ulang berkali-kali.
    ![image](https://github.com/user-attachments/assets/afbe3b7c-6617-4765-9a43-459f408d9f63)

### 2. Array
Array mempunyai ukuran (jumlah elemen) yang tetap (statis) selama eksekusi program, sehingga jumlah elemen array menjadi bagian dari deklarasi variabel dengan tipe array.<br/>
![image](https://github.com/user-attachments/assets/bf257fb1-0618-4938-adfc-f4421d286980)
Jumlah elemen array dapat diminta dengan fungsi len yang tersedia. Sebagai contoh len(arr) akan menghasilkan 73 untuk contoh di atas.<br/>
Indeks array dimulai dari 0, sehingga indeks array pada contoh adalah 0, 1.. len(arr)-1<br/>
Contoh:<br/>
![image](https://github.com/user-attachments/assets/9d6a3239-2ae2-4fcb-b98f-adf08c23df36)<br/>
    - Slice (Array Dinamik)
    Array dalam Go juga dapat mempunyai ukuran yang dinamik. (Tidak digunakan di kelas Algoritma Pemrograman). Deklarasinya mirip dengan deklarasi array, tetapi
    jumlah elemennya dikosongkan.<br/>
    ![image](https://github.com/user-attachments/assets/fd8d3f2a-0494-429f-82aa-f1c42870e8d0)
    Sebuah slice dapat direalokasi menggunakan fungsi built-in make:
    ![image](https://github.com/user-attachments/assets/eccdd951-b904-4162-a9de-86b03439c704)
    Fungsi built-in len dapat digunakan untuk mengetahui ukuran slice. Fungsi lain, cap, dapat digunakan untuk mengetahui total tempat yang disediakan untuk slice
    tersebut.<br/>
    ![image](https://github.com/user-attachments/assets/f910cb7c-0062-45a2-b702-0113e6e81fde)
    Fungsi built-in append dapat digunakan untuk menambahkan elemen ke suatu slice, dan bila perlu memperbesar tempat untuk slice tersebut.
    ![image](https://github.com/user-attachments/assets/f36ecae1-0f8b-4a6a-8a48-18ab7b97a883)
    Sebuah slice baru juga dapat terbentuk dengan mengambil slice dari suatu array atau slice yang lain.
    ![image](https://github.com/user-attachments/assets/ab36eea8-d2ef-4bf5-9ad8-f3d9059baaad)
    - Map<br/>
    Tipe array lain, sebuah array dinamik. Indeksnya (di sini disebut kunci) tidak harus berbentuk integer. Indeks dapat berasal dari tipe apa saja. Struktur ini
    disebut map.
    ![image](https://github.com/user-attachments/assets/7916e324-fdc1-48f4-a304-0f805bd3c920)
    ![image](https://github.com/user-attachments/assets/7a9c1cdb-dc40-4083-ae3d-22210e26ebce)

## II. GUIDED











