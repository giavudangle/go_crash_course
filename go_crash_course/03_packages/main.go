package main

import (
	"fmt"
	"math"

	"github.com/giavudangle/go_crash_course/03_packages/string_utils"
)

func main() {
	var a float64 = 31.14
	b := 34.3

	fmt.Println("Test packages")
	fmt.Println(math.Floor(a))
	fmt.Println(math.Ceil(b))
	fmt.Println(string_utils.Reverse("elgnaduvaig"))

}
