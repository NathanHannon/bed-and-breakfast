package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/nathanhannon/bed-and-breakfast/internal/config"
	"github.com/nathanhannon/bed-and-breakfast/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

// A dummy writer to satisfy response headers for testing
type dummyWriter struct{}

func (tw *dummyWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *dummyWriter) WriteHeader(i int) {

}

func (tw *dummyWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}

func TestMain(m *testing.M) {
	// What is going to be put into the session
	gob.Register(models.Reservation{})

	// Change this to true when in production
	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = testApp.InProduction

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}
