package main

import "fmt"

func main() {
	// Declare Array
	var friendArr [3]string

	// Assign value
	friendArr[0] = "Phan Anh"
	friendArr[1] = "Khang Hy"
	friendArr[2] = "Minh Truong"

	// Declare and assign shorthand
	foodArr := [3]string{"Burger", "Pho", "Pizza"}
	bookArr := []string{"The five of rings", "World", "Life of Demon"}

	fmt.Println(foodArr)
	fmt.Println(len(bookArr))
	fmt.Println(friendArr[0])

	// Slice Array
	fmt.Println(bookArr[0:1])
}
