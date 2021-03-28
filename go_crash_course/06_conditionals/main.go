package main

import "fmt"

const MALE = "male"
const FEMALE = "female"

func pickGender(gender string) string {
	switch gender {
	case MALE:
		return "Gender is " + MALE
	case FEMALE:
		return "Gender is " + FEMALE
	default:
		return "You are a dog ^^"
	}
}

func main() {
	a := 110
	b := 20

	if a > b {
		fmt.Printf("%d is greater than %d\n", a, b)
	} else {
		fmt.Printf("%d is lesser than %d\n", a, b)
	}

	flag := false
	if flag {
		fmt.Println("Flag On")
	} else {
		fmt.Println("Flag Off")
	}

	// Switch

	fmt.Println(pickGender("haha"))

}
