package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/config"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/handlers"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/helpers"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/models"
	"github.com/vladyslavpavlenko/dbms-lab2/internal/render"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func setup(app *config.AppConfig) error {
	// Get environment variables
	env, err := loadEvnVariables()
	if err != nil {
		return err
	}

	app.Env = env

	// Connect to the database and run migrations
	db, err := connectToPostgresAndMigrate(env)
	if err != nil {
		return err
	}

	app.DB = db

	// Run database migrations
	err = runDatabaseMigrations(db)
	if err != nil {
		return err
	}

	repo := handlers.NewRepo(app)
	handlers.NewHandlers(repo)
	render.NewRenderer(app)
	helpers.NewHelpers(app)

	return nil
}

func runDatabaseMigrations(db *gorm.DB) error {
	// Create tables
	err := db.AutoMigrate(&models.Category{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Course{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Enrollment{})
	if err != nil {
		return err
	}

	// Populate tables with initial data
	err = createInitialUsers(db)
	if err != nil {
		return fmt.Errorf("error creating initial users: %v", err)
	}

	err = createInitialCourseCategories(db)
	if err != nil {
		return fmt.Errorf("error creating initial categories: %v", err)
	}

	err = createInitialCourses(db)
	if err != nil {
		return fmt.Errorf("error creating initial courses: %v", err)
	}

	err = createInitialEnrollments(db)
	if err != nil {
		return fmt.Errorf("error creating initial enrollments: %v", err)
	}

	return nil
}

// connectToPostgresAndMigrate initializes a PostgreSQL db session and runs GORM migrations.
func connectToPostgresAndMigrate(env *config.EnvVariables) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
		env.PostgresHost, env.PostgresUser, env.PostgresDBName, env.PostgresPass)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("could not connect: ", err)
	}

	return db, nil
}

// createInitialUsers creates initial users types in users table.
func createInitialUsers(db *gorm.DB) error {
	var count int64

	if err := db.Model(&models.User{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	initialData := []models.User{
		{
			FirstName: "Vladyslav",
			LastName:  "Pavlenko",
			Email:     "vpavlenko@mail.com",
		},
		{
			FirstName: "Vadym",
			LastName:  "Ripa",
			Email:     "vripa@mail.com",
		},
		{
			FirstName: "Dmytro",
			LastName:  "Ostapenko",
			Email:     "dostapenko@mail.com",
		},
		{
			FirstName: "Valerii",
			LastName:  "Levchuk",
			Email:     "vlevchuk@mail.com",
		},
		{
			FirstName: "Volodymyr",
			LastName:  "Kravchuk",
			Email:     "vkravchuk@mail.com",
		},
	}

	if err := db.Create(&initialData).Error; err != nil {
		return err
	}

	return nil
}

// createInitialEnrollments creates initial course enrollments in enrollments table.
func createInitialEnrollments(db *gorm.DB) error {
	var count int64

	if err := db.Model(&models.Enrollment{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	initialData := []models.Enrollment{
		{
			UserID:   1,
			CourseID: 5,
		},
		{
			UserID:   2,
			CourseID: 4,
		},
		{
			UserID:   2,
			CourseID: 3,
		},
		{
			UserID:   4,
			CourseID: 5,
		},
		{
			UserID:   1,
			CourseID: 4,
		},
		{
			UserID:   1,
			CourseID: 3,
		},
		{
			UserID:   2,
			CourseID: 5,
		},
	}

	if err := db.Create(&initialData).Error; err != nil {
		return err
	}

	return nil
}

// createInitialCourseCategories creates initial categories in categories table.
func createInitialCourseCategories(db *gorm.DB) error {
	var count int64

	if err := db.Model(&models.Category{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	initialData := []models.Category{
		{Title: "Go"},
		{Title: "C++"},
		{Title: "C#"},
		{Title: "C"},
		{Title: "Rust"},
		{Title: "Ruby"},
		{Title: "Python"},
	}

	if err := db.Create(&initialData).Error; err != nil {
		return err
	}

	return nil
}

// createInitialCourses creates initial course categories in categories table.
func createInitialCourses(db *gorm.DB) error {
	var count int64

	if err := db.Model(&models.Course{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	initialData := []models.Course{
		{
			Title:        "Building Modern Web Applications with Go (Golang)",
			Description:  "Learn to program in Go from an award winning university professor.",
			InstructorID: 5,
			Categories:   []models.Category{{ID: 1}},
		},
		{
			Title:        "Programming For Beginners - Master the C, C#, C++ Languages",
			Description:  "C Programming will increase career options. Become a better dev in other languages by learning C. Pointers explained.",
			InstructorID: 4,
			Categories:   []models.Category{{ID: 2}, {ID: 3}, {ID: 4}},
		},
		{
			Title:        "Go vs. Rust: Building High-Performance Applications",
			Description:  "Dive into the world of system programming with Go and Rust. This course explores both languages' syntax, memory management, concurrency models, and ecosystem to determine the best fit for various high-performance applications.",
			InstructorID: 5,
			Categories:   []models.Category{{ID: 1}, {ID: 5}},
		},
		{
			Title:        "C++ and C#: Object-Oriented Programming Masterclass",
			Description:  "Unravel the complexities of object-oriented programming (OOP) with C++ and C#. This comprehensive course covers the foundations of OOP, including classes, inheritance, polymorphism, and encapsulation.",
			InstructorID: 4,
			Categories:   []models.Category{{ID: 2}, {ID: 3}},
		},
		{
			Title:        "From Rust to Ruby: System to Scripting",
			Description:  "Explore the spectrum of programming from system-level to web scripting with Rust and Ruby.",
			InstructorID: 3,
			Categories:   []models.Category{{ID: 6}, {ID: 5}},
		},
	}

	if err := db.Create(&initialData).Error; err != nil {
		return err
	}

	return nil
}

// loadEvnVariables loads variables from the .env file.
func loadEvnVariables() (*config.EnvVariables, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error getting environment variables: %v", err)
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPass := os.Getenv("POSTGRES_PASS")
	postgresDBName := os.Getenv("POSTGRES_DBNAME")

	return &config.EnvVariables{
		PostgresHost:   postgresHost,
		PostgresUser:   postgresUser,
		PostgresPass:   postgresPass,
		PostgresDBName: postgresDBName,
	}, nil
}
