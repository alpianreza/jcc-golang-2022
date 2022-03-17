package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// struct iface func soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

func (pp persegiPanjang) keliling() int {
	return 2 * (pp.panjang + pp.lebar)
}

func (b balok) volume() float64 {
	return float64(b.lebar) * float64(b.panjang) * float64(b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return (2 * (float64(b.lebar) + float64(b.panjang))) +
		(2 * (float64(b.lebar) + float64(b.tinggi))) +
		(2 * (float64(b.tinggi) + float64(b.panjang)))
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

// struct iface func soal 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type smarthphone interface {
	phone() string
}

func (sp phone) phone() string {
	hasil := "name : " + sp.name + "\n" + "brand : " + sp.brand + "\n" + "year : " + strconv.Itoa(sp.year) + "\n" + "colors : " + strings.Join(sp.colors, ", ")

	return hasil
}

func main() {
	// soal 1
	var bangunDatar hitungBangunDatar
	var bangunRuang hitungBangunRuang

	bangunDatar = segitigaSamaSisi{4, 4}
	fmt.Println("======== Segitiga Sama Sisi")
	fmt.Println("Luas:", bangunDatar.luas())
	fmt.Println("Keliling:", bangunDatar.keliling())

	bangunDatar = persegiPanjang{3, 5}
	fmt.Println("======== Persegi Panjang")
	fmt.Println("Luas:", bangunDatar.luas())
	fmt.Println("Keliling:", bangunDatar.keliling())

	bangunRuang = balok{2, 3, 4}
	fmt.Println("======== Balok")
	fmt.Println("Volume:", bangunRuang.volume())
	fmt.Println("Luas Permukaan:", bangunRuang.luasPermukaan())

	bangunRuang = tabung{1, 2}
	fmt.Println("======== Tabung")
	fmt.Println("Volume:", fmt.Sprintf("%.2f", bangunRuang.volume()))
	fmt.Println("Luas Permukaan:", fmt.Sprintf("%.2f", bangunRuang.luasPermukaan()))
	fmt.Println()

	// soal 2
	var colors = []string{"Mystic Bronze", "Mystic White", "Mystic Black"}
	var sentence smarthphone = phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung Galaxy Note 20",
		year:   2020,
		colors: colors,
	}
	fmt.Println(sentence.phone())
	fmt.Println()

	// soal 3
	var result interface{}
	var luasPersegi = func(n int, d bool) interface{} {
		switch {
		case n > 0 && d:
			result = "luas persegi dengan sisi " +
				strconv.Itoa(n) + " cm adalah " +
				strconv.Itoa(n*n) + " cm"
		case n > 0 && !d:
			result = n * n
		case n == 0 && d:
			result = "Maaf anda belum menginput sisi dari persegi"
		case n == 0 && !d:
			result = nil
		}
		return result
	}
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))
	fmt.Println()

	// soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	var angkaPertama = strconv.Itoa(kumpulanAngkaPertama.([]int)[0]) + " + " + strconv.Itoa(kumpulanAngkaPertama.([]int)[1])
	var angkaKedua = strconv.Itoa(kumpulanAngkaKedua.([]int)[0]) + " + " + strconv.Itoa(kumpulanAngkaKedua.([]int)[1])
	var jumlah int = kumpulanAngkaPertama.([]int)[0] + kumpulanAngkaPertama.([]int)[1] + kumpulanAngkaKedua.([]int)[0] + kumpulanAngkaKedua.([]int)[1]

	fmt.Println(prefix.(string)+angkaPertama+" + "+angkaKedua, "=", jumlah)
}
