// Array berbentuk "[...]" atau memiliki angka di dalam tanda []
// Slice berbentuk "[]" tanpa angka (kosong)

package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// MENGUBAH array MAKA AKAN MENGUBAH slice JUGA!
	//months[5] = "Diubah"
	//fmt.Println(slice1)

	// MENGUBAH slice MAKA AKAN MENGUBAH array JUGA!
	//slice1[0] = "Mei diubah"
	//fmt.Println(months)

	var names = [15]string{
		"Endik",
		"Kevin",
		"Dani",
		"Rizki",
		"Reva",
		"Mita",
		"Silvi",
		"Lis",
		"Iqoh",
		"Rosa",
		"Luthfa",
		"Dewi",
		"Desi",
		"Mufi",
		"Linda",
	}

	var slice2 = names[0:5]
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	var slice3 = months[11:]
	fmt.Println(slice3)

	var slice4 = append(slice3, "Hutomo")
	fmt.Println(slice4)
	slice4[1] = "Desember Diubah"
	fmt.Println(slice4)

	fmt.Println(slice3)
	fmt.Println(months)

	var slice5 = months[2:6]
	fmt.Println(slice5)

	var slice6 = append(slice5, "Villy")
	fmt.Println(slice6)
	slice6[1] = "Desember Diubah"
	fmt.Println(slice6)

	fmt.Println(slice5)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Hutomo"
	newSlice[1] = "Hendriawan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}
