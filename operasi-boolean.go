package main

import "fmt"

func main() {
	var nilaiAkhir = 70
	var nilaiAbsensi = 100
	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusNilaiAbsensi bool = nilaiAbsensi > 80

	var lulus bool = lulusNilaiAkhir || lulusNilaiAbsensi

	fmt.Println(lulus)
}
