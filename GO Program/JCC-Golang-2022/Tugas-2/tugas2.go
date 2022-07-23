package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Soal 1
	var text1 = "Jabar"
	fmt.Println(text1)

	var text2 = "Coding"
	fmt.Println(text2)

	var text3 = "Camp"
	fmt.Println(text3)

	var text4 = "GOlang"
	fmt.Println(text4)

	var text5 = "2022"
	fmt.Println(text5)

	// Soal 2
	halo := "Halo Dunia"
	find := "Dunia"
	replaceWith := "Golang"

	newHalo := strings.Replace(halo, find, replaceWith, 1)
	fmt.Println(newHalo)

	// Soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	kataPertama = "saya "
	kataKedua = "Senang "
	kataKetiga = "belajaR "
	kataKeempat = "GOLANG"

	var sentence = kataPertama + kataKedua + kataKetiga + kataKeempat
	fmt.Println()
	fmt.Println(sentence)

	// Soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	var num1, err1 = strconv.Atoi(angkaPertama)
	var num2, err2 = strconv.Atoi(angkaKedua)
	var num3, err3 = strconv.Atoi(angkaKetiga)
	var num4, err4 = strconv.Atoi(angkaKeempat)

	if err1 == nil || err2 == nil || err3 == nil || err4 == nil {
		fmt.Println(num1 + num2 + num3 + num4)
	}

	// Soal 5
	kalimat := "halo halo bandung"
	angka := 2022

	find2 := "halo"
	replaceWith2 := "Hi"

	newKalimat := strings.Replace(kalimat, find2, replaceWith2, 2)
	fmt.Println(`"`+newKalimat+`"`, "-", angka)
}
