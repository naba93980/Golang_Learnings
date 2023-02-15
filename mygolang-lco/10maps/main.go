package main

import (
	"fmt"
)

func main() {
	language := make(map[string]string)
	fmt.Println((language))
	language["js"] = "javascript"
	language["py"] = "python"
	fmt.Println(language["js"])
	// delete(language, "js")
	fmt.Println((language))

	// loops - comma ok syntax
	for key, value := range language {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
