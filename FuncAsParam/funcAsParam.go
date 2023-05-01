package main

import (
	"fmt"
)

func content(name string, filter func(string) string) {
	textFiltered := filter(name)
	fmt.Println("Hello ", textFiltered)
}

func contentFilter(name string) string {
	if name == "Anjing" || name == "Kodok" {
		return "..."
	} else {
		return name
	}
}

func main() {
	content("Data Baru", contentFilter)
	content("kodok", contentFilter)
}
