package main

import "fmt"

func main() {
	type noKtp string

	var ktpMaul noKtp = "123123123"
	var contoh string = "5426"
	var contohKtp noKtp = noKtp(contoh)
	fmt.Println(ktpMaul)
	fmt.Println(contohKtp)
}
