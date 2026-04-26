package main
import (
	"fmt"
	"math"
)

func main() {
	var n, x, idx, target int
	fmt.Print("N: ")
	fmt.Scan(&n)

	// Inisialisasi slice dengan panjang N
	data := make([]int, n)
	sum := 0.0
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
		sum += float64(data[i])
	}

	// a, b, c. Menampilkan isi, indeks ganjil, dan indeks genap
	fmt.Println("Isi:", data)
	fmt.Print("Ganjil: ")
	for i := 1; i < n; i += 2 { fmt.Printf("%d ", data[i]) }
	fmt.Print("\nGenap: ")
	for i := 0; i < n; i += 2 { fmt.Printf("%d ", data[i]) }

	// d. Kelipatan x
	fmt.Print("\nMasukkan x: ")
	fmt.Scan(&x)
	for i := 0; x > 0 && i < n; i++ {
		if i%x == 0 { fmt.Printf("%d ", data[i]) }
	}

	// f & g. Statistik (Rata-rata & Standar Deviasi)
	mean := sum / float64(n)
	var v float64
	for _, vVal := range data { v += math.Pow(float64(vVal)-mean, 2) }
	fmt.Printf("\nMean: %.2f\nStdDev: %.2f", mean, math.Sqrt(v/float64(n)))

	// h. Frekuensi
	fmt.Print("\nCari angka: ")
	fmt.Scan(&target)
	count := 0
	for _, v := range data {
		if v == target { count++ }
	}
	fmt.Printf("Frekuensi %d: %d", target, count)

	// e. Hapus elemen (Teknik Slice Truncation)
	fmt.Print("\nHapus indeks: ")
	fmt.Scan(&idx)
	if idx >= 0 && idx < len(data) {
		data = append(data[:idx], data[idx+1:]...)
		fmt.Println("Setelah hapus:", data)
	}
}