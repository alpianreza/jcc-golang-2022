package models

import "time"

type NilaiMahasiswa struct {
	ID          uint      `json:"id"`
	Nama        string    `json:"nama"`
	MataKuliah  string    `json:"mata_kuliah"`
	Nilai       uint      `json:"nilai"`
	IndeksNilai string    `json:"index_nilai"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"update_at"`
}

func GetIndeks(n uint) string {
	indeks := ""
	switch {
	case n >= 80:
		indeks = "A"
	case n >= 70:
		indeks = "B"
	case n >= 60:
		indeks = "C"
	case n >= 50:
		indeks = "D"
	default:
		indeks = "E"
	}
	return indeks
}
