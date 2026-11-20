package main

import "fmt"

func main() {
	// Operator Aritmatika
	a := 10
	b := 3

	fmt.Println("Penjumlahan:", a+b)
	fmt.Println("Pengurangan:", a-b)
	fmt.Println("Perkalian:", a*b)
	fmt.Println("Pembagian:", a/b)
	fmt.Println("Modulus:", a%b)

	// Operator Perbandingan
	fmt.Println("Apakah a > b:", a > b)
	fmt.Println("Apakah a == b:", a == b)

	// Operator Logika
	x := true
	y := false

	fmt.Println("AND:", x && y)
	fmt.Println("OR:", x || y)
	fmt.Println("NOT:", !x)
}