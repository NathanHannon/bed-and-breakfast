package handlers

import (
	"net/http"

	"github.com/nathanhannon/bed-and-breakfast/pkg/config"
	"github.com/nathanhannon/bed-and-breakfast/pkg/models"
	"github.com/nathanhannon/bed-and-breakfast/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again."
	// sends data to the template
	render.Template(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the reservation page and the accompanying form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Posted to search availabaility"))
}

// Generals renders the General's Quarters room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the Major's Suite room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "majors.page.tmpl", &models.TemplateData{})
}

// Contact renders the Contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "contact.page.tmpl", &models.TemplateData{})
}
