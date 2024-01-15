package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// cara 1
	// address := Address{}
	// changeCountryToIndonesia(&address)

	// cara 2
	var address *Address = &Address{}
	changeCountryToIndonesia(address)

	fmt.Println(address)
}
