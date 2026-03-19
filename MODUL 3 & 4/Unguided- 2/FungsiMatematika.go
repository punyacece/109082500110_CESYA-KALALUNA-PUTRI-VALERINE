package main
import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	fmt.Print("Masukkan nilai a : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b : ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai c : ")
	fmt.Scan(&c)

	// Baris 1: 
	hasilA := f(g(h(a)))
	fmt.Printf("f(g(h( %d ))) : %d\n", a, hasilA)

	// Baris 2: 
	hasilB := g(h(f(b)))
	fmt.Printf("g(h(f( %d ))) : %d\n", b, hasilB)

	// Baris 3: 
	hasilC := h(f(g(c)))
	fmt.Printf("h(f(g( %d ))) : %d\n", c, hasilC)
}