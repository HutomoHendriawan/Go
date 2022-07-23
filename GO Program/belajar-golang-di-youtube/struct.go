package main

import "fmt"

type Customer struct {
	Name   string
	Adress string
	Age    int
}

func main() {
	var endik Customer
	endik.Name = "Hutomo"
	endik.Adress = "Indonesia"
	endik.Age = 24

	fmt.Println(endik.Name)
	fmt.Println(endik.Adress)
	fmt.Println(endik.Age)
}
