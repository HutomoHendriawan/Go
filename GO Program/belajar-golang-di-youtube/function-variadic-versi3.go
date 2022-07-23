package main

import "fmt"

func sumAll3(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll3(10, 10, 10, 10, 10)
	fmt.Println(total)

	slice := []int{20, 20, 20, 20, 20}
	total = sumAll3(slice...)
	fmt.Println(total)
}
