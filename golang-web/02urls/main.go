package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=yhgdfg"

func main() {
	fmt.Println(myurl)

	// parsing from string to url
	result, _ := url.Parse(myurl)
	fmt.Println(result.Host)
	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.RawPath)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Println(qparams)


	// from URL instance to string
	partsOfUrl := &url.URL {
		Scheme: "https",
		Host: "lco.dev",
		Path: "/learn",
		RawQuery: "user=naba",
	}

	anotherUrl := partsOfUrl.String();
	fmt.Println(anotherUrl)
}
