package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Villy"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := "Endik"
	fmt.Println(result) // cara 1

	fmt.Println(getHello("Endik")) // cara 2
}
