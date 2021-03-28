package main

import "fmt"

func main() {

	ids := []int{1, 54, 3, 123, 51412, 123}

	// i-> index
	// id -> value
	for i, id := range ids {
		fmt.Printf("%d - ID : %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID : %d\n", id)
	}

	// Loop sum in range ids
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// Loop in map
	cars := map[string]int64{"Mercedes": 12000, "Mazada": 5000, "RollRoyces": 99999}

	for k, v := range cars {
		fmt.Printf("Model : %s -> Price : %d\n", k, v)
	}

}
