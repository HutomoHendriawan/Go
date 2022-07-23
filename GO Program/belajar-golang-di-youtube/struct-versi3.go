package main

import "fmt"

type Customer3 struct {
	Name, Adress string
	Age          int
}

func sayHello(customer Customer3, name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var endik Customer3
	endik.Name = "Endik"
	endik.Adress = "Indonesia"
	endik.Age = 24

	sayHello(endik, "Hutomo")
}
