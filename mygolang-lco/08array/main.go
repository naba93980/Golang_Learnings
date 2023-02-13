package main

import "fmt"

func main() {

	var fruitList [4]string
	fmt.Println("Fruit list is: ", fruitList)
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", fruitList[2])
	fmt.Println("Fruit list is: ", len(fruitList[2]))
	fmt.Printf("Fruit list is: %T \n", fruitList[2])
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"potato", "apple", "mushroom"};
	fmt.Println(vegList)
	fmt.Printf("%T", vegList)
}


