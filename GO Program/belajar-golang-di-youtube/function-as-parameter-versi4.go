package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter4(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter4(name string) string {
	if name == "Jorok" {
		return "***"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter4("Endik", spamFilter4)
	sayHelloWithFilter4("Jorok", spamFilter4)

	filtered := spamFilter4
	sayHelloWithFilter4("Villy", filtered)
	sayHelloWithFilter4("Jorok", filtered)
}
