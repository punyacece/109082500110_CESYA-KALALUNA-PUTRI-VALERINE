package main
import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var pemenang []string // Slice (dynamic array) untuk menyimpan nama pemenang

	// 1. Input nama klub
	fmt.Print("Klub A : ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scanln(&klubB)

	i := 1
	for {
		var skorA, skorB int
		fmt.Printf("Pertandingan %d : ", i)
		fmt.Scan(&skorA, &skorB)

		// 2. Cek kondisi berhenti (skor negatif)
		if skorA < 0 || skorB < 0 {
			break
		}

		// 3. Logika penentuan pemenang
		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}
		i++
	}

	// 4. Menampilkan hasil yang tersimpan di array
	fmt.Println()
	for j, hasil := range pemenang {
		fmt.Printf("Hasil %d : %s\n", j+1, hasil)
	}
	fmt.Println("Pertandingan selesai")
}