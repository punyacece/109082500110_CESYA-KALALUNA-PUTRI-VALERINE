package main
import (
	"fmt"
	"math"
)

// Tipe bentukan titik untuk menyimpan koordinat
type titik struct {
	x, y int
}

// Tipe bentukan lingkaran untuk menyimpan titik pusat dan radius
type lingkaran struct {
	pusat  titik
	radius int
}

// Mengembalikan jarak antara titik p(x,y) dan titik q(x,y)
func jarak(p, q titik) float64 {
	dx := float64(p.x - q.x)
	dy := float64(p.y - q.y)
	return math.Sqrt(dx*dx + dy*dy)
}

// Mengembalikan true apabila titik p berada di dalam lingkaran c
func didalam(c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {
	// Menggunakan array untuk menyimpan 2 lingkaran sesuai soal
	var c [2]lingkaran
	var p titik

	// Masukan Baris 1: Lingkaran 1
	fmt.Scan(&c[0].pusat.x, &c[0].pusat.y, &c[0].radius)

	// Masukan Baris 2: Lingkaran 2
	fmt.Scan(&c[1].pusat.x, &c[1].pusat.y, &c[1].radius)

	// Masukan Baris 3: Titik sembarang
	fmt.Scan(&p.x, &p.y)

	// Cek posisi menggunakan fungsi didalam
	inL1 := didalam(c[0], p)
	inL2 := didalam(c[1], p)

	// Logika Keluaran
	if inL1 && inL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}