package models

// User is the user model.
type User struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string `gorm:"unique;"`
}
