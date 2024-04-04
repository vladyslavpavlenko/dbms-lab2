package database

import "github.com/vladyslavpavlenko/dbms-lab2/internal/config"

// Repo the repository used by the handlers.
var Repo *Repository

// Repository is the repository type.
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository.
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewTestRepo creates a new test repository.
func NewTestRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for handlers.
func NewHandlers(r *Repository) {
	Repo = r
}
