package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan algoritma Selection Sort
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Cari indeks dengan nilai terkecil di sisa array
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// Tukar nilai terkecil yang ditemukan dengan elemen di indeks i
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	var n int
	// Membaca banyaknya daerah (n)
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	// Mengolah data untuk setiap daerah
	for i := 0; i < n; i++ {
		var m int
		// Membaca banyaknya kerabat di daerah tersebut (m)
		if _, err := fmt.Scan(&m); err != nil {
			break
		}

		// Membuat slice untuk menampung nomor rumah sebanyak m elemen
		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		// Mengurutkan data menggunakan Selection Sort
		selectionSort(rumah)

		// Menampilkan hasil yang sudah terurut
		for k := 0; k < m; k++ {
			if k > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumah[k])
		}
		fmt.Println()
	}
}