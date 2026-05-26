package main

import (
	"fmt"
)

const nMax = 7919

// Definisi struct Buku sesuai atribut di gambar
type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

// Menggunakan array dengan kapasitas nMax fix
type DaftarBuku [nMax]Buku

// 1. Prosedur untuk membaca masukan sejumlah n data buku
func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit,
			&pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

// 2. Prosedur untuk mencetak buku terfavorit (rating tertinggi) sebelum diurutkan
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		return
	}
	fav := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].rating > fav.rating {
			fav = pustaka[i]
		}
	}
	fmt.Printf("%s %s %s %d\n", fav.judul, fav.penulis, fav.penerbit, fav.tahun)
}

// 3. Prosedur untuk mengurutkan buku menurun berdasarkan rating menggunakan Insertion Sort
func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		// Karena diurutkan menurun (mengecil), geser elemen yang ratingnya lebih kecil
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

// 4. Prosedur untuk mencetak hingga 5 judul buku dengan rating tertinggi
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}
	for i := 0; i < limit; i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(pustaka[i].judul)
	}
	fmt.Println()
}

// 5. Prosedur mencari buku dengan rating tertentu menggunakan Pencarian Biner (Binary Search)
func CariBuku(pustaka DaftarBuku, n int, r int) {
	left := 0
	right := n - 1
	foundIdx := -1

	// Binary search pada array yang sudah terurut menurun
	for left <= right {
		mid := (left + right) / 2
		if pustaka[mid].rating == r {
			foundIdx = mid
			break // Berhenti jika ditemukan
		} else if pustaka[mid].rating < r {
			right = mid - 1 // Nilai dicari lebih besar, geser ke kiri karena terurut menurun
		} else {
			left = mid + 1 // Nilai dicari lebih kecil, geser ke kanan
		}
	}

	if foundIdx != -1 {
		b := pustaka[foundIdx]
		fmt.Printf("%s %s %s %d %d %d\n", b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var nPustaka int
	var ratingDicari int

	// Panggil rangkaian subprogram sesuai skenario tugas
	DaftarkanBuku(&pustaka, &nPustaka)
	CetakTerfavorit(pustaka, nPustaka)
	UrutBuku(&pustaka, nPustaka)
	Cetak5Terbaru(pustaka, nPustaka)

	// Membaca input terakhir untuk rating yang dicari
	fmt.Scan(&ratingDicari)
	CariBuku(pustaka, nPustaka, ratingDicari)
}