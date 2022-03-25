package main

import (
	"Tugas-14/models"
	"Tugas-14/nilai"
	"Tugas-14/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/nilai", GetNilai)
	router.POST("/nilai/create", PostNilai)
	router.PUT("/nilai/id/update", UpdateNilai)
	router.DELETE("/nilai/delete", DeleteNilai)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func GetNilai(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nilaiAll, err := nilai.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}
	utils.ResponseJson(w, nilaiAll, http.StatusOK)
}

func PostNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var nilmhs models.NilaiMahasiswa
	if r.Header.Get("Content-Type") == "application/json" {
		type nilaiInput struct {
			Nama, MataKuliah string
			Nilai            int
		}
		var inputStructed nilaiInput
		json.NewDecoder(r.Body).Decode(&inputStructed)
		nilmhs.Nama = inputStructed.Nama
		nilmhs.MataKuliah = inputStructed.MataKuliah
		nilmhs.Nilai = uint(inputStructed.Nilai)
	} else {
		nilmhs.Nama = r.PostFormValue("Nama")
		nilmhs.MataKuliah = r.PostFormValue("MataKuliah")
		nilaiValue, _ := strconv.Atoi(r.PostFormValue("Nilai"))
		nilmhs.Nilai = uint(nilaiValue)
	}

	if nilmhs.Nilai > 100 {
		nilmhs.Nilai = 100
	}
	nilmhs.IndeksNilai = models.GetIndeks(nilmhs.Nilai)

	// post
	if err := json.NewDecoder(r.Body).Decode(&nilmhs); err != nil {
		utils.ResponseJson(w, err, http.StatusBadRequest)
		return
	}
	if err := nilai.Insert(ctx, nilmhs); err != nil {
		utils.ResponseJson(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJson(w, res, http.StatusCreated)
}
func UpdateNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mov models.NilaiMahasiswa

	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		utils.ResponseJson(w, err, http.StatusBadRequest)
		return
	}

	var idNilai = ps.ByName("id")

	if err := nilai.Update(ctx, mov, idNilai); err != nil {
		utils.ResponseJson(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJson(w, res, http.StatusCreated)
}
func DeleteNilai(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idNilai = ps.ByName("id")
	if err := nilai.Delete(ctx, idNilai); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJson(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJson(w, res, http.StatusOK)
}
