package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nathanhannon/bed-and-breakfast/pkg/config"
	"github.com/nathanhannon/bed-and-breakfast/pkg/handlers"
	"github.com/nathanhannon/bed-and-breakfast/pkg/render"
)

const portNumber = ":8081"

// main is the main function
func main() {
	var app config.AppConfig
	repo := handlers.NewRepo(&app)

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println((fmt.Sprintf("Starting port number on port %s", portNumber)))
	_ = http.ListenAndServe(portNumber, nil)
}
