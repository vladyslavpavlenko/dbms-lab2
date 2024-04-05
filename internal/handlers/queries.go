package handlers

import (
	"fmt"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/models"
	"net/http"
)

// CoursesInCategory queries for the courses in a specific category specified by the user.
func (m *Repository) CoursesInCategory(w http.ResponseWriter, r *http.Request) {
	categoryTitle := r.URL.Query().Get("categoryTitle")

	var courses []models.Course
	err := m.App.DB.Joins("JOIN categories_junction on categories_junction.course_id = courses.id").
		Joins("JOIN categories on categories.id = categories_junction.category_id").
		Where("categories.title = ?", categoryTitle).
		Preload("Categories").
		Preload("Instructor").
		Find(&courses).Error

	if err != nil {
		// Handle error
	}

	// Process courses
}

// CoursesByInstructor queries for the courses taught by an instructor specified by the user.
func (m *Repository) CoursesByInstructor(w http.ResponseWriter, r *http.Request) {
	instructorLastName := r.URL.Query().Get("instructorLastName")

	var courses []models.Course
	err := m.App.DB.Where("instructor_id IN (?)",
		m.App.DB.Select("id").Where("last_name = ?", instructorLastName).Table("users")).
		Preload("Categories").
		Preload("Instructor").
		Find(&courses).Error

	if err != nil {
		// Handle error
	}

	// Process courses
}

// UsersInCourse queries for the users enrolled in a specific course.
func (m *Repository) UsersInCourse(w http.ResponseWriter, r *http.Request) {
	courseTitle := r.URL.Query().Get("courseTitle")

	var users []models.User
	err := m.App.DB.Joins("JOIN enrollments on enrollments.user_id = users.id").
		Joins("JOIN courses on courses.id = enrollments.course_id").
		Where("courses.title = ?", courseTitle).
		Find(&users).Error

	if err != nil {
		// Handle error
	}

	// Process users
}

// CategoriesByInstructor queries for categories of courses taught by a specific instructor.
func (m *Repository) CategoriesByInstructor(w http.ResponseWriter, r *http.Request) {
	instructorID := r.URL.Query().Get("instructorID")

	var categories []models.Category
	err := m.App.DB.Joins("JOIN categories_junction on categories_junction.category_id = categories.id").
		Joins("JOIN courses on courses.id = categories_junction.course_id").
		Where("courses.instructor_id = ?", instructorID).
		Distinct().
		Find(&categories).Error

	if err != nil {
		// Handle error
	}

	// Process categories
}

// CoursesWithMoreThanXCategories queries for the courses that are associated with more than X categories.
func (m *Repository) CoursesWithMoreThanXCategories(w http.ResponseWriter, r *http.Request) {
	var x int
	fmt.Sscanf(r.URL.Query().Get("x"), "%d", &x)

	var courses []models.Course
	err := m.App.DB.
		Joins("JOIN categories_junction on categories_junction.course_id = courses.id").
		Group("courses.id").
		Having("COUNT(categories_junction.category_id) > ?", x).
		Preload("Categories").
		Preload("Instructor").
		Find(&courses).Error

	if err != nil {
		// Handle error
	}

	// Process courses
}

// SharedCoursesBetweenUsers queries for the courses shared by two specified users.
func (m *Repository) SharedCoursesBetweenUsers(w http.ResponseWriter, r *http.Request) {
	userID1 := r.URL.Query().Get("userID1")
	userID2 := r.URL.Query().Get("userID2")

	var courses []models.Course
	err := m.App.DB.Raw(`
        SELECT * FROM courses 
        WHERE EXISTS (
            SELECT course_id FROM enrollments WHERE user_id = ? AND course_id = courses.id
        ) AND EXISTS (
            SELECT course_id FROM enrollments WHERE user_id = ? AND course_id = courses.id
        )`, userID1, userID2).Scan(&courses).Error

	if err != nil {
		// Handle error
	}

	// Process courses
}

// UsersEnrolledInAllCoursesOfCategory queries for the users enrolled in all courses of a specific category/
func (m *Repository) UsersEnrolledInAllCoursesOfCategory(w http.ResponseWriter, r *http.Request) {
	categoryID := r.URL.Query().Get("categoryID")

	var users []models.User
	err := m.App.DB.Raw(`
		SELECT u.* FROM users u
		JOIN enrollments e ON u.id = e.user_id
		JOIN courses c ON e.course_id = c.id
		JOIN categories_junction cj ON c.id = cj.course_id
		WHERE cj.category_id = ?
		GROUP BY u.id
		HAVING COUNT(DISTINCT c.id) = (
			SELECT COUNT(DISTINCT course_id) FROM categories_junction WHERE category_id = ?
		)`, categoryID, categoryID).Scan(&users).Error

	if err != nil {
		// Handle error
	}

	// Process users
}

// CategoriesSharedByAllCoursesOfInstructor queries for the categories shared by all courses of a specific instructor.
func (m *Repository) CategoriesSharedByAllCoursesOfInstructor(w http.ResponseWriter, r *http.Request) {
	instructorID := r.URL.Query().Get("instructorID")

	var categories []models.Category
	err := m.App.DB.Raw(`
		SELECT cat.* FROM categories cat
		JOIN categories_junction cj ON cat.id = cj.category_id
		JOIN courses c ON cj.course_id = c.id
		WHERE c.instructor_id = ?
		GROUP BY cat.id
		HAVING COUNT(DISTINCT c.id) = (
			SELECT COUNT(*) FROM courses WHERE instructor_id = ?
		)`, instructorID, instructorID).Scan(&categories).Error

	if err != nil {
		// Handle error
	}

	// Process categories
}
