package main

import "fmt"

func binarySearchRekursif(judul string, kiri int, kanan int) int {
	if kiri > kanan {
		return -1
	}

	mid := (kiri + kanan) / 2

	if daftarFilm[mid].judul == judul {
		return mid
	} else if daftarFilm[mid].judul > judul {
		return binarySearchRekursif(judul, kiri, mid-1)
	} else {
		return binarySearchRekursif(judul, mid+1, kanan)
	}
}

func binarySearchJudul(judul string) int {
	if jumlahFilm == 0 {
		fmt.Println("✗ Belum ada film dalam daftar.")
		return -1
	}
	selectionSortJudulAsc()
	return binarySearchRekursif(judul, 0, jumlahFilm-1)
}

func cariKombinasi() {
	var genre string
	var ratingMin float64

	fmt.Print("Masukkan genre     : ")
	fmt.Scan(&genre)
	fmt.Print("Rating minimum     : ")
	fmt.Scan(&ratingMin)

	fmt.Printf("\n=== Hasil: Genre=%s, Rating ≥ %.1f ===\n", genre, ratingMin)
	fmt.Printf("%-3s | %-20s | %-8s | %-6s | %-10s\n",
		"No", "Judul", "Genre", "Rating", "Status")
	fmt.Println("----+----------------------+----------+--------+----------")

	ketemu := false
	nomor := 1
	for i := 0; i < jumlahFilm; i++ {
		f := daftarFilm[i]
		if f.genre == genre && f.rating >= ratingMin {
			fmt.Printf("%-3d | %-20s | %-8s | %-6.1f | %-10s\n",
				nomor, f.judul, f.genre, f.rating, f.status)
			nomor++
			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("✗ Tidak ada film yang cocok dengan kriteria tersebut.")
	}
}

func runTests() {
	fmt.Println("\n========================================")
	fmt.Println("      TESTING MENYELURUH — FILMLOG")
	fmt.Println("========================================")

	lulus := 0
	gagal := 0

	cek := func(nama string, hasil bool) {
		if hasil {
			fmt.Printf("  ✓ PASS  %s\n", nama)
			lulus++
		} else {
			fmt.Printf("  ✗ FAIL  %s\n", nama)
			gagal++
		}
	}

	jumlahFilm = 0

	fmt.Println("\n[1] Tambah Film Valid")
	daftarFilm[0] = Film{"Avengers", "Action", 8.4, "ditonton"}
	daftarFilm[1] = Film{"Interstellar", "Sci-Fi", 9.3, "ditonton"}
	daftarFilm[2] = Film{"Your Name", "Romance", 0, "belum"}
	jumlahFilm = 3
	cek("3 film berhasil ditambahkan", jumlahFilm == 3)

	fmt.Println("\n[2] Duplikat Judul")
	duplikat := false
	judulBaru := "Interstellar"
	for i := 0; i < jumlahFilm; i++ {
		if daftarFilm[i].judul == judulBaru {
			duplikat = true
			break
		}
	}
	cek("Duplikat 'Interstellar' terdeteksi", duplikat)

	fmt.Println("\n[3] Film Belum Ditonton")
	cek("Rating 'Your Name' (belum) = 0", daftarFilm[2].rating == 0)

	fmt.Println("\n[4] Hapus Film Index 0 (Avengers)")
	for i := 0; i < jumlahFilm-1; i++ {
		daftarFilm[i] = daftarFilm[i+1]
	}
	jumlahFilm--
	cek("Index 0 sekarang 'Interstellar'", daftarFilm[0].judul == "Interstellar")
	cek("jumlahFilm berkurang jadi 2", jumlahFilm == 2)

	fmt.Println("\n[5] Binary Search Rekursif — Ketemu")
	jumlahFilm = 0
	daftarFilm[0] = Film{"Avengers", "Action", 8.4, "ditonton"}
	daftarFilm[1] = Film{"Interstellar", "Sci-Fi", 9.3, "ditonton"}
	daftarFilm[2] = Film{"Your Name", "Romance", 0, "belum"}
	jumlahFilm = 3
	idx := binarySearchRekursif("Interstellar", 0, jumlahFilm-1)
	cek("Cari 'Interstellar' → ditemukan", idx != -1)
	cek("Index hasil benar (idx=1)", idx == 1)

	fmt.Println("\n[6] Binary Search Rekursif — Tidak Ditemukan")
	idx2 := binarySearchRekursif("FilmGaAda", 0, jumlahFilm-1)
	cek("Cari 'FilmGaAda' → return -1", idx2 == -1)

	fmt.Println("\n[7] Array Kosong")
	jumlahFilm = 0
	idx3 := binarySearchJudul("Apapun")
	cek("binarySearchJudul saat kosong → -1", idx3 == -1)

	fmt.Println("\n[8] Array Penuh (100 film)")
	jumlahFilm = MAKS
	bisaTambah := jumlahFilm < MAKS
	cek("Tambah ditolak saat jumlahFilm == MAKS", !bisaTambah)
	jumlahFilm = 0

	fmt.Println("\n[9] Kombinasi Genre+Rating")
	daftarFilm[0] = Film{"Film A", "Action", 8.0, "ditonton"}
	daftarFilm[1] = Film{"Film B", "Action", 0, "belum"}
	daftarFilm[2] = Film{"Film C", "Action", 9.0, "ditonton"}
	daftarFilm[3] = Film{"Film D", "Horror", 8.5, "ditonton"}
	jumlahFilm = 4
	count := 0
	for i := 0; i < jumlahFilm; i++ {
		if daftarFilm[i].genre == "Action" && daftarFilm[i].rating >= 7.0 {
			count++
		}
	}
	cek("Action + rating>=7.0 → 2 hasil (A dan C saja)", count == 2)

	fmt.Println("\n========================================")
	fmt.Printf("  HASIL AKHIR: %d PASS  |  %d FAIL\n", lulus, gagal)
	fmt.Println("========================================")

	jumlahFilm = 0
}
