package main

import "fmt"

func main() {
	// Membuat map
	nilai := map[string]int{
		"Aris":   90,
		"Desta":  85,
		"Karina": 95,
	}

	fmt.Println("Data nilai:", nilai)

	// Mengakses data map
	fmt.Println("Nilai Aris:", nilai["Aris"])

	// Menambahkan data
	nilai["Naufal"] = 88

	fmt.Println("Setelah ditambahkan:", nilai)

	// Mengubah data
	nilai["Aris"] = 100

	fmt.Println("Setelah diubah:", nilai)

	// Menghapus data
	delete(nilai, "Desta")

	fmt.Println("Setelah dihapus:", nilai)

	// Perulangan map
	for nama, skor := range nilai {
		fmt.Println(nama, ":", skor)
	}
}