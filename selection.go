package main

import "fmt"

func main() {
	var umur int

	
	//switch case
	var pilihan int
	
	fmt.Println("masukan pilihan anda : ")
	fmt.Println("1. cek umur \n2. exit")
	fmt.Scanln(&pilihan)

	switch pilihan{
	case 1:
		fmt.Print("masukan umur anda : ")
		fmt.Scanln(&umur)
	
		if umur <= 10 {
			fmt.Printf("umur anda %d, kategori : anak-anak", umur)
		} else if umur <=25 {
			fmt.Printf("umur anda %d, kategori : remaja", umur)
		} else if umur <=45 {
			fmt.Printf("umur anda %d, kategori : dewasa", umur)
		} else {
			fmt.Printf("umur anda %d, kategori : lansia", umur)
		}
	default:
	break
	}
}