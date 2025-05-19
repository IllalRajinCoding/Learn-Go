package main

import "fmt"

// angka

var usia int = 20

// float
var tinggi float64 = 1.75

// string
var name string = "John Doe"

// boolean
var isMarried bool = false


func main() {
	fmt.Println("Usia:", usia)
	fmt.Println("Tinggi:", tinggi)
	isMarried = true

	if isMarried {
		fmt.Println("Status: Menikah")
	} else {
		fmt.Println("Status: Belum Menikah")
	}


	fmt.Println("marin:", isMarried)

	// hanya bisa diakses atau jalan di dalam func
	pesan := "Selamat datang di Go!"
	fmt.Println(pesan)
}