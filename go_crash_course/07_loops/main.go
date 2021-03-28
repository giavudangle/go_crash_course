package main

import "fmt"

func main() {
	// Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//Short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	// Fizz Buzz

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("Fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
