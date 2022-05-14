package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/firstneverrest/go-web-app/internal/config"
	"github.com/firstneverrest/go-web-app/internal/models"
	"github.com/firstneverrest/go-web-app/internal/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// save user ip address in the session
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)

	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["name"] = "John Doe"

	remoteIP := m.App.Session.GetString(r.Context(), "remoteIP")
	stringMap["remoteIP"] = remoteIP

	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

type coffee struct {
	Name    string  `json:"name"`
	Price   float32 `json:"price"`
	Message string  `json:"message"`
}

func (m *Repository) BuyCoffee(w http.ResponseWriter, r *http.Request) {
	response := coffee{
		Name:    "Cappuccino",
		Price:   2.50,
		Message: "Thank you for your purchase",
	}

	out, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
