package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// receives a copy of User object
func (x User) GetStatus() {
	fmt.Println("Is user active  ", x.Status)
}

func (x User) NewMail() {
	x.Email = "1234"
	fmt.Println(x.Email)
}

func main() {
	naba := User{"Nabajyoti", "naba93980@gmail.com", true, 23}
	naba.GetStatus()        // passes a copy of naba object
	naba.NewMail()          // prints 1234
	fmt.Println(naba.Email) // prints "naba93980@gmail.com"
}
