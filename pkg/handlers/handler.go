package handlers

import (
	"github.com/ulugbek0217/template/pkg/config"
	"github.com/ulugbek0217/template/pkg/models"
	"github.com/ulugbek0217/template/pkg/render"
	"net/http"
)

var Repo *Repository

// Repository is the repository struct
type Repository struct {
	App *config.TemplatesConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.TemplatesConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{
		"text": "This is from handler home",
	}
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
