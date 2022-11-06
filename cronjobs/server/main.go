package main

import (
	"fmt"
	"net/http"
)

var count = 0

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", printHello)

	fmt.Println("Starting server")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}

	fmt.Println("Stoping server")
}

func printHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello", count)
	count++
}
