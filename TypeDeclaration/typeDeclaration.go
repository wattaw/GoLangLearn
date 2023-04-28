package main

import "fmt"

func main() {
	type (
		NoKTP    string
		Marriage bool
	)

	var (
		NoKtpEko       NoKTP    = "3515273643789393"
		StatusKawinEko Marriage = true
	)

	fmt.Println(NoKtpEko, StatusKawinEko)
}
