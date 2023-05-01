package main

import (
	"fmt"
)

type Blacklist func(string) bool

func registUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Blocked")
	} else {
		fmt.Println("Welcome", name)
	}
}

// func BlacklistAdmin(name string) bool {

// }
// func BlackListUser(name string) bool {

// }

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registUser("admin", blacklist)
	registUser("ssss", blacklist)
}
