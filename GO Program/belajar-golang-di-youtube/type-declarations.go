package main

import "fmt"

func main() {
	type TanggalLahir string
	type Married bool

	var TanggalLahirEndik TanggalLahir = "151297"
	var MarriedStatus Married = false
	fmt.Println(TanggalLahirEndik)
	fmt.Println(MarriedStatus)
}
