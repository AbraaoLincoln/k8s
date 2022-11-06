package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Runnig simple http client")
	_, err := http.Get("http://simple-http-server-service:8080/hello")

	if err != nil {
		fmt.Println(err)
	}
}
