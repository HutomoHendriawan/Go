package main

import "fmt"

func endApp6() {
	message := recover()
	if message != nil {
		fmt.Println("Error denga Message:", message)
	}
	fmt.Println("Aplikasi Selesai")
}

func runApp6(error bool) {
	defer endApp6()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp6(true)
	// Untuk memastikan apakah benar aplikasinya tidak berhenti, coba tambahkan println di bawah sini
	fmt.Println("Berhasil")
}
