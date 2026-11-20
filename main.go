package main

import "fmt"

func main() {
	// Array
	var buah [3]string

	buah[0] = "Apel"
	buah[1] = "Mangga"
	buah[2] = "Jeruk"

	fmt.Println("Daftar buah:", buah)

	// Array dengan inisialisasi langsung
	angka := [5]int{10, 20, 30, 40, 50}

	fmt.Println("Daftar angka:", angka)
	fmt.Println("Angka pertama:", angka[0])
	fmt.Println("Angka terakhir:", angka[4])

	// Mengakses array menggunakan perulangan
	for i := 0; i < len(angka); i++ {
		fmt.Println("Angka ke-", i, ":", angka[i])
	}
}