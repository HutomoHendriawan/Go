package main

import "fmt"

func getFullName() (string, string) {
	return "Hutomo", "Hendriawan"
}

func main() {
	getFullName()
	fmt.Println(getFullName())
}
