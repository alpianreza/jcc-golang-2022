package main

import (
	"fmt"
	"strconv"
)

func main() {
	// soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var panjangpersegi, _ = strconv.Atoi(panjangPersegiPanjang)
	var lebarpersegi, _ = strconv.Atoi(lebarPersegiPanjang)

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var alasSegtig, _ = strconv.Atoi(alasSegitiga)
	var tinggiSegtig, _ = strconv.Atoi(tinggiSegitiga)

	var kelilingPersegiPanjang int = panjangpersegi * lebarpersegi
	var luasSegitiga int = alasSegtig * tinggiSegtig

	fmt.Println("Keliling persegi panjang = ", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga = ", luasSegitiga)

	fmt.Println()

	// soal 2

	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("Nilai John = A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("Nilai John = B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("Nilai John = C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("Nilai John = D")
	} else if nilaiJohn < 50 {
		fmt.Println("Nilai John = E")
	}

	if nilaiDoe >= 80 {
		fmt.Println("Nilai Doe = A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("Nilai Doe = B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("Nilai Doe = C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("Nilai Doe = D")
	} else if nilaiDoe < 50 {
		fmt.Println("Nilai Doe = E")
	}

	fmt.Println()

	// Soal 3
	var date = 12
	var month = 6
	var year = 1992

	switch month {
	case 1:
		fmt.Println(date, "Januari", year)
	case 2:
		fmt.Println(date, "Februari", year)
	case 3:
		fmt.Println(date, "Maret", year)
	case 4:
		fmt.Println(date, "April", year)
	case 5:
		fmt.Println(date, "Mei", year)
	case 6:
		fmt.Println(date, "Juni", year)
	case 7:
		fmt.Println(date, "Juli", year)
	case 8:
		fmt.Println(date, "Agustus", year)
	case 9:
		fmt.Println(date, "September", year)
	case 10:
		fmt.Println(date, "Oktober", year)
	case 11:
		fmt.Println(date, "November", year)
	case 12:
		fmt.Println(date, "Desember", year)
	default:
		fmt.Println("Data yang anda inpuut salah")
	}

	fmt.Println()
	// soal 4

	var tahun = 1992
	switch {
	case tahun >= 1944 && tahun <= 1964:
		fmt.Println("Baby Shower")
	case tahun >= 1965 && tahun <= 1979:
		fmt.Println("Generasi X")
	case tahun >= 1980 && tahun <= 1994:
		fmt.Println("BGenerasi Y (Milenial")
	case tahun >= 1995 && tahun <= 2004:
		fmt.Println("Generasi Z")
	}
}
