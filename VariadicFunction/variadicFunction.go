package main

import (
	"fmt"
)

func calculateData(datas ...int) int {
	total := 0
	for _, data := range datas {
		total += data
	}
	return total
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	total := calculateData(1, 23, 4, 5, 5, 2, 5, 234)
	fmt.Println(total)
	total = calculateData(slice...)
	println(total)

}
