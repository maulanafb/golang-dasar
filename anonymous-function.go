package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Hello, " + name + " you are blocked ")
	} else {
		fmt.Println("Wellcome, " + name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("anjing", blacklist)
	registerUser("abil", func(name string) bool { return name == "maulana" })
}
