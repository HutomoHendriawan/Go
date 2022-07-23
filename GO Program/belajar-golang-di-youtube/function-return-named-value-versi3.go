package main

import "fmt"

func getNameScore2() (siswa1 string, nilai1 uint8, siswa2 string, nilai2 uint8, siswa3 string, nilai3 uint8, siswa4 string, nilai4 uint8, siswa5 string, nilai5 uint8) {
	siswa1 = "Andri"
	nilai1 = 70
	siswa2 = "Rusmana"
	nilai2 = 80
	siswa3 = "Hutomo"
	nilai3 = 96
	siswa4 = "Hendriawan"
	nilai4 = 90
	siswa5 = "Endik"
	nilai5 = 98

	return
}

func main() {
	a, b, c, d, e, f, g, h, i, j := getNameScore2()
	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(e, f)
	fmt.Println(g, h)
	fmt.Println(i, j)
}
