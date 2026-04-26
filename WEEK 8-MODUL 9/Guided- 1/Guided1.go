package main
import "fmt"

func main() {
    // --- CONTOH ARRAY (Statis) ---
    // Harus tentukan ukuran [4]
    var rakSepatu = [3]string{"Nike", "Adidas", "Jordan"}
    fmt.Println("Array:", rakSepatu)

    // --- CONTOH SLICE (Dinamis) ---
    // Ukuran kosong [], bebas tambah data
    var keranjangBelanja []string 

    // Menambah data pakai append
    keranjangBelanja = append(keranjangBelanja, "Apel")
    keranjangBelanja = append(keranjangBelanja, "Mangga", "Jeruk")

    fmt.Println("Slice awal:", keranjangBelanja)
    fmt.Println("Jumlah data (len):", len(keranjangBelanja)) // 

    // Mengambil potongan (Slicing) 
    // Ambil indeks 0 sampai sebelum 2 (yaitu 0 dan 1)
    potongan := keranjangBelanja[0:2] 
    fmt.Println("Hasil Slicing [0:2]:", potongan)
}