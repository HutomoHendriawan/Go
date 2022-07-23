package main

import "fmt"

func main() {
	name := "Endik"
	counter := 0

	increment := func() {
		name := "Villy"
		fmt.Println("increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
