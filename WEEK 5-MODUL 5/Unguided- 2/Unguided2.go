package main
import "fmt"

func tampilkanFaktor(n int, i int) {
	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Printf("%d ", i)
	}

	tampilkanFaktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)

	fmt.Printf("Faktor dari %d adalah: ", n)
	
	// Mulai rekursi dari angka 1
	tampilkanFaktor(n, 1)
	fmt.Println() 
}