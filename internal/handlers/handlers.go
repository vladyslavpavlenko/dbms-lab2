package handlers

import (
	"github.com/vladyslavpavlenko/dbms-lab2/internal/config"
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
	//DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
		//DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// NewTestRepo creates a new test repository
func NewTestRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
		//DB:  dbrepo.NewTestingRepo(a),
	}
}

// NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("rendering home.gohtml...")
	render.Template(w, r, "home.gohtml", &models.TemplateData{})
}

func (m *Repository) Courses(w http.ResponseWriter, r *http.Request) {
	log.Println("rendering courses.gohtml...")
	render.Template(w, r, "courses.gohtml", &models.TemplateData{})
}

func (m *Repository) Enrollments(w http.ResponseWriter, r *http.Request) {
	log.Println("rendering enrollments.gohtml...")
	render.Template(w, r, "enrollments.gohtml", &models.TemplateData{})
}

func (m *Repository) Categories(w http.ResponseWriter, r *http.Request) {
	log.Println("rendering categories.gohtml...")
	render.Template(w, r, "categories.gohtml", &models.TemplateData{})
}

func (m *Repository) Users(w http.ResponseWriter, r *http.Request) {
	log.Println("rendering users.gohtml...")
	render.Template(w, r, "users.gohtml", &models.TemplateData{})
}

func (m *Repository) CoursesAndCategories(w http.ResponseWriter, r *http.Request) {
	log.Println("rendering courses-and-categories.gohtml...")
	render.Template(w, r, "courses-and-categories.gohtml", &models.TemplateData{})
}
