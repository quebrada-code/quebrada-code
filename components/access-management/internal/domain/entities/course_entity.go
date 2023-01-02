package entities

import "time"

type CourseEntity struct {
	ID          int64
	Title       string
	Description string
	Featured    bool
	Category    CourseCategoryEntity
	Modules     []ModuleEntity
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
