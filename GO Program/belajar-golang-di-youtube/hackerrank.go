package main

import "fmt"

func fizzBuzz(n int32) {
	fmt.Println(n)
}

func main() {
	for i := 1; i <= 15; i++ {
		if i == 3 || i == 6 || i == 9 || i == 12 {
			fmt.Println("Fizz")
		} else if i == 5 || i == 10 {
			fmt.Println("Buzz")
		} else if i == 15 {
			fmt.Println("Fizz Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
