package main

import "fmt"

func getGoodBye(name string) string {
	return "GoodBye " + name
}

func main() {
	coba := getGoodBye
	fmt.Println(coba("anjing"))
}
