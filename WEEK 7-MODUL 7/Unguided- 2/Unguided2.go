package main
import "fmt"

type angka int

type kata string

type Buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunTerbit   angka
	jumlahHalaman angka
}

func main() {
	var b Buku

	fmt.Println("=== INPUT BIODATA BUKU ===")
	fmt.Print("Masukkan judul buku : ")
	fmt.Scanln(&b.judul)
	fmt.Print("Masukkan nama penulis : ")
	fmt.Scanln(&b.penulis)
	fmt.Print("Masukkan penerbit : ")
	fmt.Scanln(&b.penerbit)
	fmt.Print("Masukkan tahun terbit : ")
	fmt.Scanln(&b.tahunTerbit)
	fmt.Print("Masukkan jumlah halaman: ")
	fmt.Scanln(&b.jumlahHalaman)

	// Bagian Output
	fmt.Println("\n=== BIODATA BUKU ===")
	fmt.Printf("Judul Buku : %s\n", b.judul)
	fmt.Printf("Penulis    : %s\n", b.penulis)
	fmt.Printf("Penerbit   : %s\n", b.penerbit)
	fmt.Printf("Tahun Terbit : %d\n", b.tahunTerbit)
	fmt.Printf("Jumlah Halaman : %d\n", b.jumlahHalaman)
}