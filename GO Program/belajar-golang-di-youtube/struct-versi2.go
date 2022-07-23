package main

import "fmt"

type Customer2 struct {
	Name, Adress string
	Age          int
}

func main() {
	var villy Customer2
	villy.Name = "Villain"
	villy.Adress = "Indonesia"
	villy.Age = 24

	fmt.Println(villy)

	cowboy := Customer2{
		Name:   "Space Cowboy",
		Adress: "Outer Space",
		Age:    30,
	}

	fmt.Println(cowboy)

	samurai := Customer2{"Space Samurai", "Outer", 30} // cara ini paling singkat, akan tetapi jika struct dirubah/edit maka code-nya akan error
	fmt.Println(samurai)
}
