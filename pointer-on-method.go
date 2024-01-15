package main

// untuk method direkomendasikan selalu menggunakan pointer ( * )

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	maulana := Man{"Maulana"}
	maulana.Married()

	fmt.Println(maulana.Name)
}
