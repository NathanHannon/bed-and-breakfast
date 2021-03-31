package render

import (
	"net/http"
	"testing"

	"github.com/nathanhannon/bed-and-breakfast/internal/models"
)

func TestDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	session.Put(r.Context(), "flash", "123")

	result := DefaultData(&td, r)
	if result.Flash != "123" {
		t.Error("Flash value of '123' not found in session")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/test-url", nil)

	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}