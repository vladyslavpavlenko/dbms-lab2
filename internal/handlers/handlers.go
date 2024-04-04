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

func (m *Repository) Courses(w http.ResponseWriter, r *http.Request) {
	log.Println("rendering courses.page.gohtml...")

	var courses []models.Course

	err := m.App.DB.Preload("Instructor").Find(&courses).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	var users []models.User

	err = m.App.DB.Find(&users).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	data := make(map[string]any)
	data["users"] = users
	data["courses"] = courses
	data["currentPage"] = "courses"

	err = render.Template(w, r, "courses.page.gohtml", &models.TemplateData{
		Data: data,
		Form: forms.New(nil),
	})
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
}

func (m *Repository) Enrollments(w http.ResponseWriter, r *http.Request) {
	log.Println("rendering enrollments.page.gohtml...")

	var enrollments []models.Enrollment
	err := m.App.DB.Preload("Course").Preload("User").Find(&enrollments).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	var users []models.User
	err = m.App.DB.Find(&users).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	var courses []models.Course
	err = m.App.DB.Find(&courses).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	data := make(map[string]any)
	data["enrollments"] = enrollments
	data["users"] = users
	data["courses"] = courses
	data["currentPage"] = "enrollments"

	err = render.Template(w, r, "enrollments.page.gohtml", &models.TemplateData{
		Data: data,
		Form: forms.New(nil),
	})
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
}

func (m *Repository) Categories(w http.ResponseWriter, r *http.Request) {
	log.Println("rendering categories.page.gohtml...")

	var categories []models.Category

	err := m.App.DB.Find(&categories).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	data := make(map[string]any)
	data["categories"] = categories
	data["currentPage"] = "categories"

	err = render.Template(w, r, "categories.page.gohtml", &models.TemplateData{
		Data: data,
		Form: forms.New(nil),
	})
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
}

func (m *Repository) Users(w http.ResponseWriter, r *http.Request) {
	log.Println("rendering users.page.gohtml...")

	var users []models.User

	err := m.App.DB.Find(&users).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	data := make(map[string]any)
	data["users"] = users
	data["currentPage"] = "users"

	err = render.Template(w, r, "users.page.gohtml", &models.TemplateData{
		Data: data,
		Form: forms.New(nil),
	})
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
}

func (m *Repository) CategoriesJunction(w http.ResponseWriter, r *http.Request) {
	log.Println("rendering categories_junction.page.gohtml...")

	var courses []models.Course
	err := m.App.DB.Preload("Categories").Find(&courses).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	var categories []models.Category
	err = m.App.DB.Find(&categories).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	data := make(map[string]any)
	data["courses"] = courses
	data["categories"] = categories
	data["currentPage"] = "categories_junction"

	err = render.Template(w, r, "categories_junction.page.gohtml", &models.TemplateData{
		Data: data,
		Form: forms.New(nil),
	})
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
}
