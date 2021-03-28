package main

import (
	"fmt"
)

func main() {
	// Strongly using var
	var name string = "Vudang"
	var age int64 = 21
	var isHandsome = true
	const vudang = "Handsome"

	// Shorthand
	vudangName := "Dang Le Gia Vu"
	level, email := "chicken", "vudangdev@gmail.com"

	fmt.Println("Hello "+name, age, isHandsome)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isHandsome)
	fmt.Print(vudangName + "\n")
	fmt.Print(level, email)

}
