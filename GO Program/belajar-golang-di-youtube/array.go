package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Hutomo"
	names[1] = "Endik"
	names[2] = "Hendriawan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var names2 [3]string
	names2[0] = "Villy"
	names2[1] = "The"
	names2[2] = "Villain"

	var villy = names2[0]
	var the = names2[1]
	var villain = names2[2]

	fmt.Println(villy)
	fmt.Println(the)
	fmt.Println(villain)

	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)

	var values2 = [3]int{
		70,
		75,
		60,
	}

	fmt.Println(values2[0])
	fmt.Println(values2[1])
	fmt.Println(values2[2])

	var values3 = [3]int{
		50,
		55,
		40,
	}

	fmt.Println(values3)
	fmt.Println(values3[0])
	fmt.Println(values3[1])
	fmt.Println(values3[2])

	fmt.Println(len(names))
	fmt.Println(len(values))
	fmt.Println(len(names2))
	fmt.Println(len(values2))
	fmt.Println(len(values3))

	var contoh [10]string

	fmt.Println(len(contoh))

	var values4 = [3]int{
		92,
		83,
		96,
	}

	fmt.Println(values4)
	fmt.Println(len(values4))
	values4[0] = 100
	fmt.Println(values4)
}
