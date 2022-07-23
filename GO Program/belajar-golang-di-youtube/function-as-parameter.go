package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "***"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Endik", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Villy", filter)
	sayHelloWithFilter("Anjing", filter)
}
