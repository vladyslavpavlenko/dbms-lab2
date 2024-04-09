package handlers

import (
	"fmt"
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

	categoryID := r.URL.Query().Get("categoryID")
	instructorLastName := r.URL.Query().Get("instructorLastName")
	minCategoryCount := r.URL.Query().Get("minCategoryCount")
	user1ID := r.URL.Query().Get("user1ID")
	user2ID := r.URL.Query().Get("user2ID")

	query := m.App.DB.Order("id").Preload("Instructor")

	if categoryID != "" {
		var cID uint
		_, err := fmt.Sscanf(categoryID, "%d", &cID)
		if err != nil {
			helpers.ServerError(w, err)
			return
		}

		query = query.Joins("JOIN categories_junction ON categories_junction.course_id = courses.id").
			Where("categories_junction.category_id = ?", cID)
	}

	if instructorLastName != "" {
		query = query.Joins("JOIN users ON users.id = courses.instructor_id").
			Where("users.last_name = ?", instructorLastName)
	}

	if minCategoryCount != "" {
		var x int
		_, err := fmt.Sscanf(minCategoryCount, "%d", &x)
		if err != nil {
			helpers.ServerError(w, err)
			return
		}

		query = query.
			Joins("JOIN categories_junction ON categories_junction.course_id = courses.id").
			Group("courses.id").
			Having("COUNT(categories_junction.category_id) >= ?", x).
			Select("courses.*, COUNT(categories_junction.category_id) as category_count")
	}

	if user1ID != "" && user2ID != "" {
		var uID1, uID2 uint
		_, err := fmt.Sscanf(user1ID, "%d", &uID1)
		if err != nil {
			helpers.ServerError(w, err)
			return
		}
		_, err = fmt.Sscanf(user2ID, "%d", &uID2)
		if err != nil {
			helpers.ServerError(w, err)
			return
		}

		query = query.Joins("JOIN enrollments e1 ON e1.course_id = courses.id").
			Joins("JOIN enrollments e2 ON e2.course_id = courses.id").
			Where("e1.user_id = ? AND e2.user_id = ?", uID1, uID2).
			Group("courses.id")
	}

	err := query.Find(&courses).Error
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
	log.Println("Rendering categories.page.gohtml...")

	instructorID := r.URL.Query().Get("instructorID")
	instructorLastName := r.URL.Query().Get("instructorLastName")

	var categories []models.Category

	if instructorID != "" {
		var iID uint
		_, err := fmt.Sscanf(instructorID, "%d", &iID)
		if err != nil {
			helpers.ServerError(w, err)
			return
		}

		err = m.App.DB.Raw(`
            SELECT c.* 
            FROM categories c
            INNER JOIN categories_junction cj ON c.id = cj.category_id
            INNER JOIN courses co ON cj.course_id = co.id
            WHERE co.instructor_id = ?
            GROUP BY c.id
            HAVING COUNT(DISTINCT co.id) = (
                SELECT COUNT(*)
                FROM courses
                WHERE instructor_id = ?
            )
        `, iID, iID).Scan(&categories).Error

		if err != nil {
			helpers.ServerError(w, err)
			return
		}
	} else if instructorLastName != "" {
		err := m.App.DB.Raw(`
            SELECT c.* 
            FROM categories c
            INNER JOIN categories_junction cj ON c.id = cj.category_id
            INNER JOIN courses co ON cj.course_id = co.id
            INNER JOIN users u ON co.instructor_id = u.id
            WHERE u.last_name = ?
            GROUP BY c.id
        `, instructorLastName).Scan(&categories).Error
		if err != nil {
			helpers.ServerError(w, err)
			return
		}
	} else {
		err := m.App.DB.Order("id").Find(&categories).Error
		if err != nil {
			helpers.ServerError(w, err)
			return
		}
	}

	data := make(map[string]any)
	data["categories"] = categories
	data["currentPage"] = "categories"

	err := render.Template(w, r, "categories.page.gohtml", &models.TemplateData{
		Data: data,
		Form: forms.New(nil),
	})
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
}

func (m *Repository) Users(w http.ResponseWriter, r *http.Request) {
	log.Println("Rendering users.page.gohtml...")

	var users []models.User
	courseID := r.URL.Query().Get("courseID")
	categoryID := r.URL.Query().Get("categoryID")

	if courseID != "" {
		var cID uint
		_, err := fmt.Sscanf(courseID, "%d", &cID)
		if err != nil {
			helpers.ServerError(w, err)
			return
		}

		err = m.App.DB.Joins("JOIN enrollments ON enrollments.user_id = users.id").
			Where("enrollments.course_id = ?", cID).
			Order("users.id").
			Find(&users).Error
		if err != nil {
			helpers.ServerError(w, err)
			return
		}
	} else if categoryID != "" {
		var catID uint
		_, err := fmt.Sscanf(categoryID, "%d", &catID)
		if err != nil {
			helpers.ServerError(w, err)
			return
		}

		err = m.App.DB.Raw(`SELECT DISTINCT users.* FROM users
            INNER JOIN enrollments ON users.id = enrollments.user_id
            INNER JOIN courses ON enrollments.course_id = courses.id
            INNER JOIN categories_junction ON courses.id = categories_junction.course_id
            WHERE categories_junction.category_id = ?
            GROUP BY users.id
            HAVING COUNT(DISTINCT courses.id) = 
                (SELECT COUNT(*) FROM courses
                 JOIN categories_junction ON courses.id = categories_junction.course_id
                 WHERE categories_junction.category_id = ?)`, catID, catID).Scan(&users).Error

		if err != nil {
			helpers.ServerError(w, err)
			return
		}
	} else {
		err := m.App.DB.Order("id").Find(&users).Error
		if err != nil {
			helpers.ServerError(w, err)
			return
		}
	}

	data := make(map[string]any)
	data["users"] = users
	data["currentPage"] = "users"

	err := render.Template(w, r, "users.page.gohtml", &models.TemplateData{
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
