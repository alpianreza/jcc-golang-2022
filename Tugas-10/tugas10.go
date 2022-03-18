package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"
)

func main() {
	// soal 1
	printJcc()

	// soal 2
	segitiga()

	// soal 3
	numbers()

	// soal 4
	fmt.Println("============== Soal 4 ")

	phones := []string{}
	smartphones := func(s string, p *[]string) {
		phones = append(*p, s)
	}

	phoneList := []string{
		"Xiaomi",
		"Asus",
		"IPhone",
		"Samsung",
		"Oppo",
		"Realme",
		"Vivo"}
	sort.Strings(phoneList)

	for i, p := range phoneList {
		smartphones(p, &phones)
		result := strconv.Itoa(i+1) + ". " + phones[i]
		fmt.Println(result)
		time.Sleep(1 * time.Second)
	}

	// soal 5
	fmt.Println("============== Soal 5 ")

	kelilingO := func(r int) float64 {
		return math.Round(2 * math.Pi * float64(r))
	}
	luasO := func(r int) float64 {
		return math.Round(math.Pi * math.Pow(float64(r), 2))
	}
	jariJariO := []int{7, 10, 15}
	for _, n := range jariJariO {
		fmt.Println("Jari-jari Lingkaran :", n)
		fmt.Println("Luas :", luasO(n), "Keliling :", kelilingO(n))
	}

	// soal 6
	fmt.Println("============== Soal 6 ")

	kelilingpersegi := func(p int, l int) int {
		return 2 * (p + l)
	}
	luaspersegi := func(p int, l int) int {
		return p * l
	}

	p := flag.Int("panjang", 5, "Masukkan Panjang")
	l := flag.Int("lebar", 10, "Masukkan Lebar")
	flag.Parse()

	fmt.Println("Panjang:", *p, "x Lebar:", *l, "= Luas:", luaspersegi(*p, *l))
	fmt.Println("Panjang:", *p, "x Lebar::", *l, "= Keliling:", kelilingpersegi(*p, *l))
}

// func soal 1

func printJCC(kalimat string, tahun int) {
	fmt.Println(kalimat, tahun)
}

func printJcc() {
	fmt.Println("============== Soal 1 ")
	kalimat := "Golang Backend Development"
	tahun := 2021
	defer printJCC(kalimat, tahun)
}

// func soal 2
func segitigacase() {
	if s := recover(); s != nil {
		fmt.Println(s)
	}
}
func kelilingSegitigaSamaSisi(n int, d bool) (string, error) {
	switch {
	case n == 0 && d == true:
		err := errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		return "", err
	case n == 0 && d == false:
		panic("Maaf and belum menginput sisi dari segitiga sama sisi")
	case d == false:
		return strconv.Itoa(int(n * 3)), nil
	default:
		return "keliling segitiga dengan sisi " + strconv.Itoa(int(n)) + " cm adalah " + strconv.Itoa(int(n*3)) + " cm", nil
	}
}
func segitiga() {
	defer segitigacase()
	defer kelilingSegitigaSamaSisi(4, true)
	defer kelilingSegitigaSamaSisi(8, false)
	defer kelilingSegitigaSamaSisi(0, true)
	defer kelilingSegitigaSamaSisi(0, false)

	fmt.Println("============== Soal 2 ")
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))
}

// func soal 3

func numbers() {
	angka := 1

	defer num(&angka)

	tambahAngka := func(n int, angka *int) {
		*angka += n
	}

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)
}

func num(n *int) {
	fmt.Println("============== Soal 3 ")
	fmt.Println(strconv.Itoa(*n))
}
