package main

import "fmt"

type Customer5 struct {
	Name, Address string
	Age           int
}

func (customerA Customer5) sayHu(name string) {
	fmt.Println("Hu", name, "My name is", customerA.Name)
}

func (customerB Customer5) sayHe() {
	fmt.Println("Heeee from", customerB.Name)
}

func main() {
	var endik Customer5
	endik.Name = "Endik"
	endik.Address = "Indonesia"
	endik.Age = 24

	endik.sayHu("Hutomo")
	endik.sayHe()
}
