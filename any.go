package main

import "fmt"

// interface kosong

func Ups() any {
	return 1
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
