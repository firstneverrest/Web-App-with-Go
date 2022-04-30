package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	slicePortNumber := portNumber[1:]
	fmt.Printf("Server is listening on port %s\n", slicePortNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
