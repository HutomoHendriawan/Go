package main

import "fmt"

func endApp4() {
	message := recover()
	fmt.Println("Error dengan Message:", message)
	fmt.Println("Aplikasi Selesai")
}

func runApp4(error bool) {
	defer endApp4()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp4(false)
}
