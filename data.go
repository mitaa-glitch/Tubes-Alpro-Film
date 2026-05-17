package main

const MAKS = 100

type Film struct {
	judul  string
	genre  string
	rating float64
	status string
}

var daftarFilm [MAKS]Film
var jumlahFilm int = 0
