// recover untuk menangkap data panic
package main

import "fmt"

// ini kode recover yang salah
// func endApp() {
// 	fmt.Println("end app")
// }

// func runApps(error bool) {
// 	defer endApp()
// 	if error {
// 		panic("ups error")
// 	}
// 	message := recover()
// 	fmt.Println("terjadi panic", message)
// }

// ini kode recover yang benar
func endApps() {
	fmt.Println("end app")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApps(error bool) {
	defer endApps()
	if error {
		panic("ups error")
	}

}

func main() {
	runApps(true)
	fmt.Println("Maulana fb")
}
