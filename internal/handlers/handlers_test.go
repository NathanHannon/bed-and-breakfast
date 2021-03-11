package handlers

import (
	"encoding/gob"
	"log"
	"net/http"	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/nathanhannon/bed-and-breakfast/internal/models"
	"github.com/nathanhannon/bed-and-breakfast/internal/render"
)

var app config.AppConfig
var session *scs.SessionManager

func getRoutes() http.Handler {
	gob.Register(models.Reservation{})

	// Change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
}