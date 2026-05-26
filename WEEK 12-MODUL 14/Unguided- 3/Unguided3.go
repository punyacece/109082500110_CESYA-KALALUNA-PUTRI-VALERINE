package main
import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan algoritma Insertion Sort
func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		// Geser elemen yang lebih besar dari key ke posisi setelahnya
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var arr []int
	var val int

	// Membaca masukan terus-menerus sampai ditemukan bilangan negatif
	for {
		fmt.Scan(&val)
		if val < 0 {
			break
		}
		arr = append(arr, val)
	}

	// Jika tidak ada data non-negatif yang dimasukkan, program selesai
	if len(arr) == 0 {
		return
	}

	// Mengurutkan data dengan Insertion Sort
	insertionSort(arr)

	// Menampilkan data yang sudah terurut
	for i := 0; i < len(arr); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(arr[i])
	}
	fmt.Println()

	// Memeriksa status jarak (selisih) antar bilangan
	if len(arr) < 2 {
		// Jika data hanya 1, otomatis dianggap berjarak tetap (misal jarak 0 atau tidak didefinisikan)
		fmt.Println("Data berjarak 0")
		return
	}

	// Ambil selisih pertama sebagai patokan awal
	jarakAwal := arr[1] - arr[0]
	berjarakTetap := true

	for i := 1; i < len(arr)-1; i++ {
		selisih := arr[i+1] - arr[i]
		if selisih != jarakAwal {
			berjarakTetap = false
			break
		}
	}

	// Menampilkan status jarak sesuai dengan hasil pemeriksaan
	if berjarakTetap {
		fmt.Printf("Data berjarak %d\n", jarakAwal)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}