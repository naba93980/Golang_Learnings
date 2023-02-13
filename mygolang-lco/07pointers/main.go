package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println(ptr)

	myNumber := 23
	var ptr2 = &myNumber
	fmt.Println(ptr2)
	fmt.Println(*ptr2)

	*ptr2 = *ptr2 * 10
	fmt.Println("New value is: ", myNumber)
}
