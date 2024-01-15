package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Id required"}
	}

	if id != "maulana" {
		return &notFoundError{"Data Tidak ditemukan"}
	}
	return nil
}

func main() {
	err := SaveData("", nil)
	if err != nil {
		// terjadi error
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("validationError", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("notFoundError", notFoundErr.Error())
		// } else {
		// 	fmt.Println("Unknown error", err.Error())
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validationError", finalError.Error())
		case *notFoundError:
			fmt.Println("notFoundError", finalError.Error())
		default:
			fmt.Println("Unknown error", err.Error())
		}
	} else {
		// sukses
		fmt.Println("sukses")
	}
}
