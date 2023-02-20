package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello</h1>"))
}

func main() {
	fmt.Println("hello mod in golang")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000",r))
}

