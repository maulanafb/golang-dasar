package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	person := Person{Name: "Maulana"}

	SayHello(person)

	kuda := Animal{Name: "Kuda"}
	SayHello(kuda)
}
