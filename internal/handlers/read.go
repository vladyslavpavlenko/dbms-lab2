package handlers

import (
	"github.com/vladyslavpavlenko/dbms-lab2/internal/forms"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/helpers"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/models"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/render"
	"log"
	"net/http"
)

func (m *Repository) Courses(w http.ResponseWriter, r *http.Request) {
	log.Println("rendering courses.page.gohtml...")

	var courses []models.Course

	err := m.App.DB.Order("id").Preload("Instructor").Find(&courses).Error
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
	err := m.App.DB.Order("user_id").Preload("Course").Preload("User").Find(&enrollments).Error
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

	err := m.App.DB.Order("id").Find(&categories).Error
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

	err := m.App.DB.Order("id").Find(&users).Error
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
	err := m.App.DB.Order("id").Preload("Categories").Find(&courses).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	var categories []models.Category
	err = m.App.DB.Order("id").Find(&categories).Error
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
