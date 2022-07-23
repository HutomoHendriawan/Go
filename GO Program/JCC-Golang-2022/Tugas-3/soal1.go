package main

import (
	"fmt"
	"strconv"
)

func main() {
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var panjang, err1 = strconv.Atoi(panjangPersegiPanjang)
	var lebar, err2 = strconv.Atoi(lebarPersegiPanjang)

	var alas, err3 = strconv.Atoi(alasSegitiga)
	var tinggi, err4 = strconv.Atoi(tinggiSegitiga)

	var kelilingPersegiPanjang int = panjang*2 + lebar*2
	var luasSegitiga int = alas * tinggi / 2

	if err1 == nil || err2 == nil || err3 == nil || err4 == nil {
		fmt.Println(kelilingPersegiPanjang)
		fmt.Println(luasSegitiga)
	}
}
