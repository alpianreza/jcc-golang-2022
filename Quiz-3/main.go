package main

import (
	"Quiz-3/config"
	"Quiz-3/functions"
	"Quiz-3/models"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	authUser       = []string{"admin", "editor", "trainer"}
	authPass       = []string{"password", "secret", "rahasia"}
	booksTemp      = []models.Book{}
	categoriesTemp = []models.Category{}
)

func main() {
	// connect sql
	db, e := config.MySQL()
	if e != nil {
		log.Fatal(e)
	}
	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}
	fmt.Println("Success")

	// route
	bangunDatarRoute()
	booksRoute()

	// serve
	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func bangunDatarRoute() {
	http.Handle("/bangun-datar/segitiga-sama-sisi", auth(http.HandlerFunc(segitigaSamaSisiRoute)))
	http.Handle("/bangun-datar/persegi", auth(http.HandlerFunc(persegiRoute)))
	http.Handle("/bangun-datar/persegi-panjang", auth(http.HandlerFunc(persegiPanjangRoute)))
	http.Handle("/bangun-datar/lingkaran", auth(http.HandlerFunc(lingkaranRoute)))
	http.Handle("/bangun-datar/jajar-genjang", auth(http.HandlerFunc(jajarGenjangRoute)))

	// example for goroutine in persegi
	// example for channel in lingkaran
}

// books route
func booksRoute() {
	route := func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == "POST":
			PostBook(w, r)
		case r.Method == "GET":
			GetBook(w, r)
		case r.Method == "PUT":
			UpBook(w, r)
		case r.Method == "DELETE":
			DelBook(w, r)
		default:
			w.Write([]byte("Function only support POST/GET/PUT/DEL"))
		}
	}
	http.Handle("/books", auth(http.HandlerFunc(route)))
	http.Handle("/books/", auth(http.HandlerFunc(route)))
}
func categoryRoute() {
	route := func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == "POST":
			PostCategory(w, r)
		case r.Method == "GET":
			GetCategory(w, r)
		case r.Method == "PUT":
			UpdateCategory(w, r)
		case r.Method == "DELETE":
			DeleteCategory(w, r)
		default:
			w.Write([]byte("Function only support POST/GET/PUT/DEL"))
		}
	}
	http.Handle("/category", auth(http.HandlerFunc(route)))
	http.Handle("/categories/", auth(http.HandlerFunc(route)))
}

