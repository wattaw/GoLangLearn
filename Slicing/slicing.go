package main

import "fmt"

func main() {
	var bulan = [...]string{
		"January",
		"Febuary",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}
	var slicing1 = bulan[4:7]
	var slicing2 = slicing1[0:1]
	fmt.Println(len(slicing1))
	fmt.Println(cap(slicing1))
	fmt.Println(slicing2)
	fmt.Println(bulan)
	var slicing3 = append(slicing2, "Eko")
	fmt.Println(slicing3)
	slicing3[0] = "Test"
	fmt.Println(slicing3)

	newSlicing := make([]string, 2, 4)
	newSlicing[0] = "data1"
	newSlicing[1] = "data2"
	fmt.Println(newSlicing)

	copySlicing := make([]string, len(newSlicing), cap(newSlicing))
	copySlicing[0] = bulan[2]
	copySlicing[1] = newSlicing[0]
	fmt.Println(copySlicing)
}
