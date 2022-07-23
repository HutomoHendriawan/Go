package main

import "fmt"

func arrange(sentence string) {
	fmt.Println(sentence)
}

func main() {
	sentence := "Hutomo Endik"
	arrange(sentence)

	if length := len(sentence); length < 3 {
		fmt.Println(sentence)
	} else {
		fmt.Println("Endik Hutomo")
	}
}
