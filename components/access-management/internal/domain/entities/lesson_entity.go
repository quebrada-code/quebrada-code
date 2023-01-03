package entities

import "time"

type LessonEntity struct {
	ID          int64
	Title       string
	Description string
	Media       MediaEntity
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
