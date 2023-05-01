package main

import (
	"fmt"
)

func getAwal(nama string) string {
	return "Halo " + nama + "\ntest"
}

func main() {
	awal := getAwal
	fmt.Println(awal("Dino"))
}
