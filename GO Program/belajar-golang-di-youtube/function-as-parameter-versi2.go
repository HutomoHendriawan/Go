package main

import "fmt"

func sayHelloWithFilter2(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter2(name string) string {
	if name == "Babi" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter2("Endik", spamFilter2)
	sayHelloWithFilter2("Babi", spamFilter2)

	filtered := spamFilter2
	sayHelloWithFilter2("Villy", filtered)
	sayHelloWithFilter2("Babi", filtered)
}
