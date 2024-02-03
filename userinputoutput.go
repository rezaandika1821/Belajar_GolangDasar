package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nama, alamat string
	fmt.Print("Masukkan Nama Depan Anda : ")
	fmt.Scanln(&nama)                     //->Menggunakan scan
	scanner := bufio.NewScanner(os.Stdin) //->menggunakan bufio dan kelebihannya adalah bisa menggunakan spasi
	fmt.Print("Masukkan alamat lengkap anda : ")
	scanner.Scan()
	alamat = scanner.Text()
	fmt.Printf("Alamat Anda : %v", alamat)

}
