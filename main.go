package main

import "fmt"

func main() {
	// Membuat slice
	buah := []string{"Apel", "Mangga", "Jeruk"}

	fmt.Println("Daftar buah:", buah)
	fmt.Println("Buah pertama:", buah[0])

	// Menambahkan data ke slice
	buah = append(buah, "Pisang")

	fmt.Println("Setelah ditambahkan:", buah)

	// Mengakses slice
	fmt.Println("Dua buah pertama:", buah[:2])

	// Menghapus elemen dengan membuat slice baru
	buah = append(buah[:1], buah[2:]...)

	fmt.Println("Setelah dihapus:", buah)

	// Slice angka
	angka := []int{10, 20, 30, 40, 50}

	fmt.Println("Daftar angka:", angka)
	fmt.Println("Panjang slice:", len(angka))
	fmt.Println("Kapasitas slice:", cap(angka))
}