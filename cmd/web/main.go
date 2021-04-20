package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/nathanhannon/bed-and-breakfast/internal/config"
	"github.com/nathanhannon/bed-and-breakfast/internal/handlers"
	"github.com/nathanhannon/bed-and-breakfast/internal/helpers"
	"github.com/nathanhannon/bed-and-breakfast/internal/models"
	"github.com/nathanhannon/bed-and-breakfast/internal/render"
)

const portNumber = ":8081"

var app config.AppConfig
var session *scs.SessionManager
var infoLog, errorLog *log.Logger

// main is the main function
func main() {

	err := run()

	fmt.Printf("Starting port number on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	// What is going to be put into the session
	gob.Register(models.Reservation{})

	// Change this to true when in production
	app.InProduction = false

	// Logging and error handling
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// Session info
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)

	return nil
}
