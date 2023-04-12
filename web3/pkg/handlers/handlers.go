package handlers

import (
	"net/http"
	"web3/models"
	render "web3/pkg/render"

	"web3/pkg/config"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository{
		App: ac,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter,
	r *http.Request) {
	m.App.Session.Put(r.Context(),
		"userid", "gyurebalint")

	render.RenderTemplate(w, "home.page.html",
		&models.PageData{})
}

func (m *Repository) AboutHandler(w http.ResponseWriter,
	r *http.Request) {

	strMap := make(map[string]string)

	render.RenderTemplate(w, "about.page.html", &models.PageData{StrMap: strMap})
}

func (m *Repository) LoginHandler(w http.ResponseWriter,
	r *http.Request) {

	strMap := make(map[string]string)

	render.RenderTemplate(w, "login.page.html", &models.PageData{StrMap: strMap})
}

func (m *Repository) MakePostHandler(w http.ResponseWriter,
	r *http.Request) {

	strMap := make(map[string]string)

	render.RenderTemplate(w, "make-post.page.html", &models.PageData{StrMap: strMap})
}

func (m *Repository) PageHandler(w http.ResponseWriter,
	r *http.Request) {

	strMap := make(map[string]string)

	render.RenderTemplate(w, "page.page.html", &models.PageData{StrMap: strMap})
}
