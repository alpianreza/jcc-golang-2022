package main

import (
	"fmt"
	"strconv"
)

// soal 1 dan 4 funcionnya di luar main

// func soal 1
func nilaiLuasLingkaran(pi float64, r float64) float64 {
	return pi * r * r
}

func nilaiKelilingLingkaran(pi float64, r float64) float64 {
	return pi * (r + r)
}
func nilaiLingkaran(pNilai *float64, nilai float64) {
	*pNilai = nilai
}

// func soal 4
func tambahDataFilm(dataFilm *[]map[string]string, judul, durasi, genre, tahun string) {
	var data = map[string]string{}
	data["title"] = judul
	data["duration"] = durasi
	data["genre"] = genre
	data["year"] = tahun
	*dataFilm = append(*dataFilm, data)
}

func main() {
	var luasLigkaran float64
	var kelilingLingkaran float64
	var pi float64 = 3.14
	var r float64 = 10

	luasLigkaran = nilaiLuasLingkaran(pi, r)
	kelilingLingkaran = nilaiKelilingLingkaran(pi, r)
	// sebelum pakai pointer
	fmt.Println("jari-jari = ", r)
	fmt.Println("Luas lingkaran = ", luasLigkaran)
	fmt.Println("Keliling lingkaran = ", kelilingLingkaran)
	fmt.Println()
	// sesudah pakai pointer
	nilaiLingkaran(&r, 5)
	luasLigkaran = nilaiLuasLingkaran(pi, r)
	kelilingLingkaran = nilaiKelilingLingkaran(pi, r)
	fmt.Println("Jari-jari = ", r)
	fmt.Println("Luas lingkaran baru = ", luasLigkaran)
	fmt.Println("Keliling lingkaran baru = ", kelilingLingkaran)
	fmt.Println()

	// soal 2
	var sentence string
	var introduce = func(hasil *string, name, gender, occupation, age string) string {
		var title = ""
		if gender == "laki-laki" {
			title = "Pak "
		} else {
			title = "Bu "
		}
		*hasil = title + "" + name + " adalah seorang " + occupation + " yang berusia " + age + " tahun"
		return *hasil
	}
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)
	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)
	fmt.Println()

	// soal 3
	var buah = []string{}
	var namabuah = &buah
	*namabuah = append(*namabuah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
	for i, fruits := range *namabuah {
		fmt.Println(strconv.Itoa(i+1) + ". " + fruits)
	}
	fmt.Println()

	// soal 4

	var dataFilm = []map[string]string{}

	tambahDataFilm(&dataFilm, "LOTR", "2 jam", "action", "1999")
	tambahDataFilm(&dataFilm, "avenger", "2 jam", "action", "2019")
	tambahDataFilm(&dataFilm, "spiderman", "2 jam", "action", "2004")
	tambahDataFilm(&dataFilm, "juon", "2 jam", "horror", "2004")

	i := 1
	for _, item := range dataFilm {
		fmt.Print(i, " ")
		fmt.Println("title:", item["title"])
		fmt.Println(" ", "duration:", item["duration"])
		fmt.Println(" ", "genre:", item["genre"])
		fmt.Println(" ", "year:", item["year"])
		i++
	}
}
