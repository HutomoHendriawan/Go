package main

import "fmt"

func sumAll2(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll2()
	fmt.Println(total)
}
