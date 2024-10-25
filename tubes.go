package main

import (
	"fmt"
)

const NMAX = 7

type Dokter struct {
	Nama      string
	DokterId  int
	Spesialis string
}

type TabDokter [NMAX]Dokter

type Janji struct {
	Tanggal string
	Jam     string
	Pasien  string
	Dokter  int
}

type TabJanji [NMAX]Janji

func main() {
	var daftarDokter TabDokter
	var nDataDokter int
	var daftarJanji TabJanji
	var nDataJanji int

	for {
		fmt.Println(" ")
		fmt.Println("===============================================")
		fmt.Println("*** JADWAL DOKTER RUMAH SAKIT HARAPAN KITA  ***")
		fmt.Println("===============================================")
		fmt.Println("1. Tambah Dokter")
		fmt.Println("2. Lihat Daftar Dokter")
		fmt.Println("3. Buat Janji")
		fmt.Println("4. Lihat Daftar Janji")
		fmt.Println("5. Hapus Janji")
		fmt.Println("6. Cari Dokter")
		fmt.Println("7. Keluar")
		fmt.Println("===============================================")
		fmt.Println(" ")
		var pilihan int
		fmt.Print("Pilihan: ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			tambahDokter(&daftarDokter, &nDataDokter)
		case 2:
			fmt.Println(" ")
			fmt.Println("===============================================")
			fmt.Println("1. Urutkan berdasarkan ID")
			fmt.Println("2. Urutkan berdasarkan Nama")
			fmt.Println("===============================================")
			fmt.Println(" ")
			var pilihan int
			fmt.Print("Pilihan: ")
			fmt.Scanln(&pilihan)
			switch pilihan {
			case 1:
				sortirDaftarDokterById(&daftarDokter, nDataDokter)
				lihatDaftarDokter(&daftarDokter, nDataDokter)
			case 2:
				sortirDaftarDokterByName(&daftarDokter, nDataDokter)
				lihatDaftarDokter(&daftarDokter, nDataDokter)
			default:
				fmt.Println("Pilihan tidak valid.")
			}
		case 3:
			buatJanji(&daftarDokter, &daftarJanji, &nDataDokter, &nDataJanji)
		case 4:
			lihatDaftarJanji(&daftarJanji, nDataJanji, &daftarDokter)
		case 5:
			hapusJanji(&daftarJanji, &nDataJanji)
		case 6:
			fmt.Println(" ")
			fmt.Println("===============================================")
			fmt.Println("Cari Dokter oleh ID")
			fmt.Println("===============================================")
			fmt.Println(" ")
			var dokterId int
			fmt.Print("Masukkan ID dokter: ")
			fmt.Scanln(&dokterId)
			cariDokterById(&daftarDokter, nDataDokter, dokterId)
		case 7:
			keluar()
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahDokter(daftarDokter *TabDokter, nDataDokter *int) {
	fmt.Print("Masukkan ID dokter: ")
	var dokterId int
	fmt.Scanln(&dokterId)
	fmt.Print("Masukkan nama depan dokter: ")
	var firstName string
	fmt.Scanln(&firstName)
	fmt.Print("Masukkan nama belakang dokter: ")
	var lastName string
	fmt.Scanln(&lastName)
	fmt.Print("Masukkan spesialis dokter: ")
	var spesialis string
	fmt.Scanln(&spesialis)

	(*daftarDokter)[*nDataDokter] = Dokter{DokterId: dokterId, Nama: firstName + " " + lastName, Spesialis: spesialis}
	*nDataDokter++
	fmt.Println("Dokter berhasil ditambahkan.")
}

func sortirDaftarDokterByName(daftarDokter *TabDokter, nDataDokter int) {
	//Mengembalikan daftar dokter yang sudah di urutkan berdasarkan nama dengan menggunakan algoritma selection sort
	for i := 0; i < nDataDokter-1; i++ {
		indexMin := i
		for j := i + 1; j < nDataDokter; j++ {
			if (*daftarDokter)[j].Nama < (*daftarDokter)[indexMin].Nama {
				indexMin = j
			} else if (*daftarDokter)[j].Nama == (*daftarDokter)[indexMin].Nama {
				if (*daftarDokter)[j].Spesialis < (*daftarDokter)[indexMin].Spesialis {
					indexMin = j
				}
			}
		}
		if indexMin != i {
			temp := (*daftarDokter)[i]
			(*daftarDokter)[i] = (*daftarDokter)[indexMin]
			(*daftarDokter)[indexMin] = temp
		}
	}
}

func sortirDaftarDokterById(daftarDokter *TabDokter, nDataDokter int) {
	//Mengembalikan daftar dokter yang sudah di urutkan berdasarkan ID dengsfan menggunakan algoritma selection sort
	for i := 0; i < nDataDokter-1; i++ {
		indexMin := i
		for j := i + 1; j < nDataDokter; j++ {
			if (*daftarDokter)[j].DokterId < (*daftarDokter)[indexMin].DokterId {
				indexMin = j
			}
		}
		if indexMin != i {
			temp := (*daftarDokter)[i]
			(*daftarDokter)[i] = (*daftarDokter)[indexMin]
			(*daftarDokter)[indexMin] = temp
		}
	}
}

func lihatDaftarDokter(daftarDokter *TabDokter, nDataDokter int) {
	if nDataDokter == 0 {
		fmt.Println("Tidak ada dokter yang terdaftar.")
	} else {
		fmt.Println("Daftar Dokter:")
		for i := 0; i < nDataDokter; i++ {
			fmt.Printf("%d. %d. %s%s, %s\n", i+1, (*daftarDokter)[i].DokterId, "dr.", (*daftarDokter)[i].Nama, (*daftarDokter)[i].Spesialis)
		}
	}
}

func buatJanji(daftarDokter *TabDokter, daftarJanji *TabJanji, nDataDokter *int, nDataJanji *int) {
	fmt.Print("Masukkan tanggal (format: YYYY-MM-DD): ")
	var tanggal string
	fmt.Scanln(&tanggal)
	fmt.Print("Masukkan jam (format: HH:MM): ")
	var jam string
	fmt.Scanln(&jam)
	fmt.Print("Masukkan nama pasien: ")
	var pasien string
	fmt.Scanln(&pasien)

	var dokterTersedia [NMAX]int
	var count int
	cariDokterYangTersedia(daftarDokter, daftarJanji, nDataDokter, nDataJanji, &dokterTersedia, &count)

	if count == 0 {
		fmt.Println("Tidak ada dokter yang tersedia.")
	} else {
		fmt.Println("Pilih dokter yang tersedia:")
		for i := 0; i < count; i++ {
			dokter := dokterTersedia[i]
			if dokter < *nDataDokter && (*daftarDokter)[dokter].Nama != "" {
				fmt.Printf("%d. Dokter %s\n", i+1, (*daftarDokter)[dokter].Nama)
			}
		}
		var dokterPilihan int
		fmt.Print("Masukkan pilihan dokter (1 sampai ", count, "): ")
		fmt.Scanln(&dokterPilihan)
		dokterPilihan--
		if dokterPilihan < 0 || dokterPilihan >= count {
			fmt.Println("Pilihan dokter tidak valid.")
		} else {
			fmt.Println("Dokter yang dipilih:", (*daftarDokter)[dokterTersedia[dokterPilihan]].Nama)
			(*daftarJanji)[*nDataJanji] = Janji{Tanggal: tanggal, Jam: jam, Pasien: pasien, Dokter: dokterTersedia[dokterPilihan]}
			*nDataJanji++
			fmt.Println("Janji berhasil dibuat")
		}
	}
}

func cariDokterYangTersedia(daftarDokter *TabDokter, daftarJanji *TabJanji, nDataDokter *int, nDataJanji *int, dokterTersedia *[NMAX]int, count *int) {
	//Mengembalikan daftar dokter yang tersedia
	*count = 0
	for i := 0; i < *nDataDokter; i++ {
		var found bool = false
		for j := 0; j < *nDataJanji; j++ {
			if (*daftarJanji)[j].Dokter == i {
				found = true
			}
		}
		if !found {
			(*dokterTersedia)[*count] = i
			(*count)++
		}
	}
}

func sortJanjiAscendingByDokterId(daftarJanji *TabJanji, nDataJanji int, daftarDokter *TabDokter) {
	for i := 0; i < nDataJanji-1; i++ {
		minIndex := i
		for j := i + 1; j < nDataJanji; j++ {
			if (*daftarJanji)[j].Dokter < (*daftarJanji)[minIndex].Dokter {
				minIndex = j
			}
		}
		if minIndex != i {
			// Swap janji
			tempJanji := (*daftarJanji)[i]
			(*daftarJanji)[i] = (*daftarJanji)[minIndex]
			(*daftarJanji)[minIndex] = tempJanji
		}
	}
}

func cariDokterById(daftarDokter *TabDokter, nDataDokter int, dokterId int) {
	low := 0
	high := nDataDokter - 1

	for low <= high {
		mid := (low + high) / 2
		if (*daftarDokter)[mid].DokterId == dokterId {
			fmt.Printf("Dokter dengan ID %d ditemukan: %s\n", dokterId, (*daftarDokter)[mid].Nama)
			return
		} else if (*daftarDokter)[mid].DokterId < dokterId {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Println("Dokter dengan ID tersebut tidak ditemukan.")
}

func lihatDaftarJanji(daftarJanji *TabJanji, nDataJanji int, daftarDokter *TabDokter) {
	sortJanjiAscendingByDokterId(daftarJanji, nDataJanji, daftarDokter)

	if nDataJanji == 0 {
		fmt.Println("Tidak ada janji yang terdaftar.")
	} else {
		fmt.Println("Daftar Janji:")
		for i := 0; i < nDataJanji; i++ {
			dokter := (*daftarDokter)[(*daftarJanji)[i].Dokter]
			fmt.Printf("%d. %s, %s, %s, dr. %s (ID: %d)\n", i+1, (*daftarJanji)[i].Tanggal, (*daftarJanji)[i].Jam, (*daftarJanji)[i].Pasien, dokter.Nama, dokter.DokterId)
		}
	}
}

func hapusJanji(daftarJanji *TabJanji, nDataJanji *int) {
	if *nDataJanji == 0 {
		fmt.Println("Tidak ada janji yang terdaftar.")
	} else {
		fmt.Print("Masukkan nomor janji yang akan dihapus (1 sampai ", *nDataJanji, "): ")
		var nomor int
		fmt.Scanln(&nomor)
		nomor--
		if nomor < 0 || nomor >= *nDataJanji {
			fmt.Println("Nomor janji tidak valid.")
		} else {
			for i := nomor; i < *nDataJanji-1; i++ {
				(*daftarJanji)[i] = (*daftarJanji)[i+1]
			}
			*nDataJanji--
			fmt.Println("Janji berhasil dihapus.")
		}
	}
}

func keluar() {
	fmt.Println("======================================================")
	fmt.Println("          *** PROGRAM AKAN DIHENTIKAN ***")
	fmt.Println("======================================================")
	fmt.Println("Terima kasih telah menggunakan program ini.")
}
