package main

import "fmt"

func endApp3() {
	message := recover()
	fmt.Println("Error dengan Message:", message)
	fmt.Println("Aplikasi Selesai")
}

func runApp3(error bool) {
	defer endApp3()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp3(true)
}
