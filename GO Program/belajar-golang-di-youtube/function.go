package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World") //Bebas dibuat berapa banyak
}

func main() {
	for i := 0; i < 10; i++ {
		sayHello()
	}
	sayHello() //Bebas dibuat berapa banyak
}
