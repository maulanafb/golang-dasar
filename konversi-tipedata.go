package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	// number overflow , maksimal int16 32767 , maka dia akan balik ke paling bawah jadi negatif
	var nilai16 int16 = int16(nilai32)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	//konversi 2

	var name = "maulana"
	var e uint8 = name[0]
	var eString = string(e)
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
