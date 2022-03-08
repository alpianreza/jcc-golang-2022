package main

import ( 
	"fmt"
	"strings"
	"strconv"

)

func main()  {
// soal 1
	var kata1 = "Jabar";
	var kata2 = "Coding"
	var kata3 = "Camp"
	var kata4 = "Golang"
	var kata5 = "2022"
	fmt.Println(kata1, kata2, kata3, kata4, kata5);

//soal 2

	helo := "helo world";
	find := "world";
	replaceWith := "golang";

	var newText1 = strings.Replace(helo, find, replaceWith, 1);
	fmt.Print(newText1);


// soal 3
	var kataPertama = "saya";
	var kataKedua = "senang";
	var kataKetiga = "belajar";
	var kataKeempat = "golang";

    fmt.Printf("%s %s %s %s \n", kataPertama, kataKedua, kataKetiga, kataKeempat);

// soal 4
angkaPertama := "8"
angkaKedua := "5"
angkaKetiga := "6"
angkaKeempat := "7"

num1, err := strconv.Atoi(angkaPertama)

if err == nil {
	fmt.Println(num1)
}
num2, err := strconv.Atoi(angkaKedua)

if err == nil {
	fmt.Println(num2)
}
num3, err := strconv.Atoi(angkaKetiga)
if err == nil {
	fmt.Println(num3)
}
num4, err := strconv.Atoi(angkaKeempat)

if err == nil {
	fmt.Println(num4)
}
sum := num1 + num2 + num3 + num4
fmt.Println(sum)
// soal 5
kalimat := "halo halo bandung"
angka := "2022"
find = "halo";
replaceWith = "hi";
var newText2 = strings.Replace(kalimat, find, replaceWith, 2);
fmt.Print(newText2, " - ", angka)

}