package main

import "fmt"

struct adalah representasi data

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var maulana Customer
	maulana.Name = "maulana"
	maulana.Address = "Indonesia"
	maulana.Age = 21

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}

	customer := Customer{
		"Customer", "Indonesia", 20,
	}

	fmt.Println(maulana)
	fmt.Println(maulana.Name)

	fmt.Println(joko)
	fmt.Println(joko.Name)

	fmt.Println(customer)
	fmt.Println(customer.Name)
}
