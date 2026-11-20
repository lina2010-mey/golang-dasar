package main

import "fmt"

// Function tanpa parameter
func sapa() {
	fmt.Println("Halo, selamat belajar Golang!")
}

// Function dengan parameter
func salam(nama string) {
	fmt.Println("Halo,", nama)
}

// Function dengan return value
func tambah(a, b int) int {
	return a + b
}

func main() {
	// Memanggil function
	sapa()

	// Memanggil function dengan parameter
	salam("Aris")

	// Memanggil function dengan return value
	hasil := tambah(10, 5)
	fmt.Println("Hasil penjumlahan:", hasil)
}