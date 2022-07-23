package main

import "fmt"

func getFullName2() (firstName string, nickName string, lastName string) {
	firstName = "Hutomo"
	nickName = "Endik"
	lastName = "Hendriawan"

	return firstName, nickName, lastName
}

func main() {
	firstName, nickName, lastName := getFullName2() // variable firstName, nickName, dan lastName di baris ini dapat diganti dengan kata/huruf/angka apa saja
	fmt.Println(firstName)
	fmt.Println(nickName)
	fmt.Println(lastName)
}
