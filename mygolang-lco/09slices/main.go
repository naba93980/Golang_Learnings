package main

import (
	"fmt"
	"sort"
)

func main() {

	// declare and zero value in itialization

	var fruitList = []string{} // []type{} , curly bracket bcoz this systax requires iitialization
	fmt.Println(fruitList)
	fmt.Printf("Fruit list is: %T \n", fruitList)


	// append
	fruitList = append(fruitList, "apple", "orange", "Tomato")
	fmt.Println(fruitList)

	fruitList = append(fruitList, "Mango", "banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	// make
	highScores := make([]int, 3)
	fmt.Println(highScores)

	highScores[0] = 111
	highScores[1] = 222
	highScores[2] = 3
	fmt.Println(highScores)

	highScores = append(highScores, 444, 555) // new underlying array has been created
	fmt.Println(highScores)

	// sorting slice
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	// deleting slice value
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
