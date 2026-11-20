package main

import "fmt"

func main() {
	// Percabangan if
	nilai := 85

	if nilai >= 75 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}

	// Percabangan if else if
	angka := 2

	if angka == 1 {
		fmt.Println("Satu")
	} else if angka == 2 {
		fmt.Println("Dua")
	} else {
		fmt.Println("Angka lainnya")
	}

	// Switch
	hari := 3

	switch hari {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	default:
		fmt.Println("Hari lainnya")
	}

	// Perulangan for
	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}
}