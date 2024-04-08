package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/config"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)

	// create
	mux.Post("/courses/create", handlers.Repo.CreateCourse)
	mux.Post("/categories/create", handlers.Repo.CreateCategory)
	mux.Post("/categories_junction/create", handlers.Repo.CreateCategoryJunction)
	mux.Post("/enrollments/create", handlers.Repo.CreateEnrollment)
	mux.Post("/users/create", handlers.Repo.CreateUser)

	// read
	mux.Get("/courses", handlers.Repo.Courses)
	mux.Get("/categories", handlers.Repo.Categories)
	mux.Get("/categories_junction", handlers.Repo.CategoriesJunction)
	mux.Get("/enrollments", handlers.Repo.Enrollments)
	mux.Get("/users", handlers.Repo.Users)

	// update
	mux.Post("/courses/update", handlers.Repo.UpdateCourse)
	mux.Post("/categories/update", handlers.Repo.UpdateCategory)
	mux.Post("/categories_junction/update", handlers.Repo.UpdateCategoryJunction)
	mux.Post("/enrollments/update", handlers.Repo.UpdateEnrollment)
	mux.Post("/users/update", handlers.Repo.UpdateUser)

	// delete
	mux.Delete("/courses/delete/{id}", handlers.Repo.DeleteCourse)
	mux.Delete("/categories/delete/{id}", handlers.Repo.DeleteCategory)
	mux.Delete("/categories_junction/delete/{category_id}/{course_id}", handlers.Repo.DeleteCategoryJunction)
	mux.Delete("/enrollments/delete/{user_id}/{course_id}", handlers.Repo.DeleteEnrollment)
	mux.Delete("/users/delete/{id}", handlers.Repo.DeleteUser)

	// queries
	mux.Get("/queries", handlers.Repo.Queries)
	mux.Get("/queries/courses_in_category", handlers.Repo.CoursesInCategory)
	return mux
}
