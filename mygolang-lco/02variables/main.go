package main

import "fmt"

// jwtToken := "token"   -> not allowed as walrus operator is alowed only inside a function

var LoginToken string = "loginTokenString"; // public, name starts in capital

func main() {

	// types
	var username string = "Nabajyoti Modak"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T\n ", username)

	var smallVal uint8 = 255
	fmt.Printf("variable is of type : %T\n ", smallVal)

	// implicit type
	var wesbite = "google.com"
	fmt.Printf("variable is of type : %T\n ", wesbite)
	// wesbite=3 -> cannot allow as implicitly has been declared as string type

	// default value
	var name int
	fmt.Println(name)

	//walrus operator
	jwtToken := "token"
	fmt.Printf("variable is of type : %T\n ", jwtToken)
	fmt.Println(jwtToken)
}
