package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello,", name, " My Name is", customer.Name)
}

func main() {
	var maulana Customer
	maulana.Name = "maulana"
	maulana.Address = "Indonesia"
	maulana.Age = 21

	// joko := Customer{
	// 	Name:    "Joko",
	// 	Address: "Indonesia",
	// 	Age:     30,
	// }

	// customer := Customer{
	// 	"Customer", "Indonesia", 20,
	// }
	maulana.sayHello("Agus")
}
