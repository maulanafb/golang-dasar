package main

import "fmt"

func main() {
	name := "maulanaw"

	switch name {
	case "maulana":
		fmt.Println("Hello Maulana")
	case "fatih":
		fmt.Println("Hello fatih")
	case "bilqisthi":
		fmt.Println("Hello bilqisthi")
	default:
		fmt.Println("Hai Boleh Kenalan?")
	}

	// short condition switch

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu Panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// kode switch tanpa kondisi

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Mayan Panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
