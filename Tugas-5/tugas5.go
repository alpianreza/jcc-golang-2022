package main

import "fmt"

// func soal 1
func luasPersegiPanjang(panjang, lebar int) int {
	return panjang * lebar
}
func kelilingPersegiPanjang(panjang, lebar int) int {
	return 2*panjang + 2*lebar
}
func volumeBalok(panjang, lebar, tinggi int) int {
	return panjang * lebar * tinggi
}

// func soal 2
func introduce(name, gender, occupation, age string) string {
	var identity string
	if gender == "laki-laki" {
		identity = "Pak "
	} else {
		identity = "Bu "
	}
	hasil := (identity + name + " adalah seorang " + occupation + " yang berusia " + age + " tahun")
	return hasil
}

// func soal 3

func buahFavorit(fav string, buah ...string) (fruits string) {

	fmt.Printf("halo nama saya %s dan buah favorit saya adalah", fav)
	for _, fruits := range buah {
		fmt.Printf(" %+q, ", fruits)
	}
	return
}

func main() {
	// soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, tinggi)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println("luas =", luas)
	fmt.Println("keliling =", keliling)
	fmt.Println("volume =", volume)
	fmt.Println()
	// soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)
	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)
	fmt.Println()

	// soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("john", buah...)
	fmt.Println(buahFavoritJohn)
	fmt.Println()

	// soal 4

	var dataFilm = []map[string]string{}

	var tambahDataFilm = func(judul, durasi, genre, tahun string) {
		var data = map[string]string{}
		data["genre"] = genre
		data["jam"] = durasi
		data["tahun"] = tahun
		data["title"] = judul
		dataFilm = append(dataFilm, data)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

}
