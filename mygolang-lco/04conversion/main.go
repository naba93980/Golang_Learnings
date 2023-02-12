package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our app")
	fmt.Println("please rate our pizza between 1 and 5")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating: ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("your rating ",numRating)	
	}
}
