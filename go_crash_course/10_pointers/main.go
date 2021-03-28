package main

import "fmt"

func main() {
	x := 10
	y := &x

	fmt.Println(*y, y)
	fmt.Printf("%T\n", y)
	fmt.Println(*&x)

	// Change value with pointer
	fmt.Println(x)
	*y = 99
	fmt.Println(x)
}
