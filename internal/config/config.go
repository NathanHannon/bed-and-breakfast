package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"nathanhannon.dev/bed-and-breakfast/internal/models"
)

// AppConfig holds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	MailChannel   chan models.MailData
}
