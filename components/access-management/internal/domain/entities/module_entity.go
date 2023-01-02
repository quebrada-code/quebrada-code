package entities

import "time"

type ModuleEntity struct {
	ID          int64
	Title       string
	Description string
	Active      bool
	Order       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
