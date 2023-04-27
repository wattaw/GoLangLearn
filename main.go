package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func main() {
// 	var hasil bool
// 	c := 83
// 	d := 23
// 	cd := d + c
// 	if cd == 108 {
// 		hasil = true
// 	} else {
// 		hasil = false
// 	}
// 	println(hasil)
// }

func main() {
	var (
		nama               string
		alamat             string
		tempatTanggalLahir string
	)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Nama: ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan Alamat: ")
	alamat, _ = reader.ReadString('\n')
	alamat = strings.TrimSpace(alamat)

	fmt.Print("Masukkan TTL : ")
	tempatTanggalLahir, _ = reader.ReadString('\n')
	tempatTanggalLahir = strings.TrimSpace(tempatTanggalLahir)

	looping := []string{nama, alamat, tempatTanggalLahir}

	for _, value := range looping {
		fmt.Println(value)
	}

	// fmt.Println("Masukkan Nama: ")
	// fmt.Scanln(&nama)
	// fmt.Println("Masukkan Alamat: ")
	// fmt.Scanln(&alamat)
	// fmt.Println("Masukkan TTL : ")
	// fmt.Scanln(&tempatTanggalLahir)
	// looping := []string{nama, alamat, tempatTanggalLahir}

	// for _, value := range looping {
	// 	fmt.Println(value)
	// }

}
