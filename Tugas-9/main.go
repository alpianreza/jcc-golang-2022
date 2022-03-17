package main

import (
	. "Tugas-9/library"
	"fmt"
	"strconv"
)

func main() {
	// soal 1

	var bangunDatar HitungBangunDatar
	var bangunRuang HitungBangunRuang

	bangunDatar = SegitigaSamaSisi{4, 4}
	fmt.Println("======== Segitiga Sama Sisi")
	fmt.Println("Luas:", bangunDatar.Luas())
	fmt.Println("Keliling:", bangunDatar.Keliling())

	bangunDatar = PersegiPanjang{3, 5}
	fmt.Println("======== Persegi Panjang")
	fmt.Println("Luas:", bangunDatar.Luas())
	fmt.Println("Keliling:", bangunDatar.Keliling())

	bangunRuang = Balok{2, 3, 4}
	fmt.Println("======== Balok")
	fmt.Println("Volume:", bangunRuang.Volume())
	fmt.Println("Luas Permukaan:", bangunRuang.LuasPermukaan())

	bangunRuang = Tabung{1, 2}
	fmt.Println("======== Tabung")
	fmt.Println("Volume:", fmt.Sprintf("%.2f", bangunRuang.Volume()))
	fmt.Println("Luas Permukaan:", fmt.Sprintf("%.2f", bangunRuang.LuasPermukaan()))
	fmt.Println()

	// soal 2

	var colors = []string{"Mystic Bronze", "Mystic White", "Mystic Black"}
	var sentence Smartphone = Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung Galaxy Note 20",
		Year:   2020,
		Colors: colors,
	}
	fmt.Println(sentence.Phone())
	fmt.Println()

	// soal 3

	var result interface{}
	var LuasPersegi = func(n int, d bool) interface{} {
		switch {
		case n > 0 && d:
			result = "Luas Persegi dengan sisi " +
				strconv.Itoa(n) + " cm adalah " +
				strconv.Itoa(n*n) + " cm"
		case n > 0 && !d:
			result = n * n
		case n == 0 && d:
			result = "Maaf anda belum menginput sisi dari Persegi"
		case n == 0 && !d:
			result = nil
		}
		return result
	}
	fmt.Println(LuasPersegi(4, true))
	fmt.Println(LuasPersegi(8, false))
	fmt.Println(LuasPersegi(0, true))
	fmt.Println(LuasPersegi(0, false))
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
