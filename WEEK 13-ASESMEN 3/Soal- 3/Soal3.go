package main

import "fmt"

const NMAX = 1000000

// struct partai
type partai struct {
	nama  int
	suara int
}

// tipe tabPartai: array of partai dengan kapasitas NMAX
type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, nama int) int {
	/* mengembalikan indeks partai yang memiliki nama yang dicari
	   pada array t yang berisi n partai atau -1 apabila tidak
	   ditemukan, gunakan sekuensial search */
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

func main() {
	// deklarasi variabel
	var p tabPartai
	var n int = 0
	var inputNama int

	/* lakukan proses input suara secara berulang di sini, simpan
	   ke dalam array p, sehingga terdapat array p yang berisi hasil
	   peroleh suara n partai.*/
	fmt.Scan(&inputNama)
	for inputNama != -1 {
		idx := posisi(p, n, inputNama)
		if idx != -1 {
			// Jika partai sudah terdaftar, tambahkan jumlah suaranya
			p[idx].suara++
		} else {
			// Jika partai baru ditemukan, daftarkan ke dalam array
			p[n].nama = inputNama
			p[n].suara = 1
			n++
		}
		fmt.Scan(&inputNama)
	}

	/* lakukan proses pengurutan dengan insertion sort descending
	   Berdasarkan jumlah suara yang diperoleh. */
	for i := 1; i < n; i++ {
		kunci := p[i]
		j := i - 1
		for j >= 0 && p[j].suara < kunci.suara {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = kunci
	}

	// tampilkan array p
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Printf("%d(%d)\n", p[i].nama, p[i].suara)
		} else {
			fmt.Printf("%d(%d) ", p[i].nama, p[i].suara)
		}
	}
}