package main
import "fmt"

// Mendefinisikan tipe set sebagai array dengan kapasitas 2022
type set [2022]int

/* mengembalikan true apabila bilangan val ada di dalam array T 
   yang berisi sejumlah n bilangan bulat */
func exist(T set, n int, val int) bool {
	for i := 0; i < n; i++ {
		if T[i] == val {
			return true
		}
	}
	return false
}

/* I.S. data himpunan telah siap pada piranti masukan
   F.S. array T berisi sejumlah n bilangan bulat yang berasal dari masukan
   (masukan berakhir apabila bilangan ada yang duplikat, atau array penuh) */
func inputSet(T *set, n *int) {
	var val int
	*n = 0
	for *n < 2022 {
		fmt.Scan(&val)
		// Jika nilai sudah ada di dalam array (duplikat), berhenti mengisi
		if exist(*T, *n, val) {
			break
		}
		// Jika belum ada, masukkan ke array dan tambahkan counter n
		T[*n] = val
		*n++
	}
}

/* I.S. terdefinisi himpunan T1 dan T2 yang berisi sejumlah n dan m anggota
   F.S. himpunan T3 berisi sejumlah h bilangan bulat yang merupakan irisan
   dari himpunan T1 dan T2 */
func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	*h = 0
	for i := 0; i < n; i++ {
		// Jika elemen di T1 juga ada di T2, maka itu adalah irisan
		if exist(T2, m, T1[i]) {
			T3[*h] = T1[i]
			*h++
		}
	}
}

/* I.S. terdefinisi sebuah himpunan T yang berisi sejumlah n bilangan bulat
   F.S. menampilkan isi array T secara horizontal (dipisahkan oleh spasi) */
func printSet(T set, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", T[i])
	}
	fmt.Println()
}

func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int

	// Input himpunan pertama
	inputSet(&s1, &n1)
	// Input himpunan kedua
	inputSet(&s2, &n2)

	// Mencari irisan antara s1 dan s2
	findIntersection(s1, s2, n1, n2, &s3, &n3)

	// Menampilkan hasil irisan
	printSet(s3, n3)
}