package main
import (
    "fmt"
    "math"
)

const pi float64 = 3.14

func volume(r, t float64) float64 {
    return pi * r * r * t
}

func massa(r, t, p float64) float64 {
    return volume(r, t) * p
}

func display(m1, m2 float64) {
    if m1 == m2 {
        fmt.Println("\nBALANCE")
    } else {
        selisih := math.Abs(m1 - m2)
        fmt.Printf("\nSelisih massa zat cair kiri dan massa zat cair kanan : %g\n", selisih)
    }
}

func main() {
    var r float64 
    var tKiri, tKanan float64       
    var mjKiri, mjKanan float64     
    var massaKiri, massaKanan float64 

    fmt.Print("Masukkan jari-jari alas tabung : ")
    fmt.Scan(&r)

    // tinggi zat cair di tabung kiri, beserta massa jenisnya
    fmt.Print("Masukkan tinggi zat cair tabung kiri : ")
    fmt.Scan(&tKiri)
    fmt.Print("Masukkan massa jenis zat cair tabung kiri : ")
    fmt.Scan(&mjKiri)

    // tinggi zat cair di tabung kanan, beserta massa jenisnya
    fmt.Print("Masukkan tinggi zat cair tabung kanan : ")
    fmt.Scan(&tKanan)
    fmt.Print("Masukkan massa jenis zat cair tabung kanan : ")
    fmt.Scan(&mjKanan)

    // massa zat cair di tabung kiri dan kanan
    massaKiri = massa(r, tKiri, mjKiri)
    massaKanan = massa(r, tKanan, mjKanan)

    // hasil dari proses penimbangan
    display(massaKiri, massaKanan)
}