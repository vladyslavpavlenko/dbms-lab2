package handlers

import (
	"errors"
	"fmt"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/helpers"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/models"
	"gorm.io/gorm"
	"net/http"
)

func (m *Repository) CreateCourse(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	title := r.Form.Get("title")
	description := r.Form.Get("description")
	instructorID := r.Form.Get("instructorID")

	var iID uint
	_, err = fmt.Sscanf(instructorID, "%d", &iID)
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	newCourse := models.Course{
		Title:        title,
		Description:  description,
		InstructorID: iID,
	}

	err = m.App.DB.Create(&newCourse).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/courses", http.StatusSeeOther)
}

func (m *Repository) CreateCategory(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	title := r.Form.Get("title")

	var existingCategory models.Category
	result := m.App.DB.Where("title = ?", title).First(&existingCategory)

	if result.Error == nil {
		http.Error(w, "A category with the given title already exists.", http.StatusConflict)
		return
	}

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		helpers.ServerError(w, result.Error)
		return
	}

	newCategory := models.Category{
		Title: title,
	}

	err = m.App.DB.Create(&newCategory).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

func (m *Repository) CreateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	firstName := r.Form.Get("firstName")
	lastName := r.Form.Get("lastName")
	email := r.Form.Get("email")

	var existingUser models.User
	result := m.App.DB.Where("email = ?", email).First(&existingUser)

	if result.Error == nil {
		http.Error(w, "A user with the given email already exists.", http.StatusConflict)
		return
	}

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		helpers.ServerError(w, result.Error)
		return
	}

	newUser := models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	err = m.App.DB.Create(&newUser).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (m *Repository) CreateEnrollment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	courseID := r.Form.Get("courseID")
	userID := r.Form.Get("userID")

	var cID uint
	var uID uint

	_, err = fmt.Sscanf(courseID, "%d", &cID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	_, err = fmt.Sscanf(userID, "%d", &uID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	newEnrollment := models.Enrollment{
		CourseID: cID,
		UserID:   uID,
	}

	err = m.App.DB.Create(&newEnrollment).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/enrollments", http.StatusSeeOther)
}

func (m *Repository) CreateCategoryJunction(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	courseID := r.Form.Get("courseID")
	categoryID := r.Form.Get("categoryID")

	var cID, catID uint
	_, err = fmt.Sscanf(courseID, "%d", &cID)
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	_, err = fmt.Sscanf(categoryID, "%d", &catID)
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	var course models.Course
	var category models.Category

	err = m.App.DB.First(&course, cID).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.App.DB.First(&category, catID).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var existingAssociations []models.Category
	err = m.App.DB.Model(&course).Association("Categories").Find(&existingAssociations, "id = ?", catID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(existingAssociations) > 0 {
		http.Error(w, "This category is already associated with the course.", http.StatusConflict)
		return
	}

	err = m.App.DB.Model(&course).Association("Categories").Append(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/categories_junction", http.StatusSeeOther)
}
