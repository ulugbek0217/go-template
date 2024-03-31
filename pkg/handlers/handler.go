package handlers

import (
	"github.com/ulugbek0217/template/pkg/config"
	"github.com/ulugbek0217/template/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.TemplatesConfig
}

func NewRepo(a *config.TemplatesConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
