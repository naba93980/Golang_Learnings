package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	naba := User{"Nabajyoti", "naba93980@gmail.com", true, 23}
	fmt.Println(naba)
	fmt.Printf("naba details: %v\n",naba)
	fmt.Printf("naba details: %+v\n",naba)
	fmt.Printf("name is %v\n",naba.Name)
	
}
