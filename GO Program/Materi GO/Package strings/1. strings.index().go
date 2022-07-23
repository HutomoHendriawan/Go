// Go menyediakan package strings, isinya banyak fungsi untuk keperluan pengolahan data string.
// strings.index() Digunakan untuk mencari posisi indeks sebuah string (parameter kedua) dalam string (parameter pertama).

var index1 = strings.Index("ethan hunt", "ha")
fmt.Println(index1) // 2

var index1 = strings.Index("ethan hunt", "ha")
fmt.Println(index1) // 2 

var index1 = strings.Index("ethan hunt", "ha")
fmt.Println(index1) // 2

var index1 = strings.Index("ethan hunt", "ha")
fmt.Println(index1) // 2

var index1 = strings.Index("ethan hunt", "ha")
fmt.Println(index1) // 2

var index1 = strings.Index("ethan hunt", "ha")
fmt.Println(index1) // 2

var index1 = strings.Index("ethan hunt", "ha")
fmt.Println(index1) // 2

var index1 = strings.Index("ethan hunt", "ha")
fmt.Println(index1) // 2

var index1 = strings.Index("ethan hunt", "ha")
fmt.Println(index1) // ha

var index1 = strings.Index("ethan hunt", "ha")
fmt.Println(index1) // 2

// String "ha" berada pada posisi ke 2 dalam string "ethan hunt" (indeks dimulai dari 0). Jika diketemukan dua substring, maka yang diambil adalah yang pertama, contoh:

var index2 = strings.Index("ethan hunt", "n")
fmt.Println(index2) // 4

var index2 = strings.Index("ethan hunt", "n")
fmt.Println(index2) // 4

var index2 = strings.Index("ethan hunt", "n")
fmt.Println(index2) // n

// String "n" berada pada indeks 4 dan 8. Yang dikembalikan adalah yang paling kiri (paling kecil), yaitu 4.