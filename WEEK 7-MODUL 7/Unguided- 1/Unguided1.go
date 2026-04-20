package main
import "fmt"

type suhu float64

func CelciusToReamur(Celcius suhu) suhu {
	return Celcius * 0.8
}

func CelciusToFahrenheit(Celcius suhu) suhu {
	return (Celcius * 1.8) + 32
}

func CelciusToKelvin(Celcius suhu) suhu {
	return Celcius + 273.15
}

func main() {
	var input suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&input)

	fmt.Println("")
	fmt.Printf("%.0f celcius = %.1f reamur\n", input, CelciusToReamur(input))
	fmt.Printf("%.0f celcius = %.1f fahrenheit\n", input, CelciusToFahrenheit(input))
	fmt.Printf("%.0f celcius = %.2f kelvin\n", input, CelciusToKelvin(input))
}