package main

import "fmt"

func getGoodbye2(name string) string {
	return "Good Bye " + name
}

func main() {
	SayGoodBye := getGoodbye2

	result := SayGoodBye("Endik")
	fmt.Println(result)
}
