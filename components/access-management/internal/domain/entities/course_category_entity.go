package entities

import "time"

type CourseCategoryEntity struct {
	ID        int64
	Title     string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
