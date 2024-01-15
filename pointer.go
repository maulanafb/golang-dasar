package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
// pointer di buat dengan &nama_variabel
func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// copy value
	// address2 := address1

	// pointer ke address pertama
	address2 := &address1
	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)
}
