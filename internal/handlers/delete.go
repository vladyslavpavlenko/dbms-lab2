package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/helpers"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/models"
	"net/http"
)

func (m *Repository) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var cID uint
	_, err := fmt.Sscanf(id, "%d", &cID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	err = m.App.DB.Table("categories_junction").Where("course_id = ?", cID).Delete(nil).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	err = m.App.DB.Where("id = ?", cID).Delete(&models.Course{}).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/courses", http.StatusSeeOther)
}

func (m *Repository) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var cID uint
	_, err := fmt.Sscanf(id, "%d", &cID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	err = m.App.DB.Table("categories_junction").Where("category_id = ?", cID).Delete(nil).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	err = m.App.DB.Where("id = ?", cID).Delete(&models.Category{}).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

func (m *Repository) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var cID uint
	_, err := fmt.Sscanf(id, "%d", &cID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	var courseIDs []uint
	err = m.App.DB.Model(&models.Course{}).Where("instructor_id = ?", id).Pluck("id", &courseIDs).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	if len(courseIDs) > 0 {
		err = m.App.DB.Table("categories_junction").Where("course_id IN (?)", courseIDs).Delete(nil).Error
		if err != nil {
			helpers.ServerError(w, err)
			return
		}
	}

	err = m.App.DB.Where("id = ?", cID).Delete(&models.User{}).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (m *Repository) DeleteEnrollment(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	courseID := chi.URLParam(r, "course_id")

	var uID uint
	var cID uint

	_, err := fmt.Sscanf(userID, "%d", &uID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	_, err = fmt.Sscanf(courseID, "%d", &cID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	err = m.App.DB.Where("user_id = ? AND course_id = ?", uID, cID).Delete(&models.Enrollment{}).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/enrollments", http.StatusSeeOther)
}

func (m *Repository) DeleteCategoryJunction(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "category_id")
	courseID := chi.URLParam(r, "course_id")

	var caID, cID uint

	_, err := fmt.Sscanf(categoryID, "%d", &caID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	_, err = fmt.Sscanf(courseID, "%d", &cID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	err = m.App.DB.Exec("DELETE FROM categories_junction WHERE course_id = ? AND category_id = ?", courseID, categoryID).Error
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/categories_junction", http.StatusSeeOther)
}
