package main

import "fmt"

func getFullName() (string, string) {
	return "Hutomo", "Hendriawan"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	// fmt.Println(firstName)
	// fmt.Println(lastName)
}
