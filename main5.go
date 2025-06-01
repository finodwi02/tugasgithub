package main

import (
	"fmt"
)
// cons max menjadi NMAX
const NMAX = 100

type Peserta struct {
	Id     string
	Nama   string
	Nilai1 float64
	Nilai2 float64
}

var dataPeserta [NMAX]Peserta
var jumlahPeserta int = 0

// -------------------- FUNGSI & PROSEDUR --------------------

func tambahPeserta(p Peserta) {
	if jumlahPeserta < NMAX {
		dataPeserta[jumlahPeserta] = p
		jumlahPeserta++
	} else {
		fmt.Println("Data penuh!")
	}
}

func tampilkanPeserta() {
	fmt.Println("Data Peserta:")
	for i := 0; i < jumlahPeserta; i++ {
		rata2 := rataRata(dataPeserta[i])
		fmt.Printf("%d. ID: %s, Nama: %s, Nilai1: %.2f, Nilai2: %.2f, Rata-rata: %.2f\n",
			i+1, dataPeserta[i].Id, dataPeserta[i].Nama,
			dataPeserta[i].Nilai1, dataPeserta[i].Nilai2,
			rata2)
	}
}

// prosedur
func rataRata(p Peserta) float64 {
	return (p.Nilai1 + p.Nilai2) / 2
}

