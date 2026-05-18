package main
import (
	"fmt"
)

func main() {
	var input int
	var totalSuaraMasuk int
	var totalSuaraSah int
	// Array untuk menampung suara calon 1 sampai 20
	// Indeks 0 tidak digunakan agar nomor calon sesuai indeks
	counts := make([]int, 21)

	for {
		fmt.Scan(&input)

		// Berhenti jika input adalah 0
		if input == 0 {
			break
		}

		totalSuaraMasuk++

		// Validasi: suara sah jika berada di antara 1 dan 20
		if input >= 1 && input <= 20 {
			totalSuaraSah++
			counts[input]++
		}
	}

	// Menampilkan hasil ringkasan
	fmt.Printf("Suara masuk: %d\n", totalSuaraMasuk)
	fmt.Printf("Suara sah: %d\n", totalSuaraSah)

	// Menampilkan daftar calon yang mendapatkan suara (diurutkan dari nomor terkecil)
	for i := 1; i <= 20; i++ {
		if counts[i] > 0 {
			fmt.Printf("%d: %d\n", i, counts[i])
		}
	}
}