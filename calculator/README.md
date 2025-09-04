🧮 Algoritma Kalkulator Infix → Postfix
1. Tokenisasi

Input string dipecah jadi token.

Contoh:

3 + 5 * ( 2 - 1 )


Jadi:

["3", "+", "5", "*", "(", "2", "-", "1", ")"]

2. Konversi Infix ke Postfix (Shunting Yard)

Gunakan stack operator & output list.

Aturan:

Angka → masuk ke output.

Operator → keluarkan operator dengan prioritas ≥ dari stack, lalu push operator baru.

"(" → push ke stack.

")" → keluarkan operator sampai ketemu "(".

Sisa operator di stack → pindahkan ke output.

Contoh:

Infix:

3 + 5 * ( 2 - 1 )


Postfix (RPN):

3 5 2 1 - * +

3. Evaluasi Postfix

Gunakan stack angka.

Aturan:

Angka → push ke stack.

Operator → ambil 2 angka terakhir dari stack, lakukan operasi, push hasil.

Contoh evaluasi:

Token: 3 5 2 1 - * +

Langkah:

Push 3 → [3]

Push 5 → [3, 5]

Push 2 → [3, 5, 2]

Push 1 → [3, 5, 2, 1]

Token - → 2 - 1 = 1 → [3, 5, 1]

Token * → 5 * 1 = 5 → [3, 5]

Token + → 3 + 5 = 8 → [8]

👉 Hasil akhir = 8

4. Ringkasan Alur

Pecah input jadi token.

Ubah notasi infix → postfix dengan algoritma shunting yard.

Evaluasi postfix dengan stack.

Tampilkan hasil.