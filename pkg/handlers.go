package pkg

import (
	"go-web-app7/pkg"
	"net/http"
)

// the repository used by the handlers
var Repo *Repository

// is the repository type
type Repository struct {
	App *pkg.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *pkg.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the hundlers
func NewHandler(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.tmpl")
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.page.tmpl")
}