func PostBook(w http.ResponseWriter, r *http.Request) {
	var buku models.Book

	// input
	if r.Header.Get("Content-Type") == "application/json" {
		type BookInput struct {
			Title        string `json:"title"`
			Description  string `json:"description"`
			Image_url    string `json:"image_url"`
			Release_year int    `json:"release_year"`
			Price        int    `json:"price"`
			Total_page   string `json:"total_page"`
			CategoryId   string `json:"category_id"`
		}
		var bookStructed BookInput
		json.NewDecoder(r.Body).Decode(&bookStructed)

		buku.Title = bookStructed.Title
		buku.Description = bookStructed.Description
		buku.Image_url = bookStructed.Image_url
		buku.Total_page = bookStructed.Total_page
		buku.Release_year = bookStructed.Release_year
		buku.CategoryId, _ = strconv.Atoi(bookStructed.CategoryId)
		buku.Price = functions.GetPriceWithCurrency(bookStructed.Price)
	} else {
		buku.Title = r.PostFormValue("title")
		buku.Description = r.PostFormValue("description")
		buku.Image_url = r.PostFormValue("image_url")
		buku.Total_page = r.PostFormValue("total_page")
		buku.Release_year, _ = strconv.Atoi(r.PostFormValue("release_year"))
		priceInput, _ := strconv.Atoi(r.PostFormValue("price"))
		buku.Price = functions.GetPriceWithCurrency(priceInput)
		buku.CategoryId, _ = strconv.Atoi(r.PostFormValue("category_id"))
	}

	// validasi
	errMsg := ""
	if !functions.IsImageUrlValid(buku.Image_url) {
		errMsg += "image_url tidak dapat diakses"
	}
	if !functions.IsReleaseYearValid(buku.Release_year) {
		if errMsg != "" {
			errMsg += " dan "
		}
		errMsg += "release_year harus di antara 1980 - 2021"
	}

	latestIndex := 0
	for _, book := range booksTemp {
		latestIndex = book.ID
	}

	buku.ID = latestIndex + 1
	tebal, _ := strconv.Atoi(buku.Total_page)
	buku.Thickness = functions.GetKategoriKetebalan(tebal)
	buku.CreatedAt = time.Now()
	buku.UpdatedAt = time.Now()

	// push new value
	booksTemp = append(booksTemp, buku)

	// show result
	if errMsg != "" {
		errJson := "{ \"error\" : \"" + errMsg + "\" }"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(errJson))
	} else {
		bookJson, _ := json.Marshal(buku)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bookJson)
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	result := []models.Book{}
	indexToDel := []int{}

	// find indexes to delete
	qTitle := r.URL.Query().Get("title")
	if qTitle != "" {
		for _, book := range booksTemp {
			if !strings.Contains(strings.ToLower(book.Title), strings.ToLower(qTitle)) {
				indexToDel = append(indexToDel, book.ID)
			}
		}
	}

	qMinYear := r.URL.Query().Get("minYear")
	if qMinYear != "" {
		qMinYearInt, _ := strconv.Atoi(qMinYear)
		for _, book := range booksTemp {
			if book.Release_year < qMinYearInt {
				indexToDel = append(indexToDel, book.ID)
			}
		}
	}

	qMaxYear := r.URL.Query().Get("maxYear")
	if qMaxYear != "" {
		qMaxYearInt, _ := strconv.Atoi(qMaxYear)
		for _, book := range booksTemp {
			if book.Release_year > qMaxYearInt {
				indexToDel = append(indexToDel, book.ID)
			}
		}
	}

	qMinPage := r.URL.Query().Get("minPage")
	if qMinPage != "" {
		qMinPageInt, _ := strconv.Atoi(qMinPage)
		for _, book := range booksTemp {
			totalPageInt, _ := strconv.Atoi(book.Total_page)
			if totalPageInt < qMinPageInt-1 {
				indexToDel = append(indexToDel, book.ID)
			}
		}
	}

	qMaxPage := r.URL.Query().Get("maxPage")
	if qMaxPage != "" {
		qMaxPageInt, _ := strconv.Atoi(qMaxPage)
		for _, book := range booksTemp {
			totalPageInt, _ := strconv.Atoi(book.Total_page)
			if totalPageInt > qMaxPageInt+1 {
				indexToDel = append(indexToDel, book.ID)
			}
		}
	}

	qSort := r.URL.Query().Get("sort")
	isAsc := false
	isDesc := false
	if qSort == "asc" {
		isAsc = true
	}
	if qSort == "desc" {
		isDesc = true
	}

	sort.Ints(indexToDel)
	uniqueIndexes := []int{}
	for i, n := range indexToDel {
		if i != 0 {
			if indexToDel[i] != indexToDel[i-1] {
				uniqueIndexes = append(uniqueIndexes, n)
			}
		} else {
			uniqueIndexes = append(uniqueIndexes, n)
		}

	}
	indexToDel = uniqueIndexes

	// append
	for _, book := range booksTemp {
		toDel := false
		for i, n := range indexToDel {
			if book.ID == n {
				toDel = true
				indexToDel = append(indexToDel[:i], indexToDel[i+1:]...)
			}
			break
		}
		if !toDel {
			result = append(result, book)
		}
	}
	if isAsc {
		sort.Slice(result, func(i, j int) bool {
			return result[i].ID < result[j].ID
		})
	}
	if isDesc {
		sort.Slice(result, func(i, j int) bool {
			return result[i].ID > result[j].ID
		})
	}

	// Menampilkan hasil
	bukubukuJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bukubukuJson)
}

