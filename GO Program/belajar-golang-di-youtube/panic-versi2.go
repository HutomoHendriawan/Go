package main

import "fmt"

func endApp2() {
	fmt.Println("Aplikasi Selesai")
}

func runApp2(error bool) {
	defer endApp2()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp2(true)
}
