nil bukan merupakan tipe data, melainkan sebuah nilai. Variabel yang isi nilainya nil berarti memiliki nilai kosong.

Semua tipe data yang sudah dibahas di atas memiliki zero value (nilai default tipe data). Artinya meskipun nilai variabel dideklarasikan
dengan tanpa nilai awal, tetap akan ada nilai default-nya.
- Zero Value dari string adalah **(string kosong)
- Zero Value dari bool adalah false
- Zero Value dari tipe numerik non-desimal adalah 0
- Zero Value dari tipe numerik desimal adalah 0,0.

Zero Value berbeda dengan nil.nil adalah nilai kosong, benar-benar kosong.nil tidak bisa digunakan pada tipe data yang sudah
dibahas di atas. Ada beberapa tipe data yang bisa di-set nilainya dengan nil, diantaranya:
- pointer
- tipe data fungsi
- slice
- map
- channel
- interface kosong atau interface{}
