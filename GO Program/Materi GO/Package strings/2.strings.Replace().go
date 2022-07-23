// Fungsi ini digunakan untuk replace atau mengganti bagian dari string dengan string tertentu. Jumlah substring yang di-replace bisa ditentukan,
// apakah hanya 1 string pertama, 2 string, atau kesemuanya.

var text = "banana"
var find = "a"
var replaceWith = "o"

var newText1 = strings.Replace(text, find, replaceWith, 1)
fmt.Println(nexText1) // "bonana"

var newText2 = strings.Replace(text, find, replaceWith, 2)
fmt.Println(newText2) // "bonona"

var newText3 = strings.Replace(text, find, replaceWith, -1)
fmt.Println(newText3) // "bonono"

var text = "monokorobo"
var find = "o"
var replaceWith = "i"

var newText1 = strings.Replace(text, find, replaceWith, 1)
fmt.Println(newText1) // "minokorobo"

var newText2 = strings.Replace(text, find, replaceWith, 2)
fmt.Println(newText2) // "minikorobo"

var newText3 = strings.Replace(text, find, replaceWith, -1)
fmt.Println(newText3) // "minikiribi"


// Penjelasan:
// 1. Pada contoh di atas, substring "a" pada string "banana" akan di-replace dengan string "o".
// 2. Pada newText1, hanya 1 huruf o saja yang tereplace karena maksimal substring yang ingin di-replace ditentukan 1.
// 3. Angka -1 akan menjadikan proses replace berlaku pada semua substring. Contoh bisa dilihat pada newText3.