package main

import "fmt"

func random() any {
	return true
}

func main() {
	var result any = random()

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown", value)
	}
}
