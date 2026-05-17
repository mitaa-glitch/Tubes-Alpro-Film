package main

import (
	"fmt"
	"strings"
)

func tambahFilm(judul, genre string, rating float64, status string) {

	if jumlahFilm >= MAKS {
		fmt.Printf("[Gagal] Tidak dapat menambah '%s'. Array penuh!\n", judul)
		return
	}

	for i := 0; i < jumlahFilm; i++ {
		if strings.EqualFold(daftarFilm[i].judul, judul) {
			fmt.Printf("[Gagal] Film dengan judul '%s' sudah ada!\n", judul)
			return
		}
	}

	statusLower := strings.ToLower(status)

	if statusLower == "belum" {
		rating = 0.0
	} else {
		if rating < 1.0 || rating > 10.0 {
			fmt.Printf("[Gagal] Film '%s' tidak valid. Rating harus 1.0 - 10.0 untuk film yang sudah ditonton!\n", judul)
			return
		}
	}

	daftarFilm[jumlahFilm] = Film{
		judul:  judul,
		genre:  genre,
		rating: rating,
		status: status,
	}
	jumlahFilm++
	fmt.Printf("[Sukses] Film '%s' berhasil ditambahkan.\n", judul)
}

func ubahFilm(judulLama string, judulBaru, genreBaru string, ratingBaru float64, statusBaru string) {
	for i := 0; i < jumlahFilm; i++ {
		if strings.EqualFold(daftarFilm[i].judul, judulLama) {

			if strings.ToLower(statusBaru) == "belum" {
				ratingBaru = 0.0
			} else if ratingBaru < 1.0 || ratingBaru > 10.0 {
				fmt.Println("[Gagal] Ubah data gagal. Rating harus 1.0 - 10.0!")
				return
			}

			daftarFilm[i].judul = judulBaru
			daftarFilm[i].genre = genreBaru
			daftarFilm[i].rating = ratingBaru
			daftarFilm[i].status = statusBaru
			fmt.Printf("[Sukses] Film '%s' berhasil diubah.\n", judulLama)
			return
		}
	}
	fmt.Printf("[Gagal] Film '%s' tidak ditemukan.\n", judulLama)
}

func hapusFilm(judul string) {
	indexDitemukan := -1

	for i := 0; i < jumlahFilm; i++ {
		if strings.EqualFold(daftarFilm[i].judul, judul) {
			indexDitemukan = i
			break
		}
	}

	if indexDitemukan == -1 {
		fmt.Printf("[Gagal] Film '%s' tidak ditemukan.\n", judul)
		return
	}

	for i := indexDitemukan; i < jumlahFilm-1; i++ {
		daftarFilm[i] = daftarFilm[i+1]
	}

	daftarFilm[jumlahFilm-1] = Film{}
	jumlahFilm--

	fmt.Printf("[Sukses] Film '%s' berhasil dihapus.\n", judul)
}

func sequentialByGenre(genre string) {
	fmt.Printf("\n--- Hasil Pencarian Genre: %s ---\n", genre)
	found := false

	for i := 0; i < jumlahFilm; i++ {
		if strings.EqualFold(daftarFilm[i].genre, genre) {
			fmt.Printf("- %s (Rating: %.1f | Status: %s)\n", daftarFilm[i].judul, daftarFilm[i].rating, daftarFilm[i].status)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada film dengan genre tersebut.")
	}
}

func sequentialByRating(minRating float64) {
	fmt.Printf("\n--- Hasil Pencarian Rating >= %.1f ---\n", minRating)
	found := false

	for i := 0; i < jumlahFilm; i++ {
		if daftarFilm[i].rating >= minRating {
			fmt.Printf("- %s (Genre: %s | Rating: %.1f)\n", daftarFilm[i].judul, daftarFilm[i].genre, daftarFilm[i].rating)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada film dengan kriteria rating tersebut.")
	}
}
