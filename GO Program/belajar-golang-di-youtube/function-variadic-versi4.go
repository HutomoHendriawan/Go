package main

import "fmt"

func sumAll4(numbers ...int) int {
	total := 2
	for _, value := range numbers {
		total *= value
	}
	return total
}

func main() {
	total := sumAll4(10, 10, 10, 10, 10)
	fmt.Println(total)

	slice := []int{20, 20, 20, 20, 20}
	total = sumAll4(slice...)
	fmt.Println(total)

}
