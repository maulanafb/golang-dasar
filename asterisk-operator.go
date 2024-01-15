package main

import "fmt"

type Addres struct {
	City, Province, Country string
}

// pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
// pointer di buat dengan &nama_variabel
func main() {
	address1 := Addres{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Addres{"Jakarta", "DKI JAKARTA", "INDONESIA"}

	fmt.Println(address1)
	fmt.Println(address2)
}
