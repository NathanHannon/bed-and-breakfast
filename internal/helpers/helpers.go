package helpers

import "github.com/nathanhannon/bed-and-breakfast/internal/config"

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}
