package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var username = "admin"
var password = "admin"

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

// middleware
func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// post auth
		if r.Method == "POST" {
			user, pass, ok := r.BasicAuth()
			if !ok {
				w.Write([]byte("Username atau Password tidak boleh kosong"))
				return
			}
			if user != username || pass != password {
				w.Write([]byte("Username atau Password tidak sesuai"))
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func nilai(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "POST":
		{
			postNilai(w, r)
		}
	case r.Method == "GET":
		{
			getNilai(w)
		}
	default:
		w.Write([]byte("Fungsi hanya mendukung metode GET dan POST"))
	}
}

// post nilai
func postNilai(w http.ResponseWriter, r *http.Request) {
	namaValue := ""
	mataKuliahValue := ""
	indeksValue := ""
	nilaiValue := 0
	idValue := 0

	// get input
	if r.Header.Get("Content-Type") == "application/json" {
		type nilaiInput struct {
			Nama, MataKuliah string
			Nilai            int
		}
		var inputStructed nilaiInput
		json.NewDecoder(r.Body).Decode(&inputStructed)

		namaValue = inputStructed.Nama
		mataKuliahValue = inputStructed.MataKuliah
		nilaiValue = inputStructed.Nilai
	} else {
		namaValue = r.PostFormValue("Nama")
		mataKuliahValue = r.PostFormValue("MataKuliah")
		nilaiValue, _ = strconv.Atoi(r.PostFormValue("Nilai"))
	}

	indeksValue = getIndeks(nilaiValue)
	idValue = len(nilaiNilaiMahasiswa)

	// limit nilai
	if nilaiValue > 100 {
		nilaiValue = 100
	}
	if nilaiValue < 0 {
		nilaiValue = 0
	}

	newValue := NilaiMahasiswa{
		namaValue,
		mataKuliahValue,
		indeksValue,
		uint(nilaiValue),
		uint(idValue)}
	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, newValue)

	// show result
	nilaiJson, _ := json.Marshal(nilaiNilaiMahasiswa[idValue])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(nilaiJson)
}

func getNilai(w http.ResponseWriter) {
	nilaiJson, _ := json.Marshal(nilaiNilaiMahasiswa)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(nilaiJson)
}

func getIndeks(n int) string {
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

func main() {
	http.Handle("/", auth(http.HandlerFunc(nilai)))

	fmt.Println("server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
