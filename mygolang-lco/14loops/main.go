package main

import "fmt"

func main() {
	days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}

	fmt.Println(days)

	// loop type 1
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// loop type 2
	for i := range days {
		fmt.Println(days[i])
	}

	// loop type 3
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	// loop type 4 - kinda while loop
	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto jump
		}

		if rougueValue == 5 {
			rougueValue++
			// continue
			break
		}

		fmt.Println(rougueValue)
		rougueValue++
	}

	jump:
	fmt.Println("Jumped")
}
