package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

// defer akan tetap berjalan walaupun program di hentikan dengan panic

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Oops Error")
	}
}

func main() {
	runApp(false)
}
