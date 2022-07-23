package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter3(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter3(name string) string {
	if name == "Kasar" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter3("Endik", spamFilter3)
	sayHelloWithFilter3("Kasar", spamFilter3)

	filtered := spamFilter3
	sayHelloWithFilter3("Villy", filtered)
	sayHelloWithFilter3("Kasar", filtered)
}
