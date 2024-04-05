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

	mux.Get("/queries", handlers.Repo.Queries)
	mux.Get("/courses", handlers.Repo.Courses)
	mux.Get("/categories", handlers.Repo.Categories)
	mux.Get("/categories_junction", handlers.Repo.CategoriesJunction)
	mux.Get("/enrollments", handlers.Repo.Enrollments)
	mux.Get("/users", handlers.Repo.Users)

	mux.Post("/courses/update", handlers.Repo.UpdateCourse)
	mux.Post("/categories/update", handlers.Repo.UpdateCategory)
	mux.Post("/categories_junction/update", handlers.Repo.UpdateCategoriesJunction)
	mux.Post("/enrollments/update", handlers.Repo.UpdateEnrollment)
	mux.Post("/users/update", handlers.Repo.UpdateUser)

	return mux
}
