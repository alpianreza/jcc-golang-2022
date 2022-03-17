package library

import (
	"math"
	"strconv"
	"strings"
)

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

func (pp PersegiPanjang) Luas() int {
	return pp.Panjang * pp.Lebar
}

func (pp PersegiPanjang) Keliling() int {
	return 2 * (pp.Panjang + pp.Lebar)
}

func (b Balok) Volume() float64 {
	return float64(b.Lebar) * float64(b.Panjang) * float64(b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return (2 * (float64(b.Lebar) + float64(b.Panjang))) +
		(2 * (float64(b.Lebar) + float64(b.Tinggi))) +
		(2 * (float64(b.Tinggi) + float64(b.Panjang)))
}

func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

// struct iface func soal 2
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type Smartphone interface {
	Phone() string
}

func (sp Phone) Phone() string {
	hasil := "name : " + sp.Name + "\n" + "brand : " + sp.Brand + "\n" + "year : " + strconv.Itoa(sp.Year) + "\n" + "colors : " + strings.Join(sp.Colors, ", ")

	return hasil
}
