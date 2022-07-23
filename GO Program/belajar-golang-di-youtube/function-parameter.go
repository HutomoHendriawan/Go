package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("Hutomo", "Hendriawan")
	sayHelloTo("Kamoltip", "Pla")
	firstName := "Villy"
	lastName := "Villain"
	sayHelloTo(firstName, lastName)
}
