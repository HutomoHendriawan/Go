package main

import "fmt"

func getHello(name string) string {
	result := "Hello "
	fmt.Println(result, name)
	return result
}

func main() {
	getHello("Endik")
}