func UpBook(w http.ResponseWriter, r *http.Request) {
	var buku models.Book
	errMsg := ""
	buku.ID, _ = strconv.Atoi(strings.Split(r.URL.Path, "/")[2])

	// get input
	if r.Header.Get("Content-Type") == "application/json" {
		type BookInput struct {
			Title        string `json:"title"`
			Description  string `json:"description"`
			Image_url    string `json:"image_url"`
			Release_year int    `json:"release_year"`
			Price        int    `json:"price"`
			Total_page   string `json:"total_page"`
			CategoryId   string `json:"category_id"`
		}
		var bookStructed BookInput
		json.NewDecoder(r.Body).Decode(&bookStructed)

		buku.Title = bookStructed.Title
		buku.Description = bookStructed.Description
		buku.Image_url = bookStructed.Image_url
		buku.Total_page = bookStructed.Total_page
		buku.Release_year = bookStructed.Release_year
		buku.CategoryId, _ = strconv.Atoi(bookStructed.CategoryId)
		buku.Price = functions.GetPriceWithCurrency(bookStructed.Price)
	} else {
		buku.Title = r.PostFormValue("title")
		buku.Description = r.PostFormValue("description")
		buku.Image_url = r.PostFormValue("image_url")
		buku.Total_page = r.PostFormValue("total_page")
		buku.Release_year, _ = strconv.Atoi(r.PostFormValue("release_year"))
		priceInput, _ := strconv.Atoi(r.PostFormValue("price"))
		buku.Price = functions.GetPriceWithCurrency(priceInput)
		buku.CategoryId, _ = strconv.Atoi(r.PostFormValue("category_id"))
	}

	chosenOneIndex := 0
	for i, book := range booksTemp {
		if book.ID == buku.ID {
			chosenOneIndex = i
		}
	}

	if chosenOneIndex == 0 {
		errMsg = "Update failed, book not found!"
	} else {
		// validation
		if !functions.IsImageUrlValid(buku.Image_url) {
			errMsg += "image_url tidak dapat diakses"
		}
		if !functions.IsReleaseYearValid(buku.Release_year) {
			if errMsg != "" {
				errMsg += " dan "
			}
			errMsg += "release_year harus di antara 1980 - 2021"
		}

		tebal, _ := strconv.Atoi(buku.Total_page)
		buku.Thickness = functions.GetKategoriKetebalan(tebal)
		buku.CreatedAt = booksTemp[buku.ID-1].CreatedAt
		buku.UpdatedAt = time.Now()

		for i, book := range booksTemp {
			if book.ID == buku.ID {
				booksTemp[i] = buku
			}
		}
	}

	if errMsg != "" {
		errJson := "{ \"error\" : \"" + errMsg + "\" }"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(errJson))
	} else {
		bookJson, _ := json.Marshal(buku)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bookJson)
	}
}

func DelBook(w http.ResponseWriter, r *http.Request) {
	resultMsg := "{ \"result\" : \"Delete failed, book not found!\" }"
	delId, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])

	for i, book := range booksTemp {
		if book.ID == delId {
			booksTemp = append(booksTemp[:i], booksTemp[i+1:]...)
			resultMsg = "{ \"result\" : \"Delete success!\" }"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resultMsg))
}

// post category

