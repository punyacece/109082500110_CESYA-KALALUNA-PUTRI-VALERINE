package main
import "fmt"

func factorial(n int) int {
	var hasil int = 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func kombinasi(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Print("masukkan nilai a : ")
	fmt.Scan(&a)
	fmt.Print("masukkan nilai b : ")
	fmt.Scan(&b)
	fmt.Print("masukkan nilai c : ")
	fmt.Scan(&c)
	fmt.Print("masukkan nilai d : ")
	fmt.Scan(&d)

	fmt.Printf("hasil permutasi %d dan %d adalah : %d\n", a, c, permutasi(a, c))
	fmt.Printf("hasil kombinasi %d dan %d adalah : %d\n", a, c, kombinasi(a, c))

	fmt.Printf("hasil permutasi %d dan %d adalah : %d\n", b, d, permutasi(b, d))
	fmt.Printf("hasil kombinasi %d dan %d adalah : %d\n", b, d, kombinasi(b, d))
}