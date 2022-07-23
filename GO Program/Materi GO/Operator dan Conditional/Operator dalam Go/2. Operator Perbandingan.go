// Operator perbandingan digunakan untuk menentukan kebenaran suatu kondisi. Hasilnya berupa nilai boolean, true atau false.

// Tabel di bawah ini berisikan operator perbandingan yang bisa digunakan di Go.

// Operator		// Keterangan
// >			// lebih dari
// <			// kurang dari
// >=	 		// lebih dari sama dengan
// <=			// kurang dari sama dengan
// ==			// sama dengan
// !=			// tidak sama dengan

// contoh penggunaan:

var angka = 8
fmt.Println(angka > 5)  // true
fmt.Println(angka < 5)  // false
fmt.Println(angka >= 5) // true
fmt.Println(angka <= 5) // false
fmt.Println(angka == 5) // false
fmt.Println(angka != 5) // true

var angka1 = 10
fmt.Println(angka1 > 11)  // false
fmt.Println(angka1 < 11)  // true
fmt.Println(angka1 >= 11) // false
fmt.Println(angka1 <= 11) // true
fmt.Println(angka1 == 11) // false
fmt.Println(angka1 != 11) // true

var angka2 = 5
fmt.Println(angka2 > 1)  // true
fmt.Println(angka2 < 1)  // false
fmt.Println(angka2 >= 1) // true
fmt.Println(angka2 <= 1) // false
fmt.Println(angka2 == 1) // false
fmt.Println(angka2 != 1) // true