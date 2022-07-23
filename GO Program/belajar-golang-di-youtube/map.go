package main

import "fmt"

func main() {

	person := map[string]string{
		"address": "Sukabumi",
		"name":    "Endik",
	}

	person["title"] = "Advertiser"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var person2 map[string]string = map[string]string{
		"address2": "Depok",
		"name2":    "Hutomo",
	}

	person2["title2"] = "Digital Marketing"

	fmt.Println(person2)
	fmt.Println(person2["name2"])
	fmt.Println(person2["title2"])
	fmt.Println(person2["address2"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Endik"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
