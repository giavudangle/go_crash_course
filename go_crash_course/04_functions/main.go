package main

import "fmt"

func hello(name string) string {
	return "Hello " + name
}

func getSum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(hello("Vudang"))
	fmt.Printf("Result of a + b is %d\n", getSum(10, 20))
}
