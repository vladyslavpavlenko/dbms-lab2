package handlers

import (
	"github.com/vladyslavpavlenko/dbms-lab2/internal/config"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/forms"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/helpers"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/models"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/render"
	"log"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// TableName struct to scan the result into
type TableName struct {
	Name string `gorm:"column:table_name"`
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("rendering home.page.gohtml...")

	var tables []TableName
	err := m.App.DB.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").
		Scan(&tables).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	data := make(map[string]any)
	data["tables"] = tables

	err = render.Template(w, r, "home.page.gohtml", &models.TemplateData{
		Data: data,
	})
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
}

func (m *Repository) Queries(w http.ResponseWriter, r *http.Request) {
	log.Println("rendering queries.page.gohtml...")

	var categories []models.Category

	err := m.App.DB.Order("id").Find(&categories).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	data := make(map[string]any)
	data["categories"] = categories
	data["currentPage"] = "queries"

	err = render.Template(w, r, "queries.page.gohtml", &models.TemplateData{
		Data: data,
		Form: forms.New(nil),
	})
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
}
