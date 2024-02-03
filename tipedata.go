package main

import "fmt"

func main() {
	var char rune = 'A'
	var strings string = "Saya Mahasiswa"
	var ints int = 2
	var floats float32 = 2.5
	var benar bool = true

	fmt.Printf("Ini character %c\n", char)
	fmt.Printf("Ini string %v\n", strings)
	fmt.Printf("Ini integer %d\n", ints)
	fmt.Printf("Ini float %2.f\n", floats)
	fmt.Println(benar)
}
