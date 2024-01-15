package main

import "fmt"

type Alamat struct {
	City, Province, Country string
}

func main() {
	alamat1 := Alamat{City: "Pontianak", Province: "Kalbar", Country: "Indo"}
	alamat2 := new(Alamat)

	alamat2.City = "Singkawang"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
