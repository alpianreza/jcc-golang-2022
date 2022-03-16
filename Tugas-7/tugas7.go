package main

import (
	"fmt"
)

// struct soal 2

type segitiga struct {
	alas, tinggi int
}
type persegi struct {
	sisi int
}
type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) luas() int {
	return s.alas * s.tinggi / 2
}
func (p persegi) luas() int {
	return p.sisi * p.sisi
}
func (pj persegiPanjang) luas() int {
	return pj.panjang * pj.lebar
}

// struct soal 3

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (hp *phone) addColor(color string) {
	hp.colors = append(hp.colors, color)
}

// func soal 4
type movie struct {
	title, genre   string
	year, duration int
}

func main() {
	// soal 1
	type buah struct {
		nama, warna string
		adaBijinya  bool
		harga       int
	}
	buahan := [...]buah{
		{"Nanas", "Kuning", false, 9000},
		{"Jeruk", "Oranye", true, 8000},
		{"Semangka", "Hijau & Merah", true, 10000},
		{"Pisang", "Kuning", false, 5000},
	}
	for _, item := range buahan {
		fmt.Println(item)
	}
	fmt.Println()

	// soal 2

	iSegitiga := segitiga{alas: 10, tinggi: 5}
	iPersegi := persegi{sisi: 4}
	iPersegiPanjang := persegiPanjang{panjang: 15, lebar: 5}

	fmt.Println("Luas segitiga =", iSegitiga.luas())
	fmt.Println("Luas persegi =", iPersegi.luas())
	fmt.Println("Luas persegi panjang =", iPersegiPanjang.luas())
	fmt.Println()

	// soal 3

	ponsel := phone{"Redmi Note", "Xioami", 2020, []string{}}
	ponsel.addColor("Biru")
	fmt.Println(ponsel)
	fmt.Println()

	//soal 4
	var dataFilm = []movie{}
	tambahDataFilm := func(title string, duration int, genre string, year int, df *[]movie) {
		data := movie{title, genre, duration, year}
		*df = append(*df, data)
	}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	i := 1
	for _, item := range dataFilm {
		fmt.Print(i, " ")
		fmt.Println("title:", item.title)
		fmt.Println(" ", "duration:", item.duration)
		fmt.Println(" ", "genre:", item.genre)
		fmt.Println(" ", "year:", item.year)
		i++
	}
}
