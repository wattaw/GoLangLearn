package main

import (
	"fmt"
)

func dataAwal() (string, string) {
	bungkus := [2]string{"Dino", "Rumah"}
	return bungkus[0], bungkus[1]
}

func main() {
	x, y := dataAwal()
	fmt.Println(x, y)
}