func PostCategory(w http.ResponseWriter, r *http.Request) {
	var kategori models.Category

	// input category
	if r.Header.Get("Content-Type") == "application/json" {
		type CategoryInput struct {
			Name string `json:"name"`
		}
		var CategoryStructed CategoryInput
		json.NewDecoder(r.Body).Decode(&CategoryStructed)

		kategori.Name = CategoryStructed.Name
	} else {
		kategori.Name = r.PostFormValue("name")
	}

	// push new value
	categoriesTemp = append(categoriesTemp, kategori)

	// show result
	errMsg := ""
	if errMsg != "" {
		errJson := "{ \"error\" : \"" + errMsg + "\" }"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(errJson))
	} else {
		bookJson, _ := json.Marshal(kategori)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bookJson)
	}
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	result := []models.Book{}
	indexToDel := []int{}

	// find indexes to delete
	kNama := r.URL.Query().Get("name")
	if kNama != "" {
		for _, category := range categoriesTemp {
			if !strings.Contains(strings.ToLower(category.Name), strings.ToLower(kNama)) {
				indexToDel = append(indexToDel, category.ID)
			}
		}
	}
	kMinYear := r.URL.Query().Get("minYear")
	if kMinYear != "" {
		qMinYearInt, _ := strconv.Atoi(kMinYear)
		for _, book := range booksTemp {
			if book.Release_year < qMinYearInt {
				indexToDel = append(indexToDel, book.ID)
			}
		}
	}

	kMaxYear := r.URL.Query().Get("maxYear")
	if kMaxYear != "" {
		kMaxYearInt, _ := strconv.Atoi(kMaxYear)
		for _, book := range booksTemp {
			if book.Release_year > kMaxYearInt {
				indexToDel = append(indexToDel, book.ID)
			}
		}
	}

	kMinPage := r.URL.Query().Get("minPage")
	if kMinPage != "" {
		kMinPageInt, _ := strconv.Atoi(kMinPage)
		for _, book := range booksTemp {
			totalPageInt, _ := strconv.Atoi(book.Total_page)
			if totalPageInt < kMinPageInt-1 {
				indexToDel = append(indexToDel, book.ID)
			}
		}
	}

	kMaxPage := r.URL.Query().Get("maxPage")
	if kMaxPage != "" {
		kMaxPageInt, _ := strconv.Atoi(kMaxPage)
		for _, book := range booksTemp {
			totalPageInt, _ := strconv.Atoi(book.Total_page)
			if totalPageInt > kMaxPageInt+1 {
				indexToDel = append(indexToDel, book.ID)
			}
		}
	}

	kSort := r.URL.Query().Get("sort")
	isAsc := false
	isDesc := false
	if kSort == "asc" {
		isAsc = true
	}
	if kSort == "desc" {
		isDesc = true
	}

	sort.Ints(indexToDel)
	uniqueIndexes := []int{}
	for i, n := range indexToDel {
		if i != 0 {
			if indexToDel[i] != indexToDel[i-1] {
				uniqueIndexes = append(uniqueIndexes, n)
			}
		} else {
			uniqueIndexes = append(uniqueIndexes, n)
		}

	}
	indexToDel = uniqueIndexes

	for _, book := range booksTemp {
		toDel := false
		for i, n := range indexToDel {
			if book.ID == n {
				toDel = true
				indexToDel = append(indexToDel[:i], indexToDel[i+1:]...)
			}
			break
		}
		if !toDel {
			result = append(result, book)
		}
	}
	if isAsc {
		sort.Slice(result, func(i, j int) bool {
			return result[i].ID < result[j].ID
		})
	}
	if isDesc {
		sort.Slice(result, func(i, j int) bool {
			return result[i].ID > result[j].ID
		})
	}

	// show result
	bukubukuJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bukubukuJson)
}

// update Category
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var kategori models.Category
	kategori.ID, _ = strconv.Atoi(strings.Split(r.URL.Path, "/")[2])

	// get input
	if r.Header.Get("Content-Type") == "application/json" {
		type BookInput struct {
			Name string `json:"name"`
		}
		var categoryStructed BookInput
		json.NewDecoder(r.Body).Decode(&categoryStructed)

		kategori.Name = categoryStructed.Name

	} else {
		kategori.Name = r.PostFormValue("title")
	}

	for i, category := range booksTemp {
		if category.ID == category.ID {
			booksTemp[i] = category
		}
	}
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	resultMsg := "{ \"result\" : \"Delete failed, book not found!\" }"
	delId, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])

	for i, book := range categoriesTemp {
		if book.ID == delId {
			categoriesTemp = append(categoriesTemp[:i], categoriesTemp[i+1:]...)
			resultMsg = "{ \"result\" : \"Delete success!\" }"
		}
	}

	// show result
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resultMsg))
}

