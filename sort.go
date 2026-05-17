package main

// SELECTION SORT

// 1. By rating - ascending (terkecil duluan)
func selectionSortRatingAsc() {
	for i := 0; i < jumlahFilm-1; i++ {
		idxMin := i
		for j := i + 1; j < jumlahFilm; j++ {
			if daftarFilm[j].rating < daftarFilm[idxMin].rating {
				idxMin = j
			}
		}
		if idxMin != i {
			daftarFilm[i], daftarFilm[idxMin] = daftarFilm[idxMin], daftarFilm[i]
		}
	}
}

// 2. By rating - descending (terbesar duluan)
func selectionSortRatingDesc() {
	for i := 0; i < jumlahFilm-1; i++ {
		idxMax := i
		for j := i + 1; j < jumlahFilm; j++ {
			if daftarFilm[j].rating > daftarFilm[idxMax].rating {
				idxMax = j
			}
		}
		if idxMax != i {
			daftarFilm[i], daftarFilm[idxMax] = daftarFilm[idxMax], daftarFilm[i]
		}
	}
}

// 3. By judul - ascending (A → Z)
func selectionSortJudulAsc() {
	for i := 0; i < jumlahFilm-1; i++ {
		idxMin := i
		for j := i + 1; j < jumlahFilm; j++ {
			if daftarFilm[j].judul < daftarFilm[idxMin].judul {
				idxMin = j
			}
		}
		if idxMin != i {
			daftarFilm[i], daftarFilm[idxMin] = daftarFilm[idxMin], daftarFilm[i]
		}
	}
}

// 4. By judul - descending (Z → A)
func selectionSortJudulDesc() {
	for i := 0; i < jumlahFilm-1; i++ {
		idxMax := i
		for j := i + 1; j < jumlahFilm; j++ {
			if daftarFilm[j].judul > daftarFilm[idxMax].judul {
				idxMax = j
			}
		}
		if idxMax != i {
			daftarFilm[i], daftarFilm[idxMax] = daftarFilm[idxMax], daftarFilm[i]
		}
	}
}

// INSERTION SORT

// 1. By rating - ascending
func insertionSortRatingAsc() {
	for i := 1; i < jumlahFilm; i++ {
		key := daftarFilm[i]
		j := i - 1
		for j >= 0 && daftarFilm[j].rating > key.rating {
			daftarFilm[j+1] = daftarFilm[j]
			j--
		}
		daftarFilm[j+1] = key
	}
}

// 2. By rating - descending
func insertionSortRatingDesc() {
	for i := 1; i < jumlahFilm; i++ {
		key := daftarFilm[i]
		j := i - 1
		for j >= 0 && daftarFilm[j].rating < key.rating {
			daftarFilm[j+1] = daftarFilm[j]
			j--
		}
		daftarFilm[j+1] = key
	}
}

// 3. By judul - ascending (A → Z)
func insertionSortJudulAsc() {
	for i := 1; i < jumlahFilm; i++ {
		key := daftarFilm[i]
		j := i - 1
		for j >= 0 && daftarFilm[j].judul > key.judul {
			daftarFilm[j+1] = daftarFilm[j]
			j--
		}
		daftarFilm[j+1] = key
	}
}

// 4. By judul - descending (Z → A)
func insertionSortJudulDesc() {
	for i := 1; i < jumlahFilm; i++ {
		key := daftarFilm[i]
		j := i - 1
		for j >= 0 && daftarFilm[j].judul < key.judul {
			daftarFilm[j+1] = daftarFilm[j]
			j--
		}
		daftarFilm[j+1] = key
	}
}
