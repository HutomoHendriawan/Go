// Operator ini digunakan untuk mencari benar tidaknya kombinasi data bertipe bool (bisa berupa variabel bertipe bool, atau hasil dari operator perbandingan).

// Beberapa operator logika standar yang bisa digunakan:

// Operator		// Keterangan

// &&			// dan
// ||			// atau
// !			// negasi / nilai kebalikan

// contoh penggunaan:
var a = true
var b = false
var c = true
var d = false

fmt.Println(a && c)   // true
fmt.Println(a && b)   // false
fmt.Println(a || b)   // true
fmt.Println(b || d)   // false
fmt.Println(!b && !d) // true
fmt.Println(!a || b)  // false

var a = true
var b = false
var c = true
var d = false

fmt.Println(a && b)   // false
fmt.Println(a && c)   // true
fmt.Println(a || b)   // true
fmt.Println(b || d)   // false
fmt.Println(!a && !c) // true
fmt.Println(!a || b)  // false

