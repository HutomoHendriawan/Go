package main

import "fmt"

func getGoodbye(name string) string {
	return "Good Bye " + name
}

func main() {
	Goodbye := getGoodbye
	fmt.Println(Goodbye("Endik"))
}
