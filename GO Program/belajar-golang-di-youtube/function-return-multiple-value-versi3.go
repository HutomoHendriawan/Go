package main

import "fmt"

func getFullName() (string, string) {
	return "Villy", "Villain"
}

func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
