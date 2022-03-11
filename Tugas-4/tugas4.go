package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// maaf baru ngumpulin

	// soal 1
	for i := 1; i <= 20; i++ {
		kata := ""
		switch {
		case i%2 == 1:
			if i%3 == 0 {
				kata = "I Love Coding"
			} else {
				kata = "JCC"
			}
		default:
			kata = "Candradimuka"
		}
		fmt.Println(i, "-", kata)
	}

	fmt.Println()

	// soal 2
	for i := 1; i <= 7; i++ {
		fmt.Println(strings.Repeat("#", i))
	}

	fmt.Println()

	// soal 3
	kalimat := [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2:5])

	fmt.Println()

	// soal 4
	sayuran := []string{}
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")
	for i, sayur := range sayuran {
		fmt.Println(strconv.Itoa(i+1) + ". " + sayur)
	}

	fmt.Println()

	// soal 5
	satuan := map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	satuan["Volume Balok"] = satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	for satu := range satuan {
		fmt.Println(satu, "=", satuan[satu])
	}
}
