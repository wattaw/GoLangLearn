package main

import "fmt"

func forLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func main() {
	loop := forLoop(2)
	fmt.Println(loop)
	fmt.Println(forLoop(5))
}
