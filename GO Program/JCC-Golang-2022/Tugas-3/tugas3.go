package main

import (
	"fmt"
	"strconv"
)

func main() {
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var panjang, err1 = strconv.Atoi(panjangPersegiPanjang)
	var lebar, err2 = strconv.Atoi(lebarPersegiPanjang)

	var kelilingPersegiPanjang int = panjang*2 + lebar*2

	if err1 == nil || err2 == nil {
		fmt.Println(kelilingPersegiPanjang)
	}
}
