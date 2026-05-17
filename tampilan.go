package main

import "fmt"

func menuUtama() {
	fmt.Println("\n══════════════════════════")
	fmt.Println("        🎬 FILMLOG")
	fmt.Println("══════════════════════════")
	fmt.Println("  1. 🎥 Lihat Semua Film")
	fmt.Println("  2. 🍿 Tambah Film")
	fmt.Println("  3. 📽️  Ubah Film")
	fmt.Println("  4. ✂️  Hapus Film")
	fmt.Println("  5. 🔍 Cari Film")
	fmt.Println("  6. ⭐ Urutkan Film")
	fmt.Println("  0. 🚪 Keluar")
	fmt.Println("══════════════════════════")
	fmt.Print("Pilih: ")
}

func tampilSemua() {
	fmt.Println("\n══════════════════════════════════════════════════════")
	fmt.Println("                   🎥 DAFTAR FILM")
	fmt.Println("══════════════════════════════════════════════════════")

	if jumlahFilm == 0 {
		fmt.Println("  Belum ada film.")
		fmt.Println("══════════════════════════════════════════════════════")
		return
	}

	fmt.Println(" No ║ Judul                ║ Genre    ║ Rating ║ Status")
	fmt.Println("────╫──────────────────────╫──────────╫────────╫──────────")

	for i := 0; i < jumlahFilm; i++ {
		f := daftarFilm[i]

		rating := fmt.Sprintf("%.1f", f.rating)
		if f.rating == 0 {
			rating = "-"
		}

		fmt.Printf(" %-3d║ %-21s║ %-9s║ %-7s║ %s\n",
			i+1, f.judul, f.genre, rating, f.status)
	}

	fmt.Println("══════════════════════════════════════════════════════")
}

func tampilDetail(idx int) {
	if idx < 0 || idx >= jumlahFilm {
		fmt.Println("[Error] Index film tidak valid.")
		return
	}

	f := daftarFilm[idx]

	rating := fmt.Sprintf("%.1f", f.rating)
	if f.rating == 0 {
		rating = "-"
	}

	status := f.status
	if f.status == "ditonton" {
		status = "✅ ditonton"
	} else {
		status = "🕐 belum"
	}

	fmt.Println("\n══════════════════════════════════════")
	fmt.Println("           🎬 DETAIL FILM")
	fmt.Println("══════════════════════════════════════")
	fmt.Printf(" %-10s ║ %s\n", "Judul", f.judul)
	fmt.Printf(" %-10s ║ %s\n", "Genre", f.genre)
	fmt.Printf(" %-10s ║ %s\n", "Rating", rating)
	fmt.Printf(" %-10s ║ %s\n", "Status", status)
	fmt.Println("══════════════════════════════════════")
}
