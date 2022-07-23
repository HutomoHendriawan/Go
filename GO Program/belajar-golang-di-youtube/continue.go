package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke ", i)
	}

	for i2 := 0; i2 < 10; i2++ {
		if i2%2 == 1 {
			continue
		}
		fmt.Println("Perulangan ke ", i2)
	}
}
