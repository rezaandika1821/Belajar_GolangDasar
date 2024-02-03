package main

import "fmt"

func main() {
	var nama string = "santoso"                          //->deklarasi menggunakan keyword var
	alamat := "Jl. Sarirogo No.1 Jawa Timur" //->deklarasi Menggunakan type inference
	var (
		nim          int = 2210032
		ipk          float32 = 3.923
		indeks_nilai rune = 'A'
	) //deklarasi dengan model majemuk menurun atau vertikal dengan type data yang berbeda-beda

	var barang, benda string = "Selimut","Bantal" //deklarasi majemuk menggunakan model horizontal dengan type data yang sama
	//dan lain-lain
// pendeklarasian bisa langsung diberi value dan setiap value harus pas dan dicocokan dengan variabel tujuan
_ = "Ini variable yang di acuhkan atau tidak ingin digunaka" //deklarasi underscore jika variable ingin diabaikan

fmt.Println("nama : ",nama)
fmt.Println("Alamat : ", alamat)
fmt.Println("Nim : ",nim)
fmt.Printf("Nilai IPK : %2.f",ipk)
fmt.Println("Indeks Pencapaian : ",indeks_nilai)
fmt.Println("Barang : ",barang)
fmt.Println("Benda : ",benda)

}