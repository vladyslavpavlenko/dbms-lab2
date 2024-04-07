package handlers

import (
	"fmt"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/helpers"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/models"
	"net/http"
)

func (m *Repository) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	courseID := r.Form.Get("id")
	title := r.Form.Get("title")
	description := r.Form.Get("description")
	instructorID := r.Form.Get("instructorID")

	var cID uint
	var iID uint
	_, err = fmt.Sscanf(courseID, "%d", &cID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	_, err = fmt.Sscanf(instructorID, "%d", &iID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	err = m.App.DB.Model(&models.Course{}).
		Where("id = ?", cID).Updates(models.Course{
		Title:        title,
		Description:  description,
		InstructorID: iID},
	).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/courses", http.StatusSeeOther)
}

func (m *Repository) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	categoryID := r.Form.Get("id")
	title := r.Form.Get("title")

	var cID uint
	_, err = fmt.Sscanf(categoryID, "%d", &cID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	err = m.App.DB.Model(&models.Category{}).
		Where("id = ?", cID).Updates(models.Category{
		Title: title,
	},
	).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

func (m *Repository) UpdateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	userID := r.Form.Get("id")
	firstName := r.Form.Get("firstName")
	lastName := r.Form.Get("lastName")
	email := r.Form.Get("email")

	var uID uint
	_, err = fmt.Sscanf(userID, "%d", &uID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	err = m.App.DB.Model(&models.User{}).
		Where("id = ?", uID).Updates(models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	},
	).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (m *Repository) UpdateEnrollment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	oldUserID := r.Form.Get("oldUserID")
	oldCourseID := r.Form.Get("oldCourseID")
	userID := r.Form.Get("userID")
	courseID := r.Form.Get("courseID")

	var ouID uint
	var ocID uint

	_, err = fmt.Sscanf(oldUserID, "%d", &ouID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	_, err = fmt.Sscanf(oldCourseID, "%d", &ocID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	var uID uint
	var cID uint

	_, err = fmt.Sscanf(userID, "%d", &uID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	_, err = fmt.Sscanf(courseID, "%d", &cID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	err = m.App.DB.Model(&models.Enrollment{}).
		Where("user_id = ? AND course_id = ?", ouID, ocID).Updates(models.Enrollment{
		UserID:   uID,
		CourseID: cID,
	},
	).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/enrollments", http.StatusSeeOther)
}

func (m *Repository) UpdateCategoryJunction(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	oldCourseID := r.Form.Get("oldCourseID")
	oldCategoryID := r.Form.Get("oldCategoryID")
	newCourseID := r.Form.Get("courseID")
	newCategoryID := r.Form.Get("categoryID")

	var ocID, ocaID, ncID, ncaID uint

	_, err = fmt.Sscanf(oldCourseID, "%d", &ocID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	_, err = fmt.Sscanf(oldCategoryID, "%d", &ocaID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	_, err = fmt.Sscanf(newCourseID, "%d", &ncID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	_, err = fmt.Sscanf(newCategoryID, "%d", &ncaID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	var existingAssociation []models.Course
	err = m.App.DB.Joins("JOIN categories_junction ON categories_junction.course_id = courses.id").
		Where("categories_junction.course_id = ? AND categories_junction.category_id = ?", ncID, ncaID).
		Find(&existingAssociation).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	if len(existingAssociation) > 0 {
		http.Error(w, "This course and category association already exists", http.StatusBadRequest)
		return
	}

	err = m.App.DB.Exec("DELETE FROM categories_junction WHERE course_id = ? AND category_id = ?", ocID, ocaID).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	err = m.App.DB.Exec("INSERT INTO categories_junction (course_id, category_id) VALUES (?, ?)", ncID, ncaID).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/categories_junction", http.StatusSeeOther)
}
