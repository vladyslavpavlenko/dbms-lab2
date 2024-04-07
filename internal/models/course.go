package models

// Course is the course model.
type Course struct {
	ID           uint
	Title        string
	Description  string
	Categories   []Category `gorm:"many2many:categories_junction;"`
	InstructorID uint
	Instructor   User         `gorm:"constraint:OnDelete:CASCADE;"`
	Enrollments  []Enrollment `gorm:"constraint:OnDelete:CASCADE;"`
}
