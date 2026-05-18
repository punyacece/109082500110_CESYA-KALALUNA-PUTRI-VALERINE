package main
import "fmt"

const NMAX = 1000000
var data [NMAX]int

func main() {
    var n, k int
    // Membaca jumlah data (n) dan angka yang dicari (k)
    fmt.Scan(&n, &k)

    // Mengisi array menggunakan prosedur
    isiArray(n)

    // Mencari posisi menggunakan fungsi
    hasil := posisi(n, k)

    // Menampilkan output sesuai hasil pencarian
    if hasil == -1 {
        fmt.Println("TIDAK ADA")
    } else {
        fmt.Println(hasil)
    }
}

func isiArray(n int) {
    // Memasukkan n buah angka ke dalam array data
    for i := 0; i < n; i++ {
        fmt.Scan(&data[i])
    }
}

func posisi(n, k int) int {
    // Mencari angka k di dalam array secara berurutan
    for i := 0; i < n; i++ {
        if data[i] == k {
            return i // Mengembalikan indeks jika ditemukan
        }
    }
    return -1 // Mengembalikan -1 jika tidak ditemukan
}