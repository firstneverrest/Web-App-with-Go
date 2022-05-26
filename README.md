# Web Application with Go

## Start server

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

## Run Multiple Go Files

```bash
go run *.go

or

go run .

# run main.go
go run cmd/web/*.go
```

## Conventional Folder Structure

- `/cmd` - main applications for the project, keep main.go file
- `/pkg` - public code that can be imported and used bu external application
- `/internal` - private code that is not reusable and can only used in this project
- `/vendor` - application dependencies which can be managed manually or by dependency management tool
- `/api` - protocol definition files, OpenAPI/Swagger specs, JSON schema files
- `/web` -
- `/static` - static files like images, video
- `/configs` - config file
- `/docs` - documents
- `/test` - additional external test apps and test data
- `go.mod` - define module name, go version and module path
- `go.sum` - contains all dependency check sums to validate each direct and indirect dependency to confirm that none of them has been modified

## Setting wide configuration

Set global config to Go application with struct.

```go
package config

import "html/template"

// application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}

```

## Routing

### pat

[pat package](https://github.com/bmizerany/pat)

```bash
go get github.com/bmizerany/pat
```

```go
// routes.go
package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/firstneverrest/go-web-app/pkg/config"
	"github.com/firstneverrest/go-web-app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New() // multiplexer

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}

// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/firstneverrest/go-web-app/pkg/config"
	"github.com/firstneverrest/go-web-app/pkg/handlers"
	"github.com/firstneverrest/go-web-app/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	slicePortNumber := portNumber[1:]
	fmt.Printf("Server is listening on http://localhost:%s\n", slicePortNumber)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}

```

### chi

chi is a lightweight, idiomatic and composable router for building Go HTTP services. [Visit GitHub repository](https://github.com/go-chi/chi).

```bash
go get github.com/go-chi/chi/v5
```

### nosurf

nosurf is an middleware HTTP package that helps you prevent Cross-Site Request Forgery (CSRF) attacks. [Visit GitHub repository](https://github.com/justinas/nosurf).

```
go get github.com/justinas/nosurf
```

```go
// routes.go
package main

import (
	"net/http"

	"github.com/firstneverrest/go-web-app/pkg/config"
	"github.com/firstneverrest/go-web-app/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter() // multiplexer

	mux.Use(middleware.Recoverer) // recover from panics without crashing the server
	mux.Use(PrintToConsole)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}

// middleware.go
package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func PrintToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("PrintToConsole middleware")
		next.ServeHTTP(w, r) // go to the next middleware
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false, // run on https or not
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler

}

```

## Session Packages

### SCS

SCS is an HTTP session management which has automatic loading and saving session data via middleware.

```
go get github.com/alexedwards/scs
```

## Send a response

### Send text as a response

```go
func (m *Repository) BuyCoffee(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You bought a coffee!"))
}
```

### Send JSON as a response

```go
func (m *Repository) BuyCoffee(w http.ResponseWriter, r *http.Request) {
	response := coffee{
		Name:    "Cappuccino",
		Price:   2.50,
		Message: "Thank you for your purchase",
	}

	out, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

```

## Form validation

### govalidator

## Templating Engines

### CloudyKit/jet

## Tip

### When you unsure about data type

You can use interface type when you don't actually know the data type yet.

```go
type Data struct {
	SomeData map[string]interface{}
}
```

### Import Cycle

Import cycle happens when two files are imported by each other.

### Go command

- `go mod tidy` - remove unused dependencies in the project
- `go mod vendor` - download all dependencies included in go.mod file into vendor directory

## Reference

- [golang-standards-project-layout](https://github.com/golang-standards/project-layout)
-
