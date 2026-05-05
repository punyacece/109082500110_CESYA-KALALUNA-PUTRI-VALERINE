package main

import (
	"fmt"
)

// Konfigurasi konstanta dan tipe data sesuai spesifikasi soal
const nProv int = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

// Procedure untuk menginput data ke dalam array
func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

// Function untuk mencari indeks provinsi dengan pertumbuhan tercepat
func ProvinsiTercepat(tumbuh TumbuhProv) int {
	idxMax := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idxMax] {
			idxMax = i
		}
	}
	return idxMax
}

// Function untuk mencari indeks berdasarkan nama provinsi
func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

// Procedure untuk menghitung dan menampilkan prediksi penduduk (pertumbuhan > 2%)
func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("\n=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			// Rumus: Prediksi = (Angka Pertumbuhan + 1) * Populasi
			hasil := float64(pop[i]) * (tumbuh[i] + 1)
			fmt.Printf("%s %.0f\n", prov[i], hasil)
		}
	}
}

func main() {
	var provinsi NamaProv
	var populasi PopProv
	var pertumbuhan TumbuhProv
	var namaCari string

	// 1. Input data 10 provinsi
	InputData(&provinsi, &populasi, &pertumbuhan)

	// 2. Input nama provinsi yang ingin dicari indeksnya
	fmt.Scan(&namaCari)

	// 3. Output Baris Pertama: Provinsi dengan pertumbuhan tercepat
	idxCepat := ProvinsiTercepat(pertumbuhan)
	fmt.Printf("\nProvinsi dengan angka pertumbuhan tercepat : %s\n", provinsi[idxCepat])

	// 4. Output Baris Kedua: Indeks provinsi yang dicari
	idxCari := IndeksProvinsi(provinsi, namaCari)
	fmt.Printf("Data provinsi yang dicari : %s (Indeks: %d)\n", namaCari, idxCari)

	// 5. Output Baris Sisanya: Prediksi untuk pertumbuhan > 2%
	Prediksi(provinsi, populasi, pertumbuhan)
}