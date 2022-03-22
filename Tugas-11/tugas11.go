package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"sync"
	"time"
)

func main() {
	// soal 1
	fmt.Println("------------- Soal 1")

	var wg sync.WaitGroup

	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	sort.Strings(phones)

	printPhones := func(result string) {
		time.Sleep(1 * time.Second)
		fmt.Println(result)
		wg.Done()
	}

	for i, p := range phones {
		result := strconv.Itoa(i+1) + ". " + p
		wg.Add(1)
		go printPhones(result)
		wg.Wait()
	}

	// soal 2

	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	getMovies := func(ch chan string, movies ...string) {
		for i, film := range movies {
			var result = strconv.Itoa(i+1) + ". " + film
			ch <- result
		}
		close(ch)
	}

	go getMovies(moviesChannel, movies...)

	fmt.Println("------------- Soal 2")
	fmt.Println("List Movies:")
	for value := range moviesChannel {
		fmt.Println(value)
	}

	// soal 3
	fmt.Println("------------- Soal 3")

	var tinggiTabung = 10
	var jariJari = []int{8, 14, 20}

	var luas = func(r int) float64 {
		return math.Round(math.Pi * math.Pow(float64(r), 2))
	}
	var keliling = func(r int) float64 {
		return math.Round(2 * math.Pi * float64(r))
	}
	var volumeTabung = func(r int) float64 {
		return math.Round(math.Pi * math.Pow(float64(r), 2) * float64(tinggiTabung))
	}
	var inputChan = func(ch chan int, rs ...int) {
		for _, r := range rs {
			ch <- r
		}
		close(ch)
	}

	var jariJariChannel = make(chan int)
	go inputChan(jariJariChannel, jariJari...)

	for r := range jariJariChannel {
		fmt.Println("Jari-Jari Lingkaran =", r)
		fmt.Println("-------------------------")
		fmt.Println("Luas Lingkaran:", luas(r))
		fmt.Println("Keliling Lingkaran:", keliling(r))
		fmt.Println("Volume Tabung:", volumeTabung(r))
		fmt.Println("-------------------------")

	}

	// soal 4
	fmt.Println("------------- Soal 4")

	luasPersegiPanjang := func(p int, l int, ch chan int) {
		ch <- p * l
	}
	kelilingPersegiPanjang := func(p int, l int, ch chan int) {
		ch <- 2 * (p + l)
	}
	volumeBalok := func(p int, l int, t int, ch chan int) {
		ch <- p * l * t
	}

	var kelilingPp = make(chan int)
	go kelilingPersegiPanjang(4, 6, kelilingPp)

	var luasPp = make(chan int)
	go luasPersegiPanjang(4, 6, luasPp)

	var volumeB = make(chan int)
	go volumeBalok(4, 6, 8, volumeB)

	for i := 0; i < 3; i++ {
		select {
		case keliling := <-kelilingPp:
			fmt.Println("Keliling Persegi Panjang:", keliling)
		case luas := <-luasPp:
			fmt.Println("Luas Persegi Panjang:", luas)
		case volume := <-volumeB:
			fmt.Println("Volume Balok:", volume)
		}
	}
}
