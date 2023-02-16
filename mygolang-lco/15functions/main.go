package main

import (
	"fmt"
)

func main() {
	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	// comma ok syntax
	proRes, myMessage := proAdder(2, 4, 3, 7, 6, 3, 5)
	fmt.Println(proRes)
	fmt.Println(myMessage)
}

func greet() {
	fmt.Println("welcome")
}

// returns 2 types
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "hi pro result function"
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}
