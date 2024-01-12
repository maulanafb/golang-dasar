package main

import "fmt"

func main() {

	// Init statement , condition , post statement
	for counter := 1; counter < 11; counter++ {
		fmt.Println("perulangan ke ", counter)
	}

	fmt.Println("selesai")

	names := []string{"Maulana", "fatih", "bilqisthi"}
	// for manual
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for range
	// for index/key namaValue , namaArray/namaSlice
	for index, name := range names {
		fmt.Println("index", index, name)
	}

	// for range tidak butuh index
	for _, name := range names {
		fmt.Println(name)
	}
}
