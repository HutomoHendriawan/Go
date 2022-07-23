package main

import "fmt"

func endApp5() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan Message:", message)
	}
	fmt.Println("Aplikasi Selesai")
}

func runApp5(error bool) {
	defer endApp5()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp5(false)
}
