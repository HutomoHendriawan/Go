package main

import "fmt"

func main() {
	var ujian = 83
	var absensi = 90

	var lulusUjian bool = ujian > 75
	var lulusAbsensi bool = absensi > 75
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus bool = lulusUjian && lulusAbsensi
	fmt.Println(lulus)

	var ulangan = 86
	var kehadiran = 90

	fmt.Println(ulangan >= 80 && kehadiran > 60)
}
