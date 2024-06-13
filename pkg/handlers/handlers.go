package handlers

import (
	"github.com/gllkmike1445/bookings/pkg/config"
	"github.com/gllkmike1445/bookings/pkg/models"
	"github.com/gllkmike1445/bookings/pkg/render"
	"net/http"
)

// Repo the repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{a}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the about home handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic

	stringMap := make(map[string]string)
	stringMap["test"] = "testing stringMap"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	// the thing to remember here is that this value will be an empty string if there is nothing in the session named remote IP.
	stringMap["remote_ip"] = remoteIp

	//send the data to the template

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
