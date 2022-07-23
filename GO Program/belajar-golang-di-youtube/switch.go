package main

import "fmt"

func main() {
	name := "Endik"

	switch name {
	case "Endik":
		fmt.Println("Hello Endik")
		fmt.Println("Hai Endik")
	case "Hutomo":
		fmt.Println("Hello Hutomo")
		fmt.Println("Hai Hutomo")
	default:
		fmt.Println("Kenalan Donk")
		fmt.Println("Hey kamu!")
	}

	var name2 string = "Hutomo"

	switch name2 {
	case "Villy":
		fmt.Println("Hello Villy")
		fmt.Println("Hai Villy")
	case "Villain":
		fmt.Println("Hello Villain")
		fmt.Println("Hai Villain")
	default:
		fmt.Println("Siapa Kamu?")
		fmt.Println("Anda Siapa?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	switch length := len(name2); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}

}
