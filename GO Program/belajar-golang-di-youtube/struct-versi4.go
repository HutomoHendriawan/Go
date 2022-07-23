package main

import "fmt"

type Customer4 struct {
	Name, Adress string
	Age          int
}

func (customer Customer4) sayHi(name string) {
	fmt.Println("Hi", name, "My name is", customer.Name)
}

func main() {
	var villy Customer4
	villy.Name = "Villy"
	villy.Adress = "Space"
	villy.Age = 24

	villy.sayHi("Villain")
}
