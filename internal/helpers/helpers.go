package helpers

import (
	"net/http"

	"github.com/nathanhannon/bed-and-breakfast/internal/config"
)

var app *config.AppConfig

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {

}

func ServerError(w http.ResponseWriter, err error) {

}
