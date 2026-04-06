package main
import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	// Langkah rekursif: Sn = Sn-1 + Sn-2
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("Deret Fibonacci hingga suku ke-10:")
	fmt.Println("-----------------------------------")
	fmt.Printf("| %-2s | %-2s |\n", "n", "Sn")
	fmt.Println("-----------------------------------")

	// Menampilkan hasil dari suku ke-0 hingga ke-10
	for i := 0; i <= 10; i++ {
		result := fibonacci(i)
		fmt.Printf("| %-2d | %-2d |\n", i, result)
	}
	fmt.Println("-----------------------------------")
}