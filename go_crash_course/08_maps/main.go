package main

import "fmt"

func main() {
	emails := make(map[string]string)
	emails["vudang"] = "vudangdev@gmail.com"
	emails["vudangcompany"] = "codingwithvudang@gmail.com"

	fmt.Println(len(emails))
	fmt.Println(emails["vudang"])
	fmt.Println(len(emails["vudang"]))
	delete(emails, "vudang")
	fmt.Println(emails)

	cars := map[string]int64{"Mercedes": 12000, "Mazada": 5000, "RollRoyces": 99999}

	fmt.Println(cars)

}
