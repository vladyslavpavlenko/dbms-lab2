package config

import (
	"gorm.io/gorm"
	"html/template"
	"log"
)

// AppConfig holds the application config.
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	DB            *gorm.DB
	Env           *EnvVariables
}

// EnvVariables holds environment variables used in the application.
type EnvVariables struct {
	PostgresHost   string
	PostgresUser   string
	PostgresPass   string
	PostgresDBName string
	JWTSecret      string
}
