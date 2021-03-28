package main

import "fmt"

//
var sum int = 0

// Closures
func closure() func(int) int {
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	fmt.Println("Hello World")

	//sum := closure()
	for i := 0; i < 10; i++ {
		//fmt.Println(sum(i))
		fmt.Println(closure()(i))
	}
}
