package main
import "fmt"

func cekGenap(angka int) bool {
	// Jika angka habis dibagi 2 (sisa baginya 0)
	if angka%2 == 0 {
		return true
	}
	// Jika tidak, kembalikan false
	return false
}

func main() {
	angkaTes := 9
	
	// Memanggil fungsi dan menyimpan hasil kembaliannya
	hasilGenap := cekGenap(angkaTes)
	
	fmt.Printf("Apakah %d itu bilangan genap? Jawabannya: %t\n", angkaTes, hasilGenap)
}	