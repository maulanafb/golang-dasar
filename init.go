package main

import (
	"fmt"
	"golang-dasar/database"
	_ "golang-dasar/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
