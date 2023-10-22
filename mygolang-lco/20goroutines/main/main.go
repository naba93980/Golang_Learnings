package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func main() {

	fmt.Println("STARTED")
	websiteList := []string{
		"https://facebook.com",
		"http://www.studio404.net",
		"https://github.com/naba93980",
		"https://youtube.com",
	}

	for _, web := range websiteList {
		go getStatusCode((web))
		wg.Add(1)
	}

	go greeter()
	// wg.Add(1)  - deadlock recipe

	wg.Wait()
}

func greeter() {

	// wg.Wait()	- deadlock recipe
	fmt.Println("HELLO GREETING")
}
