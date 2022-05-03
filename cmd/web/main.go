package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/firstneverrest/go-web-app/cmd/pkg/config"
	"github.com/firstneverrest/go-web-app/cmd/pkg/handlers"
	"github.com/firstneverrest/go-web-app/cmd/pkg/render"
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

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	slicePortNumber := portNumber[1:]
	fmt.Printf("Server is listening on port %s\n", slicePortNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
