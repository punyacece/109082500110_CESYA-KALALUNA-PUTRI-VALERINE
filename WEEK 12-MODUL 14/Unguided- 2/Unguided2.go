package main
import (
	"fmt"
)

// Fungsi untuk mengurutkan array secara membesar (Ascending)
func sortAscending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// Fungsi untuk mengurutkan array secara mengecil (Descending)
func sortDescending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIndex] {
				maxIndex = j
			}
		}
		arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
	}
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	for i := 0; i < n; i++ {
		var m int
		if _, err := fmt.Scan(&m); err != nil {
			break
		}

		// Siapkan wadah terpisah untuk menampung nomor ganjil dan genap
		var ganjil []int
		var genap []int

		for j := 0; j < m; j++ {
			var nomor int
			fmt.Scan(&nomor)
			// Pisahkan angka berdasarkan kondisinya
			if nomor%2 != 0 {
				ganjil = append(ganjil, nomor)
			} else {
				genap = append(genap, nomor)
			}
		}

		// Urutkan angka ganjil dari kecil ke besar
		sortAscending(ganjil)
		// Urutkan angka genap dari besar ke kecil
		sortDescending(genap)

		// Gabungkan hasil urutan ganjil dan genap untuk dicetak
		hasil := append(ganjil, genap...)

		// Cetak seluruh nomor rumah yang sudah rapi
		for k := 0; k < len(hasil); k++ {
			if k > 0 {
				fmt.Print(" ")
			}
			fmt.Print(hasil[k])
		}
		fmt.Println()
	}
}