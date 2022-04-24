# Web Application with Go

## Hello World

```go
package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	slicePortNumber := portNumber[1:]
	fmt.Printf("Server is listening on port %s\n", slicePortNumber)
	_ = http.ListenAndServe(portNumber, nil)
}


```

## Create Go mod

```bash
go mod init github.com/firstneverrest/go-web-app
```

## Difference between print, sprint and fprint

- print - print character stream of data on stdout console
- sprint - save character stream of data in char buffer (not print out)
- fprint - print character stream of data on a file (not on stdout console)
- [Reference](https://www.geeksforgeeks.org/difference-printf-sprintf-fprintf/)
