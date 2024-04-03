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

	mux.Get("/courses", handlers.Repo.Courses)
	mux.Get("/categories", handlers.Repo.Categories)
	mux.Get("/courses-and-categories", handlers.Repo.CoursesAndCategories)
	mux.Get("/enrollments", handlers.Repo.Enrollments)
	mux.Get("/users", handlers.Repo.Users)

	//mux.Post("/courses", handlers.Repo.PostCourses)
	//mux.Post("/categories", handlers.Repo.PostCategories)
	//mux.Post("/courses-and-categories", handlers.Repo.PostCoursesAndCategories)
	//mux.Post("/enrollments", handlers.Repo.PostEnrollments)
	//mux.Post("/users", handlers.Repo.PostUsers)

	return mux
}
