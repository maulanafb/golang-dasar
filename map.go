package main

import (
	"fmt"
)

func main() {
	// person := map[string]string{}
	// person["name"] = "Maulana"
	// person["address"] = "Pontianak"

	// namaMap := map[typeKey]typevalue{}
	person := map[string]string{
		"name":    "Maulana",
		"address": "Pontianak",
	}
	fmt.Println(person["name"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Book Golang"
	book["author"] = "Mau"
	book["hapus"] = "Mau"

	delete(book, "hapus")

	fmt.Println(book)
}
