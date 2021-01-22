package handlers

import (
	"net/http"

	"github.com/mariojackson/go-course/pkg/config"
	"github.com/mariojackson/go-course/pkg/render"
)

// Repo is the repository used by the handlers.
var Repo *Repository

// Repository is the repository type
type Repository struct {
	AppConfig *config.AppConfig
}

// NewRepo creates a new repository.
func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository{
		AppConfig: ac,
	}
}

// NewHandlers sets the repository for the handlers.
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
