package main
import "fmt"

func main() {
	var x, y int
	var beratIkan [1000]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	var totalPerWadah float64
	var jumlahWadah int
	var totalSeluruhRerata float64

	for i := 0; i < x; i++ {
		totalPerWadah += beratIkan[i]
		// Jika wadah penuh atau ini ikan terakhir
		if (i+1)%y == 0 || i == x-1 {
			fmt.Printf("%.2f ", totalPerWadah)
			totalSeluruhRerata += totalPerWadah
			totalPerWadah = 0
			jumlahWadah++
		}
	}

	fmt.Println() // Pindah baris untuk output kedua
	if jumlahWadah > 0 {
		fmt.Printf("%.2f\n", totalSeluruhRerata/float64(jumlahWadah))
	}
}