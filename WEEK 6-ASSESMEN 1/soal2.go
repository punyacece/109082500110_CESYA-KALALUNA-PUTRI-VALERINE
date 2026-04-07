package main
import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
    if tujuan == "domestik" {
        if jumlahHari > 3 {
            return 3
        }
        return jumlahHari
    } else if tujuan == "mancanegara" {
        if jumlahHari > 8 {
            return 8
        }
        return jumlahHari
    }
    return 0
}

// menghitung biaya domestik per hari untuk seluruh mahasiswa
func biayaPerHari(jumlahMhs int) int {
    totalPerMhs := (2 * 35000) + 250000 + 300000
    return totalPerMhs * jumlahMhs
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
    hariDitanggung := tanggunganHari(lamaPerjalanan, tujuan)

    biayaDomestikPerHari := float64(biayaPerHari(jumlahMhs))

    if tujuan == "domestik" {
        *totalBiaya = biayaDomestikPerHari * float64(hariDitanggung)
    } else if tujuan == "mancanegara" {
        // Biaya mancanegara adalah 1.5 kali biaya domestik
        *totalBiaya = (biayaDomestikPerHari * 1.5) * float64(hariDitanggung)
    }
}

func main() {
    var jumlah, lama int
    var tujuan string
    var biaya float64

    fmt.Print("masukkan jumlah mahasiswa : ")
    fmt.Scan(&jumlah)
    fmt.Print("masukkan lama hari study tour : ")
    fmt.Scan(&lama)
    fmt.Print("masukkan tujuan study tour (domestik/mancanegara) : ")
    fmt.Scan(&tujuan)

    perhitunganBiaya(jumlah, lama, tujuan, &biaya)

    // biaya
    fmt.Printf("\nBiaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biaya)
}