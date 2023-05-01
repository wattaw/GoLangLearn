package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func recursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialLoop(value-1)
	}
}

func main() {
	var data int
	fmt.Scanln(&data)
	recursive := factorialLoop(data)
	fmt.Println(recursive)
}
