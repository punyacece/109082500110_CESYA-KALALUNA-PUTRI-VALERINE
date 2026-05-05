package main
import "fmt"

// Definisi konstanta dan struct sesuai permintaan soal
const nMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

// Menggunakan array dengan kapasitas nMax
type arrayMahasiswa [nMax]mahasiswa

// Fungsi untuk mencari nilai pertama seorang mahasiswa dengan NIM tertentu
func cariNilaiPertama(T arrayMahasiswa, n int, nim string) int {
	for i := 0; i < n; i++ {
		if T[i].NIM == nim {
			return T[i].nilai
		}
	}
	return -1 // Mengembalikan -1 jika NIM tidak ditemukan
}

// Fungsi untuk mencari nilai terbesar seorang mahasiswa dengan NIM tertentu
func cariNilaiTerbesar(T arrayMahasiswa, n int, nim string) int {
	max := -1
	for i := 0; i < n; i++ {
		if T[i].NIM == nim {
			if T[i].nilai > max {
				max = T[i].nilai
			}
		}
	}
	return max
}

func main() {
	var data arrayMahasiswa
	var n int
	var searchNIM string

	// a. Menerima masukan sejumlah N data mahasiswa
	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	// Validasi agar tidak melebihi kapasitas array
	if n > nMax {
		n = nMax
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&data[i].NIM, &data[i].nama, &data[i].nilai)
	}

	// Meminta NIM yang ingin dicari
	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&searchNIM)

	// b, c, d. Mencari dan menampilkan hasil
	nilaiPertama := cariNilaiPertama(data, n, searchNIM)
	nilaiTerbesar := cariNilaiTerbesar(data, n, searchNIM)

	if nilaiPertama != -1 {
		fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", searchNIM, nilaiPertama)
		fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", searchNIM, nilaiTerbesar)
	} else {
		fmt.Println("Data mahasiswa dengan NIM tersebut tidak ditemukan.")
	}
}