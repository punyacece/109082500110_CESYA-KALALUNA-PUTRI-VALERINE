package main
import (
	"fmt"
)

func main() {
	var input int
	var totalSuaraMasuk int
	var totalSuaraSah int
	counts := make([]int, 21)

	// Membaca input sampai angka 0
	for {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		totalSuaraMasuk++
		if input >= 1 && input <= 20 {
			totalSuaraSah++
			counts[input]++
		}
	}

	// Mencari Ketua (suara terbanyak 1) dan Wakil (suara terbanyak 2)
	idKetua, idWakil := -1, -1
	max1, max2 := -1, -1

	for i := 1; i <= 20; i++ {
		if counts[i] > max1 {
			// Geser juara 1 lama menjadi juara 2
			max2 = max1
			idWakil = idKetua
			// Juara 1 baru
			max1 = counts[i]
			idKetua = i
		} else if counts[i] > max2 {
			// Update juara 2 jika lebih besar dari max2 saat ini
			max2 = counts[i]
			idWakil = i
		}
	}

	// Menampilkan hasil
	fmt.Printf("Suara masuk: %d\n", totalSuaraMasuk)
	fmt.Printf("Suara sah: %d\n", totalSuaraSah)
	if idKetua != -1 {
		fmt.Printf("Ketua RT: %d\n", idKetua)
	}
	if idWakil != -1 {
		fmt.Printf("Wakil ketua: %d\n", idWakil)
	}
}