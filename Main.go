package main

import (
	"fmt"
	"strings"
)

func main() {
	var pilihan int

	for {
		menuUtama()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilSemua()
		case 2:
			menuTambah()
		case 3:
			menuUbah()
		case 4:
			menuHapus()
		case 5:
			menuCari()
		case 6:
			menuSort()
		case 9:
			runTests()
		case 0:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuTambah() {
	var judul, genre, status string
	var rating float64

	fmt.Print("Judul  : ")
	fmt.Scan(&judul)
	fmt.Print("Genre  : ")
	fmt.Scan(&genre)
	fmt.Print("Status (ditonton/belum): ")
	fmt.Scan(&status)

	if strings.ToLower(status) == "ditonton" {
		fmt.Print("Rating (1.0-10.0): ")
		fmt.Scan(&rating)
	}

	tambahFilm(judul, genre, rating, status)
}

func menuUbah() {
	var judulLama, judulBaru, genreBaru, statusBaru string
	var ratingBaru float64

	fmt.Print("Judul film yang ingin diubah: ")
	fmt.Scan(&judulLama)
	fmt.Print("Judul baru : ")
	fmt.Scan(&judulBaru)
	fmt.Print("Genre baru : ")
	fmt.Scan(&genreBaru)
	fmt.Print("Status baru (ditonton/belum): ")
	fmt.Scan(&statusBaru)

	if strings.ToLower(statusBaru) == "ditonton" {
		fmt.Print("Rating baru (1.0-10.0): ")
		fmt.Scan(&ratingBaru)
	}

	ubahFilm(judulLama, judulBaru, genreBaru, ratingBaru, statusBaru)
}

func menuHapus() {
	var judul string
	fmt.Print("Judul film yang ingin dihapus: ")
	fmt.Scan(&judul)
	hapusFilm(judul)
}

func menuCari() {
	var pilihan int
	fmt.Println("\n=== CARI FILM ===")
	fmt.Println("1. Cari by Judul (Binary Search)")
	fmt.Println("2. Cari by Genre (Sequential Search)")
	fmt.Println("3. Cari by Rating Minimum")
	fmt.Println("4. Cari Kombinasi (Genre + Rating)")
	fmt.Println("0. Kembali")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		var judul string
		fmt.Print("Masukkan judul: ")
		fmt.Scan(&judul)
		idx := binarySearchJudul(judul)
		if idx == -1 {
			fmt.Println("Film tidak ditemukan.")
		} else {
			tampilDetail(idx)
		}
	case 2:
		var genre string
		fmt.Print("Masukkan genre: ")
		fmt.Scan(&genre)
		sequentialByGenre(genre)
	case 3:
		var rating float64
		fmt.Print("Rating minimum: ")
		fmt.Scan(&rating)
		sequentialByRating(rating)
	case 4:
		cariKombinasi()
	}
}

func menuSort() {
	var byPilihan, metodePilihan, urutanPilihan int

	fmt.Println("\n=== URUTKAN FILM ===")
	fmt.Println("Urutkan by:  1. Judul  2. Rating")
	fmt.Scan(&byPilihan)

	fmt.Println("Metode:      1. Selection Sort  2. Insertion Sort")
	fmt.Scan(&metodePilihan)

	fmt.Println("Urutan:      1. Ascending  2. Descending")
	fmt.Scan(&urutanPilihan)

	switch {
	case byPilihan == 1 && metodePilihan == 1 && urutanPilihan == 1:
		selectionSortJudulAsc()
	case byPilihan == 1 && metodePilihan == 1 && urutanPilihan == 2:
		selectionSortJudulDesc()
	case byPilihan == 1 && metodePilihan == 2 && urutanPilihan == 1:
		insertionSortJudulAsc()
	case byPilihan == 1 && metodePilihan == 2 && urutanPilihan == 2:
		insertionSortJudulDesc()
	case byPilihan == 2 && metodePilihan == 1 && urutanPilihan == 1:
		selectionSortRatingAsc()
	case byPilihan == 2 && metodePilihan == 1 && urutanPilihan == 2:
		selectionSortRatingDesc()
	case byPilihan == 2 && metodePilihan == 2 && urutanPilihan == 1:
		insertionSortRatingAsc()
	case byPilihan == 2 && metodePilihan == 2 && urutanPilihan == 2:
		insertionSortRatingDesc()
	}

	fmt.Println("Film berhasil diurutkan!")
	tampilSemua()
}
