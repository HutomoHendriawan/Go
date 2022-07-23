package main

import "fmt"

func main() {
	const firstName string = "Hutomo"
	const lastName = "Hendriawan"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	const (
		namaDepan        = "Villy"
		namaBelakang     = "Villain"
		nilai        int = 1000
	)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
	fmt.Println(nilai)
}
