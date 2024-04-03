package models

// Enrollment is the enrollment model.
type Enrollment struct {
	UserID   uint `gorm:"primaryKey;autoIncrement:false;not null"`
	User     User
	CourseID uint `gorm:"primaryKey;autoIncrement:false;not null"`
	Course   Course
}
