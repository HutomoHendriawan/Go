package main

import "fmt"

func main() {

	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 5
	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b
	var g = a % b
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	var h = 10
	h += 10

	var i = 10
	i -= 10

	var j = 10
	j *= 5

	var k = 20
	k /= 10

	var l = 10
	l %= 2

	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	a++
	b--
	fmt.Println(a)
	fmt.Println(b)

	var negative = -100
	var positive = +100
	fmt.Println(negative)
	fmt.Println(positive)
}
