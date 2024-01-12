package main

import "fmt"

func main() {
	name := "Raden"

	if name == "maulana" {
		fmt.Println("Hello " + name)
	} else if name == "Hendri" {
		fmt.Println("eh ada " + name)
	} else {
		fmt.Println("Woi Penyusup")
	}

	// short statement if

	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
