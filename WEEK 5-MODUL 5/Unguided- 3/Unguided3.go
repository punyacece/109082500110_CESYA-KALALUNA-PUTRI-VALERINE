package main
import "fmt"

func tampilkanGanjil(n int, current int) {
	if current > n {
		return
	}

	if current%2 != 0 {
		fmt.Printf("%d ", current)
	}

	tampilkanGanjil(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)

	fmt.Printf("Barisan bilangan ganjil dari 1 hingga %d: ", n)
	
	tampilkanGanjil(n, 1)
	fmt.Println()
}