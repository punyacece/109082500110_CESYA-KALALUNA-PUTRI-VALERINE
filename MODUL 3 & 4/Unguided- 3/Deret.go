package main
import "fmt"

func cetakDeret(n int) {
	fmt.Print(n)

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		fmt.Printf(" %d", n)
	}
	fmt.Println()
}

func main() {
	var input int

	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&input)

	if input > 0 && input < 1000000 {
		cetakDeret(input)
	} else {
		fmt.Println("Input harus bilangan positif di bawah 1.000.000")
	}
}