package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pemain struct {
	Nama        string
	Gol, Assist int
}

func main() {
	var n int
	fmt.Println("Masukkan Data Input :")
	fmt.Scanln(&n)

	// Array statik dengan kapasiti maksimum 1001 sesuai spesifikasi soal
	var list [1001]Pemain
	scanner := bufio.NewScanner(os.Stdin)

	// Membaca data setiap pemain baris demi baris
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			baris := scanner.Text()
			elemen := strings.Fields(baris)
			panjang := len(elemen)

			// Ekstrak angka gol dan assist dari hujung baris
			assist, _ := strconv.Atoi(elemen[panjang-1])
			gol, _ := strconv.Atoi(elemen[panjang-2])
			
			// Mencantumkan semula nama jika mengandungi spasi
			nama := strings.Join(elemen[:panjang-2], " ")

			list[i] = Pemain{Nama: nama, Gol: gol, Assist: assist}
		}
	}

	// Proses mengisih (Sorting) menggunakan Insertion Sort secara Descending
	for i := 1; i < n; i++ {
		kunci := list[i]
		j := i - 1

		// Perbandingan utama: Gol terbesar. Perbandingan kedua: Assist terbesar jika gol sama.
		for j >= 0 && (list[j].Gol < kunci.Gol || (list[j].Gol == kunci.Gol && list[j].Assist < kunci.Assist)) {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = kunci
	}

	// Memaparkan hasil akhir sesuai format terminal ekspektasi
	fmt.Println("\nHasil Sorting :")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %d\n", list[i].Nama, list[i].Gol, list[i].Assist)
	}
}