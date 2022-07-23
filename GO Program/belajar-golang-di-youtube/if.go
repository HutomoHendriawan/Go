package main

import "fmt"

func main() {
	var name = "Endik"

	if name == "Endik" {
		fmt.Println("Hai Endik")
	} else {
		fmt.Println("Salah!")
	}

	name2 := "Hutomo"

	if name2 == "Hutomo" {
		fmt.Println("Halo", name2)
	} else {
		fmt.Println("Sok Tahu")
	}

	var name3 = "Hendriawan"

	if name3 == "Hutomo" {
		fmt.Println("Halo Hutomo")
	} else {
		fmt.Println("Hai Kamu, Iya kamu")
	}

	name4 := "Villy"

	if name4 == "Villain" {
		fmt.Println("Apakah kamu", name4, "?")
	} else {
		fmt.Println("Bukan!")
	}

	var length = len(name)
	if length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}

	if length := len(name2); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}

}