// Pencarian  binary search untuk peserta berdasarkan ID (data harus diurutkan)
func cariPesertaByID(id string) int {
	left := 0
	right := jumlahPeserta - 1

	for left <= right {
		mid := (left + right) / 2
		if dataPeserta[mid].Id == id {
			return mid
		} else if dataPeserta[mid].Id < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func updatePeserta(id string) {
	idx := cariPesertaByID(id)
	if idx != -1 {
		fmt.Print("Masukkan nama baru: ")
		fmt.Scan(&dataPeserta[idx].Nama)
		fmt.Print("Masukkan nilai reviewer 1 baru: ")
		fmt.Scan(&dataPeserta[idx].Nilai1)
		fmt.Print("Masukkan nilai reviewer 2 baru: ")
		fmt.Scan(&dataPeserta[idx].Nilai2)
		fmt.Println("Data berhasil diupdate.")
	} else {
		fmt.Println("Peserta tidak ditemukan.")
	}
}

func hapusPeserta(id string) {
	idx := cariPesertaByID(id)
	if idx != -1 {
		for i := idx; i < jumlahPeserta-1; i++ {
			dataPeserta[i] = dataPeserta[i+1]
		}
		jumlahPeserta--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Peserta tidak ditemukan.")
	}
}

// SSortir pilihan untuk mengurutkan berdasarkan rata-rata
func urutkanBerdasarkanRataRata(naik bool) {
	for i := 0; i < jumlahPeserta-1; i++ {
		extremeIdx := i
		for j := i + 1; j < jumlahPeserta; j++ {
			currentAvg := rataRata(dataPeserta[j])
			extremeAvg := rataRata(dataPeserta[extremeIdx])
			if (naik && currentAvg < extremeAvg) || (!naik && currentAvg > extremeAvg) {
				extremeIdx = j
			}
		}
		// Swap
		dataPeserta[i], dataPeserta[extremeIdx] = dataPeserta[extremeIdx], dataPeserta[i]
	}
}

// Insertion sort for displaying in ascending/descending order
func tampilkanPengurutan(ascending bool) {
	tempData := dataPeserta // Create a copy to avoid modifying original data

	for i := 1; i < jumlahPeserta; i++ {
		key := tempData[i]
		j := i - 1

		for j >= 0 {
			keyAvg := rataRata(key)
			currentAvg := rataRata(tempData[j])
			if (ascending && currentAvg > keyAvg) || (!ascending && currentAvg < keyAvg) {
				tempData[j+1] = tempData[j]
				j--
			} else {
				break
			}
		}
		tempData[j+1] = key
	}

	// Menampilkan data yang diurutkan
	fmt.Println("Data Peserta (Tertampil Terurut):")
	for i := 0; i < jumlahPeserta; i++ {
		rata2 := rataRata(tempData[i])
		fmt.Printf("%d. ID: %s, Nama: %s, Nilai1: %.2f, Nilai2: %.2f, Rata-rata: %.2f\n",
			i+1, tempData[i].Id, tempData[i].Nama,
			tempData[i].Nilai1, tempData[i].Nilai2,
			rata2)
	}
}

func nilaiTertinggi() float64 {
	max := rataRata(dataPeserta[0])
	for i := 1; i < jumlahPeserta; i++ {
		if rataRata(dataPeserta[i]) > max {
			max = rataRata(dataPeserta[i])
		}
	}
	return max
}

func nilaiTerendah() float64 {
	min := rataRata(dataPeserta[0])
	for i := 1; i < jumlahPeserta; i++ {
		if rataRata(dataPeserta[i]) < min {
			min = rataRata(dataPeserta[i])
		}
	}
	return min
}

// Fungsi baru untuk mengurutkan berdasarkan ID
func urutkanBerdasarkanID() {
	for i := 1; i < jumlahPeserta; i++ {
		key := dataPeserta[i]
		j := i - 1
		for j >= 0 && dataPeserta[j].Id > key.Id {
			dataPeserta[j+1] = dataPeserta[j]
			j--
		}
		dataPeserta[j+1] = key
	}
}

// -------------------- MENU --------------------

func main() {
	var pilihan int
	var id string
	for {
		fmt.Println("\n--- MENU PENILAIAN LOMBA POSTER ---")
		fmt.Println("1. Tambah Peserta")
		fmt.Println("2. Tampilkan Peserta")
		fmt.Println("3. Edit Peserta")
		fmt.Println("4. Hapus Peserta")
		fmt.Println("5. Cari Peserta")
		fmt.Println("6. Urutkan (Rata-rata) Naik")
		fmt.Println("7. Urutkan (Rata-rata) Turun")
		fmt.Println("8. Tampilkan Nilai Tertinggi & Terendah")
		fmt.Println("9. Tampilkan Peserta Terurut (Ascending)")
		fmt.Println("10. Tampilkan Peserta Terurut (Descending)")
		fmt.Println("11. Urutkan Peserta Berdasarkan ID")
		fmt.Println("12. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var p Peserta
			fmt.Print("Masukkan ID: ")
			fmt.Scan(&p.Id)
			fmt.Print("Masukkan Nama: ")
			fmt.Scan(&p.Nama)
			fmt.Print("Masukkan Nilai Reviewer 1: ")
			fmt.Scan(&p.Nilai1)
			fmt.Print("Masukkan Nilai Reviewer 2: ")
			fmt.Scan(&p.Nilai2)
			tambahPeserta(p)
			// Auto-sort by ID after adding
			urutkanBerdasarkanID()
		case 2:
			tampilkanPeserta()
		case 3:
			fmt.Print("Masukkan ID peserta yang ingin diedit: ")
			fmt.Scan(&id)
			updatePeserta(id)
		case 4:
			fmt.Print("Masukkan ID peserta yang ingin dihapus: ")
			fmt.Scan(&id)
			hapusPeserta(id)
		case 5:
			fmt.Print("Masukkan ID yang dicari: ")
			fmt.Scan(&id)
			idx := cariPesertaByID(id)
			if idx != -1 {
				p := dataPeserta[idx]
				fmt.Printf("Ditemukan: ID: %s, Nama: %s, Nilai1: %.2f, Nilai2: %.2f, Rata-rata: %.2f\n",
					p.Id, p.Nama, p.Nilai1, p.Nilai2, rataRata(p))
			} else {
				fmt.Println("Peserta tidak ditemukan.")
			}
		case 6:
			urutkanBerdasarkanRataRata(true)
			fmt.Println("Data berhasil diurutkan secara naik.")
		case 7:
			urutkanBerdasarkanRataRata(false)
			fmt.Println("Data berhasil diurutkan secara turun.")
		case 8:
			fmt.Printf("Nilai tertinggi: %.2f\n", nilaiTertinggi())
			fmt.Printf("Nilai terendah: %.2f\n", nilaiTerendah())
		case 9:
			tampilkanPengurutan(true)
		case 10:
			tampilkanPengurutan(false)
		case 11:
			urutkanBerdasarkanID()
			fmt.Println("Data berhasil diurutkan berdasarkan ID.")
		case 12:
			fmt.Println("Terima kasih. Keluar program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}