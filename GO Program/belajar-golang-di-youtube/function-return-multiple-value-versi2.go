package main

import "fmt"

func getNameScore() (string, uint8) {
	return "Andri", 80
}

func main() {
	Name, Score := getNameScore()
	fmt.Println(Name, Score)
}
