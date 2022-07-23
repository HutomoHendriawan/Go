package main

import "fmt"

func getFullName3() (firstName, nickName, lastName string) {
	firstName = "Hutomo"
	nickName = "Endik"
	lastName = "Hendriawan"

	return
}

func main() {
	a, b, c := getFullName3()
	fmt.Println(a, b, c)
}
