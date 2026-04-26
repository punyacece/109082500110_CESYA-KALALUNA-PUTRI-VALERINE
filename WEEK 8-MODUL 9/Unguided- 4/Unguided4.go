package main
import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input rune
	*n = 0
	fmt.Print("Teks       : ")
	for *n < NMAX {
		fmt.Scanf("%c", &input)
		if input == '.' || input == '\n' || input == '\r' {
			break
		}
		// Abaikan spasi jika ada di antara karakter
		if input != ' ' {
			t[*n] = input
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	// Menggunakan algoritma swap (tukar) dari ujung ke ujung
	for i := 0; i < n/2; i++ {
		temp := t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	// Simpan array asli untuk dibandingkan nanti
	asli := t
	// Balikkan array t
	balikanArray(&t, n)
	
	// Bandingkan isi array asli dengan yang sudah dibalik
	for i := 0; i < n; i++ {
		if asli[i] != t[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	// Mengisi array
	isiArray(&tab, &m)

	// Cek Palindrom (dilakukan sebelum dibalik permanen di main agar logikanya jelas)
	isPal := palindrom(tab, m)

	// Balikkan isi array tab secara permanen
	fmt.Print("Reverse    : ")
	balikanArray(&tab, m)
	cetakArray(tab, m)

	// Menampilkan hasil palindrom
	fmt.Printf("Palindrom  : %t\n", isPal)
}