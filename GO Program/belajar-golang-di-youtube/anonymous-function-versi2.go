package main

import "fmt"

type Blacklist2 func(string) bool

func registerUser2(name string, blacklist Blacklist2) {
	if blacklist(name) {
		fmt.Println("Anda telah diblokir", name)
	} else {
		fmt.Println("Selamat datang", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Admin"
	}

	registerUser2("Admin", blacklist)
	registerUser2("Endik", blacklist)

	registerUser2("Root", func(name string) bool {
		return name == "Root"
	})
	registerUser2("Villy", func(name string) bool {
		return name == "Root"
	})
}
