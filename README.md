# Sistem penilaian reviewer lomba poster

## Deskripsi
Sistem Penilaian Reviewer Lomba Poster adalah sebuah aplikasi berbasis konsol yang dirancang untuk membantu panitia lomba dalam mengelola dan merekapitulasi penilaian poster secara efisien dan akurat. Sistem ini memungkinkan reviewer (juri) untuk memberikan penilaian terhadap masing-masing poster peserta berdasarkan beberapa kriteria tertentu, seperti kreativitas, kesesuaian tema, dan estetika visual.

## fitur-fitur
- Input Data: Menyimpan data peserta dan skor penilaian dari beberapa reviewer.

- Pencarian Data: Mencari peserta berdasarkan ID atau nama menggunakan metode pencarian seperti binary search.

- Pengurutan Data: Mengurutkan data peserta berdasarkan nilai rata-rata atau nama menggunakan algoritma selection sort dan insertion sort.

- Statistik Nilai: Menghitung rata-rata nilai setiap peserta dan menentukan peserta dengan nilai tertinggi.

- CRUD (Create, Read, Update, Delete): Memungkinkan pengguna untuk menambahkan, melihat, mengubah, dan menghapus data peserta.

- Tampilan Terstruktur: Menyajikan data dalam format tabel agar mudah dibaca dan dipahami.

# struktur program

1. Konstanta dan Tipe Data
- const NMAX = 100
Maksimum jumlah peserta.

- type Peserta struct
Tipe bentukan untuk menyimpan:

- ID peserta

- Nama peserta

- Nilai dari dua reviewer

2. Variabel Global
- var dataPeserta [NMAX]Peserta
Array statis untuk menyimpan data peserta.

- var jumlahPeserta int
Menyimpan jumlah data yang telah dimasukkan.

3. Fungsi dan Prosedur
CRUD:
- tambahPeserta(p Peserta)
Menambahkan data peserta ke array.

- tampilkanPeserta()
Menampilkan semua data peserta beserta rata-rata.

- updatePeserta(id string)
Memperbarui data peserta berdasarkan ID.

- hapusPeserta(id string)
Menghapus data peserta berdasarkan ID.

Utilitas Penilaian:
- rataRata(p Peserta) float64
Menghitung rata-rata dua nilai peserta.

- nilaiTertinggi() float64
Mengembalikan nilai rata-rata tertinggi.

- nilaiTerendah() float64
Mengembalikan nilai rata-rata terendah.

Pencarian:
- cariPesertaByID(id string) int
Binary search untuk mencari peserta berdasarkan ID. Catatan: data harus terurut berdasarkan ID.

Pengurutan:
- urutkanBerdasarkanRataRata(naik bool)
Mengurutkan data menggunakan selection sort berdasarkan nilai rata-rata.

- tampilkanPengurutan(ascending bool)
Menampilkan data peserta yang telah diurutkan menggunakan insertion sort (berdasarkan rata-rata).

- urutkanBerdasarkanID()
Mengurutkan data peserta berdasarkan ID menggunakan insertion sort.

4. Menu Utama (main function)
Menu interaktif berbasis konsol:

-Tambah Peserta

- Tampilkan Peserta

- Edit Peserta

- Hapus Peserta

- Cari Peserta

- Urutkan (Rata-rata) Naik

- Urutkan (Rata-rata) Turun

- Tampilkan Nilai Tertinggi & Terendah

- Tampilkan Peserta Terurut (Ascending)

- Tampilkan Peserta Terurut (Descending)

- Urutkan Berdasarkan ID

- Keluar

5. Alur Umum Program
- Program dijalankan → menu ditampilkan

- Pengguna memilih opsi → fungsi terkait dipanggil

- Data disimpan dalam array statis (dataPeserta)

- Pengurutan otomatis setelah penambahan dilakukan (berdasarkan ID)
