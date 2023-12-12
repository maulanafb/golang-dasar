// tipe data yang tidak bisa diubah saat sudah diinisiasikan

package main

import "fmt"

func main() {
	const (
		firstName = "Maulana"
		lastName  = "Fb"
	)
	fmt.Println(firstName, lastName)
}
