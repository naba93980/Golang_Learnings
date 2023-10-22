package main

import (
	"fmt"
	"net/http"
)

func getStatusCode(endpoint string) {

	defer wg.Done();
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	}
	fmt.Printf("%d status code for website %s \n", res.StatusCode, endpoint)
}
