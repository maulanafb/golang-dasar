package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

// defer akan dijalankan di akhir function tanpa memanggil function
func runApplication() {
	defer logging()
	fmt.Println("run runApplication")
}

func main() {
	runApplication()
}
