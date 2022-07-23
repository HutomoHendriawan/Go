package main

import "fmt"

func main() {
	var counter = 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter = counter + 1
	}

	counter2 := 2

	for counter2 < 20 {
		fmt.Println("Perulangan ke", counter2)
		counter2++
	}

	for counter3 := 3; counter3 <= 30; counter3++ {
		fmt.Println("Perulangan ke", counter3)
	}

	// variable "i" di bawah adalah "index", para programmer biasanya menggunakan huruf "i" untuk "index".
	slice := []string{"Endik", "Hutomo", "Hendriawan", "Villy, Villain"}

	for i := 0; i < len(slice); i++ {
		// gunakan "fmt.Println(i)" sebelum "fmt.Println(slice[i])"
		// fmt.Println(i)
		fmt.Println(slice[i])
	}

	var slice2 = []string{"Spongebob", "Patrick", "Squidward"}

	var i2 = 0
	for i2 = 0; i2 < len(slice2); i2++ {
		fmt.Println(slice2[i2])
	}

	// for _, value := range slice {
	// fmt.Println(value)
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Endik"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
