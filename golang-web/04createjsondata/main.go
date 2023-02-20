package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int    `json:"website"`
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	EncodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"react", 19, "udemy", "123", []string{"detailed", "opinionated"}},
		{"next", 29, "udemy", "123", []string{"not-detailed", "unopinionated"}},
		{"graphql", 19, "udemy", "123", nil}}

	// package this data as json data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
