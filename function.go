package main

import "fmt"

func main() {
	sayHello := sayHello("Maulana", "fatih")
	fmt.Println(sayHello)
	getHello("Maulana")
	getHello := getHello("Maulana")
	fmt.Println(getHello)
}

func sayHello(firstName string, lastName string) string {
	var fullName string = firstName + " " + lastName
	return fullName
}

func getHello(name string) string {
	return "Hello " + name
}
