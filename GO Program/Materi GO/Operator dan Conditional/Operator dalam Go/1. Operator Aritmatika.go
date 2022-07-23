// Operator aritmatika adalah operator yang digunakan untuk operasi yang sifatnya perhitungan. Go mendukung beberapa operator aritmatika standar, list-nya bisa dilihat di tabel berikut.

// Operator	 	// Keterangan
// 	+			// penjumlahan
// 	-			// pengurangan
// 	*			// perkalian
// 	/			// pembagian
// 	%			// modulus / sisa hasil pembagian

// contoh penggunaan:

// operator penjumlahan
jumlah := 8 + 3
fmt.Println(jumlah) // 11

jumlah := 1 + 1
fmt.Println(jumlah) // 2

jumlah := 12 + 2
fmt.Println(jumlah) // 14

// opperator pengurangan
kurang := 8 - 3
fmt.Println(kurang) // 5

kurang := 10 - 5
fmt.Println(kurang) // 5

kurang := 2 - 1
fmt.Println(kurang) // 1

// operator perkalian
kali := 2 * 2
fmt.Println(kali) // 4

kali := 8 * 3
fmt.Println(kali) // 24

kali := 10 * 10
fmt.Println(kali) // 100

// operator pembagian
bagi := 24 / 3
fmt.Println(bagi) // 8

bagi := 8 / 4
fmt.Println(bagi) // 2

bagi := 50 / 10
fmt.Println(bagi) // 5

// operator modulus

modulus := 8 % 3
fmt.Println(modulus) // 2

modulus := 8 % 3
fmt.Println(modulus) // 2

modulus := 8 % 3
fmt.Println(modulus) // 2

// Operasi Matematika 		// Augmented Assignments
// a = a + 10 				// a += 10
// a = a - 10 				// a -= 10
// a = a * 10 				// a *= 10
// a = a / 10 				// a /= 10
// a = a % 10 				// a %= 10

// contoh penggunaan:
var angka = 8
fmt.Println(angka) // 8
angka += 10
fmt.Println(angka) // 18

var angka2 = 5
fmt.Println(angka2) // 5
angka2 += 5
fmt.Println(angka2) // 10

var angka3 = 11
fmt.Println(angka3) // 11
angka3 += 9
fmt.Println(angka3) // 20

var angka4 = 10
fmt.Println(angka4) // 10
angka4 += 5
fmt.Println(angka4) // 15

var angka5 = 2
fmt.Println(angka5) // 2
angka5 += 1
fmt.Pringln(angka5) // 3

var angka6 = 4
fmt.Println(angka6) // 4
angka6 += 6
fmt.Println(angka6) // 10

// selain itu terdapat juga unary operator berikut listnya:
// Operator		// Keterangan
// ++			// a = a + 1
// --			// a = a - 1
// -			// negative
// +			// positive
// !			// negasi/ kebalikan dari tipe data boolean

// contoh penggunaan:

angka := 8
fmt.Println(angka) // 8
angka++
fmt.Println(angka) // 9

angka2 := 5
fmt.Println(angka2) // 5
angka--
fmt.Println(angka) // 4

angka4 := 1
fmt.Println(angka3) // 1
angka3++
fmt.Println(angka3) // 2

angka4 := 8
fmt.Println(angka4) // 8
angka4--
fmt.Println(angka4) // 7

angka5 := 10
fmt.Println(angka5) // 10
angka5++
fmt.Println(angka5) // 11

angka6 := 11
fmt.Println(angka6) // 11
angka6--
fmt.Println(angka6) // 10