// bangun datar route
func segitigaSamaSisiRoute(w http.ResponseWriter, r *http.Request) {
	var datar functions.BangunDatar

	hitung := r.URL.Query().Get("hitung")
	alas, err := strconv.Atoi(r.URL.Query().Get("alas"))
	if err != nil {
		alas = 0
	}
	tinggi, err := strconv.Atoi(r.URL.Query().Get("tinggi"))
	if err != nil {
		tinggi = 0
	}

	switch hitung {
	case "luas":
		datar.Result = float64(alas * tinggi / 2)
	case "keliling":
		datar.Result = float64(alas * 3)
	default:
		datar.Result = 0
	}

	printDatar(datar, w)
}

func persegiRoute(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	var datar functions.BangunDatar

	printPersegi := func() {
		printDatar(datar, w)
		wg.Done()
	}

	hitung := r.URL.Query().Get("hitung")
	sisi, err := strconv.Atoi(r.URL.Query().Get("sisi"))
	if err != nil {
		sisi = 0
	}

	switch hitung {
	case "luas":
		datar.Result = math.Pow(float64(sisi), 2)
	case "keliling":
		datar.Result = float64(sisi * 4)
	default:
		datar.Result = 0
	}

	wg.Add(1)
	go printPersegi()
	wg.Wait()
}

func persegiPanjangRoute(w http.ResponseWriter, r *http.Request) {
	var datar functions.BangunDatar

	hitung := r.URL.Query().Get("hitung")
	panjang, err := strconv.Atoi(r.URL.Query().Get("panjang"))
	if err != nil {
		panjang = 0
	}
	lebar, err := strconv.Atoi(r.URL.Query().Get("lebar"))
	if err != nil {
		lebar = 0
	}

	switch hitung {
	case "luas":
		datar.Result = float64(panjang * lebar)
	case "keliling":
		datar.Result = float64(2 * (panjang + lebar))
	default:
		datar.Result = 0
	}

	printDatar(datar, w)
}

func lingkaranRoute(w http.ResponseWriter, r *http.Request) {
	var datar functions.BangunDatar

	hitung := r.URL.Query().Get("hitung")
	jariJari, err := strconv.Atoi(r.URL.Query().Get("jariJari"))
	if err != nil {
		jariJari = 0
	}

	inputChan := func(ch chan int, rs int) {
		ch <- rs
		close(ch)
	}

	jariJariChannel := make(chan int)
	go inputChan(jariJariChannel, jariJari)

	switch hitung {
	case "luas":
		datar.Result = math.Round(math.Pi * math.Pow(float64(<-jariJariChannel), 2))
	case "keliling":
		datar.Result = math.Round(2 * math.Pi * float64(<-jariJariChannel))
	default:
		datar.Result = 0
	}

	printDatar(datar, w)
}

func jajarGenjangRoute(w http.ResponseWriter, r *http.Request) {
	var datar functions.BangunDatar

	hitung := r.URL.Query().Get("hitung")
	alas, err := strconv.Atoi(r.URL.Query().Get("alas"))
	if err != nil {
		alas = 0
	}
	tinggi, err := strconv.Atoi(r.URL.Query().Get("tinggi"))
	if err != nil {
		tinggi = 0
	}
	sisi, err := strconv.Atoi(r.URL.Query().Get("sisi"))
	if err != nil {
		sisi = 0
	}

	switch hitung {
	case "luas":
		datar.Result = float64(alas * tinggi)
	case "keliling":
		datar.Result = float64((2 * alas) + (2 * sisi))
	default:
		datar.Result = 0
	}

	printDatar(datar, w)
}

func printDatar(d functions.BangunDatar, w http.ResponseWriter) {
	datarJson, _ := json.Marshal(d)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(datarJson)
}

// auth
func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// basic auth
		user, pass, ok := r.BasicAuth()

		// check which account used
		passSelected := ""
		isUserRegistered := false

		for i, u := range authUser {
			if u == user {
				passSelected = authPass[i]
				isUserRegistered = true
			}
		}

		// post auth
		if r.Method != "POST" {
			// auth not ok
			if !ok {
				w.Write([]byte("Username atau Password tidak boleh kosong"))
				return
			}

			// input invalid
			if !isUserRegistered || pass != passSelected {
				w.Write([]byte("Username atau Password tidak sesuai"))
				return
			}
		}

		// input ok
		next.ServeHTTP(w, r)
	})
}
