package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke ", i)
	}

	for i2 := 0; i2 < 10; i2++ {
		fmt.Println("Perulangan ke ", i2)
		if i2 == 5 {
			break
		}
	}
}
