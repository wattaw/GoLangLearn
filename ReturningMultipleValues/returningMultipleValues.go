package main

import (
	"fmt"
)

func dataAwal() (x string, y string) {
	x = "Dino"
	y = "Rumah"
	return
}

func main() {
	x, y := dataAwal()
	fmt.Println(x, y)
}
