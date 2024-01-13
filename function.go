package main

import "fmt"

func main() {
	sayHello := sayHello("Maulana", "fatih")
	fmt.Println(sayHello)
	getHello("Maulana")
	getHello := getHello("Maulana")
	fmt.Println(getHello)
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// multiple value harus diambil semua valuenya
	// jika ingin mengambil 1 value saja , bisa mengganti yang tidak di inginkan dengan(_) garis bawah
	shortName, _ := getFullName()
	fmt.Println(shortName)

	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}

func sayHello(firstName string, lastName string) string {
	var fullName string = firstName + " " + lastName
	return fullName
}

func getHello(name string) string {
	return "Hello " + name
}

// function return multiple values
func getFullName() (string, string) {
	return "maulana", "fatih"
}

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "maulana"
	middleName = "fatih"
	lastName = "bilqisthi"
	return firstName, middleName, lastName
}
