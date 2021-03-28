package main

import (
	"fmt"
	"strconv"
)

// Define person
type Person struct {
	firstName string
	lastName  string
	age       int
	rich      bool
	gender    bool
}

// Closure apply
func (person Person) Welcome() string {
	return "Hello Age -> " + person.lastName + strconv.Itoa(person.age)
}

// Pointer Receiver
func (person *Person) hasBirthday() {
	person.age++
}

// Pointer Receiver
func (person *Person) getMarried(husbandLastName string) {
	if person.gender {
		return
	} else {
		person.lastName = husbandLastName
	}
}

func main() {
	fmt.Println("Hello World")

	vudang := Person{
		firstName: "Gia Vu",
		lastName:  "Dang Le",
		age:       21,
		rich:      false,
		gender:    true,
	}

	// Alternative

	thanhvan := Person{
		"Bui Huynh",
		"Thanh Van",
		21,
		false,
		false,
	}

	fmt.Println(vudang)
	fmt.Println(thanhvan)
	fmt.Println(thanhvan.Welcome())

	// Pointer Test
	thanhvan.hasBirthday()
	thanhvan.hasBirthday()
	thanhvan.hasBirthday()
	thanhvan.hasBirthday()
	thanhvan.hasBirthday()

	fmt.Println(thanhvan.Welcome())

	thanhvan.getMarried("VuDang")
	fmt.Println(thanhvan.Welcome())

}